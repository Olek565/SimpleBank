package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

// var testQueries *Queries
// var testDB *sql.DB

// func TestMain(m *testing.M) {
// 	config, err := util.LoadConfig("../..")
// 	if err != nil {
// 		log.Fatal("Cannot load config: ", err)
// 	}

// 	testDB, err = sql.Open(config.DBDriver, config.DBSource)
// 	if err != nil {
// 		log.Fatal("Cannot connect to db:", err)
// 	}

// 	testQueries = New(testDB)

// 	os.Exit(m.Run())
// }

var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
