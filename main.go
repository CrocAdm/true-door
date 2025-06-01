package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/join", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Lobby: connection accepted")
    })

    http.HandleFunc("/publish", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Node published")
    })

    http.HandleFunc("/resolve", func(w http.ResponseWriter, r *http.Request) {
        domain := r.URL.Query().Get("domain")
        fmt.Fprintf(w, "Resolving domain: %s\n", domain)
    })

    log.Println("Lobby server running on :443")
    log.Fatal(http.ListenAndServe(":443", nil))
}
