package main

import (
  "fmt"
  "net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, you've requested: %s!", r.URL.Path)
}

func main() {
  http.HandleFunc("/", HelloWorld)
  http.ListenAndServe(":8080", nil)
}