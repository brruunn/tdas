package cola_prioridad

const (
	_MENSAJE_PANIC    = "La cola esta vacia"
	_CAP_INICIAL      = 2
	_FACT_REDIMENSION = 2
	_COND_REDUCCION   = 4
)

type funcCmp[T any] func(T, T) int

type heap[T any] struct {
	datos []T
	cant  int
	cmp   funcCmp[T]
}

// -------------------- FUNCIONES AUXILIARES --------------------

func swap[T any](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (heap *heap[T]) upheap(pos int) {
	for pos > 0 {
		padre := (pos - 1) / 2
		if heap.cmp(heap.datos[pos], heap.datos[padre]) <= 0 {
			return
		}
		swap(heap.datos, pos, padre)
		pos = padre
	}
}

func downheapRec[T any](arr []T, limite, pos, hijoIzq, hijoDer int, cmp funcCmp[T]) {
	if hijoIzq >= limite {
		return
	}

	if hijoDer < limite {
		if cmp(arr[pos], arr[hijoIzq]) >= 0 && cmp(arr[pos], arr[hijoDer]) >= 0 {
			return
		} else if cmp(arr[hijoIzq], arr[hijoDer]) < 0 {
			swap(arr, pos, hijoDer)
			pos = hijoDer
			downheapRec(arr, limite, pos, 2*pos+1, 2*pos+2, cmp)
		}

	} else if cmp(arr[pos], arr[hijoIzq]) >= 0 {
		return
	}

	swap(arr, pos, hijoIzq)
	pos = hijoIzq
	downheapRec(arr, limite, pos, 2*pos+1, 2*pos+2, cmp)
}

func downheap[T any](arr []T, limite, pos int, cmp funcCmp[T]) {
	downheapRec(arr, limite, pos, 2*pos+1, 2*pos+2, cmp)
}

func (heap *heap[T]) redimensionar(nuevaCap int) {
	nuevoArr := make([]T, nuevaCap)
	copy(nuevoArr, heap.datos)
	heap.datos = nuevoArr
}

func heapify[T any](arr []T, limite int, cmp funcCmp[T]) {
	for i := limite - 1; i >= 0; i-- {
		downheap(arr, limite, i, cmp)
	}
}

// -------------------- FUNCIONES PARA EL USUARIO --------------------

func CrearHeap[T any](cmp funcCmp[T]) ColaPrioridad[T] {
	return &heap[T]{datos: make([]T, _CAP_INICIAL), cmp: cmp}
}

func CrearHeapArr[T any](arr []T, cmp funcCmp[T]) ColaPrioridad[T] {
	heap := &heap[T]{
		datos: make([]T, len(arr)+_CAP_INICIAL),
		cant:  len(arr),
		cmp:   cmp,
	}
	copy(heap.datos, arr)
	heapify(heap.datos, heap.cant, heap.cmp)
	return heap
}

func HeapSort[T any](elementos []T, cmp funcCmp[T]) {
	largo := len(elementos)
	heapify(elementos, largo, cmp)
	for largo > 1 {
		slice := elementos[:largo]
		largo--
		swap(slice, 0, largo)
		downheap(slice, len(slice)-1, 0, cmp)
	}
}

// -------------------- PRIMITIVAS DE LA COLA DE PRIORIDAD --------------------

func (heap *heap[T]) EstaVacia() bool {
	return heap.cant == 0
}

func (heap *heap[T]) Encolar(elemento T) {
	if heap.cant == len(heap.datos) {
		heap.redimensionar(len(heap.datos) * _FACT_REDIMENSION)
	}
	heap.datos[heap.cant] = elemento
	heap.upheap(heap.cant)
	heap.cant++
}

func (heap *heap[T]) VerMax() T {
	if heap.EstaVacia() {
		panic(_MENSAJE_PANIC)
	}
	return heap.datos[0]
}

func (heap *heap[T]) Desencolar() T {
	elemento := heap.VerMax()
	swap(heap.datos, 0, heap.cant-1)
	heap.cant--
	downheap(heap.datos, heap.cant, 0, heap.cmp)
	if (heap.cant*_COND_REDUCCION) > _CAP_INICIAL && (heap.cant*_COND_REDUCCION) <= len(heap.datos) {
		heap.redimensionar(len(heap.datos) / _FACT_REDIMENSION)
	}
	return elemento
}

func (heap *heap[T]) Cantidad() int {
	return heap.cant
}
