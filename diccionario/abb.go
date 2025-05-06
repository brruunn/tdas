package diccionario

import (
	TDAPila "tdas/pila"
)

type funcCmp[K comparable] func(K, K) int

type nodoABB[K comparable, V any] struct {
	izq   *nodoABB[K, V]
	der   *nodoABB[K, V]
	clave K
	dato  V
}

type abb[K comparable, V any] struct {
	raiz     *nodoABB[K, V]
	cantidad int
	cmp      funcCmp[K]
}

type iterDiccoonario[K comparable, V any] struct {
	pila  TDAPila.Pila[*nodoABB[K, V]]
	desde *K
	hasta *K
	cmp   funcCmp[K]
}

func CrearABB[K comparable, V any](cmp funcCmp[K]) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{
		cmp: cmp,
	}
}

func (n *nodoABB[K, V]) iterar(visitar func(K, V) bool) bool {
	//...
}

//--------------------------------PRIMITIVAS DEL DICCIONARIO ORDENADO--------------------------------------------------//
//---------------------------------------------------------------------------------------------------------------------//

func (a *abb[K, V]) Guardar(clave K, dato V) {
	// ...
}

func (a *abb[K, V]) Obtener(clave K) V {
	// ...
}

func (a *abb[K, V]) Pertenece(clave K) bool {
	// ...
}

func (a *abb[K, V]) Borrar(clave K) V {
	// ...
}

func (a *abb[K, V]) Cantidad() int {
	// ...
}

func (a *abb[K, V]) Iterar(visitar func(K, V) bool) {
	// ...
}

func (a *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(K, V) bool) {
	// ...
}

//--------------------------------- PRIMITIVAS ITERADOR EXTERNO ----------------------------------------------------------------//
//------------------------------------------------------------------------------------------------------------------------------//

func (a *abb[K, V]) Iterador() IterDiccionario[K, V] {
	// ...
}

func (a *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	// ...
}
