package routes

import (
	"github.com/kataras/iris/v12"
	"raffles/internal/server/controllers"
)

// Router defines all routes accepted for the server, and sets it's methods and respective controllers
func Router(s *iris.Application) {
	s.Get("/", controllers.Index)
	admin := s.Party("/admin")
	{
		admin.Get("/", controllers.Home)
		admin.Get("/list", controllers.DrawList)
		admin.Get("/list/{number:int8}", controllers.DrawListSingle)
		admin.Get("/edit/{number:int8}", controllers.DrawEdit)
		admin.Get("/add", controllers.DrawAdd)
		admin.Get("/winner", controllers.DrawWinner)
		admin.Post("/add", controllers.DrawAddPost)
		admin.Post("/edit/{number:int8}", controllers.DrawEditPost)
		admin.Post("/run-draw", controllers.DrawRun)
		admin.Delete("/delete/{number:int8}", controllers.DrawDelete)
		admin.Delete("/remove-toast", func(c iris.Context) {
			c.HTML("<div id=\"notify\"></div>")
		})
	}
}
