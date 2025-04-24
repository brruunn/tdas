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
	_MENSAJE_PANIC_LISTA          = "La lista esta vacia"
	_MENSAJE_PANIC_ITER_SIGUIENTE = "No hay siguiente, ya llegaste al final de la lista"
	_MENSAJE_PANIC_ITER_BORRAR    = "No hay más elementos en la lista para borrar"
)

// Primitivas de lista

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil
}

func (lista *listaEnlazada[T]) InsertarPrimero(elemento T) {
	nuevoNodo := &nodoLista[T]{dato: elemento, siguiente: lista.primero}
	lista.primero = nuevoNodo

	if lista.Largo() == 0 {
		lista.ultimo = nuevoNodo
	}

	lista.largo++
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

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic(_MENSAJE_PANIC_LISTA)
	}
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic(_MENSAJE_PANIC_LISTA)
	}
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero
	for actual != nil {
		if !visitar(actual.dato) {
			break
		}
		actual = actual.siguiente
	}
}

// Primitivas de iterador externo

func (iter *iterListaEnlazada[T]) HaySiguiente() bool {
	if iter.actual.siguiente != nil {
		return true
	}

	return false
}

func (iter *iterListaEnlazada[T]) Siguiente() {
	if !iter.HaySiguiente() {
		panic(_MENSAJE_PANIC_ITER_SIGUIENTE)
	}
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
}

func (iter *iterListaEnlazada[T]) Borrar() T {
	if iter.actual == nil {
		panic(_MENSAJE_PANIC_ITER_BORRAR)
	}

	dato := iter.actual.dato
	siguiente := iter.actual.siguiente

	if iter.actual == iter.lista.primero {
		iter.lista.primero = siguiente
	} else {
		iter.anterior.siguiente = siguiente
	}

	if siguiente == nil {
		iter.lista.ultimo = iter.anterior
	}

	if iter.lista.primero == nil {
		iter.lista.ultimo = nil
	}

	iter.actual = siguiente
	iter.lista.largo--

	return dato
}
