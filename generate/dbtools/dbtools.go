package dbtools

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/simon66669/common/generate/conf"
	"log"
	"sync"
)

type MysqlConnectPool struct {
}

var mysqlInstance *MysqlConnectPool
var mysqlOnce sync.Once

var db *gorm.DB
var err_db error

func GetMysqlInstance() *MysqlConnectPool {
	mysqlOnce.Do(func() {
		mysqlInstance = &MysqlConnectPool{}
	})
	return mysqlInstance
}

func (m *MysqlConnectPool) InitMysqlPool() (issucc bool) {
	dbConf := conf.MasterDbConfig
	port := dbConf.Port
	db, err_db = gorm.Open("mysql", dbConf.User+":"+dbConf.Pwd+"@tcp("+dbConf.Host+":"+port+")/"+dbConf.DbName+"?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)
	if err_db != nil {
		log.Fatal(err_db)
		return false
	}
	//db.LogMode(true)
	//关闭数据库，db会被多个goroutine共享，可以不调用
	// defer db.Close()
	return true
}

func (m *MysqlConnectPool) GetMysqlPool() *gorm.DB {
	//db.LogMode(true)
	return db
}

func GetMysqlDb() (db *gorm.DB) {
	return GetMysqlInstance().GetMysqlPool()
}
