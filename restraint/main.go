package main

func main() {
	restFunc()
}

func restFunc() <-chan int {
	result := make(chan int)

	go func() {
		defer close(result)

		for i := 0; i < 5; i++ {
			result <- 1
		}
	}()

	return result
}
