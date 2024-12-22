package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello, you've requested: %s</h1>", r.URL.Path)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Starting server at port 8080")
    http.ListenAndServe(":8080", nil)
}