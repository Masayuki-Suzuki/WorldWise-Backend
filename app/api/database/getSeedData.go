package database

import (
  "encoding/json"
  "fmt"
  database "github.com/Masayuki-Suzuki/World-Wise-Backend/api/database/types"
  "os"
)

func GetCitySeedData() []*database.PostCity {
  seeds := database.PostCities{}

  seedData, err := os.Open("seeds/cities-seed.json")
  if err != nil {
    fmt.Println("Failed to open seed data")
    return nil
  }

  json.NewDecoder(seedData).Decode(&seeds)

  return seeds.Cities
}
