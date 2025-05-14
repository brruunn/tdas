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

// Función recursiva de Guardar

func (a *abb[K, V]) guardar(ppNodo **nodoABB[K, V], clave K, dato V) {
	if *ppNodo == nil {
		*ppNodo = crearNodoABB(clave, dato)
		a.cantidad++
		return
	}

	comparacion := a.cmp(clave, (*ppNodo).clave)
	if comparacion == 0 {
		(*ppNodo).dato = dato
		return
	}
	if comparacion < 0 {
		a.guardar(&(*ppNodo).izq, clave, dato)

	} else if comparacion > 0 {
		a.guardar(&(*ppNodo).der, clave, dato)

	}
}

// Función recursiva auxiliar, de Pertenece y Obtener

func (n *nodoABB[K, V]) buscar(clave K, cmp cmp[K]) (bool, V) {
	if n == nil {
		var ningunDato V
		return false, ningunDato
	}

	comparacion := cmp(clave, n.clave)
	if comparacion == 0 {
		return true, n.dato
	}
	if comparacion < 0 {
		return n.izq.buscar(clave, cmp)
	}
	return n.der.buscar(clave, cmp)
}

// Función recursiva de Iterar e IterarRango

func (n *nodoABB[K, V]) iterar(visitar func(K, V) bool, cmp cmp[K], desde *K, hasta *K) bool {
	if n == nil {
		return true
	}

	if desde == nil || cmp(n.clave, *desde) > 0 {
		if !n.izq.iterar(visitar, cmp, desde, hasta) {
			return false
		}
	}

	if (desde == nil || cmp(n.clave, *desde) >= 0) &&
		(hasta == nil || cmp(n.clave, *hasta) <= 0) {
		if !visitar(n.clave, n.dato) {
			return false
		}
	}

	if hasta == nil || cmp(n.clave, *hasta) < 0 {
		return n.der.iterar(visitar, cmp, desde, hasta)
	}

	return true
}

// Función recursiva auxiliar de Iterador, IteradorRango y Siguiente

func (iter *iterABB[K, V]) apilar(nodo *nodoABB[K, V]) {
	if nodo == nil {
		return
	}
	if (iter.desde == nil || iter.cmp(nodo.clave, *iter.desde) >= 0) &&
		(iter.hasta == nil || iter.cmp(nodo.clave, *iter.hasta) <= 0) {
		iter.pila.Apilar(nodo)
		iter.apilar(nodo.izq)

	} else if iter.hasta != nil && iter.cmp(nodo.clave, *iter.hasta) > 0 {
		iter.apilar(nodo.izq)

	} else if iter.desde != nil && iter.cmp(nodo.clave, *iter.desde) < 0 {
		iter.apilar(nodo.der)

	}
}

func (a *abb[K, V]) reemplazarNodo(ppPadre **nodoABB[K, V], ppActual **nodoABB[K, V], reemplazo *nodoABB[K, V]) {
	// si ppPadre es nil, estamos en la raiz: reasignamos a.raiz
	if ppPadre == nil {
		a.raiz = reemplazo
		return
	}
	// ppPadre no es nil, así que *ppPadre siempre apunta a un nodo existente
	padre := *ppPadre
	if padre.izq == *ppActual {
		padre.izq = reemplazo
	} else {
		padre.der = reemplazo
	}
}

func (a *abb[K, V]) borrarNodo(ppPadre **nodoABB[K, V], ppActual **nodoABB[K, V], clave K) V {
	if *ppActual == nil {
		panic(_MENSAJE_PANIC_DICCIONARIO)
	}

	comparacion := a.cmp(clave, (*ppActual).clave)
	if comparacion < 0 {
		return a.borrarNodo(ppActual, &(*ppActual).izq, clave)
	}
	if comparacion > 0 {
		return a.borrarNodo(ppActual, &(*ppActual).der, clave)
	}

	valor := (*ppActual).dato

	// Caso 1: Nodo hoja
	if (*ppActual).izq == nil && (*ppActual).der == nil {
		a.reemplazarNodo(ppPadre, ppActual, nil)
	} else if (*ppActual).izq == nil {
		// Caso 2: Solo hijo derecho
		a.reemplazarNodo(ppPadre, ppActual, (*ppActual).der)
	} else if (*ppActual).der == nil {
		// Caso 3: Solo hijo izquierdo
		a.reemplazarNodo(ppPadre, ppActual, (*ppActual).izq)
	} else {
		// Caso 4: Dos hijos
		padreSucesor := *ppActual
		sucesor := (*ppActual).der

		// Buscar el sucesor
		for sucesor.izq != nil {
			padreSucesor = sucesor
			sucesor = sucesor.izq
		}

		// Copiar datos del sucesor al nodo actual
		(*ppActual).clave = sucesor.clave
		(*ppActual).dato = sucesor.dato

		// Eliminar el sucesor (que es un nodo con 0 o 1 hijo)
		if padreSucesor == *ppActual {
			padreSucesor.der = sucesor.der
		} else {
			padreSucesor.izq = sucesor.der
		}
	}

	a.cantidad--
	return valor
}

// -------------------------------------------------------------------------------------
// -------------------- PRIMITIVAS DEL DICCIONARIO ORDENADO POR ABB --------------------
// -------------------------------------------------------------------------------------

func (a *abb[K, V]) Guardar(clave K, dato V) {
	a.guardar(&a.raiz, clave, dato)
}

func (a *abb[K, V]) Pertenece(clave K) bool {
	encontrado, _ := a.raiz.buscar(clave, a.cmp)
	return encontrado
}

func (a *abb[K, V]) Obtener(clave K) V {
	encontrado, dato := a.raiz.buscar(clave, a.cmp)
	if encontrado {
		return dato
	}
	panic(_MENSAJE_PANIC_DICCIONARIO)
}

func (a *abb[K, V]) Borrar(clave K) V {
	return a.borrarNodo(nil, &a.raiz, clave)
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
	pila := TDAPila.CrearPilaDinamica[*nodoABB[K, V]]()
	iter := iterABB[K, V]{pila: pila, cmp: a.cmp}

	iter.apilar(a.raiz)
	return &iter
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
