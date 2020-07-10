package book

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "net/http"

  "fmt"
  "encoding/json"
  "log"
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

    name := r.FormValue("name")
    soname := r.FormValue("soname")

    fmt.Println(name)
    db, err := gorm.Open("mysql", "root:123456@/test?charset=utf8")
    if err != nil {
      panic(err)
    }
    user := User{Name: name, Soname: soname}

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