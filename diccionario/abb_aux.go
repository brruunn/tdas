package diccionario

// Función recursiva auxiliar de Guardar y Borrar

func (n *nodoABB[K, V]) abbBuscarNodo(clave K, cmp cmp[K]) *nodoABB[K, V] {
	if n == nil {
		return nil
	}

	comparacion := cmp(clave, n.clave)
	if comparacion == 0 {
		return n
	} else if comparacion < 0 {
		return n.izq.abbBuscarNodo(clave, cmp)
	} else {
		return n.der.abbBuscarNodo(clave, cmp)
	}
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

	} else {
		a.guardar(&(*ppNodo).der, clave, dato)

	}
}

// Función recursiva auxiliar de Pertenece y Obtener

func (n *nodoABB[K, V]) abbBuscar(clave K, cmp cmp[K]) (bool, V) {
	if n == nil {
		var ningunDato V
		return false, ningunDato
	}

	comparacion := cmp(clave, n.clave)
	if comparacion == 0 {
		return true, n.dato
	}
	if comparacion < 0 {
		return n.izq.abbBuscar(clave, cmp)
	}
	return n.der.abbBuscar(clave, cmp)
}

// Funciones recursivas auxiliares de borrar

func (a *abb[K, V]) reemplazarNodo(ppPadre **nodoABB[K, V], ppActual **nodoABB[K, V], reemplazo *nodoABB[K, V]) {

	// Si ppPadre es nil, estamos en la raíz
	if ppPadre == nil {
		a.raiz = reemplazo
		return
	}

	// Si ppPadre no es nil, *ppPadre apunta a un nodo existente
	padre := *ppPadre
	if padre.izq == *ppActual {
		padre.izq = reemplazo
	} else {
		padre.der = reemplazo
	}
}

// Función recursiva de Borrar

func (a *abb[K, V]) borrar(ppPadre **nodoABB[K, V], ppActual **nodoABB[K, V], clave K) V {
	if *ppActual == nil {
		panic(_MENSAJE_PANIC_DICCIONARIO)
	}

	comparacion := a.cmp(clave, (*ppActual).clave)
	if comparacion < 0 {
		return a.borrar(ppActual, &(*ppActual).izq, clave)
	}
	if comparacion > 0 {
		return a.borrar(ppActual, &(*ppActual).der, clave)
	}

	valor := (*ppActual).dato

	if (*ppActual).izq == nil && (*ppActual).der == nil {
		a.reemplazarNodo(ppPadre, ppActual, nil)
	} else if (*ppActual).der == nil {
		a.reemplazarNodo(ppPadre, ppActual, (*ppActual).izq)
	} else if (*ppActual).izq == nil {
		a.reemplazarNodo(ppPadre, ppActual, (*ppActual).der)
	} else {
		// Caso 4: Dos hijos - Encontrar sucesor y borrarlo
		sucesor := (*ppActual).der
		for sucesor.izq != nil {
			sucesor = sucesor.izq
		}
		claveSucesor := sucesor.clave
		datoSucesor := sucesor.dato

		// Copiar datos al nodo actual
		(*ppActual).clave = claveSucesor
		(*ppActual).dato = datoSucesor

		// Borrar el sucesor del subárbol derecho
		a.borrar(ppActual, &(*ppActual).der, claveSucesor)
		return valor // Evitar decrementar cantidad dos veces
	}

	a.cantidad--
	return valor
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
