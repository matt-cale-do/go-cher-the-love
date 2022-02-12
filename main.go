package main

import (
    "fmt"
    "net/http"
)

func sing(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Do you believe in life after love?\n")
}

func main() {
    http.HandleFunc("/sing", sing)
    http.ListenAndServe(":7788", nil)
}
