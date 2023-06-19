package handler

import (
	"github.com/gofiber/fiber/v2"
	"test_di/service"
)

type DemoHandler struct {
	svc *service.UserService
}

func (handler *DemoHandler) Test(c *fiber.Ctx) error {
	c.Status(200)
	_ = c.JSON(&UserInfoResponse{
		Code: 0,
		Msg:  "ok11111111",
		Data: handler.svc.GetUser(),
	})
	return nil
}
func (handler *DemoHandler) RegisterRouter(app *fiber.App) {
	app.Get("/v1/demo", handler.Test)
}

func NewDemoHandler(svc *service.UserService) Handler {
	return &DemoHandler{svc: svc}
}
