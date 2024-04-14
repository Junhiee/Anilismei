package beta

import (
	"fmt"
	"net/http"

	"git.virjar.com/Junhiee/anilismei/service"
	"github.com/gin-gonic/gin"
)

// 获得多个 Anime 信息
func GetAnimes(ctx *gin.Context) {

}

// 获得一个 Anime 信息
func GetAnime(ctx *gin.Context) {

}

// 添加一个 Anime 信息
func AddAnime(ctx *gin.Context) {
	
	err := service.Add()
	fmt.Println(err)
	ctx.JSON(
		http.StatusAccepted,
		nil,
	)
}

// 更新一个 Anime 信息
func UpdateAnime(ctx *gin.Context) {

}

// 删除一个 Anime 信息
func DeleteAnime(ctx *gin.Context) {

}
