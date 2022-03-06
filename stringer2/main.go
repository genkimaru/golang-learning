package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ipaddr IPAddr) String() string {
	str := "["
	for ind, v := range ipaddr {
		str += strconv.Itoa(int(v))
		if ind != 3 {
			str += "."
		}
	}
	str += "]"
	return str
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	// var ips IPAddr = [4]byte{12, 23, 34, 45}
	// fmt.Println(ips.String())

}
