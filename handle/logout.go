package handle

import (
	"net/http"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
)

const (
	currentUserKey = "oauth2_current_user" // 세션에 저장되는 CurrentUser의 키
)

func logout(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	// 세션에서 사용자 정보 제거 후 로그인 페이지로 이동
	sessions.GetSession(req).Delete(currentUserKey)
	http.Redirect(w, req, "/login", http.StatusFound)
}
