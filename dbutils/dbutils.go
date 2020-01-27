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

type OperationType uint

const (
	defaultDB = "GDBC"
	readDB    = "READ_GDBC"
	writeDB   = "WRITE_GDBC"
)

type Database struct {
	Name string
	GDBC string
}

func GetDatabases() (databases []*Database) {
	// export GDBC="user:passwd@tcp(host:port)/dbname"
	// databases.GDBC = os.Getenv("GDBC")
	// DEFAULT GDBC
	if os.Getenv(defaultDB) != "" {
		databases = append(databases, &Database{
			Name: defaultDB,
			GDBC: os.Getenv(defaultDB),
		})
	}

	if os.Getenv(readDB) != "" {
		databases = append(databases, &Database{
			Name: readDB,
			GDBC: os.Getenv(readDB),
		})
	}

	if os.Getenv(writeDB) != "" {
		databases = append(databases, &Database{
			Name: writeDB,
			GDBC: os.Getenv(writeDB),
		})
	}

	return databases
}

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

	databases := GetDatabases()

	for name, item := range databases {
		db, err := gorm.Open(
			"mysql",
			fmt.Sprintf(
				"%s?charset=utf8&parseTime=True&loc=Local",
				item.GDBC,
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

		atomic.StoreInt32(&hasInit, 1)
	}
}

func WriteDB(name string) *gorm.DB {
	return DB(writeDB)
}

func ReadDB(name string) *gorm.DB {
	return DB(readDB)
}

func DefaultDB() (db *gorm.DB) {
	return DB(defaultDB)
}

func DB(name string) (db *gorm.DB) {
	InitDefault()

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
