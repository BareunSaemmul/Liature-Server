package server

import (
	"Liature-Server/handle"
	"Liature-Server/message"
	"Liature-Server/room"
	"Liature-Server/serversession"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

// Server 객체는 Server에 대한 정보를 담고 있습니다.
type Server struct {
	neg    *negroni.Negroni
	router *httprouter.Router
}

const (
	// 애플리케이션에서 사용할 세션의 키 정보
	sessionKey    = "simple_chat_session"
	sessionSecret = "simple_chat_session_secret"
)

// New 함수는 Server에 대한 설정을 담당합니다
func New() (*Server, error) {
	sv := new(Server)
	sv.router = httprouter.New()
	sv.neg = negroni.Classic()
	store := cookiestore.New([]byte(sessionSecret))
	sv.neg.Use(sessions.Sessions(sessionKey, store))
	sv.neg.Use(serversession.LoginRequired("/login", "/"))

	err := handle.InitMongo("mongodb://127.0.0.1:27017")
	if err != nil {
		return nil, err
	}
	err = message.InitMongo("mongodb://127.0.0.1:27017")
	if err != nil {
		return nil, err
	}

	err = room.InitMongo("mongodb://127.0.0.1:27017")
	if err != nil {
		return nil, err
	}

	roomList := []string{
		"Dajeon",
		"Daegu",
		"Gwangju",
	}

	for i := 0; i < len(roomList); i++ {
		room.CreateRoom(roomList[i])
	}

	sv.router.GET("/", handle.IndexPage)

	sv.router.GET("/login", handle.LoginPage)

	sv.router.GET("/rooms", handle.RetrieveRooms)
	sv.router.POST("/rooms/messages", handle.RetrieveMessages)

	sv.router.GET("/ws/room/:area", handle.NewClient)

	//sv.router.POST("/social/google", handle.SocialAuthGoogle)
	sv.router.GET("/auth/callback/google", handle.SocialAuthGoogle)

	sv.neg.UseHandler(sv.router)
	return sv, nil
}

// Run 함수는 Server를 실행합니다.
func (s *Server) Run(port string) {
	s.neg.Run(port)
}
