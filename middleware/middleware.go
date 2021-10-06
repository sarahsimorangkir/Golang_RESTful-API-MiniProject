package middleware

import (
	"github.com/labstack/echo/middleware"
)

var IsAutenthicated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})
