package mypackage

// CarPublic Car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

// Car Car con acceso privado
type Car struct {
	brand string
	year  int
}

func PrintMessage() {
	println("HolaCurso")
}
