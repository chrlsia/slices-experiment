package main

import "fmt"

func main() {
	var sl1 []int
	/* this does not work
	sl1[0] = 100
	because the slice sl1 is zero valued => nil
	(check it with debugger)
	*/
	fmt.Println("var sl1 []int")
	fmt.Printf("sl1 is %v, capacity of it is %v\n", sl1, cap(sl1))
	sl1 = append(sl1, 1, 2, 3, 4)
	fmt.Println(sl1)

	sl2 := []int{1, 2, 3}
	fmt.Println("sl2 := []int{1, 2, 3}")
	fmt.Printf("sl2 is %v, capacity of it is %v\n", sl2, cap(sl2))
	/*this does not work, because we exceed the capacity which is 3
	sl2[3] = 4
	*/
	sl2[0] = 100
	fmt.Println(sl2)

	/*
		The make built-in function allocates and initializes
		an object of type slice, map, or chan (only)
	*/
	sl3 := make([]int, 4)
	fmt.Println("sl3 := make([]int, 4)")
	fmt.Printf("sl3 is %v, capacity of it is %v\n", sl3, cap(sl3))
	sl3[0] = 1
	fmt.Println(sl3)

	sl4 := make([]int, 0)
	/*this does not work, because we exceed the capacity which is 0
	sl4[0] = 0
	*/
	fmt.Printf("sl4 is %v, capacity of it is %v\n", sl4, cap(sl4))
}
