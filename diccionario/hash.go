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
	_TAM_INICIAL               = 7   // Debe ser un número primo
	_MAX_FACTOR_DE_CARGA       = 2.5 // Debe estar entre 2 y 3
	_MIN_FACTOR_DE_CARGA       = 1.0
	_FACTOR_REDIMENSION        = 2
)

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	tabla := make([]TDALista.Lista[parClaveValor[K, V]], _TAM_INICIAL)
	// Se inicializan listas vacias, antes eran todas nil y por ende, producia un panic
	for i := range tabla {
		tabla[i] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
	}
	return &hashAbierto[K, V]{tabla: tabla, tam: _TAM_INICIAL}
}

func crearPar[K comparable, V any](clave K, dato V) parClaveValor[K, V] {
	return parClaveValor[K, V]{clave, dato}
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func convertirAPosicion[K comparable](clave K, tam int) int {
	claveBytes := convertirABytes(clave)
	return hashingFNV(claveBytes, tam)
}

func (hash *hashAbierto[K, V]) rehashear(nuevo_tam int) {
	nuevaTabla := make([]TDALista.Lista[parClaveValor[K, V]], nuevo_tam)
	// Lo mismo, antes eran punteros a nil. Ahora están vacías al rehashear
	for i := range nuevaTabla {
		nuevaTabla[i] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
	}

	for iter := hash.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, dato := iter.VerActual()
		pos := convertirAPosicion(clave, nuevo_tam)
		par := crearPar(clave, dato)
		nuevaTabla[pos].InsertarUltimo(par)
	}

	hash.tabla = nuevaTabla
	hash.tam = nuevo_tam
}

// --------------------------------------------------------------------------------------
// -------------------- PRIMITIVAS DEL DICCIONARIO POR TABLA DE HASH --------------------
// --------------------------------------------------------------------------------------

func (hash *hashAbierto[K, V]) Guardar(clave K, dato V) {
	pos := convertirAPosicion(clave, hash.tam)
	lista := hash.tabla[pos]

	if hash.Pertenece(clave) {
		_ = hash.Borrar(clave)
		par := crearPar(clave, dato) // Simplifiqué lógica y código: Antes se iteraba hasta  par.clave == clave y con la direccion de memoria &par se cambiaba el dato
		lista.InsertarUltimo(par)    // Ahora: Usa el mismo hash.Borrar(clave) para eliminar y luego inserto de una, en vez de usar punteros.
		hash.cantidad++

	} else {
		par := crearPar(clave, dato)
		lista.InsertarUltimo(par)
		hash.cantidad++

		if float32(hash.cantidad)/float32(hash.tam) >= _MAX_FACTOR_DE_CARGA {
			hash.rehashear(hash.tam * _FACTOR_REDIMENSION)
		}
	}
}

func (hash *hashAbierto[K, V]) Pertenece(clave K) bool {
	var (
		pos       = convertirAPosicion(clave, hash.tam)
		lista     = hash.tabla[pos]
		pertenece = false
	)

	lista.Iterar(func(par parClaveValor[K, V]) bool {
		if par.clave == clave {
			pertenece = true
			return false
		}
		return true
	})

	return pertenece
}

func (hash *hashAbierto[K, V]) Obtener(clave K) V {
	pos := convertirAPosicion(clave, hash.tam)
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
	pos := convertirAPosicion(clave, hash.tam)
	lista := hash.tabla[pos]

	if lista == nil || lista.EstaVacia() {
		panic(_MENSAJE_PANIC_DICCIONARIO)
	}

	iter := lista.Iterador()
	for iter.HaySiguiente() {
		par := iter.VerActual()
		if par.clave == clave {
			dato := par.dato
			iter.Borrar()
			hash.cantidad--
			if float32(hash.cantidad)/float32(hash.tam) <= _MIN_FACTOR_DE_CARGA && hash.tam > _TAM_INICIAL {
				hash.rehashear(hash.tam / _FACTOR_REDIMENSION)
			}
			return dato
		}
		iter.Siguiente()
	}

	panic(_MENSAJE_PANIC_DICCIONARIO)
}

func (hash *hashAbierto[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashAbierto[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for _, lista := range hash.tabla {
		iter := lista.Iterador()
		for iter.HaySiguiente() {
			par := iter.VerActual()
			if !visitar(par.clave, par.dato) {
				return
			}
			iter.Siguiente()
		}
	}
}

func (hash *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	iter := iterDiccionario[K, V]{hash: hash, posActual: 0}

	for iter.posActual < hash.tam {
		lista := hash.tabla[iter.posActual]
		if lista != nil && !lista.EstaVacia() {
			iter.actual = lista.Iterador()
			return &iter
		}
		iter.posActual++
	}

	return &iter
}

// -------------------------------------------------------------------------
// -------------------- PRIMITIVAS DEL ITERADOR EXTERNO --------------------
// -------------------------------------------------------------------------

func (iter *iterDiccionario[K, V]) HaySiguiente() bool {
	return iter.posActual != iter.hash.tam
}

func (iter *iterDiccionario[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic(_MENSAJE_PANIC_ITER)
	}
	par := iter.actual.VerActual()
	return par.clave, par.dato
}

func (iter *iterDiccionario[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic(_MENSAJE_PANIC_ITER)
	}

	iter.actual.Siguiente()
	if iter.actual.HaySiguiente() {
		return
	}

	iter.posActual++
	for iter.posActual < iter.hash.tam {
		lista := iter.hash.tabla[iter.posActual]
		if lista != nil && !lista.EstaVacia() {
			iter.actual = lista.Iterador()
			return
		}
		iter.posActual++
	}
}
