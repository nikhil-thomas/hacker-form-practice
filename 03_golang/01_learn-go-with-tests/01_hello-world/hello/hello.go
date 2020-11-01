package hello

import "fmt"

func main() {
	name := "World"
	fmt.Println(Hello(name))
}

// Hello prints greeting
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
