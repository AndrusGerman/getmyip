package main

import (
	"fmt"

	"github.com/AndrusGerman/getmyip"
)

func main() {
	ip := getmyip.GetLocalIP()
	fmt.Println(ip)
}
