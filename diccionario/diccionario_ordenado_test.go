package diccionario_test

import (
	"fmt"
	"math/rand/v2"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

// --------------------------------------------------------------------------------
// -------------------- TESTS DEL DICCIONARIO ORDENADO POR ABB --------------------
// --------------------------------------------------------------------------------

// ---------------------------------- CREAR ABBS ----------------------------------

func TestDiccionarioOrdenadoVacio(t *testing.T) {
	t.Log("Un diccionario ordenado vacío no tiene claves")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestDiccionarioOrdenadoClaveDefault(t *testing.T) {
	t.Log("Un diccionario ordenado vacío, por default, no se guarda con claves")
	dicStr := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, dicStr.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicStr.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicStr.Borrar("") })

	cmpInt := func(a, b int) int { return a - b }
	dicInt := TDADiccionario.CrearABB[int, string](cmpInt)
	require.False(t, dicInt.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicInt.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicInt.Borrar(0) })
}

// -------------------------------- GUARDAR NODOS ---------------------------------

func TestDiccionarioOrdenadoUnElemento(t *testing.T) {
	t.Log("Se puede guardar, al menos, una clave correctamente")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestDiccionarioOrdenadoGuardar(t *testing.T) {
	t.Log("Se puede guardar más de una clave, correctamente")
	claves := []string{"Gato", "Perro", "Vaca"}
	valores := []string{"miau", "guau", "moo"}

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)

	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	dic.Guardar(claves[1], valores[1])
	require.EqualValues(t, 2, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.EqualValues(t, 3, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestDiccionarioOrdenadoReemplazoDato(t *testing.T) {
	t.Log("Se puede reemplazar el dato de una clave ya existente")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave := "Gato"
	dic.Guardar(clave, "miau")
	require.EqualValues(t, 1, dic.Cantidad())
	dic.Guardar(clave, "miu")
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
}

func TestDiccionarioOrdenadoClaveVacia(t *testing.T) {
	t.Log("Se pueden guardar una clave y un valor vacíos")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar("", "")
	require.True(t, dic.Pertenece(""))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, "", dic.Obtener(""))
}

func TestDiccionarioOrdenadoValorNulo(t *testing.T) {
	t.Log("nil es un valor válido para guardar")
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	dic.Guardar("Pez", nil)
	require.True(t, dic.Pertenece("Pez"))
	require.EqualValues(t, 1, dic.Cantidad())
	require.Nil(t, dic.Obtener("Pez"))
}

// --------------------------------- BORRAR NODOS ---------------------------------

func TestDiccionarioOrdenadoBorrar(t *testing.T) {
	t.Log("Se puede borrar más de un valor, correctamente")
	claves := []string{"Gato", "Perro", "Vaca"}
	valores := []string{"miau", "guau", "moo"}
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	for i := range claves {
		dic.Guardar(claves[i], valores[i])
	}

	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.False(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 2, dic.Cantidad())

	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.False(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, 1, dic.Cantidad())

	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.False(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 0, dic.Cantidad())
}

func TestBorrarRaizSinHijos(t *testing.T) {
	t.Log("Al borrar una raíz sin hijos, el diccionario ordenado queda vacío")
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })

	dic.Guardar(10, 10)

	require.Equal(t, 10, dic.Borrar(10))
	require.Equal(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(10))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(10) })
}

func TestBorrarNodoSinHijos(t *testing.T) {
	t.Log("Al borrar un nodo sin hijos, su padre, se queda sin ése hijo")
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })

	dic.Guardar(10, 10)
	dic.Guardar(5, 5)

	require.Equal(t, 5, dic.Borrar(5))
	require.Equal(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(10))
	require.False(t, dic.Pertenece(5))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(5) })
}

func TestBorrarRaizConHijoIzquierdo(t *testing.T) {
	t.Log("Al borrar una raíz con hijo izquierdo, éste pasa a ser la raíz")
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })

	dic.Guardar(10, 10)
	dic.Guardar(5, 5)

	require.Equal(t, 10, dic.Borrar(10))
	require.Equal(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(5))
	require.False(t, dic.Pertenece(10))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(10) })
}

