# golang-assignment

package main

//mport "fmt"

// func main() {
// 	fmt.Println("Hello World!")
// }

//for loop
// func main() {
// 	for i:= 1; i <= 5; i++ {
// 		fmt.Println("Hello")
// 	}
// }

// if loop
// func main() {
// 	i := 4
// 	if i%2 == 0 {
// 		fmt.Println("Even Number")
// 	} else {
// 		fmt.Println("Odd NUmber")
// 	}
// }

//using for and if loop
// func main() {
// 	for i := 1; i <= 5; i++ {
// 		if i == 3 {
// 			continue
// 		}
// 		if i == 5 {
// 			break
// 		}
// 		fmt.Println(i)
// 	}
// }

//defer
// func main() {
// 	fmt.Println("Hello")
// 	defer fmt.Println("Sruthi")
// 	fmt.Println("GoLang")
// }

//switch case
//

//Pointers: which refers to the memory address where data is stored
//*- will hold the value
//&-hold the address

//prg for pointers
// func main() {
// 	var x int = 30
// 	//x:=5 - short hand declaration
// 	//p:=&x -short hand declaration
// 	var p *int = &x
// 	fmt.Println(*p)
// }
//explanation for the above
//x- will have some address - 10
//p - we are assigning the address to the p
//if we change the value of  x , p will automatically get updated

//Structures: struct
// type person struct {
// 	firstname string
// 	age int
// }

// func main(){
// 	var p1 person
// 	p1.firstname = "Sruthi"
// 	p1.age = 23
// }

//anonymius struct

// func main(){
// 	student:= struct{
// 		name string
// 		grade int
// 	}{
// 		name: "Go",
// 		grade:10,
// 	}
// }

// Arrays- are fixed size collection of elements of same type
//a[]-slices
//a[5]- array

//example-
// type Person struct {
// }

// func main() {
// 	var arr [3]int
// 	arr[0] = 1
// 	arr[1] = 2
// 	arr[2] = 3
// }

//slice- it has 3 components:pointers,lenght(no of elements present in a slice) and capacity(max no of elements that are added before a new array)
//syntax of slice: var arrayName[]type --> no need to mention nay size
// there are 4 ways of declaring a slice-
// s:=make([]int,length)
//s:=make([]int,length,capacity)
//s:=[]int(1,2,3)

//we can declare like this
//arr :=[5] int{1,2,3,5,4}
//s:=arr[0:2]

//example for slice
// type Person struct {
// }

// func main() {
// 	s := make([]int, 3, 5)
// 	fmt.Println(s)
// 	s = append(s, 1)
// 	fmt.Println(s, len(s), cap(s))
// 	s = append(s, 2)
// 	fmt.Println(s, len(s), cap(s))

// }

//Array-fixed, value, we have to declare a size [n]type
//slice-dynamic,reference type ,[]type{} or make{}

//Maps- stores value in a key-value pair is like a data structure where the key is always unique value

//a:=make(map[key data] value data)
//a:=make(map[int]string){
//1:"one"
//2:"two"
//}
// the zero value of map is null
//example for maps
// import "fmt"

// func main() {
// 	scores := map[string]int{
// 		"Sai":    90,
// 		"Sruthi": 80,
// 	}
// 	fmt.Println("scores", scores)
// }
