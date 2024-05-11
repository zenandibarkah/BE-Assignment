package db

import (
	"account-manager/helper"
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

type Database struct {
	db    *sql.DB
	count int
}

var (
	dbIdentity *Database
)

func InitDBConnection() error {
	if err := InitConnectionDBIdentity(); err != nil {
		return err
	}

	return nil
}

func InitConnectionDBIdentity() error {
	helper.Log.Println("start init postgres identity", nil)
	dbIdentity = new(Database)
	conn, err := CreateConnDatabase(helper.MyConfig.DBIdentity)
	if err != nil {
		return err
	}

	dbIdentity.db = conn
	dbIdentity.count = 0

	return nil
}

func CreateConnDatabase(DB string) (*sql.DB, error) {
	conn, err := sql.Open("postgres", DB)
	if err != nil {
		return nil, err
	}

	// if err := conn.Ping(); err != nil {
	// 	fmt.Println("=========")
	// 	return nil, err
	// }

	return conn, nil
}

func GetConnectionDB() (*sql.DB, error) {
	return dbIdentity.GetConnection()
}

func (db *Database) GetConnection() (*sql.DB, error) {
	if db.count > 0 {
		return nil, errors.New("server is still trying to connect to DB")
	}

	if err := db.db.Ping(); err != nil {
		go db.TryConnect()
		return nil, err
	}
	return db.db, nil
}

func (db *Database) TryConnect() {
	for {
		db.count++
		helper.Log.Info("trying to connect ", db.count, " times...")
		conn, err := CreateConnDatabase(helper.MyConfig.DBIdentity)
		if err != nil {
			db.count = 0
			break
		}
		db.db = conn
	}
}

func CloseConnectionDB() {
	helper.Log.Println("Closing Identity DB connection...")
	dbIdentity.db.Close()
}
