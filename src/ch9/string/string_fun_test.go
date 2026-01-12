/*
 * @Descripttion:
 * @version:
 * @Author: dengweiyi
 * @Date: 2026-01-09 18:22:56
 * @LastEditors: dengweiyi
 * @LastEditTime: 2026-01-09 18:30:36
 */
package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C,D,E"
	// 分割字符串
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	// 连接字符串
	t.Log(strings.Join(parts, "-"))
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	// 整数转字符串
	t.Log("string" + s)
	// 字符串转整数
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
