package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/joho/godotenv"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, Worldss!")
}

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Use environment variables
    serverPort := os.Getenv("SERVER_PORT")

    http.HandleFunc("/", helloHandler)
    fmt.Println("Starting server on :" + serverPort)
    if err := http.ListenAndServe(":"+serverPort, nil); err != nil {
        fmt.Println("Failed to start server:", err)
    }
}
