package handle

import (
	"Liature-Server/db"
	"Liature-Server/room"
	"fmt"
	"os"

	"github.com/unrolled/render"
)

// Res 는 Default 응답값인 성공여부를 담는 구조체입니다.
type Res struct {
	IsSuccess bool `json:"isSuccess"`
}

var (
	renderer *render.Render
	mongoDB  *db.MongoDB
)

var currentUserLocal string

func init() {
	currentUserLocal = "대전"
	renderer = render.New()

	roomList := []string{
		"대전",
		"대구",
		"광주",
	}

	room.InitMongo("mongodb://127.0.0.1:27017")

	for i := 0; i < len(roomList); i++ {
		room.CreateRoom(roomList[i])
	}
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
