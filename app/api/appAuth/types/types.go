package appAuth

type LoginToken struct {
  Token string `json:"token"`
}

type SignUpFormData struct {
  Email     string `json:"email"`
  Password  string `json:"password"`
  FirstName string `json:"firstName"`
  LastName  string `json:"lastName"`
}
