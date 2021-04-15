package main

import (
	"database/sql"
	"net/http"
	"goblog/app/http/middlewares"
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


func main() {
	database.Initialize()
	db = database.DB

	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	// o := new(Object)
	// o.method()
	// function()



	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
