package main

import (
	"database/sql"
	"fmt"

	"github.com/bugscatcher/messages/config"
	_ "github.com/bugscatcher/messages/migration/common/migrations"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"github.com/rs/zerolog/log"
)

func main() {
	conf, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("Read config")
	}
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable", conf.PostgreSQL.Host, conf.PostgreSQL.User, conf.PostgreSQL.Database)
	migrationsConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("Open DB connection")
	}
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal().Err(err).Msg("Set dialect")
	}
	if err := goose.Up(migrationsConn, "."); err != nil {
		log.Fatal().Err(err).Msg("Migrate up")
	}
	_ = migrationsConn.Close()
}
