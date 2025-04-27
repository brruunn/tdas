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
	_MENSAJE_PANIC_LISTA = "La lista esta vacia"
	_MENSAJE_PANIC_ITER  = "El iterador termino de iterar"
)

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func nodoCrear[T any](dato T) *nodoLista[T] {
	return &nodoLista[T]{dato: dato}
}

// -------------------------------------------------------------------------
// -------------------- PRIMITIVAS DE LA LISTA ENLAZADA --------------------
// -------------------------------------------------------------------------

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil
}

func (lista *listaEnlazada[T]) InsertarPrimero(elemento T) {
	nuevoNodo := nodoCrear(elemento)
	nuevoNodo.siguiente = lista.primero

	if lista.EstaVacia() {
		lista.ultimo = nuevoNodo
	}

	lista.primero = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nuevoNodo := nodoCrear(elemento)

	if lista.EstaVacia() {
		lista.primero = nuevoNodo
	} else {
		lista.ultimo.siguiente = nuevoNodo
	}

	lista.ultimo = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic(_MENSAJE_PANIC_LISTA)
	}

	dato := lista.primero.dato
	lista.primero = lista.primero.siguiente
	lista.largo--

	if lista.EstaVacia() {
		lista.ultimo = nil
	}

	return dato
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

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{actual: lista.primero, lista: lista}
}

// -------------------------------------------------------------------------
// -------------------- PRIMITIVAS DEL ITERADOR EXTERNO --------------------
// -------------------------------------------------------------------------

func (iter *iterListaEnlazada[T]) VerActual() T {
	if !iter.HaySiguiente() {
		panic(_MENSAJE_PANIC_ITER)
	}
	return iter.actual.dato
}

func (iter *iterListaEnlazada[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iterListaEnlazada[T]) Siguiente() {
	if !iter.HaySiguiente() {
		panic(_MENSAJE_PANIC_ITER)
	}
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
}

func (iter *iterListaEnlazada[T]) Insertar(elemento T) {
	if iter.actual == iter.lista.primero {
		iter.lista.InsertarPrimero(elemento)
		iter.actual = iter.lista.primero
	} else {
		nuevoNodo := nodoCrear(elemento)
		nuevoNodo.siguiente = iter.actual

		iter.anterior.siguiente, iter.actual = nuevoNodo, nuevoNodo

		if iter.anterior == iter.lista.ultimo {
			iter.lista.ultimo = nuevoNodo
		}

		iter.lista.largo++
	}
}

func (iter *iterListaEnlazada[T]) Borrar() T {
	if !iter.HaySiguiente() {
		panic(_MENSAJE_PANIC_ITER)
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

	iter.actual = siguiente
	iter.lista.largo--

	return dato
}
