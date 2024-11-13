package routing

import (
  "fmt"
  "github.com/Masayuki-Suzuki/World-Wise-Backend/api/appAuth"
  "github.com/Masayuki-Suzuki/World-Wise-Backend/api/database"
  "github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App, db *database.MongoDB) {
  api := app.Group("api")

  database.Routing(api, db)
  appAuth.Routing(api)

  app.Get("/", func(c fiber.Ctx) error {
    fmt.Println("Starting server")
    return c.Status(200).Render("index", fiber.Map{})
  })
}
