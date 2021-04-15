package main

import (
	"database/sql"
	"net/http"
	"strings"

	"goblog/pkg/database"
	"goblog/bootstrap"

	"github.com/gorilla/mux"
)

var router *mux.Router
var db *sql.DB

// type Object struct {

// }
// // object的方法
// func (obj *Object) method() {

// }
// // 只是一个函数
// func function() {

// }

func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 设置标头
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// 2. 继续处理请求
		next.ServeHTTP(w, r)
	})
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 除首页以外，移除所有请求路径后面的斜杆
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		// 2. 将请求传递下去
		next.ServeHTTP(w, r)
	})
}

func main() {
	database.Initialize()
	db = database.DB

	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	// o := new(Object)
	// o.method()
	// function()

	// 中间件：强制内容类型为HTML
	router.Use(forceHTMLMiddleware)



	http.ListenAndServe(":3000", removeTrailingSlash(router))
}
