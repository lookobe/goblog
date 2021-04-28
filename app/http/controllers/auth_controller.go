package controllers

import (
	"fmt"
	"goblog/app/models/user"
	"goblog/app/requests"
	"goblog/pkg/auth"
	"goblog/pkg/flash"
	"goblog/pkg/view"
	"net/http"
)

// AuthController 处理静态页面
type AuthController struct {

}

// Register 注册页面
func (*AuthController) Register(w http.ResponseWriter,r *http.Request) {
	view.RenderSimple(w, view.D{},"auth.register")
}

// DoRegister 处理注册逻辑
func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	// 1. 初始化数据
	_user := user.User{
		Name: 			 r.PostFormValue("name"),
		Email: 			 r.PostFormValue("email"),
		Password: 		 r.PostFormValue("password"),
		PasswordConfirm: r.PostFormValue("password_confirm"),
	}

	// 2. 表单规则
	errs := requests.ValidateRegistrationForm(_user)

	if len(errs) >0 {
		// 3. 表单不通过 -- 重新显示表单
		// data,_ :=json.MarshalIndent(errs,""," ")
		// fmt.Fprint(w,string(data))

		view.RenderSimple(w,view.D{
			"Errors": errs,
			"User": _user,
		},"auth.register")
	}else {
		// 4. 验证成功，创建数据
		_user.Create()
		fmt.Println(_user.ID)
		if _user.ID > 0 || _user.ID == 0 {
			// 登陆用户并跳转到首页
			flash.Success("恭喜您注册成功！")
			auth.Login(_user)
			http.Redirect(w,r,"/", http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("===这额外地方撒上")
			fmt.Fprint(w,"注册失败，请联系管理员")
		}
	}
}

// Login 显示登陆表单
func (*AuthController) Login(w http.ResponseWriter,r *http.Request) {

	// session.Put("uid","1")
	// fmt.Fprint(w,session.Get("uid"))

	// session.Forget("uid")

	// session.Flush()
	view.RenderSimple(w, view.D{},"auth.login")
}

// DoLogin 处理登录表单提交
func (*AuthController) DoLogin(w http.ResponseWriter, r *http.Request) {
	// 1. 初始化表单数据
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	// 2. 尝试登陆
	if err := auth.Attempt(email,password); err == nil {
		// 登陆成功
		flash.Success("欢迎回来！")
		http.Redirect(w,r,"/", http.StatusFound)
	} else {
		// 3. 失败，显示错误提示
		view.RenderSimple(w,view.D{
			"Error": err.Error(),
			"Email": email,
			"Password": password,
		},"auth.login")
	}
}

// Logout 退出登陆
func (*AuthController) Logout(w http.ResponseWriter,r *http.Request) {
	auth.Logout()
	flash.Success("您已退出登陆")
	http.Redirect(w,r,"/",http.StatusFound)
}