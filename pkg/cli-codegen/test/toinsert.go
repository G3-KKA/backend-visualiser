package main

import (
	"fmt"
	"log"
)

/*  options to realize are name=rand(),latest() */

//visualize:start?name=mai&remainComments=true&addDebug=true
func mai() {
	//comment
	ResultShouldBe()
}
func PrintThat() {
	//visualize:stop
	//visualize:insert?name=mai

}
func ResultShouldBe() {
	//visualize:start?name=result
	fmt.Println("hello")
	log.Println("hello")
	for _, v := range []int{1, 2, 3} {
		fmt.Println(v)
	}
	//visualize:stop
	//visualize:insert?name=result
	//visualize:insert?name=result
	//visualize:insert?name=mai
	//visualize:insert?name=result

}
