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
	// p = &a
	// t.Logf("a 的地址是 %v", p)
	// t.Logf("通过指针取值 %v", *p)
	// *p = 20
	// t.Logf("修改指针指向的值 %v", a)
	a := 1
	aPtr := &a
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}
