package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

type DB struct {
	DB *sql.DB
}

var dbConnect = &DB{}

const maxOpenDbConns = 10
const maxIdleDbConns = 5
const macDbLifetime = 5 * time.Minute

func ConnectPostgres(dsn string) (*DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDbConns)
	db.SetMaxIdleConns(maxIdleDbConns)
	db.SetConnMaxLifetime(macDbLifetime)

	err = testDb(db)
	if err != nil {
		return nil, err
	}

	dbConnect.DB = db
	return dbConnect, nil

}

func testDb(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		fmt.Println("Database error:", err)
		return err
	}
	fmt.Println("*** Pinged db successfully! ***")
	return nil
}

// {
//   "name" : "latte",
//   "roast": "Medium",
//    "image": "https://media.istockphoto.com/id/505168330/photo/cup-of-cafe-latte-with-coffee-beans-and-cinnamon-sticks.jpg?s=612x612&w=0&k=20&c=h7v8kAfWOpRapv6hrDwmmB54DqrGQWxlhP_mTeqTQPA=",
//    "region" : "Japn",
//    "price": 30,
//    "grind_unit": 5
// }
