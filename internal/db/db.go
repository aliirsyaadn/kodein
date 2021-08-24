package db

import (
	"database/sql"
	"fmt"

	"github.com/aliirsyaadn/kodein/internal/config"
	"github.com/aliirsyaadn/kodein/internal/log"
	"github.com/aliirsyaadn/kodein/model"
)

const intDBTag = "InternalDBTag"

func ConnectDB(dbConfig config.DBConfig) *model.Queries {
	sqldb, err := sql.Open("postgres", fmt.Sprintf("dbname=%s user=%s password=%s sslmode=%s", dbConfig.DBName, dbConfig.User, dbConfig.Password, dbConfig.SSLMode))

	if err != nil {
		log.ErrorDetail(intDBTag, "error connect database: %v", err)
		return nil
	}

	queries := model.New(sqldb)
	return queries
}
