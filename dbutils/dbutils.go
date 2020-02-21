package dbutils

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"

	_ "github.com/go-sql-driver/mysql" // golint
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var dbs sync.Map
var hasInit int32
var hasPend sync.Mutex

const (
	DBHOST   = "DB_HOST"
	DBPORT   = "DB_PORT"
	DBUSER   = "DB_USER"
	DBPASSWD = "DB_PASSWD"
	DBNAME   = "DB_NAME"
)

const DefaultEnvPrefix = ""

func InitDefault() {
	InitDB(DefaultEnvPrefix)
}

func InitDB(dbname string) {
	if atomic.LoadInt32(&hasInit) == 1 {
		return
	}

	hasPend.Lock()
	defer hasPend.Unlock()

	if atomic.LoadInt32(&hasInit) == 1 {
		log.Println("cocurrent_between_coroutines")
		return
	}

	viper.New()

	if dbname != DefaultEnvPrefix {
		viper.SetEnvPrefix(dbname)
	}
	viper.AutomaticEnv()

	gdbc := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString(DBUSER),
		viper.GetString(DBPASSWD),
		viper.GetString(DBHOST),
		viper.GetString(DBPORT),
		viper.GetString(DBNAME),
	)

	db, err := gorm.Open(
		"mysql",
		gdbc,
	)

	if err != nil {
		log.Println("Error: connect to db server failed, ", gdbc, err)
		panic("Error: connect to db server failed")
	} else {
		log.Println("Connect to db success")
	}

	db.LogMode(true)
	db.SingularTable(true)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName + "_tab"
	}

	dbs.Store(viper.GetString(DBNAME), db)

	atomic.StoreInt32(&hasInit, 1)
}

func WriteDB(name string) *gorm.DB {
	return DB(name)
}

func ReadDB(name string) *gorm.DB {
	return DB(name)
}

func DefaultDB() (db *gorm.DB) {
	return DB(DefaultEnvPrefix)
}

func DB(name string) (db *gorm.DB) {
	InitDB(name)

	if name == DefaultEnvPrefix {
		name = viper.GetString(DBNAME)
	}

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
