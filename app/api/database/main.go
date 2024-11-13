package database

import (
  "context"
  "fmt"
  "github.com/gofiber/fiber/v3"
  "go.mongodb.org/mongo-driver/mongo"
  "log"
)

type MongoDB struct {
  Client     *mongo.Client
  Database   *mongo.Database
  Collection *mongo.Collection
  URI        string
}

func (db *MongoDB) Init() {
  fmt.Println("Database init")
  db.GetURI()
  fmt.Printf("Connecting to %v\n", db.URI)
  db.Connect()
  fmt.Printf("Create collection\n")
  db.CreateCollection()
}

func (db *MongoDB) CreateCollection() {
  db.Collection = db.Database.Collection("cities")
}

func (db *MongoDB) Disconnect() {
  if err := db.Client.Disconnect(context.TODO()); err != nil {
    log.Fatalln("Failed to disconnect from MongoDB")
    log.Fatalln(err)
  }
}

func (db *MongoDB) GetClient() *mongo.Client {
  fmt.Println("Getting client")
  return db.Client
}

func (db *MongoDB) GetCollection() *mongo.Collection {
  return db.Collection
}

func (db *MongoDB) SetOneSeedData(c fiber.Ctx) error {
  seedCities := GetCitySeedData()
  _, err := db.AddCityToDB(seedCities[0])
  if err != nil {
    return c.Status(500).JSON(fiber.Map{
      "message": "Failed to add city",
    })
  }

  return c.Status(200).JSON(fiber.Map{
    "message": "Successfully added city",
  })
}
