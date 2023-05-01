package main

func main() {
	var anio = 2020

	if anio%4 == 0 || anio%100 != 0 || anio%400 == 0 {
		println("Es biciesto")
	} else {
		println("No es biciesto")
	}
}
