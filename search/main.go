package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func main() {

	search := "https://www.google.es/search?q="
	flag.Parse()
	for _, item := range flag.Args() {
		search += item + "+"
	}
	fmt.Println(search)
	exec.Command("firefox", search).Run()
}
