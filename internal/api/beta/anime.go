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


type GetListRequest struct {
	Page       int32     `form:"page" binding:"required"`
	Size       int32     `form:"size" binding:"required,min=5,max=20"`
	Country    string    `form:"country"`
	Genre      string    `form:"genre"`
	UpdateTime time.Time `form:"update_time"`
	Sort       string    `form:"sort"`
}

// TODO Router: GetListAnimes
func (a *AnimeRouter) GetListAnimes(ctx *gin.Context) {

	var req GetListRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, r.ErrResponse(e.INVALID_PARAMS, err.Error()))
		return
	}

	page := (req.Page - 1) * req.Size
	if req.Country == "" {
		ctx.DefaultQuery("country", "japan")
	}

	data, err := a.AnimeService.GetListAnimes(req.Size, page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, r.ErrResponse(e.INVALID_PARAMS, err.Error()))
		return
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
