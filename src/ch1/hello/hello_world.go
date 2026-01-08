/*
 * @Descripttion:
 * @version:
 * @Author: dengweiyi
 * @Date: 2026-01-08 16:35:43
 * @LastEditors: dengweiyi
 * @LastEditTime: 2026-01-08 16:48:19
 */
package main // 包，声明一个可执行程序

import (
	"fmt" // 导入格式化包
	"os"
)

// 主函数，程序执行的入口, 打印 "hello, world"，到控制台
// 运行命令: go run src/ch1/hello/hello_world.go

/*
与其他主要编程语⾔言的差异
• Go 中 main 函数不不⽀支持任何返回值
• 通过 os.Exit 来返回状态
*/
func main() {
	if len(os.Args) > 1 {
		// 如果有命令行参数，打印参数并退出: go run .\hello_world.go go
		// Arguments: [C:\Users\76008\AppData\Local\Temp\go-build3454493215\b001\exe\hello_world.exe go]
		// fmt.Println("Arguments:", os.Args)
		// Arguments: [go]
		fmt.Println("Arguments:", os.Args[1:])
		os.Exit(1)
	}
	// fmt.Println("hello, world")
	// 退出程序
	os.Exit(0)
}
