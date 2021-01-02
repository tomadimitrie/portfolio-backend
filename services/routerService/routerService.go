package routerService

import (
	"backend/services/dbService"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)


func Init() {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8080"
	}
	router := gin.New()
	initRoutes(router)
	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic("failed to bind to port")
	}
}

func initRoutes(router *gin.Engine) {
	router.GET("/about", func(c *gin.Context) {
		c.JSON(200, dbService.GetAbout())
	})
}