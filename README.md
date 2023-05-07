# digin

digin is a Go library that provides middleware for the Gin web framework, allowing integration with dependency injection using the `github.com/ns-go/di` package.

## Installation

To use digin in your Go project, you need to install it first. Run the following command:

```shell
go get github.com/ns-go/digin
```

## Usage

Import the required packages:

```go
import (
	"github.com/gin-gonic/gin"
	"github.com/ns-go/di/pkg/di"
	"github.com/ns-go/digin"
)
```

### Middleware

The `Container` function in the digin package provides middleware for Gin that sets up a dependency injection container scope and stores it in the Gin context. This allows injecting dependencies into your Gin handlers.

To use the `Container` middleware with Gin, follow these steps:

1. Create a container using `di.NewContainer()`:

   ```go
   container := di.NewContainer()
   ```

2. Register your services with the container using `di.RegisterSingleton` or other registration methods provided by `github.com/ns-go/di/pkg/di`.

   ```go
   di.RegisterSingleton[Service1](container, false)
   ```

3. Create a Gin router using `gin.Default()` or any other Gin router initialization method.

4. Add the digin middleware to the Gin router using `e.Use(digin.Container(container))`.

   ```go
   e := gin.Default()
   e.Use(digin.Container(container))
   ```

Now, the container scope will be available in your Gin handlers via the `ContextKey` constant.
   ```go
    ctnVal, _ := c.Get(digin.ContextKey) //c *gin.Context
    container, _ := ctnVal.(*di.Container)
   ```
You can use function `digin.ResolveHandlerFunc` for resolve handler function. Example
- Declare type of `Service` and `Handlers`
   ```go
    type Service1 struct {
	    container *di.Container `di.inject:"scoped"`
    }

    type TestHandlers struct {
	    service *Service1 `di.inject:""`
    }

    func (h *TestHandlers) TestResponse(c *gin.Context) {
        //do something
    }
   ```
- Register to `container`
   ```go
    di.RegisterScoped[Service1](constainer, false)
	di.RegisterScoped[TestHandlers](constainer, false)  
   ```
   
- Mapping to `router`
   ```go
    router.GET("some-path", digin.ResolveHandlerFunc(func(th *TestHandlers) gin.HandlerFunc { return th.TestResponse }))
   ```

Make sure to handle any errors that may occur during the creation of the container scope and add appropriate error handling logic in your Gin handlers.

## License

This project is licensed under the [MIT License](LICENSE).