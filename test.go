package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Port string

func (p Port) Port() string {
	return strings.Split(string(p), "/")[0]
}

// Proto returns the name of the protocol.
func (p Port) Proto() string {
	parts := strings.Split(string(p), "/")
	if len(parts) == 1 {
		return "tcp"
	}
	return parts[1]
}

func main() {
	aa := strconv.Itoa(80) + "/" + strings.ToLower("tcp")
	port := Port(aa)

	fmt.Println(port)
}
