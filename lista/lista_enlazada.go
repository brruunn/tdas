package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iterListaEnlazada[T any] struct {
	anterior *nodoLista[T]
	actual   *nodoLista[T]
	lista    *listaEnlazada[T]
}

const (
	_MENSAJE_PANIC = "La lista esta vacia"
)

// Primitivas de lista

func (lista listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil
}

func (lista *listaEnlazada[T]) InsertarPrimero(elemento T) {
	nuevoNodo := &nodoLista[T]{dato: elemento, siguiente: lista.primero}
	lista.primero = nuevoNodo
}

func (lista *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nuevoNodo := &nodoLista[T]{dato: elemento}

	if lista.EstaVacia() {
		lista.primero = nuevoNodo
	} else {
		lista.ultimo.siguiente = nuevoNodo
	}

	lista.ultimo = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	primero := lista.VerPrimero()

	if lista.largo == 1 {
		lista.ultimo = nil
	}

	lista.primero = lista.primero.siguiente
	lista.largo--
	return primero
}

func (lista listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic(_MENSAJE_PANIC)
	}
	return lista.primero.dato
}

func (lista listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic(_MENSAJE_PANIC)
	}
	return lista.ultimo.dato
}

func (lista listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista listaEnlazada[T]) Iterar(visitar func(T) bool) {
	// ...
}

// Primitivas de iterador externo
