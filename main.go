package main

import (
	"fmt"
	"shop/api"
	"shop/driver"
	"shop/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	driver.Init()
	engine.Use(middleware.HandleError(), middleware.PanicHandler())
	userRouter := engine.Group("/user")
	userRouter.POST("/register", api.InsertUser)
	userRouter.POST("/login", api.Login)
	engine.Run(fmt.Sprintf(":%d", driver.GlobalConfig.Port))
}
