package response

import (
	helper "shop/pkg/error"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string      `json:"message"`
	Err     error       `json:"error"`
	Data    interface{} `json:"data"`
}

func TestJSON(attrs ...interface{}) {
	return
}

func ResponseJSON(c *gin.Context, attrs ...interface{}) {
	if len(attrs) == 0 {
		c.JSON(200, map[string]string{})
		return
	}
	res := &Response{}
	for _, attr := range attrs {
		if err, ok := attr.(*helper.CustomError); ok {
			res.Err = err
			res.Message = err.Msg
			c.Error(err)
			return
		} else {
			res.Data = attr
		}
	}
	c.JSON(200, res)
}
