package main

import (
	"fmt"
)

//   for_temp := range
//   len_temp := len(for_temp)
//   for index_temp = 0; index_temp < len_temp; index_temp++ {
//           value_temp = for_temp[index_temp]
//           index = index_temp
//           value = value_temp
//           original body
//   }

func main() {
	copyOrquote()
}

// 死循环问题， 切边遍历前数组长度就已经确定
func deadLoop() {
	v := []int{1, 2, 3}
	for i := range v {
		fmt.Println(&v[i])
		v = append(v, i)
	}

	for i := range v {
		fmt.Println(&v[i])
	}
}

// 引用问题
func quote() {
	slice := []int{0, 1, 2, 3}
	myMap := make(map[int]*int)
	for index, value := range slice {
		// 使用引用不会保存每次的变量值，
		// &value的值是不变的，myMap的的value是同一个值
		// 使用值拷贝就不会出现
		fmt.Println(&index, &value)
		myMap[index] = &value
	}
	fmt.Println("=====new map=====")
	for k, v := range myMap {
		fmt.Printf("%d => %d\n", k, *v)
	}
}

// 值拷贝和引用
func copyOrquote() {
	slice := [1000]int{1, 2, 3}
	var k, v int
	// 遍历前会拷贝，造成内存浪费，通过引用
	for k, v = range slice {
		_ = k
		_ = v
	}
	fmt.Println(k + 1)
}