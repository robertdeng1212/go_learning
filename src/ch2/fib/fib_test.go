/*
 * @Descripttion:
 * @version:
 * @Author: dengweiyi
 * @Date: 2026-01-08 16:58:49
 * @LastEditors: dengweiyi
 * @LastEditTime: 2026-01-08 17:42:17
 */
package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	var a int = 1
	var b int = 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + b
	}
}

func TestExChange(t *testing.T) {
	a := 1
	b := 2
	// tmp := a
	// a = b
	// b = tmp

	// Go 语言支持多重赋值，可以用来交换变量的值
	a, b = b, a
	t.Log(a, b)
}
