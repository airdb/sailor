package dbutils

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"

	"github.com/airdb/sailor/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var dbs sync.Map
var hasInit int32
var hasPend sync.Mutex

type OperationType uint

const (
	OperationWrite OperationType = 1
	OperationRead  OperationType = 2
)

const defaultDB = "default"

func InitDefault() {
	if atomic.LoadInt32(&hasInit) == 1 {
		return
	}

	hasPend.Lock()
	defer hasPend.Unlock()

	if atomic.LoadInt32(&hasInit) == 1 {
		log.Println("cocurrent_between_coroutines")
		return
	}

	databases := config.GetDatabases()
	for name, item := range databases {
		db, err := gorm.Open(
			"mysql",
			fmt.Sprintf(
				"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
				item.User,
				item.Password,
				item.Address,
				item.Name,
			),
		)
		if err != nil {
			fmt.Println("Error: connect to db server failed, ", err)
			// panic("Error: connect to db server failed")
		} else {
			fmt.Println("init db success")
		}

		db.LogMode(true)
		db.SingularTable(true)
		dbs.Store(name, db)

		if item.Default {
			dbs.Store(defaultDB, db)
		}

	}

	atomic.StoreInt32(&hasInit, 1)
}

func WriteDB(name string) *gorm.DB {
	return DB(name, OperationWrite)
}

func ReadDB(name string) *gorm.DB {
	return DB(name, OperationRead)
}

func DefaultDB() (db *gorm.DB) {
	InitDefault()

	_db, ok := dbs.Load(defaultDB)
	if ok {
		db = _db.(*gorm.DB)
	}

	if db == nil {
		fmt.Println("error: ", db)
	}
	return
}

func DB(name string, typ OperationType) (db *gorm.DB) {
	InitDefault()

	var nameWithOperation string

	switch typ {
	case OperationWrite:
		nameWithOperation = fmt.Sprintf("%v.write", name)
	case OperationRead:
		nameWithOperation = fmt.Sprintf("%v.read", name)
	}

	if len(nameWithOperation) > 0 {
		_db, ok := dbs.Load(nameWithOperation)
		if ok {
			db = _db.(*gorm.DB)
		}
	}

	// Fallback to default db if could not find `nameWithOperation`.
	_db, ok := dbs.Load(name)
	if ok {
		db = _db.(*gorm.DB)
	}

	if db == nil {
		fmt.Println("error: ", db)
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
