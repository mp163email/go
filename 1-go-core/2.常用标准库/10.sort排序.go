package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4, 2, 7, 1, 6}
	sort.Ints(nums)
	fmt.Println(nums)                     //升序排序
	fmt.Println(sort.IntsAreSorted(nums)) //检查是否已经排序

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println(nums) //降序排序

	fmt.Println(sort.IntsAreSorted(nums)) //检查是否已经排序

	strs := []string{"c", "a", "b"}
	sort.Strings(strs) //字符串排序
	fmt.Println(strs)

	ints := []int{5, 8, 1, 2, 6, 9} //这里是切片，不是数组，因为没有固定长度；切片是引用类型
	bubbleSort(ints)
	fmt.Println(ints)
}

/*
冒泡排序： 每轮都把最大数放到最后
*/
func bubbleSort(num []int) {
	//第一层循环是轮次，一共进行几轮,关键词  i < len(num) - 1
	for i := 0; i < len(num)-1; i++ {
		swappend := false
		//第二层循环是每一轮的比较次数 //关键词： j < len(num) - 1 - i
		for j := 0; j < len(num)-1-i; j++ {
			if num[j] > num[j+1] {
				num[j], num[j+1] = num[j+1], num[j]
				swappend = true
			}
		}
		if !swappend {
			break
		}
	}
}
