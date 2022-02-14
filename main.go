package main

import (
    "fmt"
    "net/http"
)

func sing(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Do you believe in life after love?\n")
}

func walking(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "In Memphis\n")
}

func main() {
    http.HandleFunc("/", sing)
    http.HandleFunc("/sing", sing)
    http.HandleFunc("/walking", walking)
    http.ListenAndServe(":8080", nil)
}
