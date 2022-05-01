package api

import (
	"fmt"
	"log"
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

	err = api.Serve(&api.Backend{
		DB: db,
	})
	if err != nil {
		log.Fatalln(err)
	}
}
