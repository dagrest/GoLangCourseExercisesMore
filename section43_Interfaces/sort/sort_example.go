/*

https://pkg.go.dev/sort#Sort
https://pkg.go.dev/sort/#Interface

https://go.dev/doc/effective_go#printing

*/

package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {
	fmt.Println("Sort Example")

	fmt.Println("Sort interface:")
	studyGroup := people{"Zeno", "John", "AI", "Jenny"}
	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	//sort.Strings(studyGroup)
	fmt.Println(studyGroup)
	fmt.Println("-----------")

	fmt.Println("Sort strings:")
	s := []string{"Zeno", "John", "AI", "Jenny"}
	fmt.Println(s)
	//sort.Sort(sort.StringSlice(s)) // Also working
	sort.Strings(s)
	fmt.Println(s)
	fmt.Println("-----------")

	fmt.Println("Sort reverse strings:")
	s = []string{"Zeno", "John", "AI", "Jenny"}
	fmt.Println(s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)
	fmt.Println("-----------")

	fmt.Println("Sort ints:")
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println(n)
	sort.Ints(n)
	fmt.Println(n)
	fmt.Println("-----------")

	fmt.Println("Sort reverse ints:")
	n = []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println(n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)
	fmt.Printf("n teversed type: %T \n", n)
	fmt.Println("-----------")

	nn := sort.Reverse(sort.IntSlice(n))
	fmt.Println(nn)
	fmt.Printf("n teversed type: %T \n", n)

}
