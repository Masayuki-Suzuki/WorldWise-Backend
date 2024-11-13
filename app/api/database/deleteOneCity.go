package database

import (
  "context"
  firebaseController "github.com/Masayuki-Suzuki/World-Wise-Backend/api/firebase"
  "github.com/gofiber/fiber/v3"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *MongoDB) DeleteOneCity(c fiber.Ctx) error {
  id := c.Params("id")
  req := c.GetReqHeaders()
  reqToken := req["Authorization"][0]

  msg := firebaseController.TokenValidation(reqToken)
  if msg != "" {
    return c.Status(401).JSON(fiber.Map{
      "message": msg,
    })
  }

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

  result, err := db.Collection.DeleteOne(context.TODO(), bson.M{"_id": objectId})
  if err != nil {
    return c.Status(500).JSON(fiber.Map{
      "message": "Failed to delete city.",
    })
  }

  if result.DeletedCount == 0 {
    return c.Status(404).JSON(fiber.Map{
      "message": "City not found.",
    })
  }

  return c.Status(200).JSON(fiber.Map{
    "message": "Successfully deleted city.",
  })
}
