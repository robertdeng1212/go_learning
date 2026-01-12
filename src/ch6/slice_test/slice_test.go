/*
 * @Descripttion:
 * @version:
 * @Author: dengweiyi
 * @Date: 2026-01-09 13:49:52
 * @LastEditors: dengweiyi
 * @LastEditTime: 2026-01-09 17:29:34
 */
package slice_test

import "testing"

func TestSliceInitialization(t *testing.T) {
	var s []int
	s1 := []int{1, 2, 3}
	s2 := make([]int, 3)
	s3 := make([]int, 2, 5)
	s1[1] = 5
	t.Log(s, len(s), cap(s))
	t.Log(s1, len(s1), cap(s1))
	t.Log(s2, len(s2), cap(s2))
	t.Log(s3, len(s3), cap(s3))
}

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(s0, len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(s0, len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(s1, len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(s2, len(s2), cap(s2))
	// t.Log(s2[0], s2[1], s2[2], s2[3], s2[4]) // 超出范围会报错
	s2 = append(s2, 1, 2)
	t.Log(s2, len(s2), cap(s2))
}

// 切片实现如何可变长
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Logf("len: %d, cap: %d, value: %v", len(s), cap(s), s)
	}
}

func TestSliceShareMeory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep",
		"Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknow"
	t.Log(Q2)
	t.Log(year)
}

func TestSliceComparing(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	// invalid operation: a == b (slice can only be compared to nil)
	if a == b {
		t.Log("equal")
	}
}
