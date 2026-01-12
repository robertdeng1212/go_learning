/*
 * @Descripttion:
 * @version:
 * @Author: dengweiyi
 * @Date: 2026-01-09 11:20:27
 * @LastEditors: dengweiyi
 * @LastEditTime: 2026-01-09 13:18:27
 */
package condition_test

import "testing"

func TestIfCondition(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("Condition is true")
	}
}

func TestSwitchMultipleCases(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Logf("%d is even", i)
		case 1, 3:
			t.Logf("%d is odd", i)
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		// 条件 switch
		switch {
		case i%2 == 0:
			t.Logf("%d is even", i)
		case i%2 == 1:
			t.Logf("%d is Odd", i)
		default:
			t.Log("Unknown")
		}
	}
}
