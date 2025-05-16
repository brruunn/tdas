package diccionario_test

import (
	"fmt"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

// --------------------------------------------------------------------------------
// -------------------- TESTS DEL DICCIONARIO ORDENADO POR ABB --------------------
// --------------------------------------------------------------------------------

// CREAR

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

// GUARDAR

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

// BORRAR

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

// GUARDAR Y BORRAR

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
	for i := 0; i < 1000; i++ {
		dic.Guardar(i, i)
		dic.Borrar(i)
	}
	require.EqualValues(t, 0, dic.Cantidad())
}

// ITERADOR INTERNO

func TestDiccionarioOrdenadoIteradorInternoClaves(t *testing.T) {
	claves := []string{"Gato", "Perro", "Vaca"}
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	for _, c := range claves {
		dic.Guardar(c, nil)
	}

	cs := make([]string, 3)
	cantidad := 0
	dic.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		cantidad++
		return true
	})

	require.EqualValues(t, 3, cantidad)
	for _, c := range cs {
		require.Contains(t, claves, c)
	}
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

// --------------------------------------------------------------------
// -------------------- TESTS DEL ITERADOR EXTERNO --------------------
// --------------------------------------------------------------------

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
