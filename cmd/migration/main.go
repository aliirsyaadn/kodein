package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/aliirsyaadn/kodein/internal/config"
	"github.com/aliirsyaadn/kodein/internal/log"
)

const (
	migrationTag = "Migration"
	cmdMigrate   = "migrate"
	cmdSeed      = "seed"
	schemaPath   = "file://files/sql/schemas"
	seedPath     = "file://files/sql/seeds"
	schemaTable  = "schema_migrations"
	seedTable    = "schema_seeds"
)

var (
	cmdFlag  = flag.String("cmd", "migrate", "Migrate or Seed")
	downFlag = flag.Bool("down", false, "Up or Down")
)

func main() {
	flag.Parse()
	cfg, err := config.LoadConfig()
	if err != nil {
		log.ErrorDetail(migrationTag, "error load config: %v", err)
		return
	}
	cfgDb := cfg.DB

	var path, table string
	switch *cmdFlag {
	case cmdMigrate:
		path = schemaPath
		table = schemaTable
	case cmdSeed:
		path = seedPath
		table = seedTable
	default:
		log.Fatal("Error invalid command")
	}

	ctx := context.Background()
	m, err := migrate.New(path, fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&x-migrations-table=%s", cfgDb.User, cfgDb.Password, cfgDb.Host, cfgDb.Port, cfgDb.DBName, cfgDb.SSLMode, table))

	if err != nil {
		log.ErrorDetail(migrationTag, "error create new migrate: %v", err)
		return
	}

	err = Execute(ctx, m)
	if err != nil {
		log.ErrorDetail(migrationTag, "error execute migration: %v", err)
		return
	}

}

func Execute(ctx context.Context, m *migrate.Migrate) (err error) {
	if *downFlag {
		err = m.Down()
	} else {
		err = m.Up()
	}
	if err == migrate.ErrNoChange {
		err = nil
	}
	return
}
