package config

import (
	"github.com/airdb/sailor/enum"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"path/filepath"
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

func GetEnv() (env string) {
	env = os.Getenv("ENV")
	if env == "" {
		env = enum.FromEnv(enum.EnvDev)
	}
	env = strings.ToLower(env)
	return env
}

func GetPort() (port string) {
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return
}

func GetDatabases() (databases map[string]*Database) {
	Init()
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

func Init() {
	binPath := filepath.Dir(os.Args[0])

	workDir, err := os.Getwd()
	if !strings.Contains(binPath, "go-build") {
		workDir, err = filepath.Abs(binPath)
	}

	if err != nil {
		log.Fatal(err)
	}

	viper.AddConfigPath(path.Join(workDir, "conf"))
	viper.AddConfigPath(path.Join(workDir, "config"))
	viper.AddConfigPath(path.Join(workDir, "configs"))
	viper.SetConfigName(GetEnv())

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
