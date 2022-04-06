package api

import (
	"shop/driver"
	"shop/dto"
	"shop/internal/user"
	"shop/pkg/response"

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
	response.ResponseJSON(c, c.ShouldBindJSON(&req))
	srv := inject()
	err := srv.Insert(&req)
	response.ResponseJSON(c, err)
}

func Login(c *gin.Context) {
	req := dto.FindOneByUser{}
	response.ResponseJSON(c, c.ShouldBindJSON(&req))
	srv := inject()
	res, err := srv.Login(&req)
	response.ResponseJSON(c, res, err)
}
