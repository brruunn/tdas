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

func crearNodoABB[K comparable, V any](clave K, dato V) *nodoABB[K, V] {
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
	var padre *nodoABB[K, V]
	nodo := a.raiz
	var direccion *(*nodoABB[K, V]) // Puntero al puntero del nodo en el padre

	// Buscar el nodo a borrar
	for nodo != nil {
		comparacion := a.cmp(clave, nodo.clave)
		if comparacion == 0 {
			// Caso 1: nodo sin hijos
			if nodo.izq == nil && nodo.der == nil {
				if padre == nil {
					a.raiz = nil
				}
				if padre != nil && direccion == &padre.izq {
					padre.izq = nil
				}
				if padre != nil && direccion == &padre.der {
					padre.der = nil
				}
			}

			// Caso 2: nodo con un solo hijo (izquierdo)
			if nodo.izq != nil && nodo.der == nil {
				if padre == nil {
					a.raiz = nodo.izq
				}
				if padre != nil && direccion == &padre.izq {
					padre.izq = nodo.izq
				}
				if padre != nil && direccion == &padre.der {
					padre.der = nodo.izq
				}
			}

			// Caso 2: nodo con un solo hijo (derecho)
			if nodo.izq == nil && nodo.der != nil {
				if padre == nil {
					a.raiz = nodo.der
				}
				if padre != nil && direccion == &padre.izq {
					padre.izq = nodo.der
				}
				if padre != nil && direccion == &padre.der {
					padre.der = nodo.der
				}
			}

			// Caso 3: nodo con dos hijos
			if nodo.izq != nil && nodo.der != nil {
				// buscar sucesor inorder (menor del subárbol derecho)
				sucesorPadre := nodo
				sucesor := nodo.der

				for sucesor.izq != nil {
					sucesorPadre = sucesor
					sucesor = sucesor.izq
				}

				// copiar datos del sucesor
				nodo.clave = sucesor.clave
				nodo.dato = sucesor.dato

				// eliminar el sucesor
				if sucesorPadre == nodo {
					sucesorPadre.der = sucesor.der
				}
				if sucesorPadre != nodo {
					sucesorPadre.izq = sucesor.der
				}
			}

			a.cantidad--
			return nodo.dato
		}

		// Continuar búsqueda
		padre = nodo
		if comparacion < 0 {
			direccion = &nodo.izq
			nodo = nodo.izq
		}
		if comparacion > 0 {
			direccion = &nodo.der
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
	pila := TDAPila.CrearPilaDinamica[*nodoABB[K, V]]()
	nodo := a.raiz

	for nodo != nil {
		if a.cmp(nodo.clave, *desde) < 0 {
			nodo = nodo.der
		} else if a.cmp(nodo.clave, *hasta) > 0 {
			nodo = nodo.izq
		} else {
			pila.Apilar(nodo)
			nodo = nodo.izq
		}
	}

	return &iterABB[K, V]{pila: pila, cmp: a.cmp, desde: desde, hasta: hasta}
}

func (iter *iterABB[K, V]) HaySiguiente() bool {
	// ...
}

func (iter *iterABB[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic(_MENSAJE_PANIC_ITER)
	}
	tope := iter.pila.VerTope()
	return tope.clave, tope.dato
}

func (iter *iterABB[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic(_MENSAJE_PANIC_ITER)
	}

	nodo := iter.pila.Desapilar()

	actual := nodo.der
	for actual != nil {
		iter.pila.Apilar(actual)
		actual = actual.izq
	}
}
