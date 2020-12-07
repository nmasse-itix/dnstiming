package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s hostname\n", os.Args[0])
		os.Exit(1)
	}

	start := time.Now()
	addrs, err := net.LookupHost(os.Args[1])
	end := time.Now()
	elapsed := end.Sub(start)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	for _, addr := range addrs {
		fmt.Println(addr)
	}

	fmt.Println()
	fmt.Printf("Elapsed: %d ms\n", elapsed.Milliseconds())
}
