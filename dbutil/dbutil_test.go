package dbutil_test

import (
	"os"
	"testing"

	"github.com/airdb/sailor/dbutil"
)

func TestInitDB(t *testing.T) {
	if !testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	os.Setenv("testdb", "root:airdb.dev@tcp(127.0.0.1:3306)/")
	os.Setenv("testdb1", "root:airdb.dev@tcp(127.0.0.1:3306)/")

	dbs := []string{"testdb", "testdb1"}
	dbutil.InitDB(dbs)
}

/*
	data := make(map[string]interface{})
	data["user_id"] = 0
	data["status"] = "ok"
	data["updated_at"] = uint(time.Now().Unix())
 	db := tx.Updates(data)
*/
