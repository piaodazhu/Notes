package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// f, _ := os.Create("tmp")
	// f.Write([]byte("file"))
	// f.Close()

	// finfo, _ := os.Stat("tmp")
	// fmt.Println(finfo)
	// fmt.Println(finfo.IsDir(), finfo.ModTime(), finfo.Mode(), finfo.Size())

	// os.Mkdir("./dir", os.ModePerm)
	os.Remove("./dir")
	
	fileName := "tmp"
	// f, _ := os.Open(fileName)
	f, _ := os.OpenFile(fileName, os.O_APPEND|os.O_RDWR, os.ModePerm)
	defer f.Close()
	n, err := f.Write([]byte("testwrite"))
	if err != nil {
		fmt.Println("write", n, "bytes, err=", err)
		return
	}
	buf := make([]byte, 5)
	f.Seek(0,0)
	n, err = f.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("read %d bytes: %s\n", n, buf)
	all, err := io.ReadAll(f)
	fmt.Println(all)
	
	f.Seek(0,0)
	// io.Copy(os.Stdout, f)
	io.CopyBuffer(os.Stdout, f, buf)

	io.WriteString(f, "ok")
}
