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
	var draw database.Participant
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
	p := c.PostValue("paid")
	paid, err := strconv.Atoi(p)
	if err != nil {
		c.StopWithStatus(http.StatusBadRequest)
	}
	if paid != 1 {
		paid = 0
	}
	status := database.DrawEdit(oldNumber, number, name, contact, paid)
	if status == http.StatusAccepted {
		draw = database.DrawSingle(number)
	} else {
		draw = database.DrawSingle(oldNumber)
	}
	c.RenderComponent(parts.DrawSingle(draw))
	setNotification(status, c)
}
