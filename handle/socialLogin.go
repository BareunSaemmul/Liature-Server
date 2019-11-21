package handle

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// SocialAuthNaver 는 Social 로그인 처리를 담당하는 핸들러입니다.
func SocialAuthNaver(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if req.Method == "POST" {
		authnaver.AuthNaver(w, r, "login", "naver")
	} else req.Method == "GET" {
		authnaver.AuthNaver(w, r, "callback", "naver")
	}
}

// SocialAuthFacebook 는 Social 로그인 처리를 담당하는 핸들러입니다.
func SocialAuthFacebook(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if req.Method == "POST" {
		authfacebook.AuthFacebook(w, r, "login", "facebook")
	} else req.Method == "GET" {
		authfacebook.AuthFacebook(w, r, "callback", "facebook")
	}
}

// SocialAuthGoogle 는 Social 로그인 처리를 담당하는 핸들러입니다.
func SocialAuthGoogle(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if req.Method == "POST" {
		authgoogle.AuthGoogle(w, r, "login", "facebook")
	} else req.Method == "GET" {
		authgoogle.AuthGoogle(w, r, "callback", "facebook")
	}
}
