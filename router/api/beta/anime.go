package beta

import (
	"fmt"
	"net/http"

	"git.virjar.com/Junhiee/anilismei/pkg/e"
	s "git.virjar.com/Junhiee/anilismei/service"
	resp "git.virjar.com/Junhiee/anilismei/utils/response"
	"github.com/gin-gonic/gin"
)

type AnimeRouter struct{}

type GetListRequest struct {
	Page int32 `form:"page" binding:"required"`
	Size int32 `form:"size" binding:"required,min=5,max=20"`
}

func (a *AnimeRouter) GetListAnimes(ctx *gin.Context) {

	var req GetListRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		resp.Response(ctx, http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	data, err := s.Server.GetList(req.Size, req.Page*5)
	if err != nil {
		fmt.Println(err)
	}

	resp.Response(
		ctx,
		http.StatusOK,
		e.SUCCESS,
		data,
	)
}

type GetAnimeRequest struct {
	AnimeID int64 `uri:"anime_id" binding:"required"`
}

func (a *AnimeRouter) GetAnime(ctx *gin.Context) {

	var req GetAnimeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		resp.Response(ctx, http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	data, err := s.Server.Get(req.AnimeID)

	if err != nil {
		fmt.Println(err)
	}

	resp.Response(
		ctx,
		http.StatusOK,
		e.SUCCESS,
		data,
	)
}

// 添加一个 Anime 信息
func (a *AnimeRouter) AddAnime(ctx *gin.Context) {
	
}

// 更新一个 Anime 信息
func (a *AnimeRouter) UpdateAnime(ctx *gin.Context) {

}

// 删除一个 Anime 信息
func (a *AnimeRouter) DeleteAnime(ctx *gin.Context) {

}
