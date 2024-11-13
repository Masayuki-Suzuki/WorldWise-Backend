package database

import "github.com/gofiber/fiber/v3"

func Routing(api fiber.Router, db *MongoDB) {
  api.Get("cities", func(c fiber.Ctx) error {
    return db.GetCities(c)
  })

  api.Get("cities/:id", func(c fiber.Ctx) error {
    return db.GetOneCityByID(c)
  })

  api.Post("reset-database", func(c fiber.Ctx) error {
    return db.ResetDatabase(c)
  })

  api.Post("add-one-seed-city", func(c fiber.Ctx) error {
    return db.SetOneSeedData(c)
  })

  api.Post("city", func(c fiber.Ctx) error {
    return db.AddOneCity(c)
  })

  api.Delete("cities/:id", func(c fiber.Ctx) error {
    return db.DeleteOneCity(c)
  })
}