func TestBorrarNodoConHijoIzquierdo(t *testing.T) {
	t.Log("Al borrar un nodo con hijo izquierdo, su padre se queda con el nieto")
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })

	dic.Guardar(10, 10)
	dic.Guardar(5, 5)
	dic.Guardar(3, 3)

	require.Equal(t, 5, dic.Borrar(5))
	require.Equal(t, 2, dic.Cantidad())
	require.True(t, dic.Pertenece(10))
	require.True(t, dic.Pertenece(3))
	require.False(t, dic.Pertenece(5))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(5) })
}

func TestBorrarRaizConHijoDerecho(t *testing.T) {
	t.Log("Al borrar una raíz con hijo derecho, éste pasa a ser la raíz")
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })

	dic.Guardar(10, 10)
	dic.Guardar(20, 20)

	require.Equal(t, 10, dic.Borrar(10))
	require.Equal(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(20))
	require.False(t, dic.Pertenece(10))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(10) })
}

func TestBorrarNodoConHijoDerecho(t *testing.T) {
	t.Log("Al borrar un nodo con hijo derecho, su padre se queda con el nieto")
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })

	dic.Guardar(10, 10)
	dic.Guardar(20, 20)
	dic.Guardar(25, 25)

	require.Equal(t, 20, dic.Borrar(20))
	require.Equal(t, 2, dic.Cantidad())
	require.True(t, dic.Pertenece(10))
	require.True(t, dic.Pertenece(25))
	require.False(t, dic.Pertenece(20))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(20) })
}

func TestBorraRaizDosHijos(t *testing.T) {
	t.Log("Al borrar una raíz con dos hijos, uno de los dos la reemplaza")
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })

	dic.Guardar(10, 10)
	dic.Guardar(5, 5)
	dic.Guardar(15, 15)

	require.Equal(t, 10, dic.Borrar(10))
	require.Equal(t, 2, dic.Cantidad())
	require.True(t, dic.Pertenece(5))
	require.True(t, dic.Pertenece(15))
	require.False(t, dic.Pertenece(10))
}

func TestBorraNodoDosHijos(t *testing.T) {
	t.Log("Al borrar un nodo con dos hijos, uno de los dos lo reemplaza")
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })

	dic.Guardar(10, 10)
	dic.Guardar(5, 5)
	dic.Guardar(15, 15)
	dic.Guardar(13, 13)
	dic.Guardar(17, 17)

	require.Equal(t, 15, dic.Borrar(15))
	require.Equal(t, 4, dic.Cantidad())
	require.True(t, dic.Pertenece(13))
	require.True(t, dic.Pertenece(17))
	require.False(t, dic.Pertenece(15))
}

func TestBorrarNodoDosHijosSucesorProfundo(t *testing.T) {
	t.Log("Se puede borrar un nodo con dos hijos, cuyo reemplazo esté más oculto")
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })

	dic.Guardar(10, 10)
	dic.Guardar(5, 5)
	dic.Guardar(15, 15)
	dic.Guardar(20, 20)
	dic.Guardar(19, 19)
	dic.Guardar(18, 18)
	dic.Guardar(17, 17)
	dic.Guardar(16, 16)

	require.Equal(t, 15, dic.Borrar(15))
	require.Equal(t, 7, dic.Cantidad())
	require.True(t, dic.Pertenece(16))
	require.True(t, dic.Pertenece(20))
	require.False(t, dic.Pertenece(15))
}

func TestBorrarNodoDosHijosSucesorProfundoConHijo(t *testing.T) {
	t.Log("Se puede borrar un nodo con dos hijos, de reemplazo oculto, sin que afecte a su hijo")
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })

	dic.Guardar(10, 10)
	dic.Guardar(5, 5)
	dic.Guardar(15, 15)
	dic.Guardar(20, 20)
	dic.Guardar(19, 19)
	dic.Guardar(18, 18)
	dic.Guardar(16, 16)
	dic.Guardar(17, 17)

	require.Equal(t, 15, dic.Borrar(15))
	require.Equal(t, 7, dic.Cantidad())
	require.True(t, dic.Pertenece(16))
	require.True(t, dic.Pertenece(17))
	require.False(t, dic.Pertenece(15))
}

