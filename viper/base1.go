package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"time"
)

func main() {
	c := Config{DB: Database{Host:"default-localhost"}}

	viper.SetConfigFile("config.yml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("VP")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error in reading config %v", err)
	}

	viper.Unmarshal(&c)

	fmt.Printf("environment.env: %s\n", c.Environment.Env)
	fmt.Printf("environment.loglevel: %s\n", c.Environment.LogLevel)
	fmt.Printf("environment.loglevel is warning: %v\n", c.Environment.LogLevel == LogLevelWarn)
	fmt.Printf("environment.loglevel is error: %v\n", c.Environment.LogLevel == LogLevelError)
	fmt.Println()
	fmt.Printf("database.port: %d\n", c.DB.Port)
	fmt.Printf("database.port via GetInt: %d\n", viper.GetInt("database.port"))
	fmt.Printf("database.username: %s\n", c.DB.Password)
	fmt.Printf("database.database: %s\n", c.DB.DatabaseName)
	fmt.Printf("database.host from default value: %s\n", c.DB.Host)
	fmt.Println()
	fmt.Printf("games: %v\n", c.Games)
	fmt.Printf("game 2 via GetString: %s\n", viper.GetString("games.2"))
	fmt.Printf("game 4 via GetString: %s\n", viper.GetString("games.4"))
	fmt.Println()
	fmt.Printf("until: %s\n", c.Until.Format(time.RFC3339))
	fmt.Printf("until via GetTime: %s\n", viper.GetTime("until").Format(time.RFC3339))
	fmt.Println()
	fmt.Printf("since: %s\n", c.Since.String())
	fmt.Printf("since via GetDuration: %s\n", viper.GetDuration("since").String())
	fmt.Println()
	fmt.Printf("all: %v\n", viper.AllSettings())
	fmt.Println()
	fmt.Printf("config variable only exists in env: %s\n", viper.GetString("IDK"))

}

type Config struct {
	Environment Environment
	DB Database `mapstructure:"database"`
	Games []string
	Until time.Time
	Since time.Duration
}

type Environment struct {
	Env string
	LogLevel LogLevel
}

type LogLevel string
const (
	LogLevelWarn = LogLevel("warn")
	LogLevelError = LogLevel("err")
)

type Database struct {
	Port uint
	Username string
	Password string
	DatabaseName string `mapstructure:"database"`
	Host string
}