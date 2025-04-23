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
	// ...
}

// Primitivas de lista

func (lista listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil
}

func (lista listaEnlazada[T]) InsertarPrimero(elemento T) {
	nuevoNodo := &nodoLista[T]{dato: elemento, siguiente: lista.primero}
	lista.primero = nuevoNodo
}

func (lista listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

func (lista listaEnlazada[T]) Iterar(visitar func(T) bool) {
	// ...
}

// Primitivas de iterador externo
