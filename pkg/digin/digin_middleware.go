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
			c.Errors = append(c.Errors, &gin.Error{Err: err, Type: gin.ErrorTypeAny})
		} else {
			err := di.RegisterFactory(scoped, di.Scoped, func(c di.Container) di.Container { return *scoped }, true)
			if err != nil {
				c.Errors = append(c.Errors, &gin.Error{Err: err, Type: gin.ErrorTypeAny})
			} else {
				c.Set(ContextKey, scoped)
			}

		}
		c.Next()
	}
}
