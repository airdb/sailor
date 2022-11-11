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

	os.Setenv("DSN_MAIN_WRITE", "airdb:airdb@tcp(127.0.0.1:3306)/test")
	os.Setenv("DSN_MAIN_READ", "airdb:airdb@tcp(127.0.0.1:3306)/test")

	dbutil.InitDefaultDB()
	var users []string
	// select * from information_schema.user_privileges;

	db := dbutil.WriteDB(dbutil.DSNMainWrite).Table("information_schema.user_privileges").
		Select("GRANTEE").Distinct("GRANTEE").Find(&users).Debug()
	if db.Error != nil {
		panic(db.Error)
	}

	for _, user := range users {
		t.Log("user: ", user)
	}
}

/*
	data := make(map[string]interface{})
	data["user_id"] = 0
	data["status"] = "ok"
	data["updated_at"] = uint(time.Now().Unix())
 	db := tx.Updates(data)
*/
