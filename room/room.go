package room

import (
	"Liature-Server/db"
	"fmt"
	"os"

	"github.com/unrolled/render"
	"gopkg.in/mgo.v2/bson"
)

// Room 은 채팅방 정보를 담고 있습니다.
type Room struct {
	ID   bson.ObjectId `bson:"_id" json:"id"`
	Name string        `bson:"name" json:"name"`
}

var (
	renderer *render.Render
	mongoDB  *db.MongoDB
)

func init() {
	renderer = render.New()
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

// CreateRoom 은 받은 이름으로 방을 생성합니다.
func CreateRoom(roomName string) error {
	r := new(Room)

	session := mongoDB.Session.Copy()
	defer session.Close()

	r.ID = bson.NewObjectId()
	r.Name = roomName
	c := session.DB("test").C("rooms")

	if err := c.Insert(r); err != nil {
		return err
	}

	return nil
}

func retrieveRooms() ([]Room, error) {
	// 몽고DB 세션 생성
	session := mongoDB.Session.Copy()
	// 몽고DB 세션을 닫는 코드를 defer로 등록
	defer session.Close()

	var rooms []Room
	// 모든 room 정보 조회
	err := session.DB("test").C("rooms").Find(nil).All(&rooms)
	if err != nil {
		return nil, err
	}

	// room 조회 결과 반환
	return rooms, nil
}
