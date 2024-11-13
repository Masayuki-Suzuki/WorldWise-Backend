package database

import (
  "github.com/joho/godotenv"
  "log"
  "os"
)

func (db *MongoDB) GetURI() {
  if err := godotenv.Load(); err != nil {
    log.Fatalln("No .env file found")
  }

  db.URI = os.Getenv("MONGODB_URI")
  if db.URI == "" {
    log.Fatalln("MONGODB_URI is not set in environment variables.")
  }
}
