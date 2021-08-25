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

	os.Setenv("MAIN_DSN_WRITE", "root:root@tcp(127.0.0.1:3306)/test")
	os.Setenv("MAIN_DSN_READ", "root:root@tcp(127.0.0.1:3306)/test")

	dbutil.InitDefaultDB()
	var users []string
	// select * from information_schema.user_privileges;

	db := dbutil.WriteDB(dbutil.MainDSNWrite).Table("information_schema.user_privileges").Select("GRANTEE").Distinct("GRANTEE").Find(&users)
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
