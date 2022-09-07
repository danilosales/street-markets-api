package strmarket

import (
	"net/http"

	"github.com/danilosales/api-street-markets/internal/database"
	"github.com/danilosales/api-street-markets/internal/model"
	"github.com/gin-gonic/gin"
)

func CreateStreetMarket(c *gin.Context) {

}

func SearchStreetMarket(c *gin.Context) {

}

func DeleteStreetMarket(c *gin.Context) {

}

func UpdateStreetMarket(c *gin.Context) {

}

func GetStreetMarket(c *gin.Context) {
	var market model.StreetMarket
	code := c.Params.ByName("code")
	database.DB.Where(&model.StreetMarket{Registro: code}).First(&market)

	if market.Id == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, &market)
}
