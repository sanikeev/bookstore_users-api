package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status int `json:"status"`
	Error string `json:"error"`
}

func NewBadRequestError(messge string) *RestErr {
	return &RestErr{
		Message: messge,
		Status: http.StatusBadRequest,
		Error: "bad_request",
	}
}

func NewNotFoundError(messge string) *RestErr {
	return &RestErr{
		Message: messge,
		Status: http.StatusNotFound,
		Error: "not_found",
	}
}