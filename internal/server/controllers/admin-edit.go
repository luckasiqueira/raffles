package controllers

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"raffles/internal/frontend/components/parts"
	"raffles/internal/server/database"
)

// DrawEdit handles /admin/edit/{number} (GET)
func DrawEdit(c iris.Context) {
	number, err := c.Params().GetInt("number")
	if err != nil {
		c.StopWithStatus(http.StatusBadRequest)
	}
	draw := database.DrawSingle(number)
	c.RenderComponent(parts.DrawEditing(draw))
}

// DrawEditPost handles /admin/edit/{number} (POST)
func DrawEditPost(c iris.Context) {
	//
}
