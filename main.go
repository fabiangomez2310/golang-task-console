package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	nombre     string
	desc       string
	completado bool
}

type ListaTareas struct {
	tareas []Tarea
}

func (l *ListaTareas) showAllTask() {
	if len(l.tareas) != 0 {
		for index, value := range l.tareas {
			completado := "NO"
			if value.completado {
				completado = "SI"
			}
			fmt.Println(index, ".", value.nombre, " - Descripcion: ", value.desc, " - Completado: ", completado)
		}
	} else {
		fmt.Println("-----No existen tareas")
	}
}

func (l *ListaTareas) addTask(t Tarea) {
	l.tareas = append(l.tareas, t)
	fmt.Println("-----Tarea agregada")
}

func (l *ListaTareas) editTask(index int, t Tarea) {
	l.tareas[index] = t
	fmt.Println("-----Tarea editada")
}

func (l *ListaTareas) deleteTask(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
	fmt.Println("-----Tarea eliminada")
}

func (l *ListaTareas) markCompleted(index int) {
	l.tareas[index].completado = true
	fmt.Println("-----Tarea completada")
}

func main() {

	leer := bufio.NewReader(os.Stdin)

	l := ListaTareas{}

	for {
		var op int

		fmt.Println("Bienvenido, Seleccione una opcion.")

		fmt.Println("1. Listar todas las tareas")
		fmt.Println("2. Agregar una tarea")
		fmt.Println("3. Editar una tarea")
		fmt.Println("4. Eliminar una tarea")
		fmt.Println("5. Marcar una tarea como completada")
		fmt.Println("6. Salir")

		fmt.Scanln(&op)

		switch op {
		case 1:
			l.showAllTask()
		case 2:
			t := Tarea{}
			fmt.Println("====== Agregar una nueva tarea ======")
			fmt.Println("Ingrese el nombre: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Println("Ingrese la descripcion: ")
			t.desc, _ = leer.ReadString('\n')
			l.addTask(t)
		case 3:
			t := Tarea{}
			var index int
			fmt.Println("====== Editar una tarea ======")
			l.showAllTask()
			fmt.Println("Digite el index de la tarea a editar:")
			fmt.Scanln(&index)

			fmt.Println("Ingrese el nombre: ")
			t.nombre, _ = leer.ReadString('\n')

			fmt.Println("Ingrese la descripcion: ")
			t.desc, _ = leer.ReadString('\n')
			l.editTask(index, t)
		case 4:
			var index int
			var op string
			fmt.Println("====== Eliminar una tarea ======")
			fmt.Println("Digite el index de la tarea a eliminar")
			fmt.Scanln(&index)

			for {
				fmt.Println("Esta seguro que desea eliminar la tarea: ", l.tareas[index].nombre, " (s/n)")
				fmt.Scanln(&op)
				if op == "s" {
					l.deleteTask(index)
					break
				} else if op == "n" {
					fmt.Println("Se cancela la eliminacion")
					break
				} else {
					fmt.Println("Por favor seleccione una opcion valida (s/n)")
				}
			}
		case 5:
			var index int
			var op string
			fmt.Println("====== Marcar una tarea como completada ======")
			l.showAllTask()
			fmt.Println("Digite el index de la tarea a completar")
			fmt.Scanln(&index)
			for {
				fmt.Println("Esta seguro que desea marcar la tarea ", l.tareas[index].nombre, " como completado (s/n)")
				fmt.Scanln(&op)
				if op == "s" {
					l.markCompleted(index)
					break
				} else if op == "n" {
					fmt.Println("Se cancela la marcacion de la tarea")
					break
				} else {
					fmt.Println("Por favor seleccione una opcion valida (s/n)")
				}
			}
		case 6:
			fmt.Println("Fin del programa")
			return
		default:
			fmt.Println("Seleccione una opcion valida")
			continue

		}
	}

}
