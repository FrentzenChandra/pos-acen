package helper

import (
	"net/http"
	"pos-acen/internal/modules/response/entity"

	"github.com/thedevsaddam/renderer"
)

const (
	SUCCESS_MESSSAGE string = "Success"
)

func HandleResponse(w http.ResponseWriter, render *renderer.Render, statusCode int, message interface{}, data interface{}) {
	response := entity.BaseResponse{
		Message: message,
		Data:    data,
	}

	render.JSON(w, statusCode, response)
}
