package database

import (
  "context"
  "fmt"
  "os"
)

func (db *MongoDB) ConnectionTest() bool {
  fmt.Println("Testing database connection...")
  client := db.GetClient()

  if client == nil {
    fmt.Println("Client is nil")
    return false
  }

  fmt.Println("Pinging database...")
  fmt.Printf("Connecting to %v...\n", os.Getenv("MONGODB_URI"))
  if err := client.Ping(context.Background(), nil); err != nil {
    fmt.Println("Failed to ping database")
    fmt.Println(err)
    return false
  }

  fmt.Println("Connected to database properly.")
  return true
}
