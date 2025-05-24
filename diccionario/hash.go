package diccionario

import (
	TDALista "tdas/lista"
)

type (
	listaPares[K comparable, V any]     TDALista.Lista[*parClaveValor[K, V]]
	iterListaPares[K comparable, V any] TDALista.IteradorLista[*parClaveValor[K, V]]
)

type parClaveValor[K comparable, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K comparable, V any] struct {
	tabla    []listaPares[K, V]
	tam      int
	cantidad int
}

type iterHashAbierto[K comparable, V any] struct {
	hash      *hashAbierto[K, V]
	posActual int
	actual    iterListaPares[K, V]
}

const (
	_MENSAJE_PANIC_DICCIONARIO = "La clave no pertenece al diccionario"
	_MENSAJE_PANIC_ITER        = "El iterador termino de iterar"
	_TAM_INICIAL               = 7
	_MAX_FACTOR_DE_CARGA       = 3.0
	_MIN_FACTOR_DE_CARGA       = 2.0
	_FACTOR_REDIMENSION        = 2
)

func crearTabla[K comparable, V any](tam int) []listaPares[K, V] {
	tabla := make([]listaPares[K, V], tam)
	for i := range tabla {
		tabla[i] = TDALista.CrearListaEnlazada[*parClaveValor[K, V]]()
	}
	return tabla
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	tabla := crearTabla[K, V](_TAM_INICIAL)
	return &hashAbierto[K, V]{tabla: tabla, tam: _TAM_INICIAL}
}

func crearPar[K comparable, V any](clave K, dato V) *parClaveValor[K, V] {
	return &parClaveValor[K, V]{clave, dato}
}

// -------------------- PRIMITIVAS DEL DICCIONARIO --------------------

func (hash *hashAbierto[K, V]) Guardar(clave K, dato V) {
	iter := hash.hashBuscar(clave)

	if iter.HaySiguiente() {
		iter.Borrar()
		hash.cantidad--
	}

	iter.Insertar(crearPar(clave, dato))
	hash.cantidad++

	if float32(hash.cantidad)/float32(hash.tam) >= _MAX_FACTOR_DE_CARGA {
		hash.rehashear(hash.tam * _FACTOR_REDIMENSION)
	}
}

func (hash *hashAbierto[K, V]) Pertenece(clave K) bool {
	iter := hash.hashBuscar(clave)
	return iter.HaySiguiente()
}

func (hash *hashAbierto[K, V]) Obtener(clave K) V {
	iter := hash.hashBuscar(clave)
	if iter.HaySiguiente() {
		return iter.VerActual().dato
	}
	panic(_MENSAJE_PANIC_DICCIONARIO)
}

func (hash *hashAbierto[K, V]) Borrar(clave K) V {
	iter := hash.hashBuscar(clave)
	if iter.HaySiguiente() {
		par := iter.Borrar()

		hash.cantidad--
		if float32(hash.cantidad)/float32(hash.tam) <= _MIN_FACTOR_DE_CARGA && hash.tam > _TAM_INICIAL {
			hash.rehashear(hash.tam / _FACTOR_REDIMENSION)
		}

		return par.dato
	}
	panic(_MENSAJE_PANIC_DICCIONARIO)
}

func (hash *hashAbierto[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashAbierto[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for _, lista := range hash.tabla {
		iteraProxLista := true

		lista.Iterar(func(par *parClaveValor[K, V]) bool {
			if !visitar(par.clave, par.dato) {
				iteraProxLista = false
				return false
			}
			return true
		})

		if !iteraProxLista {
			return
		}
	}
}

func (hash *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	iter := iterHashAbierto[K, V]{hash: hash}
	iter.buscarLista()
	return &iter
}

// -------------------- PRIMITIVAS DEL ITERADOR --------------------

func (iter *iterHashAbierto[K, V]) HaySiguiente() bool {
	return iter.posActual != iter.hash.tam
}

func (iter *iterHashAbierto[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic(_MENSAJE_PANIC_ITER)
	}
	par := iter.actual.VerActual()
	return par.clave, par.dato
}

func (iter *iterHashAbierto[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic(_MENSAJE_PANIC_ITER)
	}

	iter.actual.Siguiente()
	if iter.actual.HaySiguiente() {
		return
	}

	iter.posActual++
	iter.buscarLista()
}
