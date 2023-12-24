package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/swikars1/got/model"
	"github.com/swikars1/got/view/user"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Email: "swikarsharma@gmail.com",
	}
	return render(c, user.Show(u))
}
