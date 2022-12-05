package main
import (
	"fmt"
	"net"
)

func main() {
	name1 := "192.168.1.3"
	addr1 := net.ParseIP(name1)
	fmt.Println(addr1.String())

}