package tools

import (
	"github.com/gin-gonic/gin"
	"tools/internal/service/copy_property"
)

func CopyProperty(c *gin.Context) {
	copy_property.Copy(c)
}
