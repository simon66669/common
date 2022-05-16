package conf

const ModelPath = "./models/"

const ModelReplace = true

const DriverName = "mysql"

type DbConf struct {
	Host   string
	Port   string
	User   string
	Pwd    string
	DbName string
}

var MasterDbConfig = DbConf{
	Host:   "127.0.0.1",
	Port:   "3306",
	User:   "root",
	Pwd:    "123456",
	DbName: "exchange",
}
