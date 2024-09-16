package slice

import (
	"fmt"
	"testing"
)

func TestDeleteAt(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	idx := 5
	// 下标异常
	s1, err := DeleteAt(nums, idx)
	if err != nil {
		fmt.Println(err, ToStringIdxErr(len(nums), idx))
	}
	fmt.Println(s1)

	// 正常下标
	idx = 2
	s2, err := DeleteAt(nums, idx)
	if err != nil {
		fmt.Println(err, ToStringIdxErr(len(nums), idx))
	}
	fmt.Println(s2)
}

func TestDeleteAtV2(t *testing.T) {
	// 泛型实现
	nums := []int{1, 2, 3, 4, 5}
	idx := 2
	s1, err := DeleteAtV2[int](nums, idx)
	if err != nil {
		fmt.Println(err, ToStringIdxErr(len(nums), idx))
	}
	fmt.Println(s1)

	fls := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	idx = 3
	s2, err := DeleteAtV2[float64](fls, idx)
	if err != nil {
		fmt.Println(err, ToStringIdxErr(len(fls), idx))
	}
	fmt.Println(s2)

	ss := []string{"a", "b"}
	idx = 3
	// 下标异常
	s3, err := DeleteAtV2[string](ss, idx)
	if err != nil {
		fmt.Println(err, ToStringIdxErr(len(ss), idx))
	}
	fmt.Println(s3)
}
