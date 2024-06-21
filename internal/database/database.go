package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)
type DB struct {
	DB *sql.DB
}

var dbConn = &DB{}

const maxOpenConns = 10
const maxIdleConns = 5
const maxDbLifeTime = 5 * time.Minute

func ConnectPostgresDB(dsn string) (*DB, error){
	d, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	d.SetMaxOpenConns(maxOpenConns)
	d.SetMaxIdleConns(maxIdleConns)
	d.SetConnMaxLifetime(maxDbLifeTime)

	err = testDB(d)
	if err != nil {
		return nil, err
	}
	dbConn.DB = d
	return dbConn, nil
}

func testDB( d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		fmt.Println("Error pinging database", err)
		return err
	}
	fmt.Println("*** Pinged database successfully ***")
	return nil
}