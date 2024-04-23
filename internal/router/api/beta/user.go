package beta

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	s "git.virjar.com/Junhiee/anilismei/internal/service"
	e "git.virjar.com/Junhiee/anilismei/pkg/errors"
	resp "git.virjar.com/Junhiee/anilismei/tools/response"
)

type UserRouter struct{}

type GetUserRequest struct {
	UserID int64 `uri:"user_id" binding:"required"`
}

func (u *UserRouter) GetUser(ctx *gin.Context) {

	var req GetUserRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		fmt.Println(err)
		resp.Response(ctx, http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	data, err := s.Server.GetUser(req.UserID)

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

func (u *UserRouter) AddUser(ctx *gin.Context) {
	data := s.User{
		// UserID:    10001,
		UserName:  "Jack",
		Email:     "dasuaige68@gmail.com",
		UserPwd:   "123456",
		AvatarUrl: "http://avataurl.example.com",
	}

	err := s.Server.AddUser(data)

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
