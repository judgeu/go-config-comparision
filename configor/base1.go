package main

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/judgeu/go-config-comparision/common"
	"time"
)

func main() {
	var c Config

	configor.Load(&c, "config.yml")

	results := map[string]interface{}{
		common.KEY_environment_env:                  c.Environment.Env,
		common.KEY_environment_loglevel:             c.Environment.LogLevel,
		common.KEY_environment_loglevel_is_warning:  c.Environment.LogLevel == common.LogLevelWarn,
		common.KEY_environment_loglevel_is_error:    c.Environment.LogLevel == common.LogLevelError,
		common.KEY_database_port:                    c.DB.Port,
		common.KEY_database_username:                c.DB.Password,
		common.KEY_database_database:                c.DB.DatabaseName,
		common.KEY_database_host_from_default_value: c.DB.Host,
		common.KEY_games:                            c.Games,
		common.KEY_until:                            c.Until.Format(time.RFC3339),
		common.KEY_since:                            c.Since.String(),
	}

	for _, k := range common.KEYS {
		fmt.Printf(" %v | ", results[k])
	}

}

type Config struct {
	Environment Environment
	DB Database `yaml:"database"`
	Games []string
	Until time.Time
	Since time.Duration
}

type Environment struct {
	Env string
	LogLevel common.LogLevel
}

type Database struct {
	Port uint
	Username string
	Password string
	DatabaseName string `yaml:"database"`
	Host string `default:"default-localhost"`
}