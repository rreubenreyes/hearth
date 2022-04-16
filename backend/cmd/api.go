package api

import (
  "net/http"
  "log"
  "github.com/rreubenreyes/hearth/internal/api"
)

func main() {
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Fatalln(err)
  }
}
