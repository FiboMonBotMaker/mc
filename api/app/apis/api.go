/*
Moneyデータをやり取りするためのAPI
*/
package apis

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"

	"mc.com/mc/db"
)

var pass = os.Getenv("API_PASSWORD")

type MoneyJson struct {
	Uuid  string  `json:"uuid"`
	Money float64 `json:"money"`
}

type ReceiveJson struct {
	Users []MoneyJson `json:"users"`
	Pass  string      `json:"pass"`
}

type Message struct {
	Message string `json:"message"`
}

func SetApis(e echo.Echo) {
	e.POST("/adding", addingMoney)
	e.GET("/economy", getEcoList)
}

func addingMoney(c echo.Context) error {
	json := new(ReceiveJson)
	if err := c.Bind(json); err != nil {
		return err
	}
	if json.Pass != pass {
		return echo.ErrForbidden
	}
	for _, v := range json.Users {
		err := db.AddMoney(v.Uuid, v.Money)
		if err != nil {
			return c.JSON(514, Message{Message: err.Error()})
		}
	}

	return c.JSON(http.StatusOK, Message{Message: "Adding success"})
}

func getEcoList(c echo.Context) error {
	uuid := c.QueryParam("uuid")
	json := make([]MoneyJson, 0, 20)

	for _, v := range db.GetData(uuid) {
		json = append(json, MoneyJson{Uuid: v.Uuid, Money: v.Money})
	}
	return c.JSON(http.StatusOK, json)
}
