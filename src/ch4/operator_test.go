package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 2, 4}
	//c := [...] int{1, 2, 3}
	//会比较数组中的元素
	t.Log(a == b)
	//长度不同的数组不能比较
	//t.Log(a == c)
}
