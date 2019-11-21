package handle

import (
	"Liature/mainServer/src/Liature-Server/db"
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
