package main

import(
	"net/http"
	"log"
	"fmt"
)

func main() {
    fmt.Println("Starting rest server...")
    http.HandleFunc("/reviews/0", func(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "dummy", 502)
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
}
