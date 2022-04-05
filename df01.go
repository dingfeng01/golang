package main

import "fmt"

func main() {
	var name string = "kk"

	fmt.Println(name)

	name = "selecd" //更新变量的值（赋值）

	fmt.Println(name)

	{
		//不是定义 是更新的 这是赋值
		name = "aaaaaaaaaaa"
	}
	fmt.Println(name)
}
