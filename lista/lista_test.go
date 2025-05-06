package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	_MENSAJE_PANIC_LISTA = "La lista esta vacia"
	_MENSAJE_PANIC_ITER  = "El iterador termino de iterar"
)

// --------------------------------------------------------------------
// -------------------- TESTS DE LA LISTA ENLAZADA --------------------
// --------------------------------------------------------------------

// Test para verificar el comportamiento de una lista vacía
func TestListaEstaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())

	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.VerUltimo() })
	require.Panics(t, func() { lista.BorrarPrimero() })
}

// Test para verificar la inserción y verificación del primer elemento
func TestInsertarPrimeroYVer(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(10)
	require.Equal(t, 10, lista.VerPrimero())

	lista.InsertarPrimero(20)
	require.Equal(t, 20, lista.VerPrimero())
	require.Equal(t, 10, lista.VerUltimo())
}

// Test para verificar la inserción y verificación del último elemento
func TestInsertarUltimoYVer(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(10)
	require.Equal(t, 10, lista.VerUltimo())

	lista.InsertarUltimo(20)
	require.Equal(t, 20, lista.VerUltimo())
	require.Equal(t, 10, lista.VerPrimero())
}

func TestInsertarPrimeroUltimoYVer(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for n := range 5 {
		lista.InsertarPrimero(5 - n) // De la mitad para abajo
		lista.InsertarUltimo(6 + n)  // De la mitad para arriba
	}

	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 10, lista.VerUltimo())
}

// Test para verificar el borrado del primer elemento
func TestBorrarPrimeroLista(t *testing.T) {
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
	require.PanicsWithValue(t, _MENSAJE_PANIC_LISTA, func() { lista.VerUltimo() })
	require.PanicsWithValue(t, _MENSAJE_PANIC_LISTA, func() { lista.BorrarPrimero() })
}

// Test de volumen para verificar el comportamiento con muchos elementos
func TestPruebaDeVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	n := 10000

	for i := range n {
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
	arr := []int{0, 0, 1, 3, 5, 246, 7, -246, 100, -100, 13, 15}

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

func TestIteradorExternoItera(t *testing.T) {
	arr := []int{5, 10, 15, 20, 25}
	lista := TDALista.CrearListaEnlazada[int]()
	for _, v := range arr {
		lista.InsertarUltimo(v)
	}

	iter := lista.Iterador()
	var resultado []int
	for iter.HaySiguiente() {
		resultado = append(resultado, iter.VerActual())
		iter.Siguiente()
	}

	// Comprobar que se recorrió como se esperaba
	require.Equal(t, arr, resultado)
}

// Test para verificar el método VerActual del iterador
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
	require.PanicsWithValue(t, _MENSAJE_PANIC_ITER, func() { iter.Siguiente() })
	require.PanicsWithValue(t, _MENSAJE_PANIC_ITER, func() { iter.Borrar() })
}

// Test para verificar el método HaySiguiente del iterador
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

// Test para verificar el método Siguiente del iterador
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

// Test para verificar la inserción al principio usando el iterador
func TestIteradorInsertarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	iter := lista.Iterador() // El iter apunta a nil.

	iter.Insertar("Primero") // El iter apunta a "Primero", y éste, a nil.
	require.Equal(t, lista.VerPrimero(), "Primero")

	iter.Insertar("Anterior a Primero") // El iter apunta a "Anterior a Primero", y éste, a "Primero".
	require.Equal(t, lista.VerPrimero(), "Anterior a Primero")
}

// Test para verificar la inserción en medio de la lista usando el iterador
func TestIteradorInsertarEnMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	arr := []string{"Primero", "Cuarto"}

	for _, str := range arr {
		lista.InsertarUltimo(str)
	}

	iter := lista.Iterador() // El iter apunta a "Primero".

	iter.Siguiente()         // El iter apunta a "Cuarto".
	iter.Insertar("Segundo") // "Primero" y el iter apuntan a "Segundo", y éste, a "Cuarto".
	require.Equal(t, iter.VerActual(), "Segundo")

	iter.Siguiente()         // El iter vuelve a apuntar a "Cuarto".
	iter.Insertar("Tercero") // "Segundo" y el iter apuntan a "Tercero", y éste, a "Cuarto".
	require.Equal(t, iter.VerActual(), "Tercero")
}

// Test para verificar la inserción al final usando el iterador
func TestIteradorInsertarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	arr := []string{"Primero", "Segundo", "Tercero"}

	for _, str := range arr {
		lista.InsertarUltimo(str)
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

// Test para borrar el primer elemento usando el iterador externo
func TestIteradorBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(100)
	lista.InsertarUltimo(200)

	iter := lista.Iterador()
	require.Equal(t, 100, iter.Borrar()) // Borra el primero

	require.Equal(t, 200, lista.VerPrimero()) // Verifica que 200 ahora es el primero
	require.Equal(t, 1, lista.Largo())
}

// Test para borrar un elemento del medio usando el iterador externo
func TestIteradorBorrarEnMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(100)
	lista.InsertarUltimo(200)
	lista.InsertarUltimo(300)

	iter := lista.Iterador()
	iter.Siguiente() // Avanzamos al segundo elemento (200)

	require.Equal(t, 200, iter.Borrar())    // Borra el del medio
	require.Equal(t, 300, iter.VerActual()) // Verifica que ahora esta en 300

	require.Equal(t, 100, lista.VerPrimero()) // El primero sigue siendo 100
	require.Equal(t, 300, lista.VerUltimo())  // El ultimo sigue siendo 300
	require.Equal(t, 2, lista.Largo())
}

// Test para borrar el ultimo elemento usando el iterador externo
func TestIteradorBorrarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(100)
	lista.InsertarUltimo(200)

	iter := lista.Iterador()
	iter.Siguiente() // Avanzamos al ultimo elemento

	require.Equal(t, 200, iter.Borrar())  // Borra el ultimo
	require.False(t, iter.HaySiguiente()) // Verifica que no hay mas elementos

	require.Equal(t, 100, lista.VerPrimero()) // El primero sigue siendo 100
	require.Equal(t, 100, lista.VerUltimo())  // Ahora 100 tambien es el ultimo
	require.Equal(t, 1, lista.Largo())
}

func TestIteradorInsertarIntercalado(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()

	iter.Insertar(1) // [1]
	iter.Siguiente() // apuntamos a nil
	iter.Insertar(2) // [1,2]

	iter = lista.Iterador()
	iter.Insertar(0) // [0,1,2]

	for iter.HaySiguiente() {
		iter.Siguiente()
	}

	iter.Insertar(3) // [0,1,2,3]

	require.Equal(t, 0, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())
	require.Equal(t, 4, lista.Largo())
}
