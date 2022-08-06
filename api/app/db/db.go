package db

import (
	"fmt"
	"log"

	"xorm.io/xorm"
)

type LiteEco struct {
	ID    int     `xorm:"id pk"`
	UUID  string  `xorm:"uuid varchar(36)"`
	Money float64 `xorm:"money"`
}

var (
	USER     string
	PASSWORD string
	ADDRESS  string
	DB       string
)

func SetInit(password string, user string, address string, db string) {
	PASSWORD = password
	USER = user
	ADDRESS = address
	DB = db
}

func getEngine() *xorm.Engine {
	st := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=true", USER, PASSWORD, ADDRESS, DB)
	engine, err := xorm.NewEngine("mysql", st)
	if err != nil {
		log.Fatal(err)
	}
	return engine
}

func GetData(uuids []string) []LiteEco {
	engine := getEngine()
	liteEcos := make([]LiteEco, 0, 20)
	err := xorm.ErrObjectIsNil
	if len(uuids) == 0 {
		err = engine.Desc("money").Find(&liteEcos)
	} else {
		err = engine.In("uuid", uuids).Desc("money").Find(&liteEcos)
	}
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	return liteEcos
}

func AddMoney(uuid string, money float64) {
	engine := getEngine()
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
Retake:
	liteEco := LiteEco{}
	result, err := engine.Where("uuid = ?", uuid).Get(&liteEco)
	if err != nil {
		session.Rollback()
		log.Fatalf("err: %v", err)
	}
	if !result {
		liteEco.Money = 3000
		liteEco.UUID = uuid
		result, err := engine.InsertOne(liteEco)
		if err != nil {
			session.Rollback()
			log.Fatalf("err: %v", err)
		}
		if result == 0 {
			session.Rollback()
			log.Fatalf("missing value %d result", result)
		}
		goto Retake
	}
	liteEco.Money += money
	bresult, err := engine.ID(liteEco.ID).Cols("money").Update(&liteEco)
	if err != nil {
		session.Rollback()
		log.Fatalf("err: %v", err)
	}
	if bresult == 0 {
		session.Rollback()
		log.Fatalf("missing value %d result", bresult)
	}
	err = session.Commit()
	if err != nil {
		log.Fatalf("err: %v", err)
	}
}
