package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const defaultPort = "80"

func sing(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Do you believe in life after love?\n")
}

func walking(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "In Memphis\n")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	r := mux.NewRouter()
	r.HandleFunc("/", sing)
	r.HandleFunc("/sing", sing)
	r.HandleFunc("/walking", walking)
	fmt.Printf("==> Server listening at :%s 🚀\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
