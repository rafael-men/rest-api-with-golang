package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/rafael-men/rest-api-with-golang/config"
)

func OpenConnection() (*sql.DB, error) {
	conf := config.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgres", sc)

	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
