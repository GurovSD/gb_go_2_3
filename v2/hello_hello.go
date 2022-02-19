package hello_hello

import "fmt"

func SayHello() {

	fmt.Println("Hello from external function!")

}

//allow to say hello as many times as you want
func SaySuperHello(amount int) {

	for _ = range make([]int, amount) {
		SayHello()
	}

}
