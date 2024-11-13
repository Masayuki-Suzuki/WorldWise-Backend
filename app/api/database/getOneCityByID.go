package database

import (
  "context"
  "fmt"
  database "github.com/Masayuki-Suzuki/World-Wise-Backend/api/database/types"
  firebaseController "github.com/Masayuki-Suzuki/World-Wise-Backend/api/firebase"
  "github.com/gofiber/fiber/v3"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *MongoDB) GetOneCityByID(c fiber.Ctx) error {
  req := c.GetReqHeaders()
  rawToken := req["Authorization"]

  reqToken := ""

  if len(rawToken) > 0 {
    reqToken = rawToken[0]
  }

  if reqToken == "" {
    return c.Status(400).JSON(fiber.Map{
      "message": "Token is required.",
      "cities":  nil,
    })
  }

  msg := firebaseController.TokenValidation(reqToken)
  if msg != "" {
    fmt.Println("Failed to validate token.")
    return c.Status(401).JSON(fiber.Map{
      "message": msg,
    })
  }

  id := c.Params("id")
  if id == "" {
    return c.Status(400).JSON(fiber.Map{
      "message": "City ID is required.",
    })
  }

  objectId, err := primitive.ObjectIDFromHex(id)
  if err != nil {
    return c.Status(400).JSON(fiber.Map{
      "message": "Invalid city ID.",
    })
  }

  result := db.Collection.FindOne(context.TODO(), bson.M{"_id": objectId})

  var city database.City
  if err := result.Decode(&city); err != nil {
    return c.Status(500).JSON(fiber.Map{
      "message": "Failed to decode city.",
    })
  }

  return c.Status(200).JSON(fiber.Map{
    "message": "City is retrieved.",
    "city":    city,
  })
}
