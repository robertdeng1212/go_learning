/*
 * @Descripttion:
 * @version:
 * @Author: dengweiyi
 * @Date: 2026-01-09 13:23:09
 * @LastEditors: dengweiyi
 * @LastEditTime: 2026-01-09 13:41:49
 */
package array_test

import "testing"

func TestArrayInitialization(t *testing.T) {
	var arr [3]int
	arr1 := [3]int{1, 2, 3}
	// 不指定元素
	arr2 := [...]int{4, 5, 6, 7}
	arr3 := [2][2]int{{1, 2}, {3, 4}}
	arr1[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr2)
	t.Log(arr3)
}

func TestArrayTraversal(t *testing.T) {
	arr := [3]string{"a", "b", "c"}
	// 传统for循环
	for i := 0; i < len(arr); i++ {
		t.Logf("index: %d, value: %s", i, arr[i])
	}
	// for range
	for index, value := range arr {
		t.Logf("index: %d, value: %s", index, value)
	}
}

func TestArrayCopy(t *testing.T) {
	arr1 := [3]int{1, 2, 3}
	arr2 := arr1
	arr2[0] = 10
	t.Log("arr1:", arr1)
	t.Log("arr2:", arr2)
}

// 数据截取
func TestArraySection(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	arr_sec := arr[:2]
	t.Log(arr_sec)
}
