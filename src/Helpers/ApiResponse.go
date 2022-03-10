package helpers

import (
	"net/http"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Errors  interface{} `json:"errors"`
}

func UnauthorizedRes(message string, code int, data interface{}) Meta {
	status := "failed"
	if code == http.StatusOK {
		status = "success"
	}

	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
		Errors:  nil,
	}

	return meta
}
