package server

import (
	"embed"
	"github.com/kataras/iris/v12"
	"raffles/internal/server/routes"
	"raffles/utils/info"
	"raffles/utils/logger"
)

//go:embed assets
var fs embed.FS

// StartServer runs the iris web server and initialize all the defined routes
func StartServer() {
	s := iris.New()
	s.Logger().SetLevel("debug")
	s.HandleDir("/assets", fs)
	routes.Router(s)
	err := s.Listen(":" + info.Env["SERVER_PORT"])
	if err != nil {
		logger.NewLog(err.Error())
	}
}
