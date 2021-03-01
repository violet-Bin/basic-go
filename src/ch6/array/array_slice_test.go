package slice_test

import "testing"

func TestArray(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}

	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7}
	t.Log(arr[1], arr1[1], arr3[5])
}

func TestArrTravel(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}

	for index, e := range arr {
		t.Log(index, e)
	}
	for _, e := range arr {
		t.Log(e)
	}
}

//数组截取
func TestArrSelection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5, 6}
	newArr := arr[:3]

	for _, e := range newArr {
		t.Log(e)
	}
}