// ---------------------------- GUARDAR Y BORRAR NODOS ----------------------------

func TestDiccionarioOrdenadoConClavesNumericas(t *testing.T) {
	cmp := func(a, b int) int { return a - b }
	dic := TDADiccionario.CrearABB[int, string](cmp)
	clave := 10
	dic.Guardar(clave, "Gatito")
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, "Gatito", dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestDiccionarioOrdenadoConClavesStructs(t *testing.T) {
	type basico struct {
		a string
		b int
	}
	type avanzado struct {
		w int
		x basico
		y basico
		z string
	}

	cmp := func(a, b avanzado) int {
		if a.w != b.w {
			return a.w - b.w
		}
		if a.x.a != b.x.a {
			return strings.Compare(a.x.a, b.x.a)
		}
		if a.x.b != b.x.b {
			return a.x.b - b.x.b
		}
		if a.y.a != b.y.a {
			return strings.Compare(a.y.a, b.y.a)
		}
		if a.y.b != b.y.b {
			return a.y.b - b.y.b
		}
		return strings.Compare(a.z, b.z)
	}

	dic := TDADiccionario.CrearABB[avanzado, int](cmp)
	a1 := avanzado{w: 10, z: "hola", x: basico{"mundo", 8}, y: basico{"!", 10}}
	dic.Guardar(a1, 0)
	require.True(t, dic.Pertenece(a1))
	require.EqualValues(t, 0, dic.Borrar(a1))
	require.False(t, dic.Pertenece(a1))
}

func TestDiccionarioOrdenadoReutlizacionDeBorrados(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave := "hola"
	dic.Guardar(clave, "mundo!")
	dic.Borrar(clave)
	dic.Guardar(clave, "mundooo!")
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, "mundooo!", dic.Obtener(clave))
}

func TestDiccionarioOrdenadoGuardarYBorrarRepetidasVeces(t *testing.T) {
	cmp := func(a, b int) int { return a - b }
	dic := TDADiccionario.CrearABB[int, int](cmp)
	limite := 1000

	for i := 0; i < limite; i++ {
		n := rand.IntN(limite)
		dic.Guardar(n, i)
		dic.Borrar(n)
	}

	require.EqualValues(t, 0, dic.Cantidad())

	// Verificar que el árbol mantiene el orden después de operaciones
	dic.Guardar(5, 5)
	dic.Guardar(3, 3)
	dic.Guardar(7, 7)
	dic.Guardar(2, 2)
	dic.Guardar(4, 4)

	var claves []int
	dic.Iterar(func(clave int, valor int) bool {
		claves = append(claves, clave)
		return true
	})

	require.Equal(t, []int{2, 3, 4, 5, 7}, claves)
}

// ------------------------------- ITERADOR INTERNO -------------------------------

func TestDiccionarioOrdenadoIteradorInternoClaves(t *testing.T) {
	claves := []string{"Gato", "Perro", "Vaca"}
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	for _, c := range claves {
		dic.Guardar(c, nil)
	}

	cs := make([]string, 0, 3)
	dic.Iterar(func(clave string, dato *int) bool {
		cs = append(cs, clave)
		return true
	})

	require.EqualValues(t, []string{"Gato", "Perro", "Vaca"}, cs)
}

func TestDiccionarioOrdenadoIteradorInternoValores(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	claves := []string{"Gato", "Perro", "Vaca", "Burrito", "Hamster"}
	valores := []int{6, 2, 3, 4, 5}
	for i := range claves {
		dic.Guardar(claves[i], valores[i])
	}

	factorial := 1
	dic.Iterar(func(_ string, dato int) bool { factorial *= dato; return true })
	require.EqualValues(t, 720, factorial)
}

func TestDiccionarioOrdenadoIteradorInternoValoresConBorrados(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar("Elefante", 7)
	dic.Guardar("Gato", 6)
	dic.Guardar("Perro", 2)
	dic.Borrar("Elefante")

	factorial := 1
	dic.Iterar(func(_ string, dato int) bool { factorial *= dato; return true })
	require.EqualValues(t, 12, factorial)
}

func TestIteradorInternoSumarTodos(t *testing.T) {
	cmp := func(a, b int) int { return a - b }
	dic := TDADiccionario.CrearABB[int, int](cmp)
	valores := []int{5, 3, 7, 2, 4, 6, 8}

	for _, v := range valores {
		dic.Guardar(v, v)
	}

	suma := 0
	dic.Iterar(func(clave int, valor int) bool {
		suma += valor
		return true
	})

	require.Equal(t, 35, suma) // 5+3+7+2+4+6+8 = 35
}

func TestIteradorInternoCorteTemprano(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	claves := []string{"A", "B", "C", "D", "E"}
	for i, c := range claves {
		dic.Guardar(c, i)
	}

	suma := 0
	dic.Iterar(func(clave string, dato int) bool {
		suma += dato
		return suma < 3 // corta si ya sumamos 3 o más
	})
	require.LessOrEqual(t, suma, 3)
}

func TestIteradorInternoSumarPares(t *testing.T) {
	cmp := func(a, b int) int { return a - b }
	dic := TDADiccionario.CrearABB[int, int](cmp)
	valores := []int{5, 3, 7, 2, 4, 6, 8}

	for _, v := range valores {
		dic.Guardar(v, v)
	}

	suma := 0
	dic.Iterar(func(clave int, valor int) bool {
		if valor%2 == 0 {
			suma += valor
		}
		return true
	})

	require.Equal(t, 20, suma) // 2+4+6+8 = 20
}

func TestIteradorInternoConRango(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	// Árbol resultante:
	//       A
	//        \
	//         B
	//          \
	//           C
	//            \
	//             D
	//              \
	//               E
	claves := []string{"A", "B", "C", "D", "E"}
	for i, c := range claves {
		dic.Guardar(c, i)
	}

	var resultado []string

	desde := "B"
	hasta := "D"

	dic.IterarRango(&desde, &hasta, func(clave string, dato int) bool {
		resultado = append(resultado, clave)
		return true
	})

	require.Equal(t, []string{"B", "C", "D"}, resultado)
}

func TestIteradorInternoRangoInOrder(t *testing.T) {
	cmp := func(a, b int) int { return a - b }
	dic := TDADiccionario.CrearABB[int, string](cmp)
	// Árbol resultante:
	//       10
	//     /    \
	//    5     15
	//   / \    / \
	//  3   7 12  17
	claves := []int{10, 5, 15, 3, 7, 12, 17}
	valores := []string{"A", "B", "C", "D", "E", "F", "G"}

	for i, k := range claves {
		dic.Guardar(k, valores[i])
	}

	desde := 3
	hasta := 17
	var resultado []string

	dic.IterarRango(&desde, &hasta, func(clave int, valor string) bool {
		resultado = append(resultado, valor)
		return true
	})

	// In-order (claves):      3,   5,   7,  10,  12,  15,  17
	require.Equal(t, []string{"D", "B", "E", "A", "F", "C", "G"}, resultado)
}

func TestIteradorInternoRangoParcial(t *testing.T) {
	cmp := func(a, b int) int { return a - b }
	dic := TDADiccionario.CrearABB[int, string](cmp)
	// Árbol más completo:
	//              10
	//         /          \
	//        5           15
	//      /   \       /    \
	//     3    7      12    17
	//    /\    /\    / \    / \
	//   1  4  6  8  11 13  16 18
	claves := []int{10, 5, 15, 3, 7, 12, 17, 1, 4, 6, 8, 11, 13, 16, 18}

	for _, k := range claves {
		dic.Guardar(k, fmt.Sprintf("%d", k))
	}

	desde := 5
	hasta := 15
	var resultado []string

	dic.IterarRango(&desde, &hasta, func(clave int, valor string) bool {
		resultado = append(resultado, valor)
		return true
	})

	expected := []string{"5", "6", "7", "8", "10", "11", "12", "13", "15"}
	require.Equal(t, expected, resultado)
}

func TestIteradorInternoRangoConCorte(t *testing.T) {
	cmp := func(a, b int) int { return a - b }
	dic := TDADiccionario.CrearABB[int, int](cmp)
	claves := []int{10, 5, 15, 3, 7, 12, 17}

	for _, k := range claves {
		dic.Guardar(k, k)
	}

	desde := 5
	hasta := 15
	suma := 0

	dic.IterarRango(&desde, &hasta, func(clave int, valor int) bool {
		suma += valor
		return suma < 30 // Corta cuando la suma alcanza o supera 30
	})

	require.Equal(t, suma, 34) // 5+7+10+12 = 34
}

func TestIteradorInternoRangoVacio(t *testing.T) {
	cmp := func(a, b int) int { return a - b }
	dic := TDADiccionario.CrearABB[int, string](cmp)
	dic.Guardar(5, "A")
	dic.Guardar(3, "B")
	dic.Guardar(7, "C")

	// Rango donde no hay elementos
	desde := 6
	hasta := 6
	var resultado []string
	dic.IterarRango(&desde, &hasta, func(clave int, valor string) bool {
		resultado = append(resultado, valor)
		return true
	})
	require.Empty(t, resultado)

	// Rango fuera de los límites
	desde = 8
	hasta = 10
	resultado = nil
	dic.IterarRango(&desde, &hasta, func(clave int, valor string) bool {
		resultado = append(resultado, valor)
		return true
	})
	require.Empty(t, resultado)
}

func TestIteradorInternoRangoUnicoElemento(t *testing.T) {
	cmp := func(a, b int) int { return a - b }
	dic := TDADiccionario.CrearABB[int, string](cmp)
	dic.Guardar(5, "A")

	// Rango que incluye exactamente el único elemento
	desde := 5
	hasta := 5
	var resultado []string
	dic.IterarRango(&desde, &hasta, func(clave int, valor string) bool {
		resultado = append(resultado, valor)
		return true
	})
	require.Equal(t, []string{"A"}, resultado)
}

// -------------------- PRUEBA DE VOLUMEN DEL ITERADOR INTERNO --------------------

func TestDiccionarioOrdenadoIteradorInternoVolumen(t *testing.T) {
	t.Log("Prueba de volumen del iterador interno")
	cmp := func(a, b int) int { return a - b }
	dic := TDADiccionario.CrearABB[int, int](cmp)

	cantidad := 10000
	posiblesNums := 1000000

	dic.Guardar(posiblesNums/2, posiblesNums) // Para no desbalancear el árbol

	/* Guardar elementos */

	guardados := 1
	for guardados < cantidad {
		n := rand.IntN(posiblesNums)
		if !dic.Pertenece(n) {
			dic.Guardar(n, n*2)
			guardados++
		}
	}

	require.Equal(t, guardados, cantidad, "No se guardó la cantidad pedida de elementos")

	/* Verificar iteración completa */

	var menorClave int

	// Devuelve la primer clave a iterar, o sea, la menor de todas
	dic.Iterar(func(clave int, valor int) bool { menorClave = clave; return false })

	ok := true
	contador := 0
	dic.Iterar(func(clave int, valor int) bool {
		require.Equal(t, clave, valor/2, "La clave no coincide")
		require.Equal(t, valor, clave*2, "El valor no coincide")
		ok = clave >= menorClave
		menorClave = clave
		contador++
		return ok
	})

	require.True(t, ok, "Las claves no están ordenadas ascendentemente")
	require.Equal(t, contador, guardados, "No se iteraron todos los elementos")

	/* Borrar la mitad, y volver a iterar */

	borrados := 0
	dic.Iterar(func(clave int, valor int) bool {
		require.Equal(t, valor, dic.Borrar(clave))
		borrados++
		return borrados != guardados/2
	})

	require.Equal(t, borrados, guardados/2, "No se borró la cantidad pedida de elementos")

	var menorClaveMitad int
	dic.Iterar(func(clave int, valor int) bool { menorClaveMitad = clave; return false })

	okMitad := true
	contadorMitad := 0
	dic.Iterar(func(clave int, valor int) bool {
		require.Equal(t, clave, valor/2, "Clave incorrecta después de borrados")
		require.Equal(t, valor, clave*2, "Valor incorrecto después de borrados")
		okMitad = clave >= menorClaveMitad
		menorClaveMitad = clave
		contadorMitad++
		return okMitad
	})

	require.True(t, okMitad, "Las claves no están ordenadas ascendentemente, después de borrados")
	require.Equal(t, contadorMitad, guardados/2, "No se iteraron todos los elementos, después de borrados")
}

