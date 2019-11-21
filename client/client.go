package client

import (
	"Liature-Server/message"
	"Liature-Server/serversession"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

// 현재 접속 중인 전체 클라이언트 리스트
var clients []*Client

// Client 정보를 설정합니다.
type Client struct {
	Conn *websocket.Conn       // 웹소켓 커넥션
	Send chan *message.Message // 메시지 전송용 채널

	RoomID string                     // 현재 접속한 채팅방 아이디
	User   *serversession.SessionUser // 현재 접속한 사용자 정보
}

const messageBufferSize = 256

// NewClient 는 클라이언트를 생성합니다.
func NewClient(conn *websocket.Conn, roomID string, u *serversession.SessionUser) {
	// 새로운 클라이언트 생성
	c := &Client{
		Conn:   conn,
		Send:   make(chan *message.Message, messageBufferSize),
		RoomID: roomID,
		User:   u,
	}

	// clients 목록에 새로 생성한 클라이언트 추가
	clients = append(clients, c)

	// 메시지 수신/전송 대기
	go c.ReadLoop()
	go c.WriteLoop()
}

// Close 메서드는 웹 소켓을 닫습니다.
func (c *Client) Close() {
	// clients 목록에서 종료된 클라이언트 제거
	for i, client := range clients {
		if client == c {
			clients = append(clients[:i], clients[i+1:]...)
			break
		}
	}

	// send 채널 닫음
	close(c.Send)

	// 웹소켓 커넥션 종료
	c.Conn.Close()
	log.Printf("close connection. addr: %s", c.Conn.RemoteAddr())
}

// ReadLoop 메서드는 메시지를 수신하기 위한 메서드입니다.
func (c *Client) ReadLoop() {
	// 메시지 수신 대기
	for {
		m, err := c.Read()
		if err != nil {
			// 오류가 발생하면 메시지 수신 루프 종료
			log.Println("read message error: ", err)
			break
		}

		// 메시지가 수신되면 수신된 메시지를 몽고DB에 생성하고 모든 clients에 전달
		m.Create()
		broadcast(m)
	}
	c.Close()
}

// WriteLoop 메서드는 메시지 전달을 위한 메서드입니다.
func (c *Client) WriteLoop() {
	// 클라이언트의 send 채널 메시지 수신 대기
	for msg := range c.Send {
		// 클라이언트의 채팅방 아이디와 전달된 메시지의 채팅방 아이디가 일치하면 웹소켓에 메시지 전달
		if c.RoomID == msg.RoomID.Hex() {
			c.Write(msg)
		}
	}
}

func broadcast(m *message.Message) {
	// 모든 클라이언트의 send 채널에 메시지 전달
	for _, client := range clients {
		client.Send <- m
	}
}

// Read 메서드는 메시지를 읽습니다.
func (c *Client) Read() (*message.Message, error) {
	var msg *message.Message

	// 웹소켓 커넥션에 JSON 형태의 메시지가 전달되면 Message 타입으로 메시지를 읽음
	if err := c.Conn.ReadJSON(&msg); err != nil {
		return nil, err
	}

	// Message에 현재 시간과 사용자 정보 세팅
	msg.CreatedAt = time.Now()
	msg.User = c.User

	if msg.User == nil {
		msg.User = new(serversession.SessionUser)
		msg.User.UID = "guest_" + time.Now().Format("2006-01-02-15:04:05")
		msg.User.Name = "guest"
		msg.User.Email = "guest@guest.com"
		msg.User.AvatarURL = "https://cdn.pixabay.com/photo/2016/08/31/11/54/user-1633249_960_720.png"
	}

	log.Println("read from websocket:", msg)

	// 메시지 정보 반환
	return msg, nil
}

// Write 메서드는 웹 소켓 커넥션에 메시지를 전달합니다.
func (c *Client) Write(m *message.Message) error {
	log.Println("write to websocket:", m)

	// 웹소켓 커넥션에 JSON 형태로 메시지 전달
	return c.Conn.WriteJSON(m)
}
