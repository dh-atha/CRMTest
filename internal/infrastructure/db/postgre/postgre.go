package postgre

import (
	"database/sql"
	"fmt"
	"net/url"
	"strings"

	"github.com/dh-atha/EmployeeAbsenceKNTest/pkg/config"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func NewPostgres() (*sqlx.DB, error) {
	cfg := config.Configuration.Postgres

	url := (&url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.Username, cfg.Password),
		Host:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Path:     cfg.Database,
		RawQuery: strings.Join(cfg.Options, "&"),
	}).String()

	dbConn, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	db := sqlx.NewDb(dbConn, "postgres")
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