// --------------------------------------------------------------------------------
// -------------------------- TESTS DEL ITERADOR EXTERNO --------------------------
// --------------------------------------------------------------------------------

// ------------------------------- ITERAR SIN RANGO -------------------------------

func TestDiccionarioOrdenadoIterarVacio(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestDiccionarioOrdenadoIterar(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	claves := []string{"Gato", "Perro", "Vaca"}
	for i, c := range claves {
		dic.Guardar(c, fmt.Sprint(i))
	}

	iter := dic.Iterador()
	require.True(t, iter.HaySiguiente())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}

func TestDiccionarioOrdenadoIterarEnOrden(t *testing.T) {
	t.Log("Validar que el iterador externo recorre las claves en orden")
	claves := []string{"Vaca", "Gato", "Perro", "Burrito", "Hamster"}
	valores := []int{5, 2, 3, 1, 4}

	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	for i, clave := range claves {
		dic.Guardar(clave, valores[i])
	}

	iter := dic.IteradorRango(nil, nil) // Como no se itera con rango, desde y hasta son nil
	clavesOrdenadas := make([]string, 0, len(claves))
	valoresOrdenados := make([]int, 0, len(valores))

	for iter.HaySiguiente() {
		clave, valor := iter.VerActual()
		clavesOrdenadas = append(clavesOrdenadas, clave)
		valoresOrdenados = append(valoresOrdenados, valor)
		iter.Siguiente()
	}

	ordenEsperadoClaves := []string{"Burrito", "Gato", "Hamster", "Perro", "Vaca"}
	ordenEsperadoValores := []int{1, 2, 4, 3, 5}

	require.Equal(t, ordenEsperadoClaves, clavesOrdenadas, "Las claves no están en orden")
	require.Equal(t, ordenEsperadoValores, valoresOrdenados, "Los valores no están en orden")
}

func TestDiccionarioOrdenadoIteradorNoLlegaAlFinal(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	claves := []string{"A", "B", "C"}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")

	dic.Iterador()

	iter2 := dic.Iterador()
	iter2.Siguiente()

	iter3 := dic.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	tercero, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())

	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
}

func TestDiccionarioOrdenadoPruebaIterarTrasBorrados(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	claves := []string{"Gato", "Perro", "Vaca"}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")

	dic.Borrar(claves[0])
	dic.Borrar(claves[1])
	dic.Borrar(claves[2])

	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	dic.Guardar(claves[0], "A")

	iter = dic.Iterador()
	require.True(t, iter.HaySiguiente())
	c1, v1 := iter.VerActual()
	require.EqualValues(t, claves[0], c1)
	require.EqualValues(t, "A", v1)
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}

// ------------------------------- ITERAR CON RANGO -------------------------------

func TestDiccionarioOrdenadoIterarConRango(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	claves := []string{"A", "B", "C", "D", "E"}
	for i, c := range claves {
		dic.Guardar(c, i)
	}

	var resultado []string

	desde := "B"
	hasta := "D"

	iter := dic.IteradorRango(&desde, &hasta)
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		resultado = append(resultado, clave)
		iter.Siguiente()
	}

	require.Equal(t, []string{"B", "C", "D"}, resultado)
}

func TestDiccionarioOrdenadoIterarConRangoInOrder(t *testing.T) {
	cmp := func(a, b int) int { return a - b }
	dic := TDADiccionario.CrearABB[int, string](cmp)

	claves := []int{10, 5, 15, 3, 7, 12, 17}
	valores := []string{"A", "B", "C", "D", "E", "F", "G"}

	for i, k := range claves {
		dic.Guardar(k, valores[i])
	}

	desde := 3
	hasta := 17
	var resultado []string

	iter := dic.IteradorRango(&desde, &hasta)
	for iter.HaySiguiente() {
		_, valor := iter.VerActual()
		resultado = append(resultado, valor)
		iter.Siguiente()
	}

	// In-order (claves):      3,   5,   7,  10,  12,  15,  17
	require.Equal(t, []string{"D", "B", "E", "A", "F", "C", "G"}, resultado)
}

