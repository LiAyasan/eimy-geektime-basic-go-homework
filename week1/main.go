package main

import (
	"eimy-geektime-basic-go-homework/week1/slice"
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	nums2 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	// 普通实现
	//RunDeleteV1(nums)
	// 泛型实现
	RunDeleteV2[int](nums, 3)
	RunDeleteV2[float64](nums2, 3)

	RunDeleteV3[int](nums, 0)
}

//func RunDeleteV1(nums []int) {
//	nums, err := slice.DeleteV1(nums, 3)
//	if err != nil {
//		fmt.Printf("DeleteAndReturn error: %v\n", err)
//	}
//	fmt.Printf("ans: %v\n", nums)
//}

func RunDeleteV2[T any](nums []T, idx int) {
	nums, err := slice.DeleteV2[T](nums, idx)
	if err != nil {
		fmt.Printf("DeleteAndReturn error: %v\n", err)
	}
	fmt.Printf("ans: %v\n", nums)
}

func RunDeleteV3[T any](nums []T, idx int) {
	if err := slice.DeleteV3[T](nums, idx); err != nil {
		fmt.Printf("DeleteByIdx error: %v\n", err)
	}
	fmt.Printf("ans: %v\n", nums)

}
