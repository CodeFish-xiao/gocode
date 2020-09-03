package main
//不同二叉树的叶节点上可以保存相同的值序列。例如，以下两个二叉树都保存了序列 1，1，2，3，5，8，13。
//在大多数语言中，检查两个二叉树是否保存了相同序列的函数都相当复杂。 我们将使用 Go 的并发和信道来编写一个简单的解法。
//本例使用了 tree 包，它定义了类型：
//type Tree struct {
//    Left  *Tree
//    Value int
//    Right *Tree
//}
//点击下一页继续。



import "golang.org/x/tour/tree"
import "fmt"

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int){
	if t.Left != nil{
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right !=nil{
		Walk(t.Right, ch)
	}

}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	total := 10
	ch1, ch2 := make(chan int, total), make(chan int, total)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i:=0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(2)
	fmt.Println(Same(t1, t1))
	fmt.Println(Same(t1,t2))
}