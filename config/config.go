package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

var (
	Get                     = viper.Get
	GetBool                 = viper.GetBool
	GetDuration             = viper.GetDuration
	GetFloat64              = viper.GetFloat64
	GetInt                  = viper.GetInt
	GetInt32                = viper.GetInt32
	GetInt64                = viper.GetInt64
	GetSizeInBytes          = viper.GetSizeInBytes
	GetString               = viper.GetString
	GetStringMap            = viper.GetStringMap
	GetStringMapString      = viper.GetStringMapString
	GetStringMapStringSlice = viper.GetStringMapStringSlice
	GetStringSlice          = viper.GetStringSlice
	GetTime                 = viper.GetTime
	IsSet                   = viper.IsSet
	AllSettings             = viper.AllSettings
)

type Database struct {
	User     string
	Password string
	Address  string
	Name     string
}

func GetEnv() string {
	if os.Getenv("ENV") == "" {
		// default return dev
		return "dev"
	}
	return os.Getenv("ENV")
}

func GetDatabases() (databases map[string]*Database) {
	err := viper.UnmarshalKey("databases", &databases)
	if err != nil {
		log.Fatal("could not parse config for databases")
	}

	for name := range databases {
		splits := strings.SplitN(name, ".", 2)
		operationType := splits[len(splits)-1]
		if operationType == "read" || operationType == "write" {
			databases[name].Name = splits[0]
		} else {
			databases[name].Name = name
		}
	}

	return databases
}

func init() {
	viper.AddConfigPath("conf")
	viper.SetConfigName(GetEnv())

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
