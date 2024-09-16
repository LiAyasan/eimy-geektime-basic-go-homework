package slice

import (
	"errors"
	"fmt"
)

var ErrIndexOutOfRange = errors.New("下标超出范围")

// DeleteAt 删除指定位置的元素
// 如果下标不是合法的下标，返回 ErrIndexOutOfRange
func DeleteAt(s []int, idx int) ([]int, error) {
	if IsOutOfIndex(idx, len(s)) {
		return s, ErrIndexOutOfRange
	}
	tmp := make([]int, len(s)-1)
	tmpIdx := 0
	for i, num := range s {
		if i != idx {
			tmp[tmpIdx] = num
			tmpIdx++
		}
	}
	return tmp, nil
}

// DeleteAtV2
// 时间:O(n) 空间:O(n)
func DeleteAtV2[T any](elems []T, idx int) ([]T, error) {
	if IsOutOfIndex(idx, len(elems)) {
		return elems, ErrIndexOutOfRange
	}
	tmp := make([]T, len(elems)-1)
	tmpIdx := 0
	for i, num := range elems {
		if i != idx {
			tmp[tmpIdx] = num
			tmpIdx++
		}
	}
	return tmp, nil
}

// DeleteV1
// 时间:O(n) 空间:O(n)
func DeleteV1(nums []int, idx int) ([]int, error) {
	if IsOutOfIndex(idx, len(nums)) {
		return nums, errors.New("下标不合法")
	}
	tmp := make([]int, len(nums)-1)
	tmpIdx := 0
	for i, num := range nums {
		if i != idx {
			tmp[tmpIdx] = num
			tmpIdx++
		}
	}
	return tmp, nil
}

// DeleteV3
// 时间: O(n), 空间: O(1)
func DeleteV2[T any](elems []T, idx int) error {
	if IsOutOfIndex(idx, len(elems)) {
		return errors.New("下标不合法")
	}

	// 没有缩容 len = cap - 1
	elems = append(elems[:idx], elems[idx+1:]...)

	return nil
}

func IsOutOfIndex(idx, length int) bool {
	return idx < 0 || idx >= length
}

func ToStringIdxErr(length, idx int) string {
	return fmt.Sprintf("数组长度为( %d ), 删除下标为( %d )", length, idx)
}
