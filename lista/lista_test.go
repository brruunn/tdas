package lista_test

import (
	"testing"
	TDALista "tp_lista/lista"

	"github.com/stretchr/testify/require"
)

const (
	_MENSAJE_PANIC_LISTA = "La lista esta vacia"
	_MENSAJE_PANIC_ITER  = "El iterador termino de iterar"
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

func TestInsertaryVerUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(10)
	require.Equal(t, 10, lista.VerUltimo())

	lista.InsertarUltimo(20)
	require.Equal(t, 20, lista.VerUltimo())
}

func TestBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for n := range 10 {
		lista.InsertarUltimo(n)
	}

	for n := range 9 {
		require.Equal(t, lista.VerPrimero(), lista.BorrarPrimero()) // En efecto, se borra el primero.
		require.Equal(t, lista.VerPrimero(), n+1)                   // El primero, pasa a ser el siguiente.
		require.Equal(t, lista.Largo(), 9-n)                        // El largo va disminuyendo.
	}

	require.Equal(t, lista.BorrarPrimero(), 9)
	require.PanicsWithValue(t, _MENSAJE_PANIC_LISTA, func() { lista.VerPrimero() })
	require.Equal(t, lista.Largo(), 0)
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

// TESTS DEL ITERADOR INTERNO

// Itera y usa todos los elementos.
func TestSumarTodos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	arr := []int{0, 10, 20, 30, 40, -50}

	for _, n := range arr {
		lista.InsertarUltimo(n)
	}

	var suma int
	lista.Iterar(func(n int) bool {
		suma += n
		return true
	})

	require.Equal(t, suma, 50)
}

// Itera todos los elementos, y usa algunos.
func TestSumarPares(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	arr := []int{0, 10, 15, 17, 20, 21, 29, -30, 50, -53}

	for _, n := range arr {
		lista.InsertarUltimo(n)
	}

	var suma int
	lista.Iterar(func(n int) bool {
		if n%2 == 0 {
			suma += n
		}
		return true
	})

	require.Equal(t, suma, 50)
}

// Itera y usa todos los elementos, hasta una condición de corte.
func TestSumarTodosHastaSiete(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	arr := []int{0, 0, 1, 1, 2, 7, -4}

	for _, n := range arr {
		lista.InsertarUltimo(n)
	}

	var suma int
	lista.Iterar(func(n int) bool {
		if n != 7 {
			suma += n
			return true
		}
		return false
	})

	require.Equal(t, suma, 4)
}

// Itera todos los elementos, y usa algunos, hasta una condición de corte.
func TestSumarPrimerosCincoPares(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	arr := []int{0, 0, 1, 3, 5, 246, 7, -246, 100, 11, 13, 15}

	for _, n := range arr {
		lista.InsertarUltimo(n)
	}

	var suma, contador int
	lista.Iterar(func(n int) bool {
		if contador < 5 {
			if n%2 == 0 {
				suma += n
				contador++
			}
			return true
		}
		return false
	})

	require.Equal(t, suma, 100)
}

// --------------------------------------------------------------------
// -------------------- TESTS DEL ITERADOR EXTERNO --------------------
// --------------------------------------------------------------------

func TestVerActual(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for n := range 10 {
		lista.InsertarUltimo(n)
	}

	iter := lista.Iterador()
	var num int

	for iter.HaySiguiente() {
		require.Equal(t, iter.VerActual(), num)
		iter.Siguiente()
		num++
	}

	require.PanicsWithValue(t, _MENSAJE_PANIC_ITER, func() { iter.VerActual() })
}

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

func TestInsertarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	iter := lista.Iterador() // El iter apunta a nil.

	iter.Insertar("Primero") // El iter apunta a "Primero", y éste, a nil.
	require.Equal(t, lista.VerPrimero(), "Primero")

	iter.Insertar("Anterior a Primero") // El iter apunta a "Anterior a Primero", y éste, a "Primero".
	require.Equal(t, lista.VerPrimero(), "Anterior a Primero")
}

func TestInsertarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	arr := []string{"Primero", "Segundo", "Tercero"}

	for _, str := range arr {
		lista.InsertarPrimero(str)
	}

	iter := lista.Iterador() // El iter apunta a "Primero".
	for iter.HaySiguiente() {
		iter.Siguiente()
	} // Al final, el iter apunta a nil.

	iter.Insertar("Cuarto") // "Tercero" y el iter apuntan a "Cuarto", y éste, a nil.
	require.Equal(t, lista.VerUltimo(), "Cuarto")

	iter.Siguiente()                    // El iter vuelve a apuntar a nil.
	iter.Insertar("Siguiente a Cuarto") // "Cuarto" y el iter apuntan a "Siguiente a Cuarto", y éste, a nil.
	require.Equal(t, lista.VerUltimo(), "Siguiente a Cuarto")
}

func TestInsertarEnElMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	arr := []string{"Primero", "Cuarto"}

	for _, str := range arr {
		lista.InsertarPrimero(str)
	}

	iter := lista.Iterador() // El iter apunta a "Primero".

	iter.Siguiente()         // El iter apunta a "Cuarto".
	iter.Insertar("Segundo") // "Primero" y el iter apuntan a "Segundo", y éste, a "Cuarto".
	require.Equal(t, iter.VerActual(), "Segundo")

	iter.Siguiente()         // El iter vuelve a apuntar a "Cuarto".
	iter.Insertar("Tercero") // "Segundo" y el iter apuntan a "Tercero", y éste, a "Cuarto".
	require.Equal(t, iter.VerActual(), "Tercero")
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
