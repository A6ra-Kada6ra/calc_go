package main

import (
    "log"
    "net/http"

    "github.com/pashapdev/calc_go/internal/application"
)

func main() {
    app := application.New()
    if err := app.RunServer(); err != nil {
        log.Fatalf("Could not start server: %s\n", err.Error())
    }
}
