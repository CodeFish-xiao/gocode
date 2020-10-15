package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//读原图片
	ff, _ := os.Open("xx.png")
	defer ff.Close()
	sourcebuffer := make([]byte, 500000)
	n, _ := ff.Read(sourcebuffer)
	//base64压缩
	sourcestring := base64.StdEncoding.EncodeToString(sourcebuffer[:n])
	fmt.Print(sourcestring)
	//写入临时文件
	ioutil.WriteFile("a.png.txt", []byte(sourcestring), 0667)
	//读取临时文件
	cc, _ := ioutil.ReadFile("a.png.txt")

	//解压
	dist, _ := base64.StdEncoding.DecodeString(string(cc))
	//写入新文件
	f, _ := os.OpenFile("xx11.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	f.Write(dist)

}
