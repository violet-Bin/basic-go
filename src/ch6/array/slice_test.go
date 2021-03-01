package slice

import "testing"

func TestSlice(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5) //只初始化了3个空间
	t.Log(len(s2), cap(s2))
	//t.Log(s2[0], s2[1], s2[2], s2[3])//index out of range
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 3)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2], s2[3])

}

//新创建空间
func TestSliceGrowing(t *testing.T) {
	var s = []int{}
	for i := 0; i < 10; i++ {
		s = append(s, 1)
		t.Log(len(s), cap(s))
	}
}

//共享存储空间
func TestShareMemory(t *testing.T) {
	years := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}
	Q2 := years[3:6]
	t.Log(Q2, len(Q2), cap(Q2))

	Q2[2] = "99"
	t.Log(years, Q2)
}

func TestSliceCompare(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	if a == b {
		t.Log("equal")
	}
}
