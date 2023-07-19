package request

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"reflect"
)

// BindParams method is used to automatically bind the parameters in the request to the specified structure.
// When using it, you only need to use the official `uri`, `form` and `json` tags that come with gin in the structure.
func BindParams(f interface{}, c *gin.Context) error {
	v := reflect.ValueOf(f)
	t := v.Type().Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		switch field.Type.Kind() {
		case reflect.Ptr:
			elem := field.Type.Elem()
			if elem.Kind() == reflect.Struct {
				s := reflect.New(elem)
				err := c.ShouldBindJSON(s.Interface())
				if err != nil {
					return errors.WithStack(err)
				}
				v.Elem().Field(i).Set(s)
			}
		case reflect.Struct:
			s := reflect.New(field.Type)
			err := c.ShouldBindJSON(s.Interface())
			if err != nil {
				return errors.WithStack(err)
			}
			v.Elem().Field(i).Set(s.Elem())
		}
	}
	if err := c.ShouldBindQuery(v.Interface()); err != nil {
		return errors.WithStack(err)
	}
	if err := c.ShouldBindUri(v.Interface()); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
