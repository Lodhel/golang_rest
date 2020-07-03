package book

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "net/http"

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

    name := r.FormValue("name")
    soname := r.FormValue("soname")

    fmt.Println(name)
    db, err := gorm.Open("mysql", "root:123456@/test?charset=utf8")
    if err != nil {
      panic(err)
    }

    _ = db.Exec("insert into test.users (name, soname) values (?, ?)", 
    name, soname)

    w.Header().Set("Content-Type", "application/json")
}