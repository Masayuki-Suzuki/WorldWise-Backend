package appAuth

import (
  appAuth "github.com/Masayuki-Suzuki/World-Wise-Backend/api/appAuth/types"
  firebaseController "github.com/Masayuki-Suzuki/World-Wise-Backend/api/firebase"

  "github.com/gofiber/fiber/v3"
)

func Login(c fiber.Ctx) error {
  data := new(appAuth.LoginToken)

  if err := c.Bind().Body(data); err != nil {
    return c.Status(400).JSON(fiber.Map{
      "message": "Invalid data.",
    })
  }

  if data.Token == "" {
    return c.Status(400).JSON(fiber.Map{
      "message": "Token is required.",
    })
  }

  token, err := firebaseController.ValidateToken(data.Token)
  if err != nil {
    return c.Status(401).JSON(fiber.Map{
      "message": "Invalid token.",
    })
  }

  firebaseAuth := firebaseController.GetFirebaseAuth()
  user, err := firebaseAuth.GetUser(c.Context(), token.UID)
  if err != nil {
    return c.Status(401).JSON(fiber.Map{
      "message": "Failed to get user.",
    })
  }

  return c.Status(200).JSON(fiber.Map{
    "message": "Login successful.",
    "user": fiber.Map{
      "uid":         user.UID,
      "email":       user.Email,
      "displayName": user.DisplayName,
    },
  })
}
