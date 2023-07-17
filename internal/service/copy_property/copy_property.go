package copy_property

import (
	"strings"
	"text/template"
)

func CopyProperty(daoStr, dtoStr string) string {
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
	t.Execute(&builder, data)
	return builder.String()
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
