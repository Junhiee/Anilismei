package response

import (
	"git.virjar.com/Junhiee/anilismei/pkg/e"
	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, HttpCode, errCode int, data interface{}) {
	ctx.JSON(
		HttpCode, gin.H{
			"code": errCode,
			"msg":  e.GetMsg(errCode),
			"data": data,
		})
}
