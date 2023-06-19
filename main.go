package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
	"test_di/config"
	"test_di/dao"
	"test_di/db"
	"test_di/handler"
	"test_di/service"
)

type Server struct {
	app *fiber.App
	hs  []handler.Handler
}
type ServerParams struct {
	dig.In
	App      *fiber.App
	Handlers []handler.Handler `group:"handlers"`
}

func NewServer(p ServerParams) *Server {
	return &Server{
		app: p.App,
		hs:  p.Handlers,
	}
}

func (s *Server) Run() {
	for _, h := range s.hs {
		h.RegisterRouter(s.app)
	}
	if err := s.app.Listen(":3000"); err != nil {
		panic("start failed")
	}
}

func NewFiberApp() *fiber.App {
	return fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
}

func main() {
	// manual DI
	//c := config.NewConfig()
	//_db, err := db.ConnectDatabase(c)
	//if err != nil {
	//	panic(err)
	//}
	//userDao := dao.NewUserDao(_db)
	//userService := service.NewUserService(userDao)
	//userHandler := handler.NewUserHandler(userService)
	//server := NewServer(NewFiberApp(), []handler.Handler{userHandler})
	//server.Run()

	// uber DI
	container := dig.New()
	_ = container.Provide(config.NewConfig)
	_ = container.Provide(db.ConnectDatabase)
	_ = container.Provide(dao.NewUserDao)
	_ = container.Provide(service.NewUserService)
	_ = container.Provide(handler.NewUserHandler, dig.Group("handlers"))
	_ = container.Provide(handler.NewDemoHandler, dig.Group("handlers"))

	_ = container.Provide(NewFiberApp)
	err := container.Provide(NewServer)
	fmt.Printf("%v", err)
	err = container.Invoke(func(server *Server) {
		server.Run()
	})
	fmt.Printf("%v", err)
	if err != nil {
		panic(err)
	}
}
