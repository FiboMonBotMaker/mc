package db

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type LiteEco struct {
	ID    int     `xorm:"id pk"`
	Uuid  string  `xorm:"uuid varchar(36) notnull"`
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

func GetData(uuid string) []LiteEco {
	engine := getEngine()
	liteEcos := make([]LiteEco, 0, 20)
	err := xorm.ErrObjectIsNil
	if uuid == "" {
		err = engine.Table("lite_eco").Desc("money").Find(&liteEcos)
	} else {
		err = engine.Table("lite_eco").Where("uuid = ?", uuid).Desc("money").Find(&liteEcos)
	}
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	return liteEcos
}

func AddMoney(uuid string, money float64) error {
	engine := getEngine()
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
Retake:
	liteEco := LiteEco{}
	result, err := engine.Where("uuid = ?", uuid).Get(&liteEco)
	if err != nil {
		session.Rollback()
		return err
	}
	if !result {
		liteEco.Money = 3000
		liteEco.Uuid = uuid
		result, err := engine.InsertOne(liteEco)
		if err != nil {
			session.Rollback()
			return err
		}
		if result == 0 {
			session.Rollback()
			return err
		}
		goto Retake
	}
	liteEco.Money += money
	bresult, err := engine.ID(liteEco.ID).Cols("money").Update(&liteEco)
	if err != nil {
		session.Rollback()
		return err
	}
	if bresult == 0 {
		session.Rollback()
		return err
	}
	err = session.Commit()
	if err != nil {
		return err
	}
	return nil
}
