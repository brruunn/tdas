package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

type parClaveValor[K comparable, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K comparable, V any] struct {
	tabla    []TDALista.Lista[parClaveValor[K, V]]
	tam      int
	cantidad int
}

type iterDiccionario[K comparable, V any] struct {
	hash      *hashAbierto[K, V]
	posActual int
	actual    TDALista.IteradorLista[parClaveValor[K, V]]
}

const (
	_MENSAJE_PANIC_DICCIONARIO = "La clave no pertenece al diccionario"
	_MENSAJE_PANIC_ITER        = "El iterador termino de iterar"
)

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	// ...
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func convertirAPosicion[K comparable](clave K) int {
	claveBytes := convertirABytes(clave)
	return hashingFNV(claveBytes, len(claveBytes))
}

// --------------------------------------------------------------------------------------
// -------------------- PRIMITIVAS DEL DICCIONARIO POR TABLA DE HASH --------------------
// --------------------------------------------------------------------------------------

func (hash *hashAbierto[K, V]) Guardar(clave K, dato V) {
	// ...
}

func (hash *hashAbierto[K, V]) Pertenece(clave K) bool {
	// ...
}

func (hash *hashAbierto[K, V]) Obtener(clave K) V {
	pos := convertirAPosicion(clave) % hash.tam
	lista := hash.tabla[pos]

	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		par := iter.VerActual()
		if par.clave == clave {
			return par.dato
		}
	}
	panic(_MENSAJE_PANIC_DICCIONARIO)
}

func (hash *hashAbierto[K, V]) Borrar(clave K) V {
	// ...
}

func (hash *hashAbierto[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashAbierto[K, V]) Iterar(func(clave K, dato V) bool) {
	// ...
}

func (hash *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	// ...
}

// -------------------------------------------------------------------------
// -------------------- PRIMITIVAS DEL ITERADOR EXTERNO --------------------
// -------------------------------------------------------------------------

func (iter *iterDiccionario[K, V]) HaySiguiente() bool {
	// ...
}

func (iter *iterDiccionario[K, V]) VerActual() (K, V) {
	if iter.actual == nil || !iter.actual.HaySiguiente() {
		panic(_MENSAJE_PANIC_ITER)
	}
	par := iter.actual.VerActual()
	return par.clave, par.dato
}

func (iter *iterDiccionario[K, V]) Siguiente() {
	// ...
}
