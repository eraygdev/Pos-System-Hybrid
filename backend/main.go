package main

import "fmt"

func main() {
	err := initRestaurants()
	fmt.Println(err)
	fmt.Println("Successfully loaded!")
}
