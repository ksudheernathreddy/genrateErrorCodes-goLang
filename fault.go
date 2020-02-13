package main

import(
        "net/http"
        "log"
        "fmt"
        "time"
        "os"
        )

func main() {
dur, _ := time.ParseDuration(os.Getenv("SLEEP_DURATION"))
fmt.Printf("Starting rest server with %v...", dur)
http.HandleFunc("/reviews/0", func(w http.ResponseWriter, r *http.Request) {
 time.Sleep(dur)
 http.Error(w, "dummy", 502)
})
log.Fatal(http.ListenAndServe(":9080", nil))
}
