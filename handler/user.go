package handler

import (
	"github.com/blntrsz/goweb/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (uh UserHandler) HandleUserShow(c echo.Context) error {
	return render(c, user.Show())
}
