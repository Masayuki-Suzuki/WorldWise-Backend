package appAuth

import (
  "github.com/gofiber/fiber/v3"
)

func Routing(api fiber.Router) {
  auth := api.Group("auth")

  auth.Post("signup", SignUp)
  auth.Post("login", Login)
  auth.Post("token-validation", TokenValidation)
}
