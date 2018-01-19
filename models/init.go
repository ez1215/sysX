package models

import "github.com/astaxie/beego/orm"

func init()  {
	orm.RegisterModel(new(User))
}

func UserTableName()  string {
	return "t_user"
}


