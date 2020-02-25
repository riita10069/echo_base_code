package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/riita10069/echo_base_code/handler"
)

func New() *echo.Echo {
	// echo.new()
	e := echo.New()
	// using echo middleware
	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: []string{"*"},
	//	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	//	AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	//}))
	e.Validator = NewValidator()

	// routing
  e.GET("/healthcheck", func(c echo.Context) error {
    return c.String(200, "OK")
  })

  // need not login page
	e.POST("/login", handler.Login)
	e.POST("/create", handler.Create)

	// have to login page
	//api := e.Group("/api")
	//api.Use(middleware.JWTWithConfig(handler.Config))
	//api.GET("/bcp", handler.GetBcp)
	//api.POST("/gcp", handler.AddBcp)
	//api.DELETE("/bcp/:id", handler.DeleteBcp)
	//api.PUT("/gcp/:id/completed", handler.UpdateBcp)

	return e
}
