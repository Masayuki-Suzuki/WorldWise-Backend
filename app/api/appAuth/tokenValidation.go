package appAuth

import (
  appAuth "github.com/Masayuki-Suzuki/World-Wise-Backend/api/appAuth/types"
  firebaseController "github.com/Masayuki-Suzuki/World-Wise-Backend/api/firebase"
  "github.com/gofiber/fiber/v3"
)

func TokenValidation(c fiber.Ctx) error {
  var body appAuth.LoginToken

  if err := c.Bind().Body(&body); err != nil {
    return c.Status(400).JSON(fiber.Map{
      "message": "Incorrect request body.",
    })
  }

  if body.Token == "" {
    return c.Status(400).JSON(fiber.Map{
      "message": "Token is required.",
    })
  }

  _, err := firebaseController.ValidateToken(body.Token)
  if err != nil {
    return c.Status(401).JSON(fiber.Map{
      "message": "Failed to validate token.",
    })
  }

  return c.Status(200).JSON(fiber.Map{
    "message": "Token is valid.",
  })
}
