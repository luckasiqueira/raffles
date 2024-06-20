package controllers

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"raffles/internal/server/database"
	"strconv"
)

func DrawAddPost(c iris.Context) {
	name := c.FormValue("name")
	contact := c.FormValue("contact")
	n := c.FormValue("number")
	number, err := strconv.Atoi(n)
	var status int
	if err != nil {
		status = http.StatusInternalServerError
	}
	status = database.DrawAdd(number, name, contact)
	setNotification(status, c)
}
