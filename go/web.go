package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello!")
}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/", HomeHandler)
  http.Handle("/", r)
  http.ListenAndServe(":9000", nil)
}
