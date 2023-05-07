package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ns-go/di-gin/pkg/digin"
	"github.com/ns-go/di/pkg/di"
)

type Service1 struct {
}

func TestMiddleware(t *testing.T) {
	constainer := di.NewContainer()

	di.RegisterSingleton[Service1](constainer, false)
	r := gin.Default()

	r.Use(digin.Container(constainer))

	r.GET("test", func(c *gin.Context) {
		ctn, _ := c.Get(digin.ContextKey)
		if val, ok := ctn.(*di.Container); !ok {
			t.Errorf(`gin.Context.Get(digin.ContextKey).(*di.Container) = %v; want %v`, val, di.Container{})
		}

		ctn2, _ := ctn.(*di.Container)
		s1, err := di.Resolve[Service1](ctn2)
		if s1 == nil || err != nil {
			t.Errorf(`s1, err := di.Resolve[Service1](ctn2) = (%v, %v); want (%v, %v)`, s1, err, Service1{}, nil)
		}
		if s1 != nil {
			c.String(200, "success")
		}
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	r.ServeHTTP(w, req)

	if w.Body.String() != "success" {
		t.Errorf("w.Body.String() == %v; want %v", w.Body.String(), "success")
	}
}

func TestMiddlewareError(t *testing.T) {
	constainer := di.NewContainer()

	di.RegisterSingleton[Service1](constainer, false)
	w := httptest.NewRecorder()
	_, e := gin.CreateTestContext(w)
	scoped, _ := constainer.NewScope()
	e.Use(digin.Container(scoped))

	e.GET("test", func(c *gin.Context) {
		if len(c.Errors) == 0 {
			t.Errorf("len(c.Errors) == %v; want %v", len(c.Errors), 1)
		}
	})

	req, _ := http.NewRequest("GET", "/test", nil)
	e.ServeHTTP(w, req)
}
