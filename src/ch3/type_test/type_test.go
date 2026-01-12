package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	// a := 10
	// var p *int
	// &符号被称为‌取地址运算符‌或‌取址符‌，其主要作用是获取变量的‌内存地址‌，并返回一个指向该变量的‌指针‌
	// p = &a
	// t.Logf("a 的地址是 %v", p)
	// t.Logf("通过指针取值 %v", *p)
	// *p = 20
	// t.Logf("修改指针指向的值 %v", a)
	a := 1
	aPtr := &a
	// aPtr = aPtr + 1 // error 指针不能进行算术运算
	t.Log(a, aPtr)
	// int *int: 指向整型的指针
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	s := "hello, 世界"
	t.Logf("len(s) = %d", len(s))
}
