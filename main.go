package main

import (
    "io"
    "log"
    "net/http"
)

const ServerPort = ":8000"

func main() {
    http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
        io.WriteString(writer, "Hello Hearth")
    })

    log.Printf("Listening on %s", ServerPort[1:])
    log.Fatal(http.ListenAndServe(ServerPort, nil))
}
