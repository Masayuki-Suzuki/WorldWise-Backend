package firebaseController

import (
  "context"
  "firebase.google.com/go/v4"
  "firebase.google.com/go/v4/auth"
  "fmt"
  "google.golang.org/api/option"
  "log"
  "strings"
)

var firebaseAuth *auth.Client

func InitFirebase() {
  opt := option.WithCredentialsFile("./firebase-adminsdk.json")
  app, err := firebase.NewApp(context.Background(), nil, opt)
  if err != nil {
    log.Fatalf("Error initializing firebase app: %v\n", err)
  }

  firebaseAuth, err = app.Auth(context.Background())
  if err != nil {
    log.Fatalf("Error initializing firebase Auth: %v\n", err)
  }

  fmt.Println("Firebase initialized.")
}

func ValidateToken(idToken string) (*auth.Token, error) {
  token, err := firebaseAuth.VerifyIDToken(context.Background(), idToken)

  if err != nil {
    fmt.Println("Error validating token (firebase-main.go/line35): ", err)
    return nil, err
  }

  return token, nil
}

func GetFirebaseAuth() *auth.Client {
  return firebaseAuth
}

func TokenValidation(reqToken string) string {
  strMap := strings.Split(reqToken, "Bearer ")

  token := ""
  if len(strMap) > 1 {
    token = strMap[1]
  }

  if token == "" {
    return "Token is required."
  }

  _, err := ValidateToken(token)
  if err != nil {
    return "Failed to validate token."
  }

  return ""
}
