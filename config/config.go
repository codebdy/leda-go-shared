package config

import (
	"github.com/codebdy/entify-core/db"
)

const TABLE_NAME_MAX_LENGTH = 64

var fileCfg Config
var envCfg Config

type Config interface {
	getString(key string) string
	getBool(key string) bool
	getInt(key string) int
}

func GetString(key string) string {
	str := envCfg.getString(key)
	if str == "" {
		str = fileCfg.getString(key)
	}
	return str
}

func GetBool(key string) bool {
	boolValue := envCfg.getBool(key)
	if !boolValue {
		boolValue = fileCfg.getBool(key)
	}
	return boolValue
}

func GetInt(key string) int {
	intValue := envCfg.getInt(key)
	if intValue == 0 {
		intValue = fileCfg.getInt(key)
	}
	return intValue
}

func GetDbConfig() db.DbConfig {
	var cfg db.DbConfig
	cfg.Driver = GetString(DB_DRIVER)
	cfg.Database = GetString(DB_DATABASE)
	cfg.Host = GetString(DB_HOST)
	cfg.Port = GetString(DB_PORT)
	cfg.User = GetString(DB_USER)
	cfg.Password = GetString(DB_PASSWORD)
	if cfg.Driver == "" {
		cfg.Driver = "mysql"
	}
	return cfg
}

func init() {
	fileCfg = newFileConfig()
	envCfg = newEnvConfig()
}
