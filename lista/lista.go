package lista

// Lista es una interfaz que representa una lista genérica de elementos de tipo T.
type Lista[T any] interface {

	// EstaVacia devuelve true si la lista no tiene elementos, false en caso contrario.
	// Post: devuelve true si la lista está vacía, false si no.
	EstaVacia() bool

	// InsertarPrimero inserta el elemento al principio de la lista.
	// Pre: -
	// Post: Se agregó el elemento como primero de la lista. El largo de la lista aumenta en 1.
	InsertarPrimero(T)

	// InsertarUltimo inserta el elemento al final de la lista.
	// Pre: -
	// Post: Se agregó el elemento como último de la lista. El largo de la lista aumenta en 1.
	InsertarUltimo(T)

	// BorrarPrimero remueve el primer elemento de la lista y lo devuelve.
	// Pre: La lista no está vacía.
	// Post: Se removió el primer elemento de la lista. El largo de la lista disminuye en 1.
	// Devuelve el elemento que estaba en la primera posición.
	BorrarPrimero() T

	// VerPrimero devuelve el primer elemento de la lista sin removerlo.
	// Pre: La lista no está vacía.
	// Post: Devuelve el primer elemento de la lista. La lista no se modifica.
	VerPrimero() T

	// VerUltimo devuelve el último elemento de la lista sin removerlo.
	// Pre: La lista no está vacía.
	// Post: Devuelve el último elemento de la lista. La lista no se modifica.
	VerUltimo() T

	// Largo devuelve la cantidad de elementos en la lista.
	// Post: Devuelve el número de elementos en la lista (≥ 0).
	Largo() int

	// Iterar recorre la lista aplicando la función visitar a cada elemento hasta que:
	// - Se recorren todos los elementos, o
	// - visitar devuelve false
	// Pre: visitar no es nil.
	// Post: Se aplicó visitar a cada elemento hasta que devolvió false o se terminó la lista.
	Iterar(visitar func(T) bool)

	// Iterador devuelve un iterador para recorrer la lista.
	// Post: Devuelve un iterador posicionado antes del primer elemento de la lista.
	Iterador() IteradorLista[T]
}

// IteradorLista es una interfaz que permite iterar sobre una lista y modificarla.
type IteradorLista[T any] interface {

	// VerActual devuelve el elemento actual en la iteración.
	// Pre: Hay un elemento actual (se ha llamado a Siguiente al menos una vez y no se ha borrado).
	// Post: Devuelve el elemento actual. No modifica la lista.
	VerActual() T

	// HaySiguiente indica si hay un elemento siguiente para ver.
	// Post: Devuelve true si existe un elemento siguiente, false si está al final de la lista.
	HaySiguiente() bool

	// Siguiente avanza el iterador al siguiente elemento.
	// Pre: Hay un elemento siguiente (HaySiguiente() == true).
	// Post: El iterador avanzó al siguiente elemento.
	Siguiente()

	// Insertar agrega un elemento en la posición actual del iterador.
	// Pre: -
	// Post: Se insertó el elemento en la posición actual. El iterador queda posicionado en el nuevo elemento.
	// El largo de la lista aumenta en 1.
	Insertar(T)

	// Borrar remueve el elemento actual de la lista y lo devuelve.
	// Pre: Hay un elemento actual (se ha llamado a Siguiente al menos una vez y no se ha borrado).
	// Post: Se removió el elemento actual. El iterador queda posicionado en el siguiente elemento.
	// El largo de la lista disminuye en 1. Devuelve el elemento borrado.
	Borrar() T
}
