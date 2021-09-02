package db

import (
	"database/sql"

	"github.com/aliirsyaadn/kodein/internal/config"
	"github.com/aliirsyaadn/kodein/internal/log"
	"github.com/aliirsyaadn/kodein/model"
)

const intDBTag = "InternalDBTag"

func ConnectDB(dbConfig config.DBConfig) *model.Queries {
	sqldb, err := sql.Open("postgres", ParseDSN(dbConfig))

	if err != nil {
		log.ErrorDetail(intDBTag, "error connect database: %v", err)
		return nil
	}

	queries := model.New(sqldb)
	return queries
}
