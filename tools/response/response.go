package response

import (
	"github.com/gin-gonic/gin"

	e "git.virjar.com/Junhiee/anilismei/pkg/errors"
)

func Response(ctx *gin.Context, HttpCode, errCode int, data interface{}) {
	ctx.JSON(
		HttpCode, gin.H{
			"code": errCode,
			"msg":  e.GetMsg(errCode),
			"data": data,
		})
}
