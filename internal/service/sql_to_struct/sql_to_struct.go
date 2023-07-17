package sql_to_struct

import (
	"github.com/pkg/errors"
	"html/template"
	"strings"
	"unicode"
)

var noneFieldName = map[string]struct{}{
	"primary":    {},
	"unique":     {},
	"key":        {},
	"foreign":    {},
	"constraint": {},
}

var typeMap = map[string]string{
	"bool":      "bool",
	"tinyint":   "int8",
	"smallint":  "int16",
	"int":       "int",
	"bigint":    "int64",
	"varchar":   "string",
	"char":      "string",
	"timestamp": "time.Time",
	"date":      "time.Time",
}

func SqlToStruct(sql string) error {
	start := strings.Index(sql, "(")
	end := strings.LastIndex(sql, ")")
	part1 := sql[:start]
	part2 := sql[start+1 : end]
	tableName := strings.Fields(part1)[2]
	tableName = SnakeToPascalCase(tableName)
	fields := ToFieldsMap(part2)
	t, _ := template.ParseFiles("./template/struct.tmpl")
	data := map[string]interface{}{
		"StructName": tableName,
		"Fields":     fields,
	}
	var body strings.Builder
	if err := t.Execute(&body, data); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func ToFieldsMap(fieldStr string) map[string]string {
	lines := strings.Split(fieldStr, ",")
	fields := map[string]string{}
	for _, line := range lines {
		words := strings.Fields(strings.ToLower(line))
		fieldName := words[0]
		fieldType := words[1]
		if _, ok := noneFieldName[fieldName]; ok {
			continue
		}
		index := strings.Index(fieldType, "(")
		if index != -1 {
			fieldType = fieldType[:index]
		}
		fieldName = SnakeToPascalCase(fieldName)
		fields[fieldName] = typeMap[fieldType]
	}
	return fields
}

func findIndex(source []string, target string) int {
	for index, str := range source {
		if str == target {
			return index
		}
	}
	return -1
}

func SnakeToPascalCase(name string) string {
	var buffer strings.Builder
	capNext := true
	for _, r := range name {
		if r == '_' {
			capNext = true
			continue
		}
		if capNext {
			buffer.WriteRune(unicode.ToUpper(r))
			capNext = false
		} else {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}
