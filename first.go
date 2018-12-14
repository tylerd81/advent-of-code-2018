package main

import "fmt"

type Cat struct {
	Age  int
	Name string
}

type CatInterface interface {
	HelloCat()
}

func main() {
	sarge := Cat{4, "Sarge"}
	// var sarge Cat
	// sarge.Age = 4
	// sarge.Name = "sarge"

	//sarge.HelloCat()
	var cat CatInterface
	cat = sarge
	cat.HelloCat()
	fmt.Println(sarge)
}

func (c Cat) String() string {
	return "I am a cat. Meow."
}

func (c Cat) HelloCat() {
	fmt.Printf("Hello, my name is %s. I am %d years old.\n", c.Name, c.Age)
}
func greet(name string) func() {
	return func() {
		fmt.Println("Hello,", name)
	}
}

func display(s string, fn func(string)) {
	fn(s)
}

func printer(s string) {
	fmt.Println(s)
}

func printSlices(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
