package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type Position struct {
  Lat float64 `json:"lat"`
  Lng float64 `json:"lng"`
}

type PostCity struct {
  CityName string   `json:"cityName"`
  Country  string   `json:"country"`
  Emoji    string   `json:"emoji"`
  Date     string   `json:"date"`
  Notes    string   `json:"notes"`
  Position Position `json:"position"`
  UserID   string   `json:"userId"`
}

type PostCities struct {
  Cities []*PostCity `json:"cities"`
}

type City struct {
  ID       primitive.ObjectID `bson:"_id" json:"id"`
  CityName string             `json:"cityName"`
  Country  string             `json:"country"`
  Emoji    string             `json:"emoji"`
  Date     string             `json:"date"`
  Notes    string             `json:"notes"`
  Position Position           `json:"position"`
  UserID   string             `json:"userId"`
}

type Cities struct {
  Cities []*City `json:"cities"`
}

type RequestAddCity struct {
  Token    string `json:"token"`
  PostCity `json:"city"`
}

type RequestDeleteCity struct {
  Token string `json:"token"`
}
