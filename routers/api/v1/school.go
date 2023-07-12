package v1

import (
	"github.com/cindyyangcaixia/ApplicationOfNEMT/models"
	"github.com/cindyyangcaixia/ApplicationOfNEMT/pkg/app"
	"github.com/gin-gonic/gin"
)

type CreateSchoolsForm struct {
	Name     string `json:"name" valid:"Required;MaxSize(100);MinSize(2)"`
	SerialNo string `json:"serialNo" valid:"Required;MaxSize(100);MinSize(2)"`
}

// @Summary		Create a school
// @Description	Create a school
// @Tags			schools
// @Accept			json
// @Produce		json
// @Param			name	body	string	true	"School name"
// @Param			serialNo	body	string	true	"School serial number"
// @Success		200	{object}	models.School
// @Router			/schools [post]

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
