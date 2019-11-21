package message

import (
	"Liature-Server/db"
	"Liature-Server/serversession"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2/bson"
)

const messageFetchSize = 10

var (
	renderer *render.Render
	mongoDB  *db.MongoDB
)

// Message 는 메시지 정보를 담습니다
type Message struct {
	ID        bson.ObjectId              `bson:"_id" json:"id"`
	RoomID    bson.ObjectId              `bson:"room_id" json:"room_id"`
	Content   string                     `bson:"content" json:"content"`
	CreatedAt time.Time                  `bson:"created_at" json:"created_at"`
	User      *serversession.SessionUser `bson:"user" json:"user"`
}

// InitMongo 는 몽고DB의 초기 설정을 하는 함수입니다.
func InitMongo(addr string) error {
	var dbID, dbPw string
	fi, err := os.Open("db_account.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fmt.Fscan(fi, &dbID, &dbPw)

	m, err := db.NewMongoDB(addr)
	if err != nil {
		return err
	}
	mongoDB = m
	if err := mongoDB.Session.DB("admin").Login(dbID, dbPw); err != nil {
		panic(err)
	}
	return nil
}

// Create 는 메시지 생성
func (m *Message) Create() error {
	// 몽고DB 세션 생성
	session := mongoDB.Session.Copy()
	// 몽고DB 세션을 닫는 코드를 defer로 등록
	defer session.Close()
	// 몽고DB 아이디 생성
	m.ID = bson.NewObjectId()
	// 메시지 생성 시간 기록
	m.CreatedAt = time.Now()
	// message 정보 저장을 위한 몽고DB 컬렉션 객체 생성
	c := session.DB("test").C("messages")

	// messages 컬렉션에 message 정보 저장
	if err := c.Insert(m); err != nil {
		return err
	}
	return nil
}

// RetrieveMessages 는 메시지를 조회합니다.
func RetrieveMessages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 몽고DB 세션 생성
	session := mongoDB.Session.Copy()
	// 몽고DB 세션을 닫는 코드를 defer로 등록
	defer session.Close()

	// 쿼리 매개변수로 전달된 limit 값 확인
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		// 정상적인 limit 값이 전달되지 않으면 limit를 messageFetchSize으로 세팅
		limit = messageFetchSize
	}

	var messages []Message
	// _id 역순으로 정렬하여 limit 수만큼 message 조회
	err = session.DB("test").C("messages").
		Find(bson.M{"room_id": bson.ObjectIdHex(ps.ByName("id"))}).
		Sort("-_id").Limit(limit).All(&messages)
	if err != nil {
		// 오류 발생 시 500 에러 반환
		renderer.JSON(w, http.StatusInternalServerError, err)
		return
	}
	// message 조회 결과 반환
	renderer.JSON(w, http.StatusOK, messages)
}
