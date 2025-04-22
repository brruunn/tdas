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

func (lista listaEnlazada[T]) Iterar(visitar func(T) bool) {
	// ...
}

// Primitivas de iterador externo
