package appAuth

import (
  "context"
  "firebase.google.com/go/v4/auth"
  appAuth "github.com/Masayuki-Suzuki/World-Wise-Backend/api/appAuth/types"
  firebaseController "github.com/Masayuki-Suzuki/World-Wise-Backend/api/firebase"
  "github.com/gofiber/fiber/v3"
)

func SignUp(c fiber.Ctx) error {
  firebaseAuth := firebaseController.GetFirebaseAuth()
  req := new(appAuth.SignUpFormData)

  if err := c.Bind().Body(req); err != nil {
    return c.Status(400).JSON(fiber.Map{
      "message": "Incorrect request body.",
      "details": err.Error(),
    })
  }

  if req.Email == "" || req.Password == "" {
    return c.Status(400).JSON(fiber.Map{
      "message": "Email and password are required.",
    })
  }

  if req.FirstName == "" || req.LastName == "" {
    return c.Status(400).JSON(fiber.Map{
      "message": "First name and last name are required.",
    })
  }

  displayName := req.FirstName + " " + req.LastName
  params := (&auth.UserToCreate{}).Email(req.Email).Password(req.Password).DisplayName(displayName)
  userRecord, err := firebaseAuth.CreateUser(context.Background(), params)

  if err != nil {
    return c.Status(401).JSON(fiber.Map{
      "message": "Failed to create user.",
      "details": err.Error(),
    })
  }

  token, err := firebaseAuth.CustomToken(context.Background(), userRecord.UID)
  if err != nil {
    return c.Status(401).JSON(fiber.Map{
      "message": "Failed to create custom token.",
      "details": err.Error(),
    })
  }

  c.Status(200).JSON(fiber.Map{
    "id":        userRecord.UID,
    "email":     userRecord.Email,
    "firstName": req.FirstName,
    "lastName":  req.LastName,
    "token":     token,
  })

  return nil
}
