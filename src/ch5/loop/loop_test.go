/*
 * @Descripttion:
 * @version:
 * @Author: dengweiyi
 * @Date: 2026-01-09 10:50:45
 * @LastEditors: dengweiyi
 * @LastEditTime: 2026-01-09 10:52:20
 */
package loop_test

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}
