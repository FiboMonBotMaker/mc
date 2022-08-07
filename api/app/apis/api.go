/*
Moneyデータをやり取りするためのAPI
*/
package apis

import (
	"os"
	"net/http"

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
		return 
	}

	return c.JSON(http.StatusOK, "adding money")
}

func getEcoList(c echo.Context) error {
	uuid := c.QueryParam("uuid")
	json := make([]MoneyJson, 0, 20)

	for _, v := range db.GetData(uuid) {
		json = append(json, MoneyJson{Uuid: v.Uuid, Money: v.Money})
	}
	return c.JSON(http.StatusOK, json)
}
