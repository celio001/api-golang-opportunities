package router

import "github.com/gin-gonic/gin"

func Initialize(){
	// Inicialize Router
	router := gin.Default()
	//Inicialize Routes
	inicializeRoutes(router)
	// Run the server
	router.Run(":8080") // oiça e sirva na `0.0.0.0:8080`
}