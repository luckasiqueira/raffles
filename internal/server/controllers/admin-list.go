package controllers

import (
	"github.com/kataras/iris/v12"
	"raffles/internal/frontend/components"
	"raffles/internal/server/database"
)

// DrawList handles /admin/list
func DrawList(c iris.Context) {
	draws := database.DrawList()
	c.RenderComponent(components.AdminList(draws))
}

func AdminListSingle(c iris.Context) {

}
