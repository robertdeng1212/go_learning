package map_ext

import "testing"

func TestDeleteMapKey(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Logf("before delete, len m1=%d", len(m1))
	delete(m1, 2)
	t.Logf("after delete, len m1=%d", len(m1))
	if v, ok := m1[2]; ok {
		t.Logf("key 2's value is %d", v)
	}
}

// Map 的 value 可以是⼀一个⽅方法
// 与 Go 的 Dock type 接⼝口⽅方式⼀一起，可以⽅方便便的实现单⼀一⽅方法对象的⼯工⼚厂模式
func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

// Go 的内置集合中没有 Set 实现， 可以 map[type]bool
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		t.Logf("%d is in the set", n)
	} else {
		t.Logf("%d is not in the set", n)
		// map[1:true]
		t.Log(mySet)
	}
	mySet[3] = true
	delete(mySet, 1)
	n = 1
	if mySet[n] {
		t.Logf("%d is in the set", n)
	} else {
		t.Logf("%d is not in the set", n)
	}
}
