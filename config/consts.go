package config

const (
	DB_DRIVER   = "driver"
	DB_USER     = "username"
	DB_PASSWORD = "password"
	DB_HOST     = "host"
	DB_PORT     = "port"
	DB_DATABASE = "database"

	// 模型迁移选项，service指model.json，app指默认app：defaultApp.json
	// 取值范围：
	// sync 强制同步
	// install 未安装时安装
	MIGRATION = "migration"
)

const (
	MIGRATION_SYNC    = "sync"
	MIGRATION_INSTALL = "install"
)
