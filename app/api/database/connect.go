package database

import (
  "context"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

func (db *MongoDB) Connect() error {
  serverAPI := options.ServerAPI(options.ServerAPIVersion1)
  opts := options.Client().ApplyURI(db.URI).SetServerAPIOptions(serverAPI)

  client, err := mongo.Connect(context.TODO(), opts)

  db.Client = client
  db.Database = client.Database("worldwise")
  db.ConnectionTest()

  return err
}
