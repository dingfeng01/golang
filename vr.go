package main

import (
	"fmt"
)

func main() {
	// 以上三个都可以正常使用的
	//这个定义变量 是比较全的
	var name1 string = "kk"
	// 可以定义 变量没有值 默认值是零值 ""
	var zestring string
	//这个定义是没有类型的
	var de01 = "kk"
	//短声明 通过多应的值类型推到的变量的类型
	// （必须在函数内包含函数内子块使用，不能在包级别后用）
	short01 := "kk01"
	fmt.Println(name1, zestring, de01, short01)
}
