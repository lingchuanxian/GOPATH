package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
	"strconv"
)

type User struct {
	Id      int  `orm:"auto"`
	Name    string `form:"name" valid:"Required;MaxSize(20);MinSize(6)"`
	Pass    string `form:"pass" valid:"Required;MaxSize(20);MinSize(6)"`
	Email   string `form:"email" valid:"Required;Email"`
	Phone   string `form:"phone" valid:"Required;Mobile"`
	Image   string `form:"image" valid:"MaxSize(50);MinSize(6)"`
	Addr    string `form:"addr" valid:"MaxSize(30)" `
	Regtime string `form:"regtime" `
	Birth   string `form:"birth"`
	Remark  string `valid:"MaxSize(200)" form:"remark"`
}

func init() {
	orm.RegisterModel(new(User))
}

func (u *User) TableName() string {
	return "t_user"
}

func (this *User) ToString() string {
	return fmt.Sprintf("Id:%d\tName:%s", this.Id, this.Name)
}

func GetUserList() (userlist []*User) {
	o := orm.NewOrm()
	qs := o.QueryTable("t_user")
	qs.OrderBy("-Id").All(&userlist)
	/*count, _ = qs.Count()*/
	return userlist
}

func AddUser(u *User) (int64, error) {
	o := orm.NewOrm()
	user := new(User)
	fmt.Println("Test:"+u.Name)
	user.Name = u.Name
	user.Email = u.Email
	user.Phone = u.Phone
	user.Image = u.Image
	user.Addr = u.Addr

	tm2, _ := time.Parse("2006-02-02", u.Birth) //日期字符串转为时间戳
	user.Birth = strconv.FormatInt(tm2.Unix(), 10)

	user.Remark = u.Remark
	user.Regtime = strconv.FormatInt(time.Now().Unix(), 10) //获取当前时间戳
	id, err := o.Insert(user)
	return id, err
}
