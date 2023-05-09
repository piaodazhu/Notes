package main

import (
	"fmt"
	"io/fs"
	"os"
	"runtime"
	"sync"
	"time"
)

func visitFile(dir string, f fs.DirEntry) {
	if f.IsDir() {
		path := dir + "/" + f.Name()
		dir, _ := os.ReadDir(path)
		for _, file := range dir {
			visitFile(path, file)
		}
	} else {
		fmt.Printf("%s |  %s\n", dir, f.Name())
	}
}

func main() {
	fmt.Println(runtime.GOROOT(), runtime.GOARCH, runtime.GOOS, runtime.NumCPU())
	runtime.Gosched()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer fmt.Println("exit..")
		fmt.Println("sleeping")
		time.Sleep(time.Second * 1)
		wg.Done()
		runtime.Goexit()
		fmt.Println("foobar")
	}()
	wg.Wait()
	dirname := "."
	dir, err := os.ReadDir(dirname)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, f := range dir {
		visitFile(dirname, f)
	}
}
