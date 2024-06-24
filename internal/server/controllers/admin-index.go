package controllers

import (
	"github.com/kataras/iris/v12"
	"raffles/internal/frontend/components"
	"raffles/internal/server/database"
)

// DrawList handles /admin/
func Home(c iris.Context) {
	status := database.DrawStatus()
	c.RenderComponent(components.Home(status))
}
