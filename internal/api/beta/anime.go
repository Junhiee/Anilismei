package beta

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Junhiee/anilismei/internal/service"
	e "github.com/Junhiee/anilismei/pkg/errors"
	r "github.com/Junhiee/anilismei/pkg/resp"
)

type AnimeRouter struct {
	AnimeService service.AnimeService
}

func NewAnimeRouter(animeService service.AnimeService) *AnimeRouter {
	return &AnimeRouter{
		AnimeService: animeService,
	}
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

	data, err := a.AnimeService.GetAnimeByID(req.AnimeID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, r.ErrResponse(e.ERROR_DB, err.Error()))
		return
	}

	ctx.JSON(
		http.StatusOK,
		r.Response(e.SUCCESS, data),
	)
}

type GetListRequest struct {
	Page        int32     `form:"page" binding:"required"`
	Size        int32     `form:"size" binding:"required,min=5,max=20"`
	Country     string    `form:"country"`
	Type        string    `form:"type"`
	ReleaseDate time.Time `form:"release_date"`
	UpdateTime  time.Time `form:"update_time"`
	Sort        string    `form:"sort"`
}


// FIXME 删除 Valid 字段
func (a *AnimeRouter) GetListAnimes(ctx *gin.Context) {

	var req GetListRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, r.ErrResponse(e.INVALID_PARAMS, err.Error()))
		return
	}

	page := (req.Page - 1) * req.Size

	data, err := a.AnimeService.GetListAnimes(req.Size, page)

	// 获取按国家分类的结果
	if req.Country != "" {
		data, err = a.AnimeService.GetAnimesByCountry(req.Country, req.Size, page)
	}

	// 根据动画推出日期获得数据
	if !req.ReleaseDate.IsZero() {
		data, err = a.AnimeService.GetAnimesByRelease(req.ReleaseDate, req.Size, page)
	}

	// 根据动画类型获得数据
	if req.Type != "" {
		data, err = a.AnimeService.GetAnimesByType(req.Type, req.Size, page)

	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, r.ErrResponse(e.INVALID_PARAMS, err.Error()))
		return
	}

	ctx.JSON(
		http.StatusOK,
		r.Response(e.SUCCESS, data),
	)

}
