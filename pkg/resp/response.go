package res

import (
	"github.com/gin-gonic/gin"

	e "github.com/Junhiee/anilismei/pkg/errors"
)

func Response(errCode int, data interface{}) gin.H {
	return gin.H{
		"code": errCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	}
}

func ErrResponse(errCode int, data interface{}) gin.H {
	return gin.H{
		"code": errCode,
		"msg":  e.GetMsg(errCode),
		"err":  data,
	}
}
