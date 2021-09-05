package main

import "fmt"

func main() {
	gen1, gen2 := make(chan int), make(chan int)

	select {
	case num := <-gen1: // gen1から受信できるとき
		fmt.Println(num)
	case num := <-gen2: // gen2から受信できるとき
		fmt.Println(num)
	default: // どっちも受信できないとき
		fmt.Println("neither chan cannot use")
	}
}