func TestDiccionarioOrdenadoIterarConRangoParcial(t *testing.T) {
	cmp := func(a, b int) int { return a - b }
	dic := TDADiccionario.CrearABB[int, string](cmp)
	claves := []int{10, 5, 15, 3, 7, 12, 17, 1, 4, 6, 8, 11, 13, 16, 18}

	for _, k := range claves {
		dic.Guardar(k, fmt.Sprintf("%d", k))
	}

	desde := 5
	hasta := 15
	var resultado []string

	iter := dic.IteradorRango(&desde, &hasta)
	for iter.HaySiguiente() {
		_, valor := iter.VerActual()
		resultado = append(resultado, valor)
		iter.Siguiente()
	}

	expected := []string{"5", "6", "7", "8", "10", "11", "12", "13", "15"}
	require.Equal(t, expected, resultado)
}

func TestDiccionarioOrdenadoIterarConRangoConCorte(t *testing.T) {
	cmp := func(a, b int) int { return a - b }
	dic := TDADiccionario.CrearABB[int, int](cmp)
	claves := []int{10, 5, 15, 3, 7, 12, 17}

	for _, k := range claves {
		dic.Guardar(k, k)
	}

	desde := 5
	hasta := 15
	suma := 0

	iter := dic.IteradorRango(&desde, &hasta)
	for iter.HaySiguiente() {
		_, valor := iter.VerActual()
		suma += valor
		if suma >= 30 {
			break
		}
		iter.Siguiente()
	}

	require.Equal(t, suma, 34) // 5+7+10+12 = 34
}

func TestDiccionarioOrdenadoIterarConRangoVacio(t *testing.T) {
	cmp := func(a, b int) int { return a - b }
	dic := TDADiccionario.CrearABB[int, string](cmp)
	dic.Guardar(5, "A")
	dic.Guardar(3, "B")
	dic.Guardar(7, "C")

	// Rango donde no hay elementos
	desde := 6
	hasta := 6
	var resultado []string

	iter1 := dic.IteradorRango(&desde, &hasta)
	for iter1.HaySiguiente() {
		_, valor := iter1.VerActual()
		resultado = append(resultado, valor)
		iter1.Siguiente()
	}
	require.Empty(t, resultado)

	// Rango fuera de los límites
	desde = 8
	hasta = 10
	resultado = nil

	iter2 := dic.IteradorRango(&desde, &hasta)
	for iter2.HaySiguiente() {
		_, valor := iter2.VerActual()
		resultado = append(resultado, valor)
		iter2.Siguiente()
	}
	require.Empty(t, resultado)
}

func TestDiccionarioOrdenadoIterarConRangoUnicoElemento(t *testing.T) {
	cmp := func(a, b int) int { return a - b }
	dic := TDADiccionario.CrearABB[int, string](cmp)
	dic.Guardar(5, "A")

	// Rango que incluye exactamente el único elemento
	desde := 5
	hasta := 5
	var resultado []string

	iter := dic.IteradorRango(&desde, &hasta)
	for iter.HaySiguiente() {
		_, valor := iter.VerActual()
		resultado = append(resultado, valor)
		iter.Siguiente()
	}
	require.Equal(t, []string{"A"}, resultado)
}

// --------------------------------------------------------------------------------
// ---------------------------------- BENCHMARKS ----------------------------------
// --------------------------------------------------------------------------------

// ---------------------------------- PARA EL ABB ----------------------------------

