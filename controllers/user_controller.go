package controllers

import (
	"github.com/astaxie/beego"
	"log"
	"sysX/utils"
	user "sysX/models"
	"os"
	"net"
	"runtime"
	"path/filepath"
	"strings"
	"fmt"
	json "encoding/json"
	"time"
	"strconv"
)

type UserController struct {
	beego.Controller
}


func (c *UserController) Login()  {
	var username = c.GetString("username")
	var password = c.GetString("password")
	if(utils.IsBlank(username) || utils.IsBlank(password)){
		c.TplName = "login.tpl"
		c.Data["result"] = "用户名或密码错误"
		c.Data["status"] = true
		return
	}
	log.Println("用户登录，用户名:"+username+"  密码:"+password)

	user,err :=user.Login(username,password)
	if err != nil{
		log.Println(err)
		return
	}
	if user != nil{
		c.SetSession("user",user)
		log.Println("登录成功")
		c.Ctx.Redirect(302,"/index")
	}
}

func (c *UserController) Index()  {
	u := c.GetSession("user")
	if u == nil{
		c.TplName = "login.tpl"
		return
	}
	c.TplName = "index.tpl"
	log.Println(u)
	c.Data["user"] = u
}

func (c *UserController) Welcom() {
	u := c.GetSession("user")
	if u == nil{
		c.TplName = "login.tpl"
		return
	}
	c.TplName = "welcome.tpl"
}

func (c *UserController)LoginOut()  {
	u := c.GetSession("user")
	log.Println("注销用户",u)
	if u == nil{
		c.TplName = "login.tpl"
		return
	}
	c.DelSession("user")
	c.TplName = "login.tpl"
}

func (c *UserController)SysInfo()  {
	info := make(map[string]string)
	//计算机名称
	host,_ := os.Hostname()
	info["pcname"] = host

	//ip
	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				info["ip"] = ipnet.IP.String()
				break
			}
		}
	}
	//域名
	//port
	info["port"] = beego.AppConfig.String("httpport")

	//版本
	info["version"] = runtime.Version()

	//操作系统
	info["os"] = runtime.GOOS

	//当前时间
	curent_time := time.Now().Format("2006-01-02 15:04:05")
	info["time"] = curent_time

	//路径
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	info["path"] = strings.Replace(dir, "\\", "/", -1)

	//CPU数
	info["cpunum"] =  fmt.Sprintf("%d", runtime.NumCPU())

	//线程数
	info["threadnum"] =  fmt.Sprintf("%d", runtime.NumGoroutine())
	json ,_:= json.Marshal(info)
	c.Ctx.WriteString(string(json))
}

func (c *UserController) Register()  {
	var username = c.GetString("username")
	var password = c.GetString("password")
	var nickname = c.GetString("nickname")
	if(utils.IsBlank(username) || utils.IsBlank(password)){
		c.Data["json"] = "用户名或密码不能为空"
		c.ServeJSON()
	}
	u := user.User{LoginName:username,Password:password,NickName:nickname}
	id ,err := user.Add(&u)
	if(err != nil){
		log.Println("新增用户失败")
		c.Data["json"] = "新增用户失败"
		c.ServeJSON()
		return
	}
	log.Println("新增用户成功:id={}",id)
	c.Data["json"] = 1
	c.ServeJSON()
}

func (c *UserController) List()  {
	//user.ListByPage(utils.Page{})
	data := user.ListByUser()
	c.Data["data"] = data
	c.TplName="sysx/user-list.tpl"
}

func (c *UserController) AddPage()  {
	c.TplName="sysx/user-add.tpl"
}

func (c *UserController) DeleteUser()  {
	id,err := strconv.Atoi(c.GetString("id"))
	if(err != nil){
		log.Println(err)
		c.Data["json"] = 0
		c.ServeJSON()
		return
	}
	if user.DeleteUser(id){
		c.Data["json"] = 1
		c.ServeJSON()
		return
	}
	c.Data["json"] = 0
	c.ServeJSON()
}