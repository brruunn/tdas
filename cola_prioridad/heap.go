package cola_prioridad

const (
	_MENSAJE_PANIC    = "La cola esta vacia"
	_CAP_INICIAL      = 2
	_FACT_REDIMENSION = 2
	_COND_REDUCCION   = 4
)

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

// -------------------- FUNCIONES AUXILIARES --------------------

func swap[T any](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (h *colaConPrioridad[T]) upheap(pos int) {
	for pos > 0 {
		padre := (pos - 1) / 2
		if h.cmp(h.datos[pos], h.datos[padre]) <= 0 {
			return
		}
		swap(h.datos, pos, padre)
		pos = padre
	}
}

func downheap[T any](arr []T, limite, pos int, funcCmp func(T, T) int) {
	for pos < limite {
		hIzq, hDer := 2*pos+1, 2*pos+2

		if hIzq < limite {
			if hDer < limite {
				if funcCmp(arr[pos], arr[hIzq]) >= 0 && funcCmp(arr[pos], arr[hDer]) >= 0 {
					return
				}
				if funcCmp(arr[hIzq], arr[hDer]) < 0 {
					swap(arr, pos, hDer)
					pos = hDer
					continue
				}

			} else if funcCmp(arr[pos], arr[hIzq]) >= 0 {
				return

			}

			swap(arr, pos, hIzq)
			pos = hIzq
			continue
		}

		return
	}
}

func (h *colaConPrioridad[T]) redimensionar(nuevaCap int) {
	nuevoArr := make([]T, nuevaCap)
	copy(nuevoArr, h.datos)
	h.datos = nuevoArr
}

func heapify[T any](arr []T, limite int, funcCmp func(T, T) int) {
	for i := limite - 1; i >= 0; i-- {
		downheap(arr, limite, i, funcCmp)
	}
}

// -------------------- FUNCIONES PARA EL USUARIO --------------------

func CrearHeap[T any](funcCmp func(T, T) int) ColaPrioridad[T] {
	return &colaConPrioridad[T]{datos: make([]T, _CAP_INICIAL), cmp: funcCmp}
}

func CrearHeapArr[T any](arr []T, funcCmp func(T, T) int) ColaPrioridad[T] {
	h := &colaConPrioridad[T]{
		datos: make([]T, len(arr)+_CAP_INICIAL),
		cant:  len(arr),
		cmp:   funcCmp,
	}
	copy(h.datos, arr)
	heapify(h.datos, h.cant, h.cmp)
	return h
}

func HeapSort[T any](elementos []T, funcCmp func(T, T) int) {
	largo := len(elementos)
	heapify(elementos, largo, funcCmp)
	for largo > 1 {
		slice := elementos[:largo]
		largo--
		swap(slice, 0, largo)
		downheap(slice, len(slice)-1, 0, funcCmp)
	}
}

// -------------------- PRIMITIVAS DE LA COLA DE PRIORIDAD POR HEAP --------------------

func (h *colaConPrioridad[T]) EstaVacia() bool {
	return h.cant == 0
}

func (h *colaConPrioridad[T]) Encolar(elem T) {
	if h.cant == len(h.datos) {
		h.redimensionar(len(h.datos) * _FACT_REDIMENSION)
	}
	h.datos[h.cant] = elem
	h.upheap(h.cant)
	h.cant++
}

func (h *colaConPrioridad[T]) VerMax() T {
	if h.EstaVacia() {
		panic(_MENSAJE_PANIC)
	}
	return h.datos[0]
}

func (h *colaConPrioridad[T]) Desencolar() T {
	dato := h.VerMax()
	swap(h.datos, 0, h.cant-1)
	h.cant--
	downheap(h.datos, h.cant, 0, h.cmp)
	if (h.cant*_COND_REDUCCION) > _CAP_INICIAL && (h.cant*_COND_REDUCCION) <= len(h.datos) {
		h.redimensionar(len(h.datos) / _FACT_REDIMENSION)
	}
	return dato
}

func (h *colaConPrioridad[T]) Cantidad() int {
	return h.cant
}
