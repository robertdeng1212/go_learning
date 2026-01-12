/*
 * @Descripttion:
 * @version:
 * @Author: dengweiyi
 * @Date: 2026-01-09 10:40:08
 * @LastEditors: dengweiyi
 * @LastEditTime: 2026-01-09 10:43:14
 */
package operator_test

import "testing"

func TestOperatorPrecedence(t *testing.T) {
	a := [...]int{1, 2, 3}
	b := [...]int{4, 5, 6}
	// c := [...]int{7, 8, 9, 10}

	t.Log(a == b)
	// t.Log(a == c) // 编译错误：不同长度的数组不能比较
}
