package handle

import (
	"Liature-Server/client"
	"Liature-Server/serversession"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

const (
	// 소켓 버퍼 크기
	socketBufferSize = 1024
)

var (
	upgrader = &websocket.Upgrader{
		ReadBufferSize:  socketBufferSize,
		WriteBufferSize: socketBufferSize,
	}
)

// NewClient 는 새로운 클라이언트를 생성합니다.
func NewClient(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	client.NewClient(socket, ps.ByName("room_id"), serversession.GetCurrentUser(r))
}
