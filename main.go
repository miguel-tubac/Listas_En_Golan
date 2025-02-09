package main

import (
	"fmt"
)

// Este es el menu principal del programa
func menu_principal() {
	var opcion int
	for {
		fmt.Println("\n------------------- Menú principal --------------------")
		fmt.Println("\nSeleccione el tipo de lista a probar:")
		fmt.Println("1. Lista Contigua (Arreglo)")
		fmt.Println("2. Lista Ligada (Simple)")
		fmt.Println("3. Lista Doblemente Ligada")
		fmt.Println("4. Lista Indexada")
		fmt.Println("5. Salir")
		fmt.Print("\nIngrese una opcion: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			//fmt.Println("Lista_Contigua()")
			lista_contigua_menu()
		case 2:
			//fmt.Println("Lista_Ligada()")
			lista_ligada_menu()
		case 3:
			//fmt.Println("Lista_Doblemente_Ligada()")
			lista_ligada_doble_menu()
		case 4:
			//fmt.Println("Lista_Indexada()")
			lista_indexada_menu()
		case 5:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opción inválida, intente nuevamente.")
		}
	}
}

// ******************************* Lista Contigua ***********************************
func lista_contigua_menu() {
	var opcion int
	var datos []int
	for {
		fmt.Println("\n¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬ Menú opciones Lista_Contigua ¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬")
		fmt.Println("\nSeleccione la operacion a realizar:")
		fmt.Println("1. Insertar datos")
		fmt.Println("2. Eliminar datos")
		fmt.Println("3. Mostrar datos visualmente")
		fmt.Println("4. Regresar...")
		fmt.Print("\nIngrese una opcion: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			//fmt.Println("Insertar datos")
			datos = append(datos, lista_contigua_insertar(datos)...)
		case 2:
			fmt.Println("Eliminando datos...")
			datos = nil
			fmt.Scanln() // Espera a que el usuario presione Enter
		case 3:
			//fmt.Println("Mostrar datos visualmente")
			fmt.Println(datos)
			fmt.Scanln() // Espera a que el usuario presione Enter
		case 4:
			fmt.Println("Regresando...")
			return
		default:
			fmt.Println("Opción inválida, intente nuevamente.")
		}
	}
}

func lista_contigua_insertar(lista []int) []int {
	n := 0
	i := 0

	for {
		fmt.Print("\nIngrese el número de elementos (máximo 10): ")
		fmt.Scanln(&n)
		if n > 10 || n < 1 {
			print("\n!! El rango se pasa del limite, intentelo de nuevo ¡¡")
		} else {
			break
		}
	}

	fmt.Println() //Dejamos un salto de linea
	var dato int
	for i = 0; i < n; i++ {
		fmt.Print("Ingresa el elemento para la posicion ", i+1, " en la lista: ")
		fmt.Scanln(&dato)
		lista = append(lista, dato)
	}
	return lista
}

// ******************************* Fin de Lista Contigua ***********************************

// ******************************* Lista Ligada (Simple) ***********************************

// Nodo que representa un elemento en la lista ligada
type Nodo struct {
	valor int
	sig   *Nodo
}

// ListaLigada que representa la estructura de la lista
type ListaLigada struct {
	cabeza *Nodo
}

func lista_ligada_menu() {
	var opcion int
	lista := ListaLigada{}
	var valor int
	for {
		fmt.Println("\n¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬ Menú opciones Lista_Ligada ¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬")
		fmt.Println("\nSeleccione la operacion a realizar:")
		fmt.Println("1. Insertar datos")
		fmt.Println("2. Eliminar datos")
		fmt.Println("3. Mostrar datos visualmente")
		fmt.Println("4. Regresar...")
		fmt.Print("\nIngrese una opcion: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			//fmt.Println("Insertar datos")
			fmt.Println("Ingrese elementos (-1 para terminar):")
			for {
				fmt.Scan(&valor)
				if valor == -1 {
					break
				}
				lista.Insertar(valor) // Insertamos al inicio
			}
		case 2:
			fmt.Println("Eliminando datos...")
			lista.Eliminar()
			fmt.Scanln() // Espera a que el usuario presione Enter
		case 3:
			//fmt.Println("Mostrar datos visualmente")
			lista.Mostrar()
			fmt.Scanln() // Espera a que el usuario presione Enter
		case 4:
			fmt.Println("Regresando...")
			return
		default:
			fmt.Println("Opción inválida, intente nuevamente.")
		}
	}
}

// Metodos
// Insertar, agrega un nuevo nodo al inicio de la lista
func (l *ListaLigada) Insertar(valor int) {
	nuevoNodo := &Nodo{valor: valor, sig: l.cabeza}
	l.cabeza = nuevoNodo
}

// Mostrar, imprime la lista ligada
func (l *ListaLigada) Mostrar() {
	actual := l.cabeza
	fmt.Print("Lista Ligada: ")
	for actual != nil {
		fmt.Print(actual.valor, " -> ")
		actual = actual.sig
	}
	fmt.Println("NULL")
}

// Eliminar, limpiar la lista
func (l *ListaLigada) Eliminar() {
	l.cabeza = nil
}

// ******************************* Fin de Lista Ligada (Simple) ***********************************

// ******************************* Lista Doblemente Ligada ****************************************
// Nodo representa un elemento en la lista doblemente ligada
type Nodo2 struct {
	valor int
	sig   *Nodo2
	ant   *Nodo2
}

// ListaDoblementeLigada representa la lista
type ListaDoblementeLigada struct {
	cabeza *Nodo2
	tail   *Nodo2 // Último nodo para recorrer en reversa
}

func lista_ligada_doble_menu() {
	var opcion int
	lista := ListaDoblementeLigada{}
	var valor int
	for {
		fmt.Println("\n¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬ Menú opciones Lista_Ligada_Doble ¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬")
		fmt.Println("\nSeleccione la operacion a realizar:")
		fmt.Println("1. Insertar datos")
		fmt.Println("2. Eliminar datos")
		fmt.Println("3. Mostrar datos visualmente")
		fmt.Println("4. Regresar...")
		fmt.Print("\nIngrese una opcion: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			//fmt.Println("Insertar datos")
			fmt.Println("Ingrese elementos (-1 para terminar):")
			for {
				fmt.Scan(&valor)
				if valor == -1 {
					break
				}
				lista.Insertar(valor) // Insertamos al inicio
			}
		case 2:
			fmt.Println("Eliminando datos...")
			lista.Eliminar()
			fmt.Scanln() // Espera a que el usuario presione Enter
		case 3:
			//fmt.Println("Mostrar datos visualmente")
			lista.Mostrar()
			fmt.Scanln() // Espera a que el usuario presione Enter
		case 4:
			fmt.Println("Regresando...")
			return
		default:
			fmt.Println("Opción inválida, intente nuevamente.")
		}
	}
}

// Insertar agrega un nodo al inicio de la lista
func (l *ListaDoblementeLigada) Insertar(valor int) {
	nuevoNodo := &Nodo2{valor: valor, sig: l.cabeza}

	if l.cabeza != nil {
		l.cabeza.ant = nuevoNodo
	} else {
		l.tail = nuevoNodo // Si es el primer nodo, también es el último
	}

	l.cabeza = nuevoNodo
}

// Mostrar imprime la lista en orden normal
func (l *ListaDoblementeLigada) Mostrar() {
	actual := l.cabeza
	fmt.Print("Lista Doblemente Ligada (inicio a fin): ")
	for actual != nil {
		fmt.Print(actual.valor, " <-> ")
		actual = actual.sig
	}
	fmt.Println("NULL")
}

// Eliminar
func (l *ListaDoblementeLigada) Eliminar() {
	l.cabeza = nil
}

// ******************************* Fin de Lista Doblemente Ligada ***********************************

// ******************************* Lista Indexada ***********************************

func lista_indexada_menu() {
	var opcion int
	var lista []int
	var indice []int
	var n int
	for {
		fmt.Println("\n¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬ Menú opciones Lista_Indexada ¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬¬")
		fmt.Println("\nSeleccione la operacion a realizar:")
		fmt.Println("1. Insertar datos")
		fmt.Println("2. Eliminar datos")
		fmt.Println("3. Mostrar datos visualmente")
		fmt.Println("4. Regresar...")
		fmt.Print("\nIngrese una opcion: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			//fmt.Println("Insertar datos")
			listaIndexada_agragar(&lista, &indice, &n)
		case 2:
			fmt.Println("Eliminando datos...")
			lista = nil
			indice = nil
			n = 0
			fmt.Scanln() // Espera a que el usuario presione Enter
		case 3:
			//fmt.Println("Mostrar datos visualmente")
			mostrar_listaIndexada(&lista, &indice, n)
			fmt.Scanln() // Espera a que el usuario presione Enter
		case 4:
			fmt.Println("Regresando...")
			return
		default:
			fmt.Println("Opción inválida, intente nuevamente.")
		}
	}
}

func listaIndexada_agragar(lista *[]int, indice *[]int, n *int) {
	// Verificamos que n no exceda 10
	for {
		fmt.Print("\nIngrese el número de elementos (máximo 10): ")
		fmt.Scanln(&(*n))
		if *n > 10 || *n < 1 {
			print("\n!! El rango se pasa del limite, intentelo de nuevo ¡¡")
		} else {
			break
		}
	}

	// Redimensionamos las slices
	*lista = make([]int, *n)
	*indice = make([]int, *n)

	// Guardamos los valores y sus índices
	for i := 0; i < *n; i++ {
		fmt.Printf("Ingrese el elemento %d: ", i+1)
		fmt.Scan(&(*lista)[i])
		(*indice)[i] = i + 1 // Cada valor tiene un índice asociado
	}
}

func mostrar_listaIndexada(lista *[]int, indice *[]int, n int) {
	// Mostramos la lista indexada
	fmt.Println("Lista Indexada:")
	for i := 0; i < n; i++ {
		fmt.Printf("Indice: %d -> Valor: %d\n", (*indice)[i], (*lista)[i])
	}
}

// ******************************* Fin de Lista Indexada ***********************************

func main() {
	menu_principal()
}
