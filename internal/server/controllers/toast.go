package controllers

import (
	"github.com/kataras/iris/v12"
	"raffles/internal/frontend/components/parts"
)

// setNotification receives the status from every action, and creates a message to render properly the toast component
func setNotification(status int, c iris.Context) {
	var message string
	var class = "bg-green-500"
	switch status {
	case 201:
		message = "Cadastro realizado com sucesso!"
	case 202:
		message = "Palpite alterado com sucesso"
	case 204:
		message = "Palpite removido com sucesso"
	}
	c.RenderComponent(parts.Toast(class, message))
}
