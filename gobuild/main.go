package main

import (
	"flag"
	"fmt"
	"go/build"
)

func main() {
	flag.Parse()

	for _, arg := range flag.Args() {
		pkg, err := build.Default.Import(arg, arg, build.FindOnly)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(pkg)
		}
	}

	//Print global variables
	fmt.Println(build.Default.GOARCH, build.Default.GOPATH)
}
