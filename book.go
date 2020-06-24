package book

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "net/http"
  "encoding/json"
  "fmt"
)


type User struct {
  gorm.Model
  Name   string 
  Soname  string 
}

type Author struct {
  Firstname string 
  Lastname  string
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var user User
    _ = json.NewDecoder(r.Body).Decode(&user)
    json.NewEncoder(w).Encode(user)
}