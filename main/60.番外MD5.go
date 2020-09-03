package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	h := md5.New()
	h.Write([]byte("admin")) // 需要加密的字符串为 admin
	cipherStr := h.Sum(nil)
	//fmt.Println(cipherStr)
	md5str:=hex.EncodeToString(cipherStr)
	fmt.Println(md5str)
	fmt.Printf("%s\n",Md5Encrypt("admin")) // 输出加密结果
}
func Md5Encrypt(data string) string {
	md5Ctx := md5.New()                            //md5 init
	md5Ctx.Write([]byte(data))                     //md5 updata
	cipherStr := md5Ctx.Sum(nil)                   //md5 final
	encryptedData := hex.EncodeToString(cipherStr) //hex_digest
	return encryptedData
}