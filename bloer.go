package main

import "fmt"

//包级别
var packagesvar string = "package.var"

func main() {
	//函数级别
	var funcvar string = "func.var"
	{
		//块级别，函数
		var block string = "block.var"
		{
			//子块级别。函数
			var install string = "安装好了"
			fmt.Println("3", packagesvar, funcvar, block, install)
		}
		fmt.Println("2", packagesvar, funcvar, block)
	}
	fmt.Println("1", packagesvar, funcvar)
}
