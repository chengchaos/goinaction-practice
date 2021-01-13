package p3_5_1

import (
    "fmt"
    "os"
    "reflect"
)

// 使用指针传递比较大数组。
func foo(array *[1]int) {
    fmt.Println(array[0])
}

// 学习数组呀
func PracticeArrays() {
    args := os.Args
    fmt.Println("args =>", args)

    fmt.Printf("The works, %v\n", 3.14)
    // 声明一个数组
    array := [5]int{1, 2, 3, 4, 5}
    fmt.Println("array =>", array)

    // 使用 。。。 代替数组长度
    // 会根据元素长度来确定
    array2 := [...]int{1, 3, 5, 7, 9}
    fmt.Println("array2 =>", array2)
    // 结果：[5]int
    fmt.Println("Type : ", reflect.TypeOf(array2))

    // 使用索引赋值：
    // 结果：[0 10 20 0 0 ]
    array3 := [5]int{1: 10, 2: 20}
    fmt.Println("array3 =>", array3)

    // 声明 5 个指向整数指针的数组：
    array4 := [5]*int{0: new(int), 1: new(int)}
    *array4[0] = 10
    *array4[1] = 20
    // 结果： array4 => [0xc00001c0d8 0xc00001c0f0 <nil> <nil> <nil>] 0xc00001c0d8 20
    fmt.Println("array4 =>", array4, array4[0], *array4[1])

    // 使用指针传递比较大的数组
    // 一个 8 兆大的数组
    var array1e6 [1]int = [1]int{100}
    foo(&array1e6)
}

// 学习 Slice
// 函数调用时，传递的是 Slice 的值，
// 一个 Slice 的值包含 24 个字节（指针：8， 长度：8， 容量8）
// 底层数组不属于切片本身，因此不会被复制。
func PracticeSlice() {

    // 一种方法是使用 make()
    // 长度为 5， 容量也为 5.
    slice1 := make([]string, 5)
    slice1[1] = "hello"
    // 长度为 3， 容量为 5.
    slice2 := make([]int, 3, 5)

    // 另一种方式是使用字面量
    slice3 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
    slice1 = slice3[1:3] // [ Blue, Green ]

    // 制定长度，要给出最大的元素：
    slice4 := []string{10: ""}

    // 空的 Slice
    // nil, 长度：0， 容量： 0
    var slice5 []int = nil

    // 使用 make() 函数创建
    slice6 := make([]int, 0)
    var slice7 []int

    fmt.Println("slice1 =>", slice1)
    fmt.Println("slice2 =>", slice2)
    fmt.Println("slice3 =>", slice3)
    fmt.Println("slice4 =>", slice4)
    fmt.Println("slice5 =>", slice5)
    fmt.Println("slice6 =>", slice6)
    fmt.Println("slice7 =>", slice7)

    slice8 := []int{10, 20, 30, 40, 50}
    slice9 := slice8[1:3] // 20, 30
    slice9 = append(slice9, 60)

    fmt.Println("slice9 =>", slice9) // 20, 30, 60
    fmt.Println("slice8 =>", slice8) //10, 20, 30, 60, 50

    slice10 := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
    slice11 := slice10[2:3:4]
    fmt.Println("slice_11 =>", slice11) // [Plum] 但是容量到 4 的位置。

    slice12 := append(slice8, slice9...)

    fmt.Println("append ... =>", slice12) // [10 20 30 60 50 20 30 60]

    fmt.Printf("slice12 len: %d, cap: %d\n",
        len(slice12),
        cap(slice12))
    // range 创建了每个元素的副本，而不是直接返回对该元素的引用。
    for index, value := range slice12 {
        fmt.Printf("Index: %d Value: %d And Value-Addr: %X, ElemAddr: %X \n",
            index, value,
            &value, &slice12[index])
    }

}


/*
学习 Map 。

Key 可以使用任何类型，但是这个类型可以使用 == 来比较。
切片、函数以及包含切片的结构类型具有引用语义，因此不能作为 Key。
 */
func PracticeMaps() {

	// 创建和初始化
	dict1 := make(map[string]int)
	dict2 := map[string]string{"Red" : "#da1337", "Orange" : "#e95a22"}
	fmt.Println("dict1 =>", dict1)
	fmt.Println("dict2 =>", dict2)

	fmt.Printf("red => %s\n", dict2["Red"])

	// nil map 不能用于存储键值对。
	//var dict3 map[string]string = nil
	//dict3["name"] = "chengchao"
	//fmt.Printf("dict3 name => %s\n", dict3["name"])

	theBlue, exists := dict2["Blue"]
	if exists {
		fmt.Printf("theBlue => %s\n", theBlue)
	} else {
		fmt.Println("theBlue is nil")
	}

	// 或者：
	theBlue = dict2["Blue"]
	fmt.Printf("theBlue => %v ; => %q\n",
		reflect.TypeOf(theBlue),
		theBlue)
	// 它真的会返回一个空串 ""； 其实是返回零值。
	if theBlue != "" {
		fmt.Println("或者使用值来判断", theBlue)
	} else {
		fmt.Println("这会打印什么？")
	}

	for key, value := range dict2 {
		fmt.Printf("Key: %s, Value: %s\n", key,
			value)
	}

	delete(dict2, "Coral")
}
