package repository

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"github.com/vinoMamba/zero/pkg/log"
	"github.com/vinoMamba/zero/pkg/queries"
)

type Repository struct {
	queries *queries.Queries
	logger  *log.Logger
}

func NewRepository(logger *log.Logger, queries *queries.Queries) *Repository {
	return &Repository{
		logger:  logger,
		queries: queries,
	}
}

func NewDB(log *log.Logger, conf *viper.Viper) *sql.DB {
	cfg := mysql.Config{
		User:   conf.GetString("db.user"),
		Passwd: conf.GetString("db.password"),
		Net:    conf.GetString("db.net"),
		Addr:   conf.GetString("db.addr"),
		DBName: conf.GetString("db.name"),
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		panic(pingErr)
	}
	log.Info("Database connected")

	return db
}

func NewQueries(db *sql.DB) *queries.Queries {
	return queries.New(db)
}
