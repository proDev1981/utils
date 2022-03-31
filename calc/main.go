package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	if strings.Contains(flag.Arg(0), "+") {
		param := strings.Split(flag.Arg(0), "+")
		sum(param...)
	}
	if strings.Contains(flag.Arg(0), "*") {
		params := strings.Split(flag.Arg(0), "*")
		plas(params...)
	}
	if strings.Contains(flag.Arg(0), "-") {
		params := strings.Split(flag.Arg(0), "-")
		rest(params...)
	}
	if strings.Contains(flag.Arg(0), "/") {
		params := strings.Split(flag.Arg(0), "/")
		div(params...)
	}
}
func sum(params ...string) {
	var res, ints int
	for _, item := range params {
		ints, _ = strconv.Atoi(item)
		res += ints
	}
	fmt.Println(res)
}
func rest(params ...string) {
	var res, ints int
	for _, item := range params {
		ints, _ = strconv.Atoi(item)
		if res > ints {
			res = res - ints
		} else {
			res = ints - res
		}
	}
	fmt.Println(res)
}
func plas(params ...string) {
	var res, ints int
	res = 1
	for _, item := range params {
		ints, _ = strconv.Atoi(item)
		res *= ints
	}
	fmt.Println(res)
}
func div(params ...string) {
	var res, ints float32
	res = 1
	for _, item := range params {
		num, _ := strconv.Atoi(item)
		ints = float32(num)
		if res == 1 {
			res = ints
		} else {
			res = res / ints
		}
	}
	fmt.Println(res)
}
