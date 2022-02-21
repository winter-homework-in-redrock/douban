package tool

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //编译驱动
)

var DB *sql.DB

//InitMySQL 测试并连接数据库
func InitMySQL() error {
	sqlConf := "root:87nb32A6@@tcp(127.0.0.1:3306)/dou_ban"
	sqlEngine := "mysql"
	db, err := sql.Open(sqlEngine, sqlConf)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(1000)
	DB = db
	return err
}
