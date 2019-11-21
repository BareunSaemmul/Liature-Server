package handle

import (
	"Liature-Server/message"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// RetrieveMessages 핸들러는 모든 Room의 정보를 응답합니다.
func RetrieveMessages(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	m, err := message.RetrieveMessages(w, req, ps)
	if err != nil {
		renderer.JSON(w, http.StatusInternalServerError, err)
		return
	}
	renderer.JSON(w, http.StatusOK, m)
}
