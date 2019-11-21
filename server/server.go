package server

import (
	"Liature-Server/client"
	"Liature-Server/handle"
	"Liature-Server/message"
	"Liature-Server/serversession"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	sessions "github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
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

	// 소켓 버퍼 크기
	socketBufferSize = 1024
)

var (
	upgrader = &websocket.Upgrader{
		ReadBufferSize:  socketBufferSize,
		WriteBufferSize: socketBufferSize,
	}
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

	sv.router.GET("/", func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	})
	sv.router.GET("/login", func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	})
	sv.router.GET("/rooms/messages", message.RetrieveMessages)

	sv.router.GET("/ws/:room_id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		socket, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal("ServeHTTP:", err)
			return
		}
		client.NewClient(socket, ps.ByName("room_id"), serversession.GetCurrentUser(r))
	})

	//sv.router.POST("/social/google", handle.SocialAuthGoogle)
	sv.router.GET("/auth/callback/google", handle.SocialAuthGoogle)

	sv.neg.UseHandler(sv.router)
	return sv, nil
}

// Run 함수는 Server를 실행합니다.
func (s *Server) Run(port string) {
	s.neg.Run(port)
}
