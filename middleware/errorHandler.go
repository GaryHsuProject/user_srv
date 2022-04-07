package middleware

import (
	"encoding/json"
	"fmt"
	"shop/driver"
	helper "shop/pkg/error"
	"strings"

	"github.com/gin-gonic/gin"
)

func HandleError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		errs := c.Errors
		if len(errs) > 0 {
			for _, err := range errs {
				switch err.Err.(type) {
				case *helper.CustomError:
					customErr := err.Err.(*helper.CustomError)
					v, marshalErr := json.Marshal(customErr.Value)
					if marshalErr != nil {
						panic(marshalErr)
					}
					s := strings.Builder{}
					s.Write([]byte(fmt.Sprintf("StatusCode: %d,PrivateError: %s, value: %s", customErr.StatusCode, customErr.PrivateMsg, string(v))))
					if customErr.Err != nil {
						s.Write([]byte(fmt.Sprintf(", Error: %s", customErr.Err.Error())))
					}
					driver.Logger.Error(s.String())
					c.AbortWithStatusJSON(customErr.StatusCode, customErr.NewResponseMsg())
				default:
					c.AbortWithStatusJSON(500, "Unknown Error Caused.")
				}
			}
		}
	}
}

func PanicHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		defer func() {
			if err := recover(); err != nil {
				driver.Logger.Panic(err)
				return
			}
		}()
	}
}
