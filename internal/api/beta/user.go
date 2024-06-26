package beta

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Junhiee/anilismei/internal/service"
	e "github.com/Junhiee/anilismei/pkg/errors"
	r "github.com/Junhiee/anilismei/pkg/resp"
)

type UserRouter struct {
	userService service.UserService
}

func NewUserRouter(userService service.UserService) *UserRouter {
	return &UserRouter{
		userService: userService,
	}
}

type GetUserRequest struct {
	UserID int64 `uri:"user_id" binding:"required"`
}

func (u *UserRouter) GetUser(ctx *gin.Context) {

	var req GetUserRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, r.Response(e.INVALID_PARAMS, nil))
		return
	}

	data, err := u.userService.GetUser(req.UserID)

	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(http.StatusOK, r.Response(e.SUCCESS, data))
}

func (u *UserRouter) AddUser(ctx *gin.Context) {
	data := service.User{
		// UserID:    10001,
		UserName:  "Jack",
		Email:     "dasuaige68@gmail.com",
		UserPwd:   "123456",
		AvatarUrl: "http://avataurl.example.com",
	}

	err := u.userService.AddUser(data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, r.Response(e.ERROR_DB, nil))
		return
	}

	ctx.JSON(http.StatusOK, r.Response(e.SUCCESS, nil))

}
