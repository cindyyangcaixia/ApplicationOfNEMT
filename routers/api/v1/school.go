package v1

import (
	"github.com/cindyyangcaixia/gin-example/models"
	"github.com/cindyyangcaixia/gin-example/pkg/app"
	"github.com/gin-gonic/gin"
)

type CreateSchoolsForm struct {
	Name     string `json:"name" valid:"Required;MaxSize(100);MinSize(2)"`
	SerialNo string `json:"serialNo" valid:"Required;MaxSize(100);MinSize(2)"`
}

func CreateSchools(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form CreateSchoolsForm
	)

	check := app.BindAndValid(c, &form)

	if check != nil {
		appG.Response(check, nil)
		return
	}

	school, res := models.CreateSchool(form.Name, form.SerialNo)

	if res != nil {
		appG.Response(res, nil)
		return
	}

	appG.Response(nil, school)
}
