package handle

import (
	"Liature/mainServer/src/Liature-Server/server"
	"net/http"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
)

func logout(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	// 세션에서 사용자 정보 제거 후 로그인 페이지로 이동
	sessions.GetSession(req).Delete(server.GetCurrentUserKey())
	http.Redirect(w, req, "TODO-LOGIN-URL", http.StatusFound)
}
