package dbutils

import (
	"fmt"
	"sync"

	"github.com/airdb/sailor/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var dbs sync.Map

type OperationType uint

const (
	OperationWrite OperationType = 1
	OperationRead  OperationType = 2
)

func InitDefault() {
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
	}
}

func WriteDB(name string) *gorm.DB {
	return DB(name, OperationWrite)
}

func ReadDB(name string) *gorm.DB {
	return DB(name, OperationRead)
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
