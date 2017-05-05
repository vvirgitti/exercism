package hello

import "fmt"

const testVersion = 2

//A simple hello world greeting

func main() {
	fmt.Print(HelloWorld("Gopher"))
}

func HelloWorld(name string) string {
	if name == "" {
		return "Hello, World!"
	} else {
		return "Hello, " + name + "!"
	}


}
