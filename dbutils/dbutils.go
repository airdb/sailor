package dbutils

import (
	"fmt"
	"log"
	"os"
	"sync"
	"sync/atomic"

	_ "github.com/go-sql-driver/mysql" // golint
	"github.com/jinzhu/gorm"
)

var dbs sync.Map
var hasInit int32
var hasPend sync.Mutex

func InitDB(dbNames []string) {
	for _, dbName := range dbNames {
		if atomic.LoadInt32(&hasInit) == 1 {
			return
		}

		hasPend.Lock()
		defer hasPend.Unlock()

		if atomic.LoadInt32(&hasInit) == 1 {
			log.Println("cocurrent_between_coroutines")
			return
		}

		gdbc := os.Getenv(dbName)

		conn, err := gorm.Open("mysql", gdbc)
		if err != nil {
			log.Println("Error: connect to db server failed, ", gdbc, err)
			panic("Error: connect to db server failed")
		} else {
			log.Println("Connect to db success")
		}

		conn.LogMode(true)
		conn.SingularTable(true)

		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			return defaultTableName + "_tab"
		}

		dbs.Store(dbName, conn)

		atomic.StoreInt32(&hasInit, 1)
	}
}

func WriteDB(name string) *gorm.DB {
	return DB(name)
}

func ReadDB(name string) *gorm.DB {
	return DB(name)
}

func DB(name string) (db *gorm.DB) {
	_db, ok := dbs.Load(name)
	if ok {
		db = _db.(*gorm.DB)
	}

	if db == nil {
		fmt.Println("error db: ", db)
	}

	return
}

// InitTestDB will init the mock DB and lock the db so that the actual db will not be required.
func InitTestDB(name string, db *gorm.DB) error {
	if !atomic.CompareAndSwapInt32(&hasInit, 0, 1) {
		return nil
	}

	dbs.Store(name, db)

	return nil
}

// ReleaseTestDB is to release the lock for other unit tests to mock db successfully.
func ReleaseTestDB() {
	atomic.CompareAndSwapInt32(&hasInit, 1, 0)
}
