/*
 * @Descripttion:
 * @version:
 * @Author: dengweiyi
 * @Date: 2026-01-12 10:30:16
 * @LastEditors: dengweiyi
 * @LastEditTime: 2026-01-12 11:27:04
 */
package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 多返回值函数
func returnMulivValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

// 函数作为参数
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Printf("time spent: %v\n", time.Since(start).Seconds())
		return ret
	}
}

// 模拟一个耗时的函数
func slowFun(op int) int {
	time.Sleep(time.Second * 2)
	return op
}

func TestFn(t *testing.T) {
	a, _ := returnMulivValues()
	t.Log(a)
	tsf := timeSpent(slowFun)
	t.Log(tsf(10))
}

// func TestAnonyFn(t *testing.T) {
// 	f := func(a, b int) int {
// 		return a + b
// 	}
// 	t.Log(f(1, 2))
// }

// func TestClosure(t *testing.T) {
// 	counter := func() func() int {
// 		i := 0
// 		return func() int {
// 			i++
// 			return i
// 		}
// 	}()
// 	t.Log(counter())
// 	t.Log(counter())
// 	t.Log(counter())
// }

// 可变参数函数
func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// 可变参数函数测试
func TestVariadicFn(t *testing.T) {
	t.Log(Sum(1, 2))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("Clear resources")
}

func TestDefer(t *testing.T) {
	// 可以用来释放资源
	defer Clear()
	fmt.Println("Start TestDefer")
	panic("Something wrong happened") // 模拟异常，触发panic，但Clear仍会被调用
}
