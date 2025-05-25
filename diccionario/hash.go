package diccionario

import (
	"fmt"
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

// ----------------------- FUNCIONES AUXILIARES -----------------------

// De hashing

func convertirABytes[K comparable](clave K) []byte {
	return fmt.Appendf(nil, "%v", clave)
}

func hashingFNV(clave []byte, tam int) int {
	var h uint64 = 14695981039346656037
	for _, c := range clave {
		h *= 1099511628211
		h ^= uint64(c)
	}
	return int(h % uint64(tam))
}

func convertirAPosicion[K comparable](clave K, tam int) int {
	claveBytes := convertirABytes(clave)
	return hashingFNV(claveBytes, tam)
}

// Del diccionario

func (hash *hashAbierto[K, V]) hashBuscar(clave K) iterListaPares[K, V] {
	pos := convertirAPosicion(clave, hash.tam)
	lista := hash.tabla[pos]

	iter := lista.Iterador()
	for iter.HaySiguiente() {
		par := iter.VerActual()
		if par.clave == clave {
			return iter
		}
		iter.Siguiente()
	}

	return iter
}

func (hash *hashAbierto[K, V]) rehashear(nuevoTam int) {
	nuevaTabla := crearTabla[K, V](nuevoTam)

	for _, lista := range hash.tabla {
		iter := lista.Iterador()
		for iter.HaySiguiente() {
			par := iter.VerActual()
			pos := convertirAPosicion(par.clave, nuevoTam)
			nuevaTabla[pos].InsertarUltimo(par)
			iter.Siguiente()
		}
	}

	hash.tabla = nuevaTabla
	hash.tam = nuevoTam
}

// Del iterador externo

func (iter *iterHashAbierto[K, V]) buscarLista() {
	for iter.HaySiguiente() {
		lista := iter.hash.tabla[iter.posActual]
		if !lista.EstaVacia() {
			iter.actual = lista.Iterador()
			return
		}
		iter.posActual++
	}
}

// ----------------------- FUNCIONES DE CREACIÓN ----------------------

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

// ----------------- PRIMITIVAS DEL ITERADOR EXTERNO ------------------

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
