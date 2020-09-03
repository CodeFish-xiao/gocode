package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)



func login(w http.ResponseWriter,r *http.Request)   {
	fmt.Printf("method:",r.Method)//获取请求的方法
	fmt.Println("进来了")
	if r.Method =="GET"{
		t,_:=template.ParseFiles("login.tpl")
		log.Println(t.Execute(w,nil))
	}else {
		//执行的是登录数据进行登录逻辑判断
		_=r.ParseForm()
		fmt.Println("username:",r.Form["username"])
		fmt.Println("password:",r.Form["password"])
		if pwd:=r.Form.Get("username"); pwd=="123456"{//验证密码
			fmt.Fprintf(w,"欢迎登录,%s!",r.Form.Get("username"))//输出到客户端信息
		}else {
			fmt.Fprintf(w,"密码啊，小扑街")//输出到客户端信息
		}
	}

}
func sayHello(w http.ResponseWriter,r *http.Request)  {
	_=r.ParseForm()//3.解析参数，默认中不会解析
	//解析url传递的参数，对于post则解析相应包的主体（Request body）
	fmt.Println(r.Form)//4：输出到服务器的打印信息
	fmt.Println("Path",r.URL.Path)//输出路径
	fmt.Println("Host",r.Host)//输出端口
	for k,v:=range r.Form{
		fmt.Println("key:",k)
		fmt.Println("val:",strings.Join(v,""))
	}
	_,_=fmt.Fprintf(w ,"HELLO WEB,%s!",r.Form.Get("name"))//5.写入到w的是输出到客户端的内容
}
func main() {
	http.HandleFunc("/",sayHello)//1.设置访问路由
	http.HandleFunc("/login",login)//设置login的访问路由为/login
	err:=http.ListenAndServe(":8080",nil)//设置监听接口
	if err!=nil{
		log.Fatal("ListenAndServe",err)
	}
}