package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
	"time"
	"github.com/judgeu/go-config-comparision/common"
)

func main() {
	c := Config{DB: Database{Host: "default-localhost"}}

	viper.SetConfigFile("config.yml")
	viper.SetEnvPrefix("VP")
	viper.SetDefault("wurst", "brot")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // need this for sub-keys / nested keys via env, e.g. VP_DATABASE_PORT
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error in reading config %v", err)
	}

	viper.Unmarshal(&c)

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

	/*
	fmt.Printf("| database.port via GetInt: %d|", viper.GetInt("database.port"))
	fmt.Printf("| database.port via Get with .: %v|", viper.Get("database.port"))
	fmt.Printf("| database.port via Get with _: %v|", viper.Get("database_port"))
	fmt.Printf("| games via Get: %s|", viper.Get("games"))
	fmt.Printf("| game 2 via GetString: %s|", viper.GetString("games.2"))
	fmt.Printf("| game 4 via GetString: %s|", viper.GetString("games.4"))
	fmt.Printf("| until via GetTime: %s|", viper.GetTime("until").Format(time.RFC3339))
	fmt.Printf("| since via GetDuration: %s|", viper.GetDuration("since").String())
	fmt.Printf("| config variable only exists in env: %s|", viper.GetString("IDK"))
	fmt.Printf("| read default: %s|", viper.GetString("wurst"))
	*/

}

type Config struct {
	Environment Environment
	DB          Database `mapstructure:"database"`
	Games       []string
	Until       time.Time
	Since       time.Duration
}

type Environment struct {
	Env      string
	LogLevel common.LogLevel
}

type Database struct {
	Port         uint
	Username     string
	Password     string
	DatabaseName string `mapstructure:"database"`
	Host         string
}
