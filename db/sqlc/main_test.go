package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/mshero7/simplebank/util"
)

// Queries 에 method 들로 붙어있다
var testQueries *Queries
var testDB *sql.DB

// test main entry point
// test할때 TestMain은 무조건 타기에 전역변수인 var에 접근가능해진다
func TestMain(m *testing.M) {
	var err error

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
