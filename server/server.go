package server

import (
	"github.com/codegangsta/negroni"
	sessions "github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
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
)

// New 함수는 Server에 대한 설정을 담당합니다
func New() (*Server, error) {
	sv := new(Server)
	sv.router = httprouter.New()
	sv.neg = negroni.Classic()
	store := cookiestore.New([]byte(sessionSecret))
	sv.neg.Use(sessions.Sessions(sessionKey, store))

	err := handle.InitMongo("mongodb://127.0.0.1:27017")
	if err != nil {
		return nil, err
	}

	sv.router.POST("/api/account/login", handle.Login)

	sv.router.POST("/social/facebook", handle.SocialAuthFacebook)
	sv.router.GET("/social/facebook/callback", handle.SocialAuthFacebook)

	sv.router.POST("/social/google", handle.SocialAuthGoogle)
	sv.router.GET("/social/google/callback", handle.SocialAuthGoogle)

	sv.router.POST("/social/naver", handle.SocialAuthNaver)
	sv.router.GET("/social/naver/callback", handle.SocialAuthNaver)

	sv.neg.UseHandler(sv.router)
	return sv, nil
}

// Run 함수는 Server를 실행합니다.
func (s *Server) Run(port string) {
	s.neg.Run(port)
}
