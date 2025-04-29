package diccionario

import (
	"fmt"
)

type comparable struct{} // Creo que falta definir un método o función para compararlos

/*

Para un hash cerrado:

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado // Usar constante
}

---------------------

Para un hash abierto:

type parClaveValor[K comparable, V any] struct {
	clave K
	dato  V
}

*/

type hash[K comparable, V any] struct {
	/*

		Para un hash cerrado (hashCerrado):

		tabla    []celdaHash[K, V]
		cantidad int
		tam      int
		borrados int

		--------------------

		Para un hash abierto (hashAbierto):

		tabla    []TDALista.Lista[parClaveValor[K, V]]
		tam      int
		cantidad int

	*/
}

type iterDiccionario[K comparable, V any] struct {
	// ...
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

/*

Algunas funciones de hashing no-criptográficas para usar (hay muchas más):

---------------------------

Sumar todos los caracteres:

func hashSuma(clave []byte, largo int) int {
	suma := 0
	for _, c := range clave {
		suma += int(c)
	}
	return suma % largo
}

------------

FNV (Fowler-Noll-Vo) Hashing:

func fnvHashing(clave []byte, largo int) int {
	h := 14695981039346656037
	n = strlen(clave)
	for _, c := range clave:
		h *= 1099511628211
		h ^= int(c)
	return h % largo
}

---------------

Jenkins One-at-a-time Hashing:

func jenkinsHash(clave []byte, largo int) int {
	var hash uint32
	for i := range clave {
		hash += uint32(clave[i])
		hash += (hash << 10)
		hash ^= (hash >> 6)
	}
	hash += (hash << 3)
	hash ^= (hash >> 11)
	hash += (hash << 15)
	return int(hash % uint32(largo))
}

*/

// --------------------------------------------------------------------------------------
// -------------------- PRIMITIVAS DEL DICCIONARIO POR TABLA DE HASH --------------------
// --------------------------------------------------------------------------------------

func (hash *hash[K, V]) Guardar(clave K, dato V) {
	// ...
}

func (hash *hash[K, V]) Pertenece(clave K) bool {
	// ...
}

func (hash *hash[K, V]) Obtener(clave K) V {
	// ...
}

func (hash *hash[K, V]) Borrar(clave K) V {
	// ...
}

func (hash *hash[K, V]) Cantidad() int {
	// ...
}

func (hash *hash[K, V]) Iterar(func(clave K, dato V) bool) {
	// ...
}

func (hash *hash[K, V]) Iterador() IterDiccionario[K, V] {
	// ...
}

// -------------------------------------------------------------------------
// -------------------- PRIMITIVAS DEL ITERADOR EXTERNO --------------------
// -------------------------------------------------------------------------

func (iter *iterDiccionario[K, V]) HaySiguiente() bool {
	// ...
}

func (iter *iterDiccionario[K, V]) VerActual() (K, V) {
	// ...
}

func (iter *iterDiccionario[K, V]) Siguiente() {
	// ...
}
