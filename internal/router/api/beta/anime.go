package beta

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	s "github.com/Junhiee/anilismei/internal/service"
	e "github.com/Junhiee/anilismei/pkg/errors"
	r "github.com/Junhiee/anilismei/pkg/resp"
)

type AnimeRouter struct{}

type GetListRequest struct {
	Page int32 `form:"page" binding:"required"`
	Size int32 `form:"size" binding:"required,min=5,max=20"`
}

func (a *AnimeRouter) GetListAnimes(ctx *gin.Context) {

	var req GetListRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, r.Response(e.INVALID_PARAMS, nil))
		return
	}

	data, err := s.Server.GetListAnimes(req.Size, req.Page*5)
	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(
		http.StatusOK,
		r.Response(e.SUCCESS, data),
	)

}

type GetAnimeRequest struct {
	AnimeID int64 `uri:"anime_id" binding:"required"`
}

func (a *AnimeRouter) GetAnime(ctx *gin.Context) {

	var req GetAnimeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, r.ErrResponse(e.INVALID_PARAMS, err.Error()))
		return
	}

	data, err := s.Server.GetAnime(req.AnimeID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, r.ErrResponse(e.ERROR_DB, err.Error()))
		return
	}

	ctx.JSON(
		http.StatusOK,
		r.Response(e.SUCCESS, data),
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
		ctx.JSON(http.StatusBadRequest, r.Response(e.INVALID_PARAMS, nil))
		return
	}

	ctx.JSON(
		http.StatusOK,
		r.Response(e.SUCCESS, data),
	)

}

// 更新一个 Anime 信息
func (a *AnimeRouter) UpdateAnime(ctx *gin.Context) {

}

// 删除一个 Anime 信息
func (a *AnimeRouter) DeleteAnime(ctx *gin.Context) {

}
