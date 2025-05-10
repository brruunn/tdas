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

type iterABB[K comparable, V any] struct {
	pila  TDAPila.Pila[*nodoABB[K, V]]
	desde *K
	hasta *K
	cmp   funcCmp[K]
}

func CrearABB[K comparable, V any](cmp funcCmp[K]) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{cmp: cmp}
}

func nodoABBCrear[K comparable, V any](clave K, dato V) *nodoABB[K, V] {
	return &nodoABB[K, V]{clave: clave, dato: dato}
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
	nodo := a.raiz
	for nodo != nil {
		comparacion := a.cmp(clave, nodo.clave)
		if comparacion == 0 {
			return nodo.dato // Encontré la clave que estaba buscando
		}
		if comparacion < 0 {
			nodo = nodo.izq
		}
		if comparacion > 0 {
			nodo = nodo.der
		}
	}
	panic(_MENSAJE_PANIC_DICCIONARIO)
}

func (a *abb[K, V]) Pertenece(clave K) bool {
	// ...
}

func (a *abb[K, V]) Borrar(clave K) V {
	nodo := a.raiz
	for nodo != nil {
		comparacion := a.cmp(clave, nodo.clave)
		if comparacion == 0 {
			// Caso 1: nodo sin hijos
			if nodo.izq == nil && nodo.der == nil {
				// ...
			}
			// Caso 2: nodo con un solo hijo
			if (nodo.izq == nil && nodo.der != nil) || (nodo.izq != nil && nodo.der == nil) {
				// ...
			}
			// Caso 3: nodo con dos hijos
			if nodo.izq != nil && nodo.der != nil {
				// ...
			}
			break
		}
		if comparacion < 0 {
			nodo = nodo.izq
		}
		if comparacion > 0 {
			nodo = nodo.der
		}
	}
	panic(_MENSAJE_PANIC_DICCIONARIO)
}

func (a *abb[K, V]) Cantidad() int {
	return a.cantidad
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
	pila := TDAPila.CrearPilaDinamica[*nodoABB[K, V]]()
	nodo := a.raiz

	for nodo != nil {
		pila.Apilar(nodo)
		nodo = nodo.izq
	}

	return &iterABB[K, V]{pila: pila, cmp: a.cmp}
}

func (a *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	// ...
}

func (iter *iterABB[K, V]) HaySiguiente() bool {
	// ...
}

func (iter *iterABB[K, V]) VerActual() (K, V) {
	// ...
}

func (iter *iterABB[K, V]) Siguiente() {
	// ...
}
