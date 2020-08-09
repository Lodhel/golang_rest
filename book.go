package book

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "net/http"

  "fmt"
  "encoding/json"
  "log"
  "regexp"
  "errors"
)


type User struct {
  gorm.Model
  Email string `json:"email" validate:"required,email"`
  Name   string 
  Soname  string 
}

type Author struct {
  Firstname string 
  Lastname  string
}

type ErrorResponse struct {
  response string
}

var (
  badformat = errors.New("invalid format")
  emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

func serilizer(email string, name string, soname string) {
  valide := emailRegexp.MatchString(email)
  request := ErrorResponse{response: "error"}
  data, err := json.Marshal(request)
  if err != nil {
      fmt.Println("bad request")
    }
  if valide == true {
      fmt.Println(data)
  }
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

    email := r.FormValue("email")
    name := r.FormValue("name")
    soname := r.FormValue("soname")

    db, err := gorm.Open("mysql", "root:123456@/test?charset=utf8")
    if err != nil {
      panic(err)
    }
    user := User{Email: email, Name: name, Soname: soname}

    serilizer(email, name, soname)

    instance := db.Create(&user)

    data, err := json.Marshal(instance)
    if err != nil {
        log.Println(err)
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Authorization", "token")
    w.WriteHeader(http.StatusOK)
    _, _ = w.Write(data)
}