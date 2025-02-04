package router

import (
	"github.com/celio001/api-golang-opportunities.git/handler"
	"github.com/gin-gonic/gin"
)

func inicializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		//Show Opening
		v1.GET("/opening", handler.ShowOpeningHandler)
		//POST Opening
		v1.POST("/opening", handler.CreateOpeningHandler)
		//Delete Opening
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		//Update Opening
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		//Update Openings
		v1.GET("/opening", handler.ListOpeningsHandler)
	}
}
