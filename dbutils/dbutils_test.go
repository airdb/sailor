package dbutils_test

import (
	"os"
	"testing"

	"github.com/airdb/sailor/dbutils"
)

func TestInitDB(t *testing.T) {
	os.Setenv("testdb", "root:airdb.dev@tcp(127.0.0.1:3306)/")

	dbs := []string{"testdb"}
	dbutils.InitDB(dbs)
}
