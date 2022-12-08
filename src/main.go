package main

import (
	"curso_goland_platzi/src/mypackage"
	"fmt"
	"log"
	"strconv"
)

type pc struct {
	ram   int
	disco int
	proce string
}

func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d GB de RAM, %d GB de disco duro y un procesador %s", myPc.ram, myPc.disco, myPc.proce)
}


func dobleRetorno(a int) (int, int) {
	return 5, 6
}

func isPalindromo(text string) {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es Palindromo")
	} else {
		fmt.Println("NO Es Palindromo")
	}

}

type car struct {
	brand string
	year  int
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

	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println(areaCuadrado)

	valor1, valor2 := dobleRetorno(5)
	fmt.Println(valor1, valor2)

	value1, _ := dobleRetorno(5)
	fmt.Println(value1)

	// For condicional
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// For while
	counter := 0
	for counter < 5 {
		fmt.Println(counter)
		counter++
	}

	// For forever
	counter2 := 0
	for {
		fmt.Println(counter2)
		counter2++
		if counter2 == 5 {
			break
		}
	}

	// convertir texto a numero
	num, err := strconv.Atoi("12")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Error numero:", num)

	// defer
	defer fmt.Println("Mundo")
	fmt.Println("Hola")

	// Arrays
	var array [5]int
	array[0] = 1
	fmt.Println(array)
	fmt.Println(array, len(array), cap(array))

	// Slices
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Metodos de los slices
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append nueva lista
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	// Recorrer slice
	slice2 := []string{"hola", "que", "hace"}

	for i, valor := range slice2 {
		fmt.Println(i, valor)
	}

	//Palindromo
	isPalindromo("amo")

	// LLave Valor
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// Recorrer map
	for llave, valor := range m {
		fmt.Println(llave, valor)
	}

	// Encontrar valor
	value, ok := m["Jose"]
	fmt.Println(value, ok)

	// Structs
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	// Structs 2
	var myCar2 car
	myCar2.brand = "Ferrari"
	myCar2.year = 2021
	fmt.Println(myCar2)

	var myCar3 mypackage.CarPublic
	myCar3.Brand = "Honda"
	myCar3.Year = 2021
	fmt.Println(myCar3)

	mypackage.PrintMessage()

	// myPc
	myPc := pc{ram: 16, disco: 500, proce: "Intel"}
	fmt.Println(myPc)

}

