package main

import "fmt"

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarAlista(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) eliminarDeLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

// * El puntero que hace referencia
func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) acttualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func (t *taskList) imprimirLista(){
	for _, tarea := range t.tasks {
		fmt.Println("Nombre", tarea.nombre)
		fmt.Println("Descrp", tarea.descripcion)
	}
}


func main() {
	//& para no crear copia sino el espacio de memoria
	t1 := &task{
		nombre:      "Completar el curso de GO",
		descripcion: "Completarlo en esta semana",
	}

	t2 := &task{
		nombre:      "Completar el curso de paython",
		descripcion: "Completarlo en esta semana2",
	}

	t3 := &task{
		nombre:      "Completar el curso de nodejs",
		descripcion: "Completarlo en esta semana3",
	}

	//% y + es patron para imprimir propiedad y valor
	fmt.Printf("%+v\n", t1)
	t1.marcarCompleta()
	t1.actualizarNombre("Finalizar mi curso de Go")
	t1.acttualizarDescripcion("Completar rapido el curso")
	fmt.Printf("%+v\n", t1)

	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	fmt.Println(lista.tasks[0])
	lista.agregarAlista(t3)

	fmt.Println(len(lista.tasks))

	lista.eliminarDeLista(1)
	fmt.Println(len(lista.tasks))

	for i := 0; i < len(lista.tasks); i++ {
		fmt.Println("Index", i, "Nombre", lista.tasks[i].nombre)
	}

	for index, tarea := range lista.tasks{
		fmt.Println("Index", index, "Descripcion", tarea.descripcion)
	}

	println("<------------- Impresion desde funcion --------->")
	lista.imprimirLista()

	mapaTareas := make(map[string]*taskList)

	mapaTareas["David"] = lista

	t4 := &task{
		nombre:      "Completar el curso de Java",
		descripcion: "Completarlo en esta semana",
	}

	t5 := &task{
		nombre:      "Completar el curso de C#",
		descripcion: "Completarlo en esta semana2",
	}

	lista2 := &taskList{
		tasks: []*task{
			t4, t5,
		},
	}

	mapaTareas["Steven"] = lista2

	fmt.Println("Tareas de David")
	mapaTareas["David"].imprimirLista()

	
	fmt.Println("Tareas de Steven")
	mapaTareas["Steven"].imprimirLista()

}
