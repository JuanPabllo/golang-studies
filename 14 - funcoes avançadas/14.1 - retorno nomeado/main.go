package main

func calcMaths(n1, n2 int) (soma, sub int) {
	soma = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	soma, sub := calcMaths(10, 20)
	println(soma, sub)
}
