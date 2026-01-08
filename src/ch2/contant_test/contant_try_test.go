package constant_test

import "testing"

// @Descripttion: 测试常量和 iota 的使用
// 连续常量的简洁写法，等同于
/*
const (
	Monday   = 1
	Tuesday  = 2
	Wensday  = 3
	Thursday = 4
	Friday   = 5
	Saturday = 6
	Sunday   = 7
)
*/
const (
	Monday = iota + 1
	Tuesday
	Wensday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	Readable   = 1 << iota // 1 << 0 which is 00000001
	Writable               // 1 << 1 which is 00000010
	Executable             // 1 << 2 which is 00000100
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wensday, Thursday, Friday, Saturday, Sunday)
}

func TestConstantTry1(t *testing.T) {
	a := 1 //0001
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
