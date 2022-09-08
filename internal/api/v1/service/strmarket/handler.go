package strmarket

import (
	"net/http"

	"github.com/danilosales/api-street-markets/config/logger"
	"github.com/danilosales/api-street-markets/internal/database"
	"github.com/danilosales/api-street-markets/internal/model"
	"github.com/gin-gonic/gin"
)

type StretMarketHandler struct {
	Logger logger.Logger
}

func New(l *logger.Logger) *StretMarketHandler {
	return &StretMarketHandler{Logger: *l}
}

// @title           Street Markets API
// @version         1.0

// CreateStreetMarket godoc
// @Summary 				Create a Street Market
// @Description 		Create a Street Market
// @Tags						Street Market
// @Produce 				json
// @Accept 					json
// @Param 					market body model.StreetMarketDto true "Street Market"
// @Success 				201 {object} model.StreetMarketDto "Created"
// @Failure 				400 "Invalid Request"
// @Router 					/street-markets [post]
func (s *StretMarketHandler) CreateStreetMarket(c *gin.Context) {
	var marketDto model.StreetMarketDto
	if err := c.ShouldBindJSON(&marketDto); err != nil {
		s.Logger.Debug().Err(err).Msg("Error to unmarshal json")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := marketDto.Validate(); err != nil {
		s.Logger.Debug().Err(err).Msg("Invalid fields")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	market := marketDto.ToModel()
	market.Registro = marketDto.Registro
	result := database.DB.Create(&market)

	if result.Error != nil {
		s.Logger.Err(result.Error).Msg("Error to save market on DB")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error})
		return
	}

	c.JSON(http.StatusCreated, market.ToDto())

}

// SearchStreetMarket godoc
// @Summary 				Search a Street Market
// @Description 		Search a Market by distrito, regiao5, nome or bairro
// @Tags						Street Market
// @Produce 				json
// @Param 					distrito query string false "Street Market Distrito"
// @Param 					regiao5 query string false "Street Market Regiao5"
// @Param 					nome query string false "Street Market Name"
// @Param 					bairro query string false "Street Market Bairro"
// @Success 				200 {object} model.StreetMarkets "ok"
// @Success 				204 "Can not find a Street Market with parameters"
// @Router 					/street-markets [get]
func (s *StretMarketHandler) SearchStreetMarket(c *gin.Context) {
	distrito := c.Query("distrito")
	regiao := c.Query("regiao5")
	nome := c.Query("nome")
	bairro := c.Query("bairro")

	var markets model.StreetMarkets

	result := database.DB.Where(&model.StreetMarket{
		Distrito:  distrito,
		Regiao5:   regiao,
		NomeFeira: nome,
		Bairro:    bairro,
	}).Find(&markets)

	if result.Error != nil {
		s.Logger.Err(result.Error).Msg("Error on search street market")
		c.Status(http.StatusInternalServerError)
		return
	}

	if len(markets) == 0 {
		c.JSON(http.StatusNoContent, []model.StreetMarketDto{})
		return
	}

	c.JSON(http.StatusOK, markets.ToDto())
}

// DeleteStreetMarket godoc
// @Summary 				Delete a Street Market by Register Code
// @Description 		Delete a Market by register Code
// @Tags						Street Market
// @Produce 				json
// @Param 					code path string true "Street Market Register Code"
// @Success 				204 {object} model.StreetMarketDto "ok"
// @Failure 				404 "Can not find a Street Market with this Register Code"
// @Router 					/street-markets/{code} [delete]
func (s *StretMarketHandler) DeleteStreetMarket(c *gin.Context) {
	code := c.Params.ByName("code")
	result := database.DB.Where(&model.StreetMarket{Registro: code}).Delete(&model.StreetMarket{})

	if result.Error != nil {
		s.Logger.Err(result.Error).Msg("Error on delete street market")
		c.Status(http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	c.Status(http.StatusNoContent)
}

// UpdateStreetMarket godoc
// @Summary 				Update a Street Market by register code
// @Description 		Update a Street Market by register code
// @Tags						Street Market
// @Produce 				json
// @Accept 					json
// @Param 					market body model.StreetMarketDto true "Street Market"
// @Param 					code path string true "Street Market Register Code"
// @Success 				200 {object} model.StreetMarketDto "ok"
// @Failure 				400 "Invalid Request"
// @Router 					/street-markets/{code} [put]
func (s *StretMarketHandler) UpdateStreetMarket(c *gin.Context) {
	var marketBD model.StreetMarket
	var marketDTO model.StreetMarketDto
	code := c.Params.ByName("code")
	database.DB.Where(&model.StreetMarket{Registro: code}).First(&marketBD)

	if marketBD.Id == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	if err := c.ShouldBindJSON(&marketDTO); err != nil {
		s.Logger.Debug().Err(err).Msg("Error to unmarshal json")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := marketDTO.Validate(); err != nil {
		s.Logger.Debug().Err(err).Msg("Invalid fields")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	marketToUpdate := marketDTO.ToModel()
	marketToUpdate.Id = marketBD.Id
	marketToUpdate.Registro = code

	result := database.DB.Save(&marketToUpdate)

	if result.Error != nil {
		s.Logger.Err(result.Error).Msg("Error to update market on DB")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error})
		return
	}

	c.JSON(http.StatusOK, &marketToUpdate)
}

// GetStreetMarket godoc
// @Summary 				Get a Street Market by Register Code
// @Description 		Search a Market by register Code
// @Tags						Street Market
// @Produce 				json
// @Param 					code path string true "Street Market Register Code"
// @Success 				200 {object} model.StreetMarketDto "ok"
// @Failure 				404 "Can not find a Street Market with this Register Code"
// @Router 					/street-markets/{code} [get]
func (s *StretMarketHandler) GetStreetMarket(c *gin.Context) {
	var market model.StreetMarket
	code := c.Params.ByName("code")
	database.DB.Where(&model.StreetMarket{Registro: code}).First(&market)

	if market.Id == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, market.ToDto())
}
