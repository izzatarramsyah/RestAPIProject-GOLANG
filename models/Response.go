package models

import (
	"fmt"
)

type Response struct {
	Code        string 	   `json:"code"`   	 	
	Message		string 	   `json:"message"`
	Object		[]Employee `json:"object"`
}

func (w *Response) Error() string {
    return fmt.Sprintf(w.Code,w.Message, w.Object)
}

func Wrap(code string, message string, object []Employee) *Response {
    return &Response{
        Code		: code,
		Message		: message,
		Object		: object,
    }
}