package controllers

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"raffles/internal/frontend/components/parts"
	"raffles/internal/server/database"
	"strconv"
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
	oldNumber, err := c.Params().GetInt("number")
	if err != nil {
		c.StopWithStatus(http.StatusBadRequest)
	}
	name := c.FormValue("name")
	n := c.FormValue("draw")
	number, err := strconv.Atoi(n)
	if err != nil {
		c.StopWithStatus(http.StatusBadRequest)
	}
	contact := c.FormValue("contact")
	status := database.DrawEdit(oldNumber, number, name, contact)
	setNotification(status, c)
}
