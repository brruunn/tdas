package pila

const (
	_MENSAJE_PANIC    = "La pila esta vacia"
	_CAP_INICIAL      = 2
	_FACT_REDIMENSION = 2
	_COND_REDUCCION   = 4
)

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int // Cantidad de elementos almacenados.
}

func CrearPilaDinamica[T any]() Pila[T] {
	datos := make([]T, _CAP_INICIAL)
	return &pilaDinamica[T]{datos: datos}
}

func (p *pilaDinamica[T]) redimensionar(tam int) {
	nuevosDatos := make([]T, tam)
	copy(nuevosDatos, p.datos)
	p.datos = nuevosDatos
}

/// ------------------------------------------------------------------------
/// -------------------- PRIMITIVAS DE LA PILA DINÁMICA --------------------
/// ------------------------------------------------------------------------

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic(_MENSAJE_PANIC)
	}
	return p.datos[(p.cantidad - 1)]
}

func (p *pilaDinamica[T]) Apilar(elemento T) {
	if p.cantidad == len(p.datos) {
		p.redimensionar(len(p.datos) * _FACT_REDIMENSION)
	}
	p.datos[p.cantidad] = elemento
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	tope := p.VerTope()
	p.cantidad--
	if (p.cantidad*_COND_REDUCCION) > _CAP_INICIAL && (p.cantidad*_COND_REDUCCION) <= len(p.datos) {
		p.redimensionar(len(p.datos) / _FACT_REDIMENSION)
	}
	return tope
}
