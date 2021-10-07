package routes

import (
	"echo-rest/controllers"
	"echo-rest/middleware"
	"net/http"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "HELLO This is Echo") //test
	})

	//Admin
	e.GET("/admin", controllers.TakeAdmin, middleware.IsAutenthicated)
	e.POST("/admin", controllers.AllAdmin, middleware.IsAutenthicated)
	e.PUT("/admin", controllers.UpdateAdmin, middleware.IsAutenthicated)
	e.DELETE("/admin", controllers.DeleteAdmin, middleware.IsAutenthicated)

	//Product
	e.GET("/product", controllers.TakeProduct)
	e.POST("/product", controllers.AllProduct, middleware.IsAutenthicated)
	e.PUT("/product", controllers.UpdateProduct, middleware.IsAutenthicated)
	e.DELETE("/product", controllers.DeleteProduct, middleware.IsAutenthicated)

	//Event
	e.GET("/event", controllers.TakeEvent)
	e.POST("/event", controllers.AllEvent, middleware.IsAutenthicated)
	e.PUT("/event", controllers.UpdateEvent, middleware.IsAutenthicated)
	e.DELETE("/event", controllers.DeleteEvent, middleware.IsAutenthicated)

	//Login
	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)

	e.GET("test-struct-validation", controllers.TestStructValidation)
	e.GET("test-variable-validation", controllers.TestVariableValidation)

	return e
}
