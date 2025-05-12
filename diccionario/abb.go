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

func (n *nodoABB[K, V]) buscar(clave K, cmp funcCmp[K]) (bool, V) {
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

func (n *nodoABB[K, V]) iterar(visitar func(K, V) bool) bool {
	if n == nil {
		return true
	}

	if !n.izq.iterar(visitar) {
		return false
	}

	if !visitar(n.clave, n.dato) {
		return false
	}

	return n.der.iterar(visitar)
}

func (n *nodoABB[K, V]) iterarRango(visitar func(K, V) bool, cmp funcCmp[K], desde *K, hasta *K) bool {
	if n == nil {
		return true
	}

	if desde == nil || cmp(n.clave, *desde) > 0 {
		if !n.izq.iterarRango(visitar, cmp, desde, hasta) {
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
		return n.der.iterarRango(visitar, cmp, desde, hasta)
	}

	return true
}

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

//--------------------------------PRIMITIVAS DEL DICCIONARIO ORDENADO--------------------------------------------------//
//---------------------------------------------------------------------------------------------------------------------//

func (a *abb[K, V]) Guardar(clave K, dato V) {
	a.guardar(&a.raiz, clave, dato) // Para modificar el nodo, paso un puntero a su puntero
}

func (a *abb[K, V]) Pertenece(clave K) bool {
	nodo := a.raiz
	encontrado, _ := nodo.buscar(clave, a.cmp)
	return encontrado
}

func (a *abb[K, V]) Obtener(clave K) V {
	nodo := a.raiz
	encontrado, dato := nodo.buscar(clave, a.cmp)

	if encontrado {
		return dato
	}
	panic(_MENSAJE_PANIC_DICCIONARIO)
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

// Iteradores internos

func (a *abb[K, V]) Iterar(visitar func(K, V) bool) {
	a.raiz.iterar(visitar)
}

func (a *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(K, V) bool) {
	a.raiz.iterarRango(visitar, a.cmp, desde, hasta)
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

//--------------------------------- PRIMITIVAS ITERADOR EXTERNO ----------------------------------------------------------------//
//------------------------------------------------------------------------------------------------------------------------------//

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
