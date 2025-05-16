package diccionario

import (
	"fmt"
)

// -------------------- AUXILIARES HASHING --------------------

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

func convertirAPosicion[K comparable](clave K, tam int) int {
	claveBytes := convertirABytes(clave)
	return hashingFNV(claveBytes, tam)
}

// -------------------- AUXILIARES DICCIONARIO --------------------

func (hash *hashAbierto[K, V]) hashBuscar(clave K, seBorraPar bool) (*parClaveValor[K, V], listaPares[K, V]) {
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
			return par, lista
		}
		iter.Siguiente()
	}

	return nil, lista
}

func (hash *hashAbierto[K, V]) rehashear(nuevoTam int) {
	nuevaTabla := crearTabla[K, V](nuevoTam)

	for _, lista := range hash.tabla {
		iter := lista.Iterador()
		for iter.HaySiguiente() {
			par := iter.VerActual()
			pos := convertirAPosicion(par.clave, nuevoTam)
			nuevaTabla[pos].InsertarUltimo(par)
			iter.Siguiente()
		}
	}

	hash.tabla = nuevaTabla
	hash.tam = nuevoTam
}

// -------------------- AUXILIARES ITERADOR --------------------

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
