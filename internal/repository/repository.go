package repository

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"github.com/vinoMamba/zero/pkg/log"
)

type Repository struct {
	db     *sql.DB
	logger *log.Logger
}

func NewRepository(logger *log.Logger, db *sql.DB) *Repository {
	return &Repository{
		logger: logger,
		db:     db,
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