func ejecutarPruebaVolumenABB(b *testing.B, n int) {
	cmp := func(a, b int) int { return a - b }
	dic := TDADiccionario.CrearABB[int, int](cmp)
	posiblesNums := 1000000

	dic.Guardar(posiblesNums/2, posiblesNums)

	/* Guardar elementos */

	guardados := 1
	for guardados < n {
		elem := rand.IntN(posiblesNums)
		if !dic.Pertenece(elem) {
			dic.Guardar(elem, elem*2)

			require.True(b, dic.Pertenece(elem))
			require.Equal(b, elem*2, dic.Obtener(elem))

			guardados++
		}
	}

	require.Equal(b, guardados, n, "No se guardó la cantidad pedida de elementos")

	/* Verificar iteración completa */

	menorClave := -1 // Se guardaron solo números positivos
	ok := true
	contador := 0

	iter1 := dic.Iterador()
	for iter1.HaySiguiente() {
		clave, valor := iter1.VerActual()
		require.Equal(b, clave, valor/2, "La clave no coincide")
		require.Equal(b, valor, clave*2, "El valor no coincide")

		ok = clave >= menorClave
		if !ok {
			break
		}

		contador++
		iter1.Siguiente()
	}

	require.True(b, ok, "Las claves no están ordenadas ascendentemente")
	require.Equal(b, contador, guardados, "No se iteraron todos los elementos")

	/* Borrar la mitad, y volver a iterar */

	borrados := 0
	iterBorrados := dic.Iterador()
	for iterBorrados.HaySiguiente() && borrados != guardados/2 {
		clave, valor := iterBorrados.VerActual()

		require.True(b, dic.Pertenece(clave))
		require.Equal(b, valor, dic.Borrar(clave))

		borrados++
		iterBorrados.Siguiente()
	}

	require.Equal(b, borrados, guardados/2, "No se borró la cantidad pedida de elementos")

	menorClave = -1 // Se reinicia el valor
	okMitad := true
	contadorMitad := 0

	iter2 := dic.Iterador()
	for iter2.HaySiguiente() {
		clave, valor := iter2.VerActual()
		require.Equal(b, clave, valor/2, "La clave no coincide")
		require.Equal(b, valor, clave*2, "El valor no coincide")

		okMitad = clave >= menorClave
		if !okMitad {
			break
		}

		contadorMitad++
		iter2.Siguiente()
	}

	require.True(b, okMitad, "Las claves no están ordenadas ascendentemente, después de borrados")
	require.Equal(b, contadorMitad, guardados/2, "No se iteraron todos los elementos, después de borrados")
}

func BenchmarkDiccionarioOrdenado(b *testing.B) {
	b.Log("Prueba de stress del Diccionario Ordenado (ABB)")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenABB(b, n)
			}
		})
	}
}

// --------------------------- PARA EL ITERADOR EXTERNO ----------------------------

func ejecutarPruebasVolumenIteradorABB(b *testing.B, n int) {
	cmp := func(a, b int) int { return a - b }
	dic := TDADiccionario.CrearABB[int, *int](cmp)

	claves := make([]int, n)
	valores := make([]int, n)

	// Generar claves en orden aleatorio
	for i := 0; i < n; i++ {
		claves[i] = i
	}
	rand.Shuffle(n, func(i, j int) {
		claves[i], claves[j] = claves[j], claves[i]
	})

	// Guardar en el ABB
	for i := 0; i < n; i++ {
		valores[i] = i
		dic.Guardar(claves[i], &valores[i])
	}

	// Prueba de iteración ordenada
	iter := dic.Iterador()
	require.True(b, iter.HaySiguiente())

	ok := true
	var i int
	var clave int
	var valor *int
	anterior := -1

	for i = 0; i < n; i++ {
		if !iter.HaySiguiente() {
			ok = false
			break
		}
		c1, v1 := iter.VerActual()
		clave = c1

		// Verificar orden ascendente
		require.Greater(b, clave, anterior, "El iterador no está ordenando correctamente")
		anterior = clave

		valor = v1
		if valor == nil {
			ok = false
			break
		}
		*valor = n // Modificar el valor
		iter.Siguiente()
	}
	require.True(b, ok, "Iteración en volumen falló")
	require.EqualValues(b, n, i, "No se iteraron todos los elementos")
	require.False(b, iter.HaySiguiente(), "El iterador debe terminar")

	// Verificar modificaciones
	ok = true
	for i = 0; i < n; i++ {
		if valores[i] != n {
			ok = false
			break
		}
	}
	require.True(b, ok, "No se actualizaron todos los valores")
}

func BenchmarkIteradorABB(b *testing.B) {
	b.Log("Prueba de stress del Iterador del ABB")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebasVolumenIteradorABB(b, n)
			}
		})
	}
}
