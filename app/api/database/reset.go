package database

import (
  "context"
  "fmt"
  database "github.com/Masayuki-Suzuki/World-Wise-Backend/api/database/types"
  "github.com/gofiber/fiber/v3"
  "go.mongodb.org/mongo-driver/bson"
)

func (db *MongoDB) ResetDatabase(c fiber.Ctx) error {
  fmt.Println("Starting to reset database")

  cities, err := db.Collection.Find(context.TODO(), bson.D{{"userid", "8mbD54fkAObYNxLqb0pSsNlImw33"}})

  if err != nil {
    fmt.Println("Failed to find cities")
    return c.Status(500).JSON(fiber.Map{
      "message": "Failed to find cities",
    })
  }

  var result []database.City
  if err = cities.All(context.TODO(), &result); err != nil {
    fmt.Println("Failed to decode cities")
    return c.Status(500).JSON(fiber.Map{
      "message": "Failed to decode cities",
    })
  }

  if len(result) != 0 {
    deleteResult, err := db.Collection.DeleteMany(context.TODO(), bson.D{{"userid", "8mbD54fkAObYNxLqb0pSsNlImw33"}})
    if err != nil {
      fmt.Println("Failed to delete test user documents")
      return c.Status(500).JSON(fiber.Map{
        "message": "Failed to drop collection",
      })
    }
    fmt.Println("Deleted test user documents: ", deleteResult.DeletedCount)
  }

  seedCities := GetCitySeedData()

  for _, city := range seedCities {
    _, err := db.AddCityToDB(city)

    if err != nil {
      fmt.Println("Failed to add city")
      return c.Status(500).JSON(fiber.Map{
        "message": "Failed to add seed city",
      })
    }
  }

  return c.Status(200).JSON(fiber.Map{
    "message": "Database reset",
  })
}
