package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)
//请求结构体
type requestBody struct {
	Key string `json:"key"`
	Info string `json:"info"`
	UserId string `json:"userId"`
}

//结果结构体
type reponseBody struct {
	Code int `json:"code"`
	Text string `json:"text"`
	List []string `json:"list"`
	Url string `json:"url"`
}
//请求机器人
func process(inputChan <-chan string,userid string)  {
	for  {
		//从通道接收输入
		input:=<-inputChan
		if input=="EOF" {
			break
		}
		//请求结构体
		reqData:=&requestBody{
			Key: "792bcf45156d488c92e9d11da494b085",
			Info: input,
			UserId: userid,
		}

		//转义为json
		byteData, _ :=json.Marshal(&reqData)
		//请求聊天机器人接口
		req,err:=http.NewRequest("POST","http://www.tuling123.com/openapi/api",bytes.NewReader(byteData))
		req.Header.Set("Content-Type","application/json;charset=UTF-8")
		client:=http.Client{}
		resp,err:=client.Do(req)
		if err!=nil {
			fmt.Println("网络错误！")
		}else {
			//将json解析并且输出到命令行
			body,_:=ioutil.ReadAll(resp.Body)
			var respData reponseBody
			json.Unmarshal(body,&respData)
			fmt.Println("AI："+respData.Text)
			}
			resp.Body.Close()
	}
	
}
func main() {
	var input string
	fmt.Println("输入’EOF‘结束对话：")
	//创建通道
	channel:=make(chan string)
	//main结束时关闭通道
	defer close(channel)
	//启动goroutine运行机器人回答线程
	go process(channel,string(rand.Int63()))
	for  {
		//从命令端读取输入
		fmt.Scanf("%s",&input)
		//将输入放到通道
		channel<-input
		//结束程序
		if input=="EOF" {
			fmt.Println("结束，滚蛋")
			break
		}
	}

}