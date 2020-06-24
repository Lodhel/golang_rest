package main

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"

  "github.com/gorilla/mux"
  "net/http"
  "log"
  "fmt"
)

import k "Lodhel/tests"

func indexHandler(w http.ResponseWriter, r *http.Request) string {
    fmt.Fprint(w, "Index Page")

    return ""
}

func main() {
    r := mux.NewRouter()
    db, err := gorm.Open("mysql", "root:123456@/test?charset=utf8")
    if err != nil {
      panic(err)
    }

    db.Debug().AutoMigrate(&k.User{})
    fmt.Println("connection")

    r.HandleFunc("/user", k.CreateUser).Methods("POST")
    log.Fatal(http.ListenAndServe(":8000", r))
}