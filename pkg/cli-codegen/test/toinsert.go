package main

import (
	"fmt"
	"log"
	"unsafe"
)

/*  options to realize are name=rand(),latest() */
type FuncHolder struct {
	string
	fPtr1 func()
	fPtr2 func(int)
	fPtr3 func(string) error
}

func main() {

	fmt.Println(unsafe.Sizeof(FuncHolder{}))
	ResultShouldBe()
}
func PrintThat() {
	fmt.Print(`type FuncHolder struct {
	string
	fPtr1 func()
	fPtr2 func(int)
	fPtr3 func(string) error
}

func main() {

	fmt.Println(unsafe.Sizeof(FuncHolder{}))
	ResultShouldBe()
}
func PrintThat() {`)

}
func ResultShouldBe() {
	fmt.Println("hello")
	log.Println("hello")
	for _, v := range []int{1, 2, 3} {
		fmt.Println(v)
	}
	fmt.Print(`	fmt.Println("hello")
	log.Println("hello")
	for _, v := range []int{1, 2, 3} {
		fmt.Println(v)
	}`)
	fmt.Print(`	fmt.Println("hello")
	log.Println("hello")
	for _, v := range []int{1, 2, 3} {
		fmt.Println(v)
	}`)
	fmt.Print(`type FuncHolder struct {
	string
	fPtr1 func()
	fPtr2 func(int)
	fPtr3 func(string) error
}

func main() {

	fmt.Println(unsafe.Sizeof(FuncHolder{}))
	ResultShouldBe()
}
func PrintThat() {`)
	fmt.Print(`	fmt.Println("hello")
	log.Println("hello")
	for _, v := range []int{1, 2, 3} {
		fmt.Println(v)
	}`)

}
