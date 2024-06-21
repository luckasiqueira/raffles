package controllers

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"raffles/internal/frontend/components/parts"
)

const success = "fa-solid fa-check"
const alert = "fa-solid fa-triangle-exclamation"
const error = "fa-solid fa-xmark"
const classSuccess = "bg-green-500"
const classAlert = "bg-orange-500"
const classError = "bg-red-500"

// setNotification receives the status from every action, and creates a message to render properly the toast component
func setNotification(status int, c iris.Context) {
	var message string
	var class = classSuccess
	var icon = success
	switch status {
	case http.StatusCreated:
		message = "Cadastro realizado com sucesso!"
	case http.StatusAccepted:
		message = "Palpite alterado com sucesso!"
	case http.StatusNoContent:
		message = "Palpite removido com sucesso!"
	case http.StatusConflict:
		message = "Esse palpite já foi cadastrado!"
		class = classAlert
		icon = alert
	default:
		message = "Houve um erro inesperado. Tente de novo!"
		class = classError
		icon = error
	}
	c.RenderComponent(parts.Toast(class, message, icon))
}
