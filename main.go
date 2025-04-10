package main

import (
	"flag"
)

func main() {
	target := flag.String("target", "", "the ip address of the target")
	if *target == "" {
		panic("must include target")
	}
}
