package main

import (
	"fmt"
	"os"

	"go.jetpack.io/typeid"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: typeid [<type_prefix>]")
		os.Exit(1)
	}

	prefix := os.Args[1]
	tid := typeid.Must(typeid.New(prefix))
	fmt.Println(tid)
}
