package main

// func display() int {
// 	var a int
// 	fmt.Print("Enter value for a : ")
// 	fmt.Scan(&a)
// 	fmt.Printf("The value of a is : %v\n", a)
// 	return a
// }

func main() {
	// // VARIABLE DECLARATION

	// // Explicit
	// var a int = 10
	// var b float64 = 1.2
	// var c string = "salutations"
	// var d bool = true
	// var a1 int32

	// // Type inference
	// var a2 = 12

	// // Short-hand
	// d1 := false

	// // Constant
	// const f int = 25

	// fmt.Printf("Explicit declaration of variables : %v, %v, %v, %v, %v \nType Inference declaration of variables : %v
	// \nShort-hand declaration of variables : %v \nConstants declaration : %v", a, b, c, d, a1, a2, d1, f)

	// // FUNCTION CALLING

	// val := display()
	// fmt.Printf("Returned value : %v", val)

	// // ARRAY AND SLICE DECLARATION

	// var arr = [3]int{1, 2, 3}
	// var arr1 = [0]int{}

	// var nums = []int{4, 5}
	// var nums1 = []int{}

	// // Appending elements in the slices
	// nums = append(nums, 8)
	// nums1 = append(nums1, 5)

	// // Removing elements from slices
	// i := 1
	// nums = append(nums[:i], nums[i+1:]...)

	// // Printing the elements
	// fmt.Println("arr : ", arr)
	// fmt.Println("arr1 : ", arr1)
	// fmt.Println("nums : ", nums)
	// fmt.Println("nums1 : ", nums1)

	// // The size and capacity of arrays and slices
	// fmt.Println("Length & Capacity:")
	// fmt.Printf("arr -> len: %d, cap: %d\n", len(arr), cap(arr))
	// fmt.Printf("arr1 -> len: %d, cap: %d\n", len(arr1), cap(arr1))
	// fmt.Printf("nums -> len: %d, cap: %d\n", len(nums), cap(nums))
	// fmt.Printf("nums1 -> len: %d, cap: %d\n", len(nums1), cap(nums1))

	// // MAPS

	// // Explicit declaration
	// var m map[int]int = make(map[int]int)

	// // Short-hand
	// m1 := make(map[int]int)

	// // Adding values
	// m[1] = 2
	// m[2] = 3
	// m[3] = 4

	// m1[1] = 0
	// m1[2] = 10
	// m1[3] = 100

	// // Printing all the values
	// for i, j := range m {
	// 	fmt.Println("The key and their respective values are : ", i, j)
	// }

	// // Getting the length
	// fmt.Println("Length : ", len(m))

	// // Accessing values
	// fmt.Println("value of key 1 : ", m[1])

	// // Deleting values
	// delete(m, m[2])

	// // Reprinting all the values
	// for i, j := range m {
	// 	fmt.Println("The new values and their pairs are : ", i, j)
	// }

	// // New size
	// fmt.Println("New size is : ", len(m))

	// LOOPS

	A := 10

	for i := range A {
		println(i)
	}

	for i := 0; i < A; i++ {
		println(A)
	}

	// For Maps
	c := make(map[string]int)

	c["B"] = 2
	c["C"] = 3
	c["D"] = 4

	for i, j := range c {
		println("The key and value pairs are : ", i, j)
	}
}
