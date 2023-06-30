package v1

import (
	"net/http"

	"github.com/cindyyangcaixia/gin-example/models"
	"github.com/cindyyangcaixia/gin-example/pkg/app"
	"github.com/cindyyangcaixia/gin-example/pkg/e"
	"github.com/gin-gonic/gin"
)

type CreateSchoolsForm struct {
	Name     string `json:"name" valid:"Required"`
	SerialNo string `json:"serialNo" valid:"Required"`
}

func CreateSchools(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form CreateSchoolsForm
	)

	status, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(status, errCode, nil)
		return
	}

	_, code := models.CreateSchool(form.Name, form.SerialNo)

	if code != e.SUCCESS {
		appG.Response(status, code, nil)
		return
	}

	appG.Response(http.StatusCreated, e.SUCCESS, nil)
}
