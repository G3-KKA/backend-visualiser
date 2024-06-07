package test

import "fmt"

//vdsdddisualize:start?name

func mai() {
	//comment
	ResultShouldBe()
}

func PrintThat() {
	fmt.Print(`func mai() {
	//comment
	ResultShouldBe()
}
`)

}
func ResultShouldBe() {
	fmt.Print(`func other() {
	//comment
	fmt.Println("Hello, World!")
}`)
}
