package main

import "fmt"

func main() {
	//4.1-4.2 声明一个包括5个元素的整型数组

	// 初始化
	var array [5]int
	// 字面量初始化并且赋值
	array2 := [5]int{12, 12, 23, 34, 435}
	fmt.Println(array, array2)
	//4.3 auto count the length of the array
	array3 := [...]int{12, 23, 34, 3, 4, 23434}
	fmt.Println(len(array3))
	// visit the aray
	array4 := [...]int{12, 23, 34, 3, 4, 23434}
	fmt.Println(array4)
	array4[4] = 1111111
	fmt.Println(array4)
	//array store point
	array5 := [...]*int{0: new(int)}
	*array5[0] = 10
	fmt.Println(*array5[0])
	// copy
	var array6 [5]string
	array7 := [5]string{"ds", "sdfsdfa", "sdfasdfasf", "sadfadf", "sdadsfads"}
	array6 = array7
	fmt.Println(array6)
	// 4.9 把一个指针数组赋值给另一个
	var array8 [3]*string
	ssP := new(string)
	*ssP = "dsf"
	array9 := [3]*string{0: ssP}
	array8 = array9
	fmt.Println(array8, *array8[0])
	fmt.Println(*array8[0])
	//TODO:多维数组************//

}
