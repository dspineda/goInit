package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(){
	inicio := time.Now()

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instahram.com",
	}

	for _, servidor := range servidores{
		revisarServidor(servidor)
	}
	tiempoPaso := time.Since((inicio))
	fmt.Printf("Tiempo de ejecucion %s", tiempoPaso)
}

func revisarServidor(servidor string){
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println("El servidos no esta disponible")
	}else{
		fmt.Println("El servidor esta funcionando normalmente")
	}
}  