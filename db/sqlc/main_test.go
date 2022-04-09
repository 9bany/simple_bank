package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Can not connect to database:", err)
	}
	testDb = conn
	testQueries = New(conn)

	os.Exit(m.Run())
}
