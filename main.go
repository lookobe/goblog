package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"strings"
)
var router = mux.NewRouter()
func homeHandler(w http.ResponseWriter,r *http.Request) {
	fmt.Fprint(w,"<h1>Hello, 欢迎来到 goblog！</h1>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w,"此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
		"<a href=\"mailto:summer@example.com\">summer@example.com</a>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w,"<h1>请求页面未找到 :(</h1><p>如有疑惑，请联系我们。</p>")
}

func articlesShowHandle(w http.ResponseWriter,r *http.Request) {
	vars := mux.Vars(r)
	id:= vars["id"]
	fmt.Fprint(w,"文章 ID:"+id)
}

func articlesIndexHandler(w http.ResponseWriter,r *http.Request) {
	fmt.Fprint(w,"访问文章列表")
}

func articlesStoreHandler(w http.ResponseWriter,r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		// 解析错误，这里应该有错误处理
		fmt.Fprint(w,"请提供正确的数据")
		return
	}
	title := r.PostForm.Get("title")
	fmt.Fprintf(w,"POST PostForm: %v <br>",r.PostForm)
	fmt.Fprintf(w,"POST Form:%v <br>", r.Form)
	fmt.Fprintf(w,"title 的值为：%v<br>",title)

	fmt.Fprintln(w,"+++++++++++++++++++++++++++<br>")
	fmt.Fprintf(w,"r.Form中的title的值为：%v <br>",r.FormValue("title"))
	fmt.Fprintf(w,"r.PostForm中的title的值为：%v <br>",r.PostFormValue("title"))
	fmt.Fprintf(w,"r.Form中的test的值为：%v <br>",r.FormValue("test"))
	fmt.Fprintf(w,"r.PostForm中的test值为：%v<br>",r.PostFormValue("test"))

}

func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		// 1. 设置标头
		w.Header().Set("Content-Type","text/html; charset=utf-8")
		// 2. 继续处理请求
		next.ServeHTTP(w,r)
	})
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter,r *http.Request) {
		// 1. 除首页以外，移除所有请求路径后面的斜杆
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path,"/")
		}
		// 2. 将请求传递下去
		next.ServeHTTP(w,r)
	})
}

func articlesCreateHandler(w http.ResponseWriter, r *http.Request) {
	html := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<title>创建文章 —— 我的技术博客</title>
	</head>
	<body>
		<form action="%s?test=data" method="post">
		<p><input type="text" name="title"></p>
        <p><textarea name="body" cols="30" rows="10"></textarea></p>
        <p><button type="submit">提交</button></p>
	</body>
	</html>
	`
	storeURL,_ := router.Get("articles.store").URL()
	fmt.Fprintf(w,html,storeURL)
}

func main()  {


	router.HandleFunc("/",homeHandler).Methods("GET").Name("home")
	router.HandleFunc("/about",aboutHandler).Methods("GET").Name("about")

	router.HandleFunc("/articles/{id:[0-9]+}",articlesShowHandle).Methods("GET").Name("articles.show")
	router.HandleFunc("/articles",articlesIndexHandler).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles",articlesStoreHandler).Methods("POST").Name("articles.store")

	router.HandleFunc("/articles/create",articlesCreateHandler).Methods("GET").Name("articles.create")

	// 自定义404页面
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	// 中间件：强制内容类型为HTML
	router.Use(forceHTMLMiddleware)


	http.ListenAndServe(":3000",removeTrailingSlash(router))
}