package main

/*
实现 Pic。它应当返回一个长度为 dy 的切片，其中每个元素是一个长度为 dx，元素类型为 uint8 的切片。当你运行此程序时，它会将每个整数解释为灰度值（好吧，其实是蓝度值）并显示它所对应的图像。
图像的选择由你来定。几个有趣的函数包括 (x+y)/2, x*y, x^y, x*log(y) 和 x%(y+1)。
（提示：需要使用循环来分配 [][]uint8 中的每个 []uint8；请使用 uint8(intValue) 在类型之间转换；你可能会用到 math 包中的函数。）
 */

//uint8 即为byte类型
import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"math/rand"
)
func Pic(dx, dy int) [][]uint8 { //dx元素长度，dy整个slice长度
	s := [][]uint8{}
	ss := []uint8{}
	for i := 0; i < dx; i++ {
		ss = append(ss, uint8(rand.Intn(43)))
	}
	for i := 0; i < dy; i++ {
		s = append(s, ss)
	}
	return s}
func Show(f func(int, int) [][]uint8) {
	const (
		dx = 256
		dy = 256 )
	data := f(dx, dy)
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
		}
	}
	ShowImage(m)
}
func ShowImage(m image.Image) {
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		panic(err)
	}
	enc := base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Println("IMAGE:" + enc)
}
func main() { Show(Pic)}