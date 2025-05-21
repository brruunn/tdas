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

func CrearHeap[T any](funcCmp func(T, T) int) ColaPrioridad[T] {
	return &colaConPrioridad[T]{
		datos: make([]T, _CAP_INICIAL),
		cmp:   funcCmp,
	}
}

func CrearHeapArr[T any](arr []T, funcCmp func(T, T) int) ColaPrioridad[T] {
	h := &colaConPrioridad[T]{
		datos: make([]T, len(arr)),
		cant:  len(arr),
		cmp:   funcCmp,
	}
	copy(h.datos, arr)

	for i := h.cant/2 - 1; i >= 0; i-- {
		h.downheap(i)
	}

	return h
}

func HeapSort[T any](elementos []T, funcCmp func(T, T) int) {
	h := CrearHeapArr(elementos, funcCmp)
	ultPos := len(elementos) - 1
	for ultPos >= 0 {
		max := h.Desencolar()
		elementos[ultPos] = max
		ultPos--
	}
}

func (h *colaConPrioridad[T]) swap(i, j int) {
	h.datos[i], h.datos[j] = h.datos[j], h.datos[i]
}

func (h *colaConPrioridad[T]) upheap(pos int) {
	for pos > 0 {
		padre := (pos - 1) / 2
		if h.cmp(h.datos[pos], h.datos[padre]) <= 0 {
			return
		}
		h.swap(pos, padre)
		pos = padre
	}
}

func (h *colaConPrioridad[T]) downheap(pos int) {
	for pos < h.cant {
		hIzq, hDer := 2*pos+1, 2*pos+2

		if hIzq < h.cant {
			if hDer < h.cant {
				if h.cmp(h.datos[pos], h.datos[hIzq]) >= 0 && h.cmp(h.datos[pos], h.datos[hDer]) >= 0 {
					return
				}
				if h.cmp(h.datos[hIzq], h.datos[hDer]) < 0 {
					h.swap(pos, hDer)
					pos = hDer
					continue
				}

			} else if h.cmp(h.datos[pos], h.datos[hIzq]) >= 0 {
				return

			}

			h.swap(pos, hIzq)
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

// -------------------------------------------------------------------------------------
// -------------------- PRIMITIVAS DE LA COLA DE PRIORIDAD POR HEAP --------------------
// -------------------------------------------------------------------------------------

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
	h.swap(0, h.cant-1)
	h.cant--
	h.downheap(0)
	if (h.cant*_COND_REDUCCION) > _CAP_INICIAL && (h.cant*_COND_REDUCCION) <= len(h.datos) {
		h.redimensionar(len(h.datos) / _FACT_REDIMENSION)
	}
	return dato
}

func (h *colaConPrioridad[T]) Cantidad() int {
	return h.cant
}
