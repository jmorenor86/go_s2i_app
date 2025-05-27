package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("SERVER_PORT")
    if port == "" {
        port = "8080"  // puerto por defecto
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hola desde Go en puerto %s!", port)
    })

    log.Printf("Servidor corriendo en puerto %s\n", port)
    log.Fatal(http.ListenAndServe(":" + port, nil))
}
