package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rreubenreyes/hearth/internal/api"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	pgHost := os.Getenv("HEARTH_PG_HOST")
	pgUser := os.Getenv("HEARTH_PG_USER")
	pgPass := os.Getenv("HEARTH_PG_PASS")
	pgDB := os.Getenv("HEARTH_PG_DB")
	pgPort := os.Getenv("HEARTH_PG_PORT")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		pgHost, pgPort, pgUser, pgPass, pgDB)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	http.Handle("/api/v1/users", api.UsersServeMux(db))
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
