package handle

import (
	"Liature-Server/socialAuth/authgoogle"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/*
// SocialAuthNaver 는 Social 로그인 처리를 담당하는 핸들러입니다.
func SocialAuthNaver(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if req.Method == "POST" {
		authnaver.AuthNaver(w, req, "login", "naver")
	} else if req.Method == "GET" {
		authnaver.AuthNaver(w, req, "callback", "naver")
	}
}
*/

/*
// SocialAuthFacebook 는 Social 로그인 처리를 담당하는 핸들러입니다.
func SocialAuthFacebook(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if req.Method == "POST" {
		authfacebook.AuthFacebook(w, req, "login", "facebook")
	} else if req.Method == "GET" {
		authfacebook.AuthFacebook(w, req, "callback", "facebook")
	}
}
*/

// SocialAuthGoogle 는 Social 로그인 처리를 담당하는 핸들러입니다.
func SocialAuthGoogle(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if req.Method == "POST" {
		authgoogle.AuthGoogle(w, req, "login", "google")
	} else if req.Method == "GET" {
		authgoogle.AuthGoogle(w, req, "callback", "google")
	}
}
