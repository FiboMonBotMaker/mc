/*
フロントエンドで利用するJSONとMP3データを提供するAPI
*/
package apis

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"hureru_button.com/hureru_button/db"
)

type MoneyJson struct {
	UUID  string  `json:"uuid"`
	Money float64 `json:"money"`
}

type ReceiveJson struct {
	Users []MoneyJson `json:"users"`
	Pass  string      `json:"pass"`
}

func SetApis(e echo.Echo) {
	e.POST("/adding", addingMoney)
	e.POST("/economy", getEcoList)
}

func addingMoney(c echo.Context) error {
	json := new(ReceiveJson)
	if err := c.Bind(json); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "adding money")
}

func getEcoList(c echo.Context) error {
	params, err := c.FormParams()
	if err != nil {
		c.Error(err)
		return err
	}
	uuids := params["uuids"]
	json := make([]MoneyJson, 0, 20)

	for _, v := range db.GetData(uuids) {
		json = append(json, MoneyJson{UUID: v.UUID, Money: v.Money})
	}
	return c.JSON(http.StatusOK, json)
}
