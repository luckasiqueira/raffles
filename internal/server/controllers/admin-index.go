package controllers

import (
	"github.com/kataras/iris/v12"
	"raffles/internal/frontend/components"
)

// DrawList handles /admin/
func Home(c iris.Context) {
	c.RenderComponent(components.Home())
}
