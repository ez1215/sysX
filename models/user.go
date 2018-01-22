package models

import (
	"github.com/astaxie/beego/orm"
	"sysX/utils"
	"log"
	"time"
	"fmt"
	"strconv"
)

func (u *User) TableName() string {
	return UserTableName()
}

type User struct {
	Id int
	NickName string
	LoginName string
	Password string
	Status int
	Phone string
	CreateTime time.Time
}

var salt_pre string = "1VGjNA0!"
var salt_suf string = "BGd2kR#O"

func Login(username string,password string) (*User,error) {
	enpwd := utils.String2md5(salt_pre+password+salt_suf)
	user := User{}
	error := orm.NewOrm().QueryTable("t_user").Filter("login_name",username).Filter("password",enpwd).One(&user)
	if error != nil{
		log.Println(error)
		return nil ,error
	}
	return &user,nil
}

//新增，第一个参数返回0表明新增失败，否则返回新增后的id
func Add(user *User) (int64,error) {
	if(utils.IsBlank(user.Password) || utils.IsBlank(user.LoginName)){
		return 0,nil
	}
	enpwd := utils.String2md5(salt_pre+user.Password+salt_suf)
	user.Password = enpwd
	//user.CreateTime = time.Now()
	user.Status = 0
	o := orm.NewOrm()
	id,err := o.Insert(user)
	if err != nil{
		log.Println(err)
		return 0,err
	}
	return id,nil
}

func ListByPage(page utils.Page) utils.Page  {
	//校验page
	query := orm.NewOrm().QueryTable(UserTableName())
	query.Limit(0,10)
	fmt.Print("数量:")
	fmt.Println(query.Count())
	return page
}

func ListByUser() []*User  {
	data := make([]*User, 0)
	orm.NewOrm().QueryTable(UserTableName()).All(&data)
	return  data
}

func DeleteUser(id int) bool {
	o := orm.NewOrm()
	if num ,err := o.Delete(&User{Id:id});err==nil{
		log.Println("删除"+strconv.FormatInt(num,10)+"行数据")
		return true
	}else {
		return false
	}
}