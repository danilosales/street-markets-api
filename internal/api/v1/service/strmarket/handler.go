package strmarket

import (
	"net/http"

	"github.com/danilosales/api-street-markets/internal/database"
	"github.com/danilosales/api-street-markets/internal/model"
	"github.com/gin-gonic/gin"
)

func CreateStreetMarket(c *gin.Context) {
	var marketDto model.StreetMarketDto
	if err := c.ShouldBindJSON(&marketDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := marketDto.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	market := marketDto.ToModel()
	market.Registro = marketDto.Registro
	result := database.DB.Create(&market)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error})
		return
	}

	c.JSON(http.StatusCreated, market.ToDto())

}

func SearchStreetMarket(c *gin.Context) {
	distrito := c.Query("distrito")
	regiao := c.Query("regiao5")
	nome := c.Query("nome")
	bairro := c.Query("bairro")

	var markets model.StreetMarkets

	database.DB.Where(&model.StreetMarket{
		Distrito:  distrito,
		Regiao5:   regiao,
		NomeFeira: nome,
		Bairro:    bairro,
	}).Find(&markets)

	if len(markets) == 0 {
		c.JSON(http.StatusNoContent, []model.StreetMarketDto{})
		return
	}

	c.JSON(http.StatusOK, markets.ToDto())
}

func DeleteStreetMarket(c *gin.Context) {
	code := c.Params.ByName("code")
	result := database.DB.Where(&model.StreetMarket{Registro: code}).Delete(&model.StreetMarket{})

	if result.RowsAffected == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	c.Status(http.StatusNoContent)
}

func UpdateStreetMarket(c *gin.Context) {
	var marketBD model.StreetMarket
	var marketDTO model.StreetMarketDto
	code := c.Params.ByName("code")
	database.DB.Where(&model.StreetMarket{Registro: code}).First(&marketBD)

	if marketBD.Id == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	if err := c.ShouldBindJSON(&marketDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := marketDTO.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	marketToUpdate := marketDTO.ToModel()
	marketToUpdate.Id = marketBD.Id
	marketToUpdate.Registro = code

	result := database.DB.Save(&marketToUpdate)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error})
		return
	}

	c.JSON(http.StatusOK, &marketToUpdate)
}

func GetStreetMarket(c *gin.Context) {
	var market model.StreetMarket
	code := c.Params.ByName("code")
	database.DB.Where(&model.StreetMarket{Registro: code}).First(&market)

	if market.Id == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, market.ToDto())
}
