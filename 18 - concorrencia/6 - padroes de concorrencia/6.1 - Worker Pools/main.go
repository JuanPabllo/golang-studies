package main

func main() {
	task := make(chan int, 45)
	results := make(chan int, 45)

	go worker(task, results)
	go worker(task, results)
	go worker(task, results)
	go worker(task, results)

	for i := 0; i < 45; i++ {
		task <- i
	}

	close(task)

	for i := 0; i < 45; i++ {
		resultsAll := <-results
		println(resultsAll)
	}

}

func worker(task <-chan int, results chan<- int) {

	for number := range task {
		results <- fibonacci(number)
	}

}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
