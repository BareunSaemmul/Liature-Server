package handle

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//IndexPage 는 메인 HTML 파일을 응답합니다.
func IndexPage(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	renderer.HTML(w, http.StatusOK, "index", nil)
}
