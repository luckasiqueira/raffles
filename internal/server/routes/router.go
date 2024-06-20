package routes

import (
	"github.com/kataras/iris/v12"
	"raffles/internal/server/controllers"
)

// Router defines all routes accepted for the server, and sets it's methods and respective controllers
func Router(s *iris.Application) {
	s.Get("/", func(c iris.Context) {
		c.WriteString("Hello World!")
	})
	admin := s.Party("/admin")
	{
		admin.Get("/list", controllers.AdminList)
		admin.Get("/edit/{number:int8}", controllers.DrawEdit)

		admin.Post("/edit/{number:int8}", controllers.DrawEditPost)
	}
}
