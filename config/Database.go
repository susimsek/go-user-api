package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"user-api/util"
)

var DB *gorm.DB

var MysqlHost = util.GetEnvStr("MYSQL_HOST", "localhost")
var MysqlPort = util.GetEnvInt("MYSQL_PORT", 3306)
var MysqlUser = util.GetEnvStr("MYSQL_USER", "user")
var MysqlPass = util.GetEnvStr("MYSQL_PASS", "root")
var MysqlDb = util.GetEnvStr("MYSQL_DB", "user")

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	fmt.Println()
	dbConfig := DBConfig{
		Host:     MysqlHost,
		Port:     MysqlPort,
		User:     MysqlUser,
		Password: MysqlPass,
		DBName:   MysqlDb,
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
