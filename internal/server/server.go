package server

import (
	"embed"
	"fmt"
	"github.com/kataras/iris/v12"
	"raffles/internal/server/routes"
	"raffles/utils/info"
)

// StartServer runs the iris web server and initialize all the defined routes
func StartServer(fs embed.FS) {
	s := iris.New()
	s.HandleDir("/assets", fs)
	routes.Router(s)
	err := s.Listen(":" + info.Env["SERVER_PORT"])
	if err != nil {
		fmt.Println(err.Error())
	}
}
