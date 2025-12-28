package main

import (
    "fmt"
    "net/http"
    "os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, CI/CD World! Version 1.0")
}

func main() {
    http.HandleFunc("/", helloHandler)
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    fmt.Printf("Server starting on port %s\n", port)
    err := http.ListenAndServe(":"+port, nil)
    if err != nil {
        fmt.Printf("Error starting server: %s\n", err)
    }
}