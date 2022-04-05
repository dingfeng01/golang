package main

import "fmt"

var version = "1.0"

func main() {
	/* 块注释
	var name string = "01"
	var name01 = "02"
	var name02 string
	*/
	// 块函数
	var (
		name01 string = "01"
		name02        = "02"
		name03 string
	)

	/*
		x := "x"
		y := "y"
	*/
	x, y := "x", "y"
	fmt.Println(name01, name02, name03, x, y)
}
