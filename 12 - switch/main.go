package main

import "fmt"

func dayOfWeek(number int) string {
	switch number {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

func dayOfWeek2(number int) string {
	var diaDaSemana string

	switch {
	case number == 1:
		diaDaSemana = "Domingo"
	case number == 2:
		diaDaSemana = "Segunda"
	case number == 3:
		diaDaSemana = "Terça"
	case number == 4:
		diaDaSemana = "Quarta"
	case number == 5:
		diaDaSemana = "Quinta"
	case number == 6:
		diaDaSemana = "Sexta"
	case number == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Número inválido"
	}

	return diaDaSemana

}

func main() {
	fmt.Println("Switch")
	day := dayOfWeek(500)
	day2 := dayOfWeek2(2)
	fmt.Println(day)
	fmt.Println(day2)
}
