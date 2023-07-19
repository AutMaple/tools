package copy_property

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"strings"
	"text/template"
	"tools/internal/utils/request"
)

type CopyForm struct {
	Struct CopyStruct
}

type CopyStruct struct {
	DaoStr string `json:"dao"`
	DtoStr string `json:"dto"`
}

func Copy(c *gin.Context) {
	var form CopyForm
	err := request.BindParams(&form, c)
	if err != nil {
		_ = c.Error(err)
		return
	}
	res, err := CopyProperty(form.Struct.DaoStr, form.Struct.DtoStr)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.String(200, res)
}

func CopyProperty(daoStr, dtoStr string) (string, error) {
	dtoWords := strings.Fields(dtoStr)
	srcName := dtoWords[1]
	TargetFields := ArrayToMap(dtoWords[4 : len(dtoWords)-1])
	daoWords := strings.Fields(daoStr)
	targetName := daoWords[1]
	SrcFields := ArrayToMap(daoWords[4 : len(daoWords)-1])
	t, err := template.ParseFiles("./template/copy_property.tmpl")
	if err != nil {
		panic(err)
	}
	data := map[string]interface{}{
		"SrcName":      srcName,
		"TargetName":   targetName,
		"TargetFields": TargetFields,
		"SrcFields":    SrcFields,
	}
	var builder strings.Builder
	err = t.Execute(&builder, data)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return builder.String(), nil
}

func ArrayToMap(arr []string) map[string]string {
	fields := make(map[string]string)
	for i := 0; i < len(arr); i += 2 {
		fieldName := arr[i]
		fieldType := arr[i+1]
		fields[fieldName] = fieldType
	}
	return fields
}
