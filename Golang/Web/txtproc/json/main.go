package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	file, err := os.Open("test.json")
	if err != nil {
		panic("cannot open target json file")
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		panic("cannot read from the file")
	}
	json.Unmarshal(data, &s)
	fmt.Println(s)

	s.Servers = append(s.Servers, Server{"GuangZhou", "127.0.0.3"})
	b, err := json.MarshalIndent(s, "", "	")
	if err != nil {
		panic("marshal error")
	}
	_, err = os.Stdout.Write(b)
}
