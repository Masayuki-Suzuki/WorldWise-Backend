package database

import (
  database "github.com/Masayuki-Suzuki/World-Wise-Backend/api/database/types"
  firebaseController "github.com/Masayuki-Suzuki/World-Wise-Backend/api/firebase"
  "github.com/gofiber/fiber/v3"
)

func (db *MongoDB) AddOneCity(c fiber.Ctx) error {
  var body database.RequestAddCity

  // Parse request body
  if err := c.Bind().Body(&body); err != nil {
    return c.Status(400).JSON(fiber.Map{
      "message": "Incorrect request body.",
    })
  }

  // Check if token is provided
  if body.Token == "" {
    return c.Status(400).JSON(fiber.Map{
      "message": "Token is required.",
    })
  }

  // Validate token
  _, err := firebaseController.ValidateToken(body.Token)
  if err != nil {
    return c.Status(401).JSON(fiber.Map{
      "message": "Failed to validate token.",
    })
  }

  city := body.PostCity

  // Check if city is provided
  if city.CityName == "" || city.Country == "" || city.Emoji == "" || city.UserID == "" {
    return c.Status(400).JSON(fiber.Map{
      "message": "Invalid city data.",
    })
  }

  // Add city to database
  addedCity, err := db.AddCityToDB(&city)
  if err != nil {
    return c.Status(500).JSON(fiber.Map{
      "message": "Failed to add city.",
    })
  }

  return c.Status(200).JSON(fiber.Map{
    "message": "Successfully added city.",
    "city":    addedCity,
  })
}
