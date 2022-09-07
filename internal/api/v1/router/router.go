package router

import (
	"github.com/danilosales/api-street-markets/config/logger"
	"github.com/danilosales/api-street-markets/internal/api/v1/router/middleware"
	"github.com/danilosales/api-street-markets/internal/api/v1/service/strmarket"
	"github.com/gin-gonic/gin"
)

func New(l *logger.Logger) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.StructuredLogger(l))
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		v1.POST("/street-markets", strmarket.CreateStreetMarket)
		v1.GET("/street-markets", strmarket.SearchStreetMarket)
		v1.DELETE("/street-markets/:code", strmarket.DeleteStreetMarket)
		v1.PUT("/street-markets/:code", strmarket.UpdateStreetMarket)
		v1.GET("/street-markets/:code", strmarket.GetStreetMarket)
	}

	return r
}
