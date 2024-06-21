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
		admin.Get("/list", controllers.DrawList)
		admin.Get("/list/{number:int8}", controllers.DrawListSingle)
		admin.Get("/edit/{number:int8}", controllers.DrawEdit)
		admin.Get("/add", controllers.DrawAdd)
		admin.Post("/add", controllers.DrawAddPost)
		admin.Post("/edit/{number:int8}", controllers.DrawEditPost)
		admin.Delete("/delete/{number:int8}", controllers.DrawDelete)
		admin.Delete("/remove-toast", func(c iris.Context) {
			c.HTML("<div id=\"notify\"></div>")
		})
	}
}
