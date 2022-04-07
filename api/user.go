package api

import (
	"shop/driver"
	"shop/dto"
	"shop/internal/user"

	"github.com/gin-gonic/gin"
)

func inject() user.Service {
	gorm := driver.InitGorm()
	repo := user.NewUserRepository(gorm)
	srv := user.NewUserService(repo)
	return srv
}

func InsertUser(c *gin.Context) {
	req := dto.InsertUser{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(err)
		return
	}
	srv := inject()
	err = srv.Insert(&req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, struct{}{})
}

func Login(c *gin.Context) {
	req := dto.FindOneByUser{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(err)
		return
	}
	srv := inject()
	res, err := srv.Login(&req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, res)
}
