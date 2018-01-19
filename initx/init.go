package initx

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func InitSysX()  {
	//读取app.conf的配置
	dbType := beego.AppConfig.String("db_type")
	dbName := beego.AppConfig.String("db_name")
	dbUrl := beego.AppConfig.String("db_url")
	dbUserName := beego.AppConfig.String("db_username")
	dbPassWord := beego.AppConfig.String("db_password")
	dbPort := beego.AppConfig.String("db_port")
	dbAlias := beego.AppConfig.String("db_alias")

	orm.RegisterDataBase(dbAlias,dbType,dbUserName+":"+dbPassWord+"@tcp("+dbUrl+":"+dbPort+")/"+dbName+"?charset=UTF8",30)
	orm.SetMaxOpenConns(dbAlias,10)
	//设置时间
	orm.DefaultTimeLoc = time.Local

	//如果是开发模式，则显示命令信息
	isDev := (beego.AppConfig.String("runmode") == "dev")
	if isDev {
		orm.Debug = isDev
	}
}