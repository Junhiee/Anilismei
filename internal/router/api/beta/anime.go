package beta

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	s "git.virjar.com/Junhiee/anilismei/internal/service"
	e "git.virjar.com/Junhiee/anilismei/pkg/errors"
	resp "git.virjar.com/Junhiee/anilismei/tools/response"
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

	data, err := s.Server.GetListAnimes(req.Size, req.Page*5)
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

	data, _ := s.Server.GetAnime(req.AnimeID)

	resp.Response(
		ctx,
		http.StatusOK,
		e.SUCCESS,
		data,
	)
}

// 添加一个 Anime 信息
func (a *AnimeRouter) AddAnime(ctx *gin.Context) {
	data := s.Animation{
		AnimeID:     10001,
		Title:       "Linux Title",
		Evaluate:    "Add Anime 1",
		GenreID:     10001,
		ReleaseDate: time.Now(),
		StudioID:    10001,
		AnimeStatus: "completed",
		Rating:      9.2,
	}
	err := s.Server.AddAnime(data)

	if err != nil {
		resp.Response(
			ctx,
			http.StatusBadRequest,
			e.ERROR_DB,
			nil,
		)
		return
	}

	resp.Response(
		ctx,
		http.StatusOK,
		e.SUCCESS,
		data,
	)

}

// 更新一个 Anime 信息
func (a *AnimeRouter) UpdateAnime(ctx *gin.Context) {

}

// 删除一个 Anime 信息
func (a *AnimeRouter) DeleteAnime(ctx *gin.Context) {

}
