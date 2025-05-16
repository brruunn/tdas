package diccionario

import (
	TDALista "tdas/lista"
)

type parClaveValor[K comparable, V any] struct {
	clave K
	dato  V
}

type listaPares[K comparable, V any] TDALista.Lista[*parClaveValor[K, V]]

type hashAbierto[K comparable, V any] struct {
	tabla    []listaPares[K, V]
	tam      int
	cantidad int
}

type iterHashAbierto[K comparable, V any] struct {
	hash      *hashAbierto[K, V]
	posActual int
	actual    TDALista.IteradorLista[*parClaveValor[K, V]]
}

const (
	_MENSAJE_PANIC_DICCIONARIO = "La clave no pertenece al diccionario"
	_MENSAJE_PANIC_ITER        = "El iterador termino de iterar"
	_TAM_INICIAL               = 7
	_MAX_FACTOR_DE_CARGA       = 3.0
	_MIN_FACTOR_DE_CARGA       = 2.0
	_FACTOR_REDIMENSION        = 2
)

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	tabla := make([]listaPares[K, V], _TAM_INICIAL)
	for i := range tabla {
		tabla[i] = TDALista.CrearListaEnlazada[*parClaveValor[K, V]]()
	}
	return &hashAbierto[K, V]{tabla: tabla, tam: _TAM_INICIAL}
}

func crearPar[K comparable, V any](clave K, dato V) *parClaveValor[K, V] {
	return &parClaveValor[K, V]{clave, dato}
}

// -------------------- PRIMITIVAS DEL DICCIONARIO --------------------

func (hash *hashAbierto[K, V]) Guardar(clave K, dato V) {
	_, lista := hash.hashBuscar(clave, true)
	lista.InsertarUltimo(crearPar(clave, dato))
	hash.cantidad++

	if float32(hash.cantidad)/float32(hash.tam) >= _MAX_FACTOR_DE_CARGA {
		hash.rehashear(hash.tam * _FACTOR_REDIMENSION)
	}
}

func (hash *hashAbierto[K, V]) Pertenece(clave K) bool {
	par, _ := hash.hashBuscar(clave, false)
	return par != nil
}

func (hash *hashAbierto[K, V]) Obtener(clave K) V {
	par, _ := hash.hashBuscar(clave, false)
	if par != nil {
		return par.dato
	}
	panic(_MENSAJE_PANIC_DICCIONARIO)
}

func (hash *hashAbierto[K, V]) Borrar(clave K) V {
	par, _ := hash.hashBuscar(clave, true)
	if par != nil {
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
