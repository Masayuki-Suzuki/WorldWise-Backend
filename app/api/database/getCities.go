package database

import (
  "context"
  "fmt"
  database "github.com/Masayuki-Suzuki/World-Wise-Backend/api/database/types"
  firebaseController "github.com/Masayuki-Suzuki/World-Wise-Backend/api/firebase"
  "github.com/gofiber/fiber/v3"
  "go.mongodb.org/mongo-driver/bson"
  "strings"
)

func (db *MongoDB) GetCities(c fiber.Ctx) error {
  req := c.GetReqHeaders()
  reqToken := req["Authorization"][0]
  uuid := req["Uuid"][0]

  if reqToken == "" {
    return c.Status(400).JSON(fiber.Map{
      "message": "Token is required.",
      "cities":  []database.City{},
    })
  }

  if uuid == "" {
    return c.Status(400).JSON(fiber.Map{
      "message": "UUID is required.",
      "cities":  []database.City{},
    })
  }

  token := strings.Split(reqToken, "Bearer ")[1]

  if token == "" {
    return c.Status(400).JSON(fiber.Map{
      "message": "Token is required.",
    })
  }

  _, err := firebaseController.ValidateToken(token)
  if err != nil {
    return c.Status(401).JSON(fiber.Map{
      "message": "Failed to validate token.",
    })
  }

  result, err := db.Collection.Find(context.TODO(), bson.D{{"userid", uuid}})
  if err != nil {
    fmt.Println(err)
    return c.Status(500).JSON(fiber.Map{
      "message": "Failed to get cities.",
    })
  }

  var cities []database.City
  err = result.All(context.TODO(), &cities)
  if err != nil {
    fmt.Println(err)
    return c.Status(500).JSON(fiber.Map{
      "message": "Failed to get cities.",
    })
  }

  if cities == nil {
    cities = []database.City{}
  }

  return c.Status(200).JSON(fiber.Map{
    "message": "Cities are retrieved.",
    "cities":  cities,
  })

}
