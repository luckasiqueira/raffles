package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"net/http"
	"raffles/internal/frontend/components"
	"raffles/internal/server/database"
	"strconv"
)

// DrawAddPost /admin/add (GET)
func DrawAdd(c iris.Context) {
	c.RenderComponent(components.AdminAdd())
}

// DrawAddPost /admin/add (POST)
func DrawAddPost(c iris.Context) {
	name := c.FormValue("name")
	contact := c.FormValue("contact")
	n := c.FormValue("number")
	number, err := strconv.Atoi(n)
	var status int
	if err != nil {
		status = http.StatusInternalServerError
	}
	p := c.FormValue("paid")
	fmt.Println("Paid antes:", p)
	paid, err := strconv.Atoi(p)
	if err != nil {
		status = http.StatusInternalServerError
	}
	if paid != 1 {
		paid = 0
	}
	fmt.Println("Paid depois:", paid)
	status = database.DrawAdd(number, name, contact, paid)
	setNotification(status, c)
}
