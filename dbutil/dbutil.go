package dbutil

import (
	"fmt"
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
	SecondDatabaseWrite = "AIRDB_SECOND_DATABASE_WRITE"
	SecondDatabaseRead  = "AIRDB_SECOND_DATABASE_READ"
	MainDSNWrite        = "MAIN_DSN_WRITE"
	MainDSNRead         = "MAIN_DSN_READ"
)

var DefaultDBs = []string{
	"MainDSNWrite",
	"MainDSNRead",
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
		conn := Connect(MainDSNWrite)

		dbs.Store(dbName, conn)
	}

	fmt.Println("xxx", dbs)

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

const (
	maxIdleConn = 2
	maxOpenConn = 5
)

// Connection gets connection of mysql database
func Connect(dbName string) (db *gorm.DB) {
	dsn := os.Getenv(dbName)
	if !strings.Contains(dsn, "?") {
		dsn += "?charset=utf8mb4&parseTime=True&loc=Local"
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tab_", // table name prefix, table for `User` would be `t_users`
			SingularTable: true,   // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()

	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetMaxOpenConns(maxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
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
