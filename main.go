package main

import (
  "log"
  "fmt"
  "github.com/lodhel/golang_rest/book"
    //"github.com/gorilla/mux"
   //  _ "github.com/lib/pq"
   // _ "github.com/mattn/go-sqlite3"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)


func indexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Index Page")
}

func main() {
  db, err := gorm.Open("mysql", "root:123456@/test?charset=utf8&parseTime=True&loc=Local")

  if err != nil {
    panic(err)
  } 
  defer db.Close()
    r := mux.NewRouter()
    books = append(book.books, book.Book{ID: "1", Title: "Война и Мир", Author: &Author{Firstname: "Лев", Lastname: "Толстой"}})
    books = append(book.books, book.Book{ID: "2", Title: "Преступление и наказание", Author: &Author{Firstname: "Фёдор", Lastname: "Достоевский"}})
    r.HandleFunc("/books", getBooks).Methods("GET")
    r.HandleFunc("/books/{id}", getBook).Methods("GET")
    r.HandleFunc("/books", createBook).Methods("POST")
    r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
    r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000", r))
}