package dbutil

import (
	"log"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	dbs     sync.Map
	hasInit int32
	hasPend sync.Mutex
)

const (
	MainDatabaseWrite   = "AIRDB_MAIN_DATABASE_WRITE"
	MainDatabaseRead    = "AIRDB_MAIN_DATABASE_READ"
	SecondDatabaseWrite = "AIRDB_SECOND_DATABASE_WRITE"
	SecondDatabaseRead  = "AIRDB_SECOND_DATABASE_READ"
)

var DefaultDBs = []string{
	"MainDatabaseWrite",
	"MainDatabaseRead",
}

func InitDefaultDB() {
	InitDB(DefaultDBs)
}

func InitDB(dbNames []string) {
	if atomic.LoadInt32(&hasInit) == 1 {
		return
	}

	hasPend.Lock()
	defer hasPend.Unlock()

	for _, dbName := range dbNames {
		dsn := os.Getenv(dbName)

		if !strings.Contains(dsn, "?") {
			dsn += "?charset=utf8&parseTime=True"
		}

		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // Slow SQL threshold
				LogLevel:      logger.Info, // Log level
				Colorful:      false,       // Disable color
			},
		)

		conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: newLogger,
			NamingStrategy: schema.NamingStrategy{
				// TablePrefix:   "t_", // table name prefix, table for `User` would be `t_users`
				SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
			},
		})

		if err != nil {
			// panic("Error: connect to db server failed")
			log.Printf("Error: connect to databases %s failed, err is %v\n", dsn, err)
		} else {
			log.Printf("Connect to %s successfully.\n", dbName)
		}

		dbs.Store(dbName, conn)
	}

	if atomic.LoadInt32(&hasInit) == 1 {
		log.Println("concurrent_between_coroutines")

		return
	}

	atomic.StoreInt32(&hasInit, 1)
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
		log.Printf("database %s is nil.\n", name)
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
