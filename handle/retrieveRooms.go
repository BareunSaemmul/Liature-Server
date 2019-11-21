package handle

import (
	"Liature-Server/room"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// RetrieveRooms 핸들러는 모든 Room의 정보를 응답합니다.
func RetrieveRooms(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	r, err := room.RetrieveRooms()
	if err != nil {
		renderer.JSON(w, http.StatusInternalServerError, err)
	}
	renderer.JSON(w, http.StatusOK, r)
}
