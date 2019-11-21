package handle

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//LoginPage 는 메인 HTML 파일을 응답합니다.
func LoginPage(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	renderer.HTML(w, http.StatusOK, "login", nil)
}
