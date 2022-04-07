package main

import (
	"fmt"
	"shop/api"
	"shop/driver"
	"shop/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	driver.Init()
	driver.InitLogger()
	engine.Use(middleware.HandleError(), middleware.PanicHandler())
	userRouter := engine.Group("/user")
	userRouter.POST("/register", api.InsertUser)
	userRouter.POST("/login", api.Login)
	err := engine.Run(fmt.Sprintf(":%d", driver.GlobalConfig.Port))
	panic(err)
}
