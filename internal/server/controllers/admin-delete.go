package controllers

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"raffles/internal/server/database"
)

// AdminDelete handles /admin/delete/{number} (DELETE)
func AdminDelete(c iris.Context) {
	number, err := c.Params().GetInt("number")
	if err != nil {
		c.StopWithStatus(http.StatusBadRequest)
	}
	status := database.DrawDelete(number)
	setNotification(status, c)
}
