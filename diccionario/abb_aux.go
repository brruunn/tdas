package diccionario

// Función recursiva auxiliar de Guardar, Pertenece, Obtener y Borrar

func (a *abb[K, V]) abbBuscar(clave K, actual **nodoABB[K, V]) **nodoABB[K, V] {
	if *actual == nil {
		return actual
	}

	comparacion := a.cmp(clave, (*actual).clave)
	if comparacion == 0 {
		return actual
	}
	if comparacion < 0 {
		return a.abbBuscar(clave, &(*actual).izq)
	}
	return a.abbBuscar(clave, &(*actual).der)
}

// Función auxiliar de Borrar

func (a *abb[K, V]) buscarMinimo(actual **nodoABB[K, V]) **nodoABB[K, V] {
	if (*actual).izq == nil {
		return actual
	}
	return a.buscarMinimo(&(*actual).izq)
}

// Función recursiva de Iterar e IterarRango

func (n *nodoABB[K, V]) iterar(visitar func(K, V) bool, cmp cmp[K], desde *K, hasta *K) bool {
	if n == nil {
		return true
	}

	if desde == nil || cmp(n.clave, *desde) >= 0 {
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

	if hasta == nil || cmp(n.clave, *hasta) <= 0 {
		return n.der.iterar(visitar, cmp, desde, hasta)
	}

	return true
}

// Función recursiva auxiliar de Iterador, IteradorRango y Siguiente

func (iter *iterABB[K, V]) apilar(actual *nodoABB[K, V]) {
	if actual == nil {
		return
	}
	if (iter.desde == nil || iter.cmp(actual.clave, *iter.desde) >= 0) &&
		(iter.hasta == nil || iter.cmp(actual.clave, *iter.hasta) <= 0) {
		iter.pila.Apilar(actual)
		iter.apilar(actual.izq)

	} else if iter.hasta != nil && iter.cmp(actual.clave, *iter.hasta) > 0 {
		iter.apilar(actual.izq)

	} else if iter.desde != nil && iter.cmp(actual.clave, *iter.desde) < 0 {
		iter.apilar(actual.der)

	}
}
