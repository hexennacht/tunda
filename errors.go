package tunda

import (
	"encoding/json"
	"fmt"
)

type ErrorKerjaan[T comparable] struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data T `json:"data"`
}

func newErrorKerjaan[T comparable](code int, message string, data T) error {
	return &ErrorKerjaan[T]{
		Code: code,
		Message: message,
		Data: data,
	}
}

func (e *ErrorKerjaan[T]) Error() string {
	data, _ := json.Marshal(e)

	return string(data)
}

func (e *ErrorKerjaan[T]) Unwrap() error {
	return fmt.Errorf("code: %d, message: %s", e.Code, e.Message)
}