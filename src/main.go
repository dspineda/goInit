package main

import "fmt"

func dobleRetorno(a int) (int, int) {
	return 5, 6
}

func main() {
	// Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415
	fmt.Println("Hello, World!", pi, pi2)

	// Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int
	
	fmt.Println(base, altura, area)	

	// Zero values
	var a int
	var b float64
	var c string
	var d bool 

	fmt.Println(a,b,c,d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("🚀 ~ file: main.go:29 ~ funcmain ~ areaCuadrado", areaCuadrado)
	fmt.Println(areaCuadrado)

	valor1, valor2 := dobleRetorno(5)
	fmt.Println(valor1, valor2)

	value1, _ := dobleRetorno(5)
	fmt.Println(value1)
}
