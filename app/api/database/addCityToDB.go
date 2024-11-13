package database

import (
  "context"
  "errors"
  database "github.com/Masayuki-Suzuki/World-Wise-Backend/api/database/types"
  "go.mongodb.org/mongo-driver/bson"
)

func (db *MongoDB) AddCityToDB(city *database.PostCity) (database.City, error) {
  result, err := db.Collection.InsertOne(context.TODO(), city)
  emptyCity := database.City{}

  if err != nil {
    return emptyCity, err
  }

  if result.InsertedID == nil {
    return emptyCity, errors.New("failed to insert city")
  }

  var addedCity database.City
  err = db.Collection.FindOne(context.TODO(), bson.M{"_id": result.InsertedID}).Decode(&addedCity)
  if err != nil {
    return emptyCity, err
  }

  return addedCity, nil
}
