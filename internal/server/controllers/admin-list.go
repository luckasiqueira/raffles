package controllers

import (
	"github.com/kataras/iris/v12"
	"raffles/internal/frontend/components"
	"raffles/internal/server/database"
)

// AdminList handles /admin/list
func AdminList(c iris.Context) {
	draws := database.DrawList()
	c.RenderComponent(components.AdminList(draws))
}
