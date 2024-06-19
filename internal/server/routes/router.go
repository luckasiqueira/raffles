package routes

import "github.com/kataras/iris/v12"

// Router defines all routes accepted for the server, and sets it's methods and respective controllers
func Router(s *iris.Application) {
	s.Get("/", func(c iris.Context) {
		c.WriteString("Hello World!")
	})
}
