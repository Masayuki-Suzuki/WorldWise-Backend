package main

import (
  "github.com/Masayuki-Suzuki/World-Wise-Backend/api/database"
  firebaseController "github.com/Masayuki-Suzuki/World-Wise-Backend/api/firebase"
  "github.com/Masayuki-Suzuki/World-Wise-Backend/api/routing"
  "github.com/gofiber/fiber/v3/middleware/cors"
  "github.com/gofiber/fiber/v3/middleware/static"
  "log"

  "github.com/gofiber/fiber/v3"
  "github.com/gofiber/template/html/v2"
)

func main() {
  // Firebase setups
  firebaseController.InitFirebase()

  // Database setups
  database := new(database.MongoDB)
  database.Init()
  defer database.Disconnect()

  //Server configurations
  engine := html.New("./views", ".html")
  app := fiber.New(fiber.Config{
    Views: engine,
  })

  app.Use(cors.New(cors.Config{
    AllowOrigins: []string{"http://localhost:3000"},
  }))
  app.Use("/", static.New("./web/static"))

  routing.Setup(app, database)

  log.Fatal(app.Listen(":4000"))
}
