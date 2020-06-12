package dbutils_test

import (
	"os"
	"testing"

	"github.com/airdb/sailor/dbutils"
)

func TestInitDB(t *testing.T) {
	os.Setenv("testdb", "root:airdb.dev@tcp(127.0.0.1:3306)/")
	os.Setenv("testdb1", "root:airdb.dev@tcp(127.0.0.1:3306)/")

	dbs := []string{"testdb", "testdb1"}
	dbutils.InitDB(dbs)
}
