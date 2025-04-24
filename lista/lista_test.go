package lista_test

import (
	"testing"
	TDALista "tp_lista/lista"

	"github.com/stretchr/testify/require"
)

// --------------------------------------------------------------------
// -------------------- TESTS DE LA LISTA ENLAZADA --------------------
// --------------------------------------------------------------------

func TestListaEstaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())

	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.VerUltimo() })
	require.Panics(t, func() { lista.BorrarPrimero() })
}

func TestInsertarYVerPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(10)
	require.Equal(t, 10, lista.VerPrimero())

	lista.InsertarPrimero(20)
	require.Equal(t, 20, lista.VerPrimero())

}
func TestPruebaDeVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	n := 10000

	for i := 0; i < n; i++ {
		lista.InsertarPrimero(i)
		require.Equal(t, i, lista.VerPrimero())
	}

	for i := n - 1; i >= 0; i-- {
		require.Equal(t, i, lista.BorrarPrimero())
	}

	require.True(t, lista.EstaVacia())
	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.BorrarPrimero() })
}

// --------------------------------------------------------------------
// -------------------- TESTS DEL ITERADOR EXTERNO --------------------
// --------------------------------------------------------------------

func TestHaySiguiente(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)

	iter := lista.Iterador()
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}

func TestSiguiente(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarUltimo("A")
	lista.InsertarUltimo("B")

	iter := lista.Iterador()
	require.Equal(t, "A", iter.VerActual())

	iter.Siguiente()
	require.Equal(t, "B", iter.VerActual())

	iter.Siguiente()
	require.Panics(t, func() { iter.VerActual() })
}

func TestBorrar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(100)
	lista.InsertarUltimo(200)
	lista.InsertarUltimo(300)

	iter := lista.Iterador()

	require.Equal(t, 100, iter.VerActual())
	require.Equal(t, 100, iter.Borrar()) // borra el primero

	require.Equal(t, 200, iter.VerActual())
	require.Equal(t, 200, iter.Borrar()) // borra el del medio

	require.Equal(t, 300, iter.VerActual())
	require.Equal(t, 300, iter.Borrar()) // borra el ultimo

	require.True(t, lista.EstaVacia())
	require.False(t, iter.HaySiguiente())
}

// --------------------------------------------------------------------
// -------------------- TESTS DEL ITERADOR INTERNO --------------------
// --------------------------------------------------------------------

func TestIterar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	elementos := []int{10, 20, 30}

	for _, elem := range elementos {
		lista.InsertarUltimo(elem)
	}

	iter := lista.Iterador()
	i := 0
	for iter.HaySiguiente() {
		require.Equal(t, elementos[i], iter.VerActual())
		iter.Siguiente()
		i++
	}

	require.Equal(t, len(elementos), i)
}
