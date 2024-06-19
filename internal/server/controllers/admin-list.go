package controllers

import (
	"github.com/kataras/iris/v12"
	"raffles/internal/frontend/components"
)

// AdminList handles /admin/list
func AdminList(c iris.Context) {
	c.RenderComponent(components.AdminList())
}
