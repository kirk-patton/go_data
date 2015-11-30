package main

import "fmt"

func main() {
	var h map[string]int
	h = make(map[string]int)
	h["hello"] = 100

	i := len(h)
	fmt.Println(i)
}
