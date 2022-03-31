package main

import (
	"flag"
	"os"
	"os/exec"
	"strings"
)

func main() {
	flag.Parse()
	str := flag.Arg(0)

	if strings.Contains(str, "/") {
		slice := strings.Split(str, "/")
		createMultiple(slice)
	} else {
		createSimple(str)
	}

}
func createSimple(str string) {
	exec.Command("mkdir", str).Run()
	os.Create("./" + str + "/main.go")
	os.Chdir(str)
	exec.Command("code", ".").Run()
}
func createMultiple(slice []string) {
	var root string
	var pack string

	root = slice[0]
	exec.Command("mkdir", root).Run()
	os.Create("./" + root + "/main.go")
	if len(slice) > 1 {
		pack = slice[1]
		exec.Command("mkdir", "./"+root+"/"+pack).Run()
	}
	if len(slice) == 3 {
		if strings.Contains(slice[2], "+") {
			files := strings.Split(slice[2], "+")
			for _, file := range files {
				os.Create("./" + root + "/" + pack + "/" + file + ".go")
			}

		} else {
			os.Create("./" + root + "/" + pack + "/" + slice[2] + ".go")
		}
	}
	os.Chdir(root)
	exec.Command("./"+root+"/code", ".").Run()
}
