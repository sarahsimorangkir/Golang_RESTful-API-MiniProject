package middleware

import (
	"echo-rest/config"
	"time"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

func LogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	coll := config.DBLog.Database("kreasikuapp").Collection("logger")

	return func(c echo.Context) error {
		request := c.Request()
		response := c.Response()
		log := bson.M{
			"time":   time.Now(),
			"method": request.Method,
			"path":   request.URL.Path,
			"status": response.Status,
		}

		err := next(c)
		coll.InsertOne(request.Context(), log)
		return err

	}
}
