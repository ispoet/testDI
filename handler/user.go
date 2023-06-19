package handler

import (
	"github.com/gofiber/fiber/v2"
	"test_di/dao"
	"test_di/service"
)

type Handler interface {
	RegisterRouter(app *fiber.App)
}

type UserInfoResponse struct {
	Code int
	Msg  string
	Data *dao.User
}

type UserHandler struct {
	svc *service.UserService
}

func (handler *UserHandler) GetUser(c *fiber.Ctx) error {
	c.Status(200)
	_ = c.JSON(&UserInfoResponse{
		Code: 0,
		Msg:  "ok",
		Data: handler.svc.GetUser(),
	})
	return nil
}
func (handler *UserHandler) RegisterRouter(app *fiber.App) {
	app.Get("/v1/getUser", handler.GetUser)
}

func NewUserHandler(svc *service.UserService) Handler {
	return &UserHandler{svc: svc}
}
