package main

import (
  "github.com/Masayuki-Suzuki/World-Wise-Backend/api/database"
  firebaseController "github.com/Masayuki-Suzuki/World-Wise-Backend/api/firebase"
  "github.com/Masayuki-Suzuki/World-Wise-Backend/api/routing"
  "github.com/gofiber/fiber/v3/middleware/cors"
  "github.com/gofiber/fiber/v3/middleware/static"
  "log"
  "os"

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

  originA := os.Getenv("ORIGIN_A")
  originB := os.Getenv("ORIGIN_B")
  originC := os.Getenv("ORIGIN_C")
  originD := os.Getenv("ORIGIN_D")

  origins := []string{originA, originB, originC, originD}

  app.Use(cors.New(cors.Config{
    AllowOrigins: origins,
  }))
  app.Use("/", static.New("./web/static"))

  routing.Setup(app, database)

  log.Fatal(app.Listen(":4000"))
}
