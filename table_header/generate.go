package main

import (
	"fmt"
	"github.com/judgeu/go-config-comparision/common"
	"strings"
)

func main() {
	fmt.Print("| |")
	for _, k := range common.KEYS {
		fmt.Printf(" %s | ", k)
	}

	fmt.Println()

	fmt.Println("|" + strings.Repeat("---|", len(common.KEYS) + 1))
}
