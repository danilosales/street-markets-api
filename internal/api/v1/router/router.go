package router

import (
	"github.com/danilosales/api-street-markets/config/logger"
	docs "github.com/danilosales/api-street-markets/docs"
	"github.com/danilosales/api-street-markets/internal/api/v1/router/middleware"
	"github.com/danilosales/api-street-markets/internal/api/v1/service/strmarket"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func New(l *logger.Logger) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.StructuredLogger(l))
	r.Use(gin.Recovery())

	docs.SwaggerInfo.BasePath = "/api/v1"

	h := strmarket.New(l)

	v1 := r.Group("/api/v1")
	{
		v1.POST("/street-markets", h.CreateStreetMarket)
		v1.GET("/street-markets", h.SearchStreetMarket)
		v1.DELETE("/street-markets/:code", h.DeleteStreetMarket)
		v1.PUT("/street-markets/:code", h.UpdateStreetMarket)
		v1.GET("/street-markets/:code", h.GetStreetMarket)
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	return r
}
