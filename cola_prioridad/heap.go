package cola_prioridad

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	// ...
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	h := &colaConPrioridad[T]{
		datos: make([]T, len(arreglo)),
		cmp:   funcion_cmp,
	}
	copy(h.datos, arreglo)

	// heapify de abajo hacia arriba
	for i := len(h.datos)/2 - 1; i >= 0; i-- {
		h.downheap(i)
	}

	return h
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	// ...
}

// ------------------------------------------
// Métodos internos (privados)
// ------------------------------------------

func upheap[T any](datos []T, pos int, cmp func(T, T) int) {
	for pos > 0 {
		padre := (pos - 1) / 2
		if cmp(datos[pos], datos[padre]) <= 0 {
			break
		}
		swap(datos, pos, padre)
		pos = padre
	}
}

func (h *colaConPrioridad[T]) downheap(indice int) {
	// ...
}

func swap[T any](datos []T, i, j int) {
	datos[i], datos[j] = datos[j], datos[i]
}

func (h *colaConPrioridad[T]) EstaVacia() bool {
	return len(h.datos) == 0
}

func (h *colaConPrioridad[T]) Encolar(elem T) {
	h.datos = append(h.datos, elem)
	upheap(h.datos, len(h.datos)-1, h.cmp)
}

func (h *colaConPrioridad[T]) VerMax() T {
	if h.EstaVacia() {
		panic(_MENSAJE_PANIC_HEAP)
	}
	return h.datos[0]
}

func (h *colaConPrioridad[T]) Desencolar() T {
	// ...
}

func (h *colaConPrioridad[T]) Cantidad() int {
	return len(h.datos)
}
