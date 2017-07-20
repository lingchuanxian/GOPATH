package controllers

import (
	m "firbee/models"
)

type UserController struct {
	BaseController
}

func (c *UserController) URLMapping() {
    c.Mapping("GetUsers", c.GetUsers)
    c.Mapping("AddUser", c.AddUser)
	c.Mapping("DelUser", c.DelUser)
	c.Mapping("GetUserById", c.GetUserById)
	c.Mapping("EditUser", c.EditUser)
}

// @Title 获取所有用户
// @Description get all of user
// @Param   query query int false "The pageNumber of user"
// @Success 200 User对象
// @Failure 400 服务器异常
// @Failure 404 接口丢失
// @router /GetUsers [get]
func (this *UserController) GetUsers() {
	nodes := m.GetUserList()
	this.RspData(true, "获取成功",&nodes)
}

// @Title 新增用户
// @Description 新增用户
// @Param   Name formData string true "用户名"
// @Param   Pass formData string true "密码"
// @Param   Email formData string true "邮箱"
// @Param   Phone formData string true "手机号"
// @Param   Image formData string true "头像"
// @Param   Addr formData string true "联系地址"
// @Param   Birth formData string true "出生日期"
// @Success 200 User对象
// @Failure 400 服务器异常
// @Failure 404 接口丢失
// @router /AddUser [post]
func (this *UserController) AddUser() {
	u := m.User{}
	u.Name = this.GetString(":name");
	if err := this.ParseForm(&u); err != nil {
		this.Rsp(false, err.Error())
		return
	}
	id, err := m.AddUser(&u)
	if err == nil && id > 0 {
		this.Rsp(true, "新增成功")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}
}

// @Title 删除用户
// @Description 删除用户
// @Param   Id query int false "The Id of user"
// @Success 200 User对象
// @Failure 400 服务器异常
// @Failure 404 接口丢失
// @router /DelUser [delete]
func (this *UserController) DelUser() {

}

// @Title 根据ID查询用户
// @Description 根据ID查询用户
// @Param   Id query int false "The Id of user"
// @Success 200 User对象
// @Failure 400 服务器异常
// @Failure 404 接口丢失
// @router /GetUserById [get]
func (this *UserController) GetUserById() {

}

// @Title 编辑用户
// @Description 编辑用户
// @Param   Id formData int false "The Id of user"
// @Param   Name formData string true "用户名"
// @Param   Pass formData string true "密码"
// @Param   Email formData string true "邮箱"
// @Param   Phone formData string true "手机号"
// @Param   Image formData string true "头像"
// @Param   Addr formData string true "联系地址"
// @Param   Birth formData string true "出生日期"
// @Success 200 User对象
// @Failure 400 服务器异常
// @Failure 404 接口丢失
// @router /EditUser [put]
func (this *UserController) EditUser() {

}


