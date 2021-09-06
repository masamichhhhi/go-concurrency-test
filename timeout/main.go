package main

import (
	"fmt"
	"time"
)

func main() {
	// timeout := time.After(1 * time.Second)
	// for {
	// 	select {
	// 	case <-timeout:
	// 		fmt.Println("time out")
	// 		return
	// 	default:
	// 		fmt.Println("default")
	// 		time.Sleep(time.Millisecond * 100)
	// 	}
	// }
	batch()
}

func batch() {
	t := time.NewTicker(time.Millisecond * 100)
	defer t.Stop()
	for i := 0; i < 5; i++ {
		select {
		case <-t.C:
			fmt.Println("tick")
		}
	}
}
