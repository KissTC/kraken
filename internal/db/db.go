package db

import (
	"database/sql"
	"fmt"
)

func GetConnection(user, password, dbname, host string, port int) (*sql.DB, error) {
	uri := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)
	return sql.Open("postgres", uri)
}
