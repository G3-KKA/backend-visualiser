package main

import "fmt"

/* possible options are name=rand(),latest() */
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
	fmt.Print(`func mai() {
	//comment
	ResultShouldBe()
}
func PrintThat() {`)
}

/*func mai() {
	//comment
	ResultShouldBe()
}
func PrintThat() {*/
