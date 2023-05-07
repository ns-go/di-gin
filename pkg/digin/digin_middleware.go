package digin

import (
	"github.com/gin-gonic/gin"
	"github.com/ns-go/di/pkg/di"
)

const ContextKey string = "container"

func Container(container *di.Container) gin.HandlerFunc {
	if container == nil {
		panic("container could not be null")
	}

	return func(c *gin.Context) {
		scoped, err := container.NewScope()
		if err != nil {
			c.AbortWithError(500, err)
		} else {
			c.Set(ContextKey, scoped)
		}
		c.Next()
	}
}
