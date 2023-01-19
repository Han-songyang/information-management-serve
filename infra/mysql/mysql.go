package mysql

import (
	"fmt"
	"github.com/han-songyang/information-management-serve/infra/conf"
	"github.com/han-songyang/information-management-serve/infra/logger"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitMysql 初始化Mysql
func InitMysql() {
	//dsn := "root:hansongyang123@tcp(127.0.0.1:3306)/MANAGER?charset=utf8mb4&parseTime=True&loc=Local"
	username := conf.Config.Section("mysql").Key("username").String()
	password := conf.Config.Section("mysql").Key("password").String()
	host := conf.Config.Section("mysql").Key("host").String()
	port := conf.Config.Section("mysql").Key("port").String()
	database := conf.Config.Section("mysql").Key("database").String()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.WarnLog("init", "InitMysql", logrus.Fields{
			"err": err,
			"msg": "数据库初始化失败",
		})
	}

	fmt.Println("mysql init success")
}
