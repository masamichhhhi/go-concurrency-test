package main

import (
	"fmt"
	"sync"
)

func main() {
	done := make(chan struct{})

	gen1 := generator(done, 1) // int 1をひたすら送信するチャネル(doneで止める)
	gen2 := generator(done, 2) // int 2をひたすら送信するチャネル(doneで止める)

	result := fanIn2(done, gen1, gen2) // 1か2を受け取り続けるチャネル
	for i := 0; i < 5; i++ {
		<-result
	}
	close(done)
	fmt.Println("main close done")

	// これを使って、main関数でcloseしている間に送信された値を受信しないと
	// チャネルがブロックされてしまってゴールーチンリークになってしまう恐れがある
	for {
		if _, ok := <-result; !ok {
			break
		}
	}
}

func fanIn1(done chan struct{}, c1, c2 <-chan int) <-chan int {
	result := make(chan int)

	go func() {
		defer fmt.Println("closed fanin")
		defer close(result)

		for {
			select {
			case <-done:
				fmt.Println("done")
				return
			case num := <-c1:
				fmt.Println("send 1")
				result <- num
			case num := <-c2:
				fmt.Println("send 2")
				result <- num
			default:
				fmt.Println("continue")
				continue
			}
		}
	}()
	return result
}

func fanIn2(done chan struct{}, cs ...<-chan int) <-chan int {
	result := make(chan int)

	var wg sync.WaitGroup

	wg.Add(len(cs))

	for i, c := range cs {
		go func(c <-chan int, i int) {
			defer wg.Done()

			for num := range c {
				select {
				case <-done:
					fmt.Println("wg.Done", i)
					return
				case result <- num:
					fmt.Println("send", i)
				}
			}
		}(c, i)
	}

	go func() {
		wg.Wait()
		fmt.Println("closing fanin")
		close(result)
	}()

	return result
}

func generator(done chan struct{}, num int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
	LOOP:
		for {
			select {
			case <-done:
				break LOOP
			case result <- num:
			}
		}
	}()
	return result
}
