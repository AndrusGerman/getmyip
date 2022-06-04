package main

import (
	"fmt"

	"github.com/AndrusGerman/getmyip"
)

func main() {
	ip := getmyip.GetPublicIP()
	fmt.Println(ip)
}
