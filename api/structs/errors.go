package structs

import "errors"

type Response struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

var (
	ErrNotFound = errors.New("not found")
	ErrInternalResponse = Response{
	Status: "Failed",
	Msg:    "Internal server error occured",
}
	NotFoundResponse = Response{
		Status: "not found",
	}

)