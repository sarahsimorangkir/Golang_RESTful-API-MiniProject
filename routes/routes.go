package routes

import (
	"echo-rest/controllers"
	"net/http"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "HELLO This is Echo")
	})

	e.GET("/admin", controllers.TakeAdmin)
	e.POST("/admin", controllers.AllAdmin)
	e.PUT("/admin", controllers.UpdateAdmin)
	e.DELETE("/admin", controllers.DeleteAdmin)

	e.GET("/product", controllers.TakeProduct)
	e.POST("/product", controllers.AllProduct)
	e.PUT("/product", controllers.UpdateProduct)
	e.DELETE("/product", controllers.DeleteProduct)

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)

	return e
}
