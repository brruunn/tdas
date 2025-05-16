package diccionario

import (
	TDAPila "tdas/pila"
)

type cmp[K comparable] func(c1, c2 K) int

type nodoABB[K comparable, V any] struct {
	izq   *nodoABB[K, V]
	der   *nodoABB[K, V]
	clave K
	dato  V
}

type abb[K comparable, V any] struct {
	raiz     *nodoABB[K, V]
	cantidad int
	cmp      cmp[K]
}

type iterABB[K comparable, V any] struct {
	pila  TDAPila.Pila[*nodoABB[K, V]]
	desde *K
	hasta *K
	cmp   cmp[K]
}

func CrearABB[K comparable, V any](cmp cmp[K]) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{cmp: cmp}
}

func crearNodoABB[K comparable, V any](clave K, dato V) *nodoABB[K, V] {
	return &nodoABB[K, V]{clave: clave, dato: dato}
}

// -------------------------------------------------------------------------------------
// -------------------- PRIMITIVAS DEL DICCIONARIO ORDENADO POR ABB --------------------
// -------------------------------------------------------------------------------------

func (a *abb[K, V]) Guardar(clave K, dato V) {
	_, nodo := a.abbBuscar(clave, nil, &a.raiz)
	if *nodo != nil {
		(*nodo).dato = dato
	} else {
		*nodo = crearNodoABB(clave, dato)
		a.cantidad++
	}
}

func (a *abb[K, V]) Pertenece(clave K) bool {
	_, nodo := a.abbBuscar(clave, nil, &a.raiz)
	return *nodo != nil
}

func (a *abb[K, V]) Obtener(clave K) V {
	_, nodo := a.abbBuscar(clave, nil, &a.raiz)
	if *nodo != nil {
		return (*nodo).dato
	}
	panic(_MENSAJE_PANIC_DICCIONARIO)
}

func (a *abb[K, V]) Borrar(clave K) V {
	padre, nodo := a.abbBuscar(clave, nil, &a.raiz)

	if *nodo != nil {
		dato := (*nodo).dato

		if (*nodo).izq == nil && (*nodo).der == nil {
			a.reemplazarNodo(padre, nodo, nil)

		} else if (*nodo).izq != nil && (*nodo).der == nil {
			a.reemplazarNodo(padre, nodo, (*nodo).izq)

		} else if (*nodo).izq == nil && (*nodo).der != nil {
			a.reemplazarNodo(padre, nodo, (*nodo).der)

		} else {
			sucesorPadre, sucesor := a.encontrarSucesor(nodo, &(*nodo).der)
			(*nodo).clave = (*sucesor).clave
			(*nodo).dato = (*sucesor).dato
			a.reemplazarNodo(sucesorPadre, sucesor, (*sucesor).der)

		}

		a.cantidad--
		return dato
	}

	panic(_MENSAJE_PANIC_DICCIONARIO)
}

func (a *abb[K, V]) Cantidad() int {
	return a.cantidad
}

// Iteradores internos

func (a *abb[K, V]) Iterar(visitar func(K, V) bool) {
	a.raiz.iterar(visitar, a.cmp, nil, nil)
}

func (a *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(K, V) bool) {
	a.raiz.iterar(visitar, a.cmp, desde, hasta)
}

// Iteradores externos

func (a *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return a.IteradorRango(nil, nil)
}

func (a *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodoABB[K, V]]()
	iter := iterABB[K, V]{pila: pila, desde: desde, hasta: hasta, cmp: a.cmp}

	iter.apilar(a.raiz)
	return &iter
}

// -------------------------------------------------------------------------
// -------------------- PRIMITIVAS DEL ITERADOR EXTERNO --------------------
// -------------------------------------------------------------------------

func (iter *iterABB[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()
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
	iter.apilar(nodo.der)
}
