package app

import (
	"net/http"

	"github.com/beego/beego/v2/core/validation"
	"github.com/cindyyangcaixia/ApplicationOfNEMT/pkg/e"
	"github.com/gin-gonic/gin"
)

func getValidationMessage(errors []*validation.Error) string {
	var message string
	for _, err := range errors {
		message += err.Message + "; "
	}
	return message
}

func BindAndValid(c *gin.Context, form interface{}) *ResponseMessage {
	err := c.Bind(form)
	if err != nil {
		// why the reponse is text not json ? todo
		return &ResponseMessage{Status: http.StatusBadRequest, Code: e.INVALID_PARAMS, Message: err.Error()}
	}

	valid := validation.Validation{}

	check, err := valid.Valid(form)
	if err != nil {
		return &ResponseMessage{Status: http.StatusInternalServerError, Code: e.ERROR, Message: err.Error()}
	}

	if !check {
		return &ResponseMessage{Status: http.StatusBadRequest, Code: e.INVALID_PARAMS, Message: getValidationMessage(valid.Errors)}

	}

	return nil
}
