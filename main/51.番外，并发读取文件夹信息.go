package main
import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)
/*
问题描述：
实现一个程序计算并打印输入的目录下所有文件的总数和总大小(以GB计算)。完成之后你将熟悉select、WaitGroup、ioutil的用法。

要点：
并发读取文件(夹)信息。
限制开启的goroutines的最大数量。
运行时每隔500ms打印当前已经统计的文件数和总大小（使用命令行参数指定此功能是否启用）。

拓展：
在执行中在有外部输入时退出程序。
 */
var verbose = flag.Bool("v", false, "show verbose progress messages")

var sema = make(chan struct{}, 50)
var done = make(chan struct{})

func dirents(dir string) []os.FileInfo{
	select {
	case sema <- struct{}{}:
	case <- done:
		return nil

	}

	defer func() {<- sema}()

	entries, err := ioutil.ReadDir(dir)

	if err != nil{
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}

	return entries

}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64){
	defer n.Done()

	if cancelled(){
		return
	}

	for _, entry := range dirents(dir){
		if entry.IsDir(){
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func printDiskUsage(nfiles, nbytes int64){
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes/1e9))
}

func cancelled() bool{
	select {
	case <- done:
		return true
	default:
		return false
	}
}

func main() {

	flag.Parse()

	roots := flag.Args()

	var tick <-chan time.Time

	if *verbose{
		tick = time.Tick(500 * time.Millisecond)
	}


	if len(roots) == 0{
		roots = []string{"."}
	}


	fileSizes := make(chan int64)
	var nfiles, nbytes int64

	var n sync.WaitGroup

	for _, root := range roots{
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

loop:
	for {
		select {
		case <-done:
			for range fileSizes{
				//
			}
		case size, ok := <- fileSizes:
			if !ok {
				break loop
			}

			nfiles++
			nbytes += size

		case <- tick:
			printDiskUsage(nfiles, nbytes)
		}
	}

	printDiskUsage(nfiles, nbytes)
}
