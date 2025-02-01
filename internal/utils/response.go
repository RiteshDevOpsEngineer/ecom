// internal/utils/utils.go
package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	HTTPStatus int         `json:"-"`
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func NewResponse(httpStatus int, status, message string, data interface{}) *Response {
	return &Response{
		HTTPStatus: httpStatus,
		Status:     status,
		Message:    message,
		Data:       data,
	}
}
func (r *Response) Send(c *gin.Context) {
	c.JSON(r.HTTPStatus, r)
}
