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
	// ...
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	// ...
}

// ------------------------------------------
// Métodos internos (privados)
// ------------------------------------------

func (h *colaConPrioridad[T]) downheap(indice int) {
	// ...
}

func (h *colaConPrioridad[T]) swap(i, j int) {
	// ...
}

// ------------------------------------------
// Implementación de las primitivas del heap
// ------------------------------------------

func (h *colaConPrioridad[T]) EstaVacia() bool {
	// ...
}

func (h *colaConPrioridad[T]) Encolar(elemento T) {
	// ...
}

func (h *colaConPrioridad[T]) VerMax() T {
	// ...
}

func (h *colaConPrioridad[T]) Desencolar() T {
	// ...
}

func (h *colaConPrioridad[T]) Cantidad() int {
	// ...
}
