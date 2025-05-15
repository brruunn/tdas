package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

// Auxiliares de convertirAPosicion

func convertirABytes[K comparable](clave K) []byte {
	return fmt.Appendf(nil, "%v", clave)
}

func hashingFNV(clave []byte, tam int) int {
	var h uint64 = 14695981039346656037
	for _, c := range clave {
		h *= 1099511628211
		h ^= uint64(c)
	}
	return int(h % uint64(tam))
}

// Auxiliar de hashBuscar

func convertirAPosicion[K comparable](clave K, tam int) int {
	claveBytes := convertirABytes(clave)
	return hashingFNV(claveBytes, tam)
}

// Auxiliar de Guardar, Pertenece, Obtener y Borrar

func (hash *hashAbierto[K, V]) hashBuscar(clave K, seBorraPar bool) (bool, V, TDALista.Lista[parClaveValor[K, V]]) {
	pos := convertirAPosicion(clave, hash.tam)
	lista := hash.tabla[pos]

	iter := lista.Iterador()
	for iter.HaySiguiente() {
		par := iter.VerActual()
		if par.clave == clave {
			if seBorraPar {
				iter.Borrar()
				hash.cantidad--
			}
			return true, par.dato, lista
		}
		iter.Siguiente()
	}

	var ningunDato V
	return false, ningunDato, lista
}

// Auxiliar de Guardar y Borrar

func (hash *hashAbierto[K, V]) rehashear(nuevo_tam int) {
	nuevaTabla := make([]TDALista.Lista[parClaveValor[K, V]], nuevo_tam)
	for i := range nuevaTabla {
		nuevaTabla[i] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
	}

	for _, lista := range hash.tabla {
		iter := lista.Iterador()
		for iter.HaySiguiente() {
			par := iter.VerActual()
			pos := convertirAPosicion(par.clave, nuevo_tam)
			nuevaTabla[pos].InsertarUltimo(par)
			iter.Siguiente()
		}
	}

	hash.tabla = nuevaTabla
	hash.tam = nuevo_tam
}

// Auxiliar de Iterador y Siguiente

func (iter *iterHashAbierto[K, V]) buscarLista() {
	for iter.HaySiguiente() {
		lista := iter.hash.tabla[iter.posActual]
		if !lista.EstaVacia() {
			iter.actual = lista.Iterador()
			return
		}
		iter.posActual++
	}
}
