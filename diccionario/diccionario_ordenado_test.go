package diccionario_test

import (
	"fmt"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

var TAMS_DE_VOLUMEN = []int{12500, 25000, 50000, 100000, 200000, 400000}

func TestDiccionarioOrdenadoVacio(t *testing.T) {
	t.Log("Comprueba que DiccionarioOrdenado vacío no tiene claves")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestDiccionarioOrdenadoClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un ABB vacío que si buscamos la clave default del tipo, no existe")
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

func TestDiccionarioOrdenadoUnElemento(t *testing.T) {
	t.Log("Comprueba que DiccionarioOrdenado con un elemento tiene esa Clave")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestDiccionarioOrdenadoGuardar(t *testing.T) {
	t.Log("Guarda elementos y verifica comportamiento")
	claves := []string{"Gato", "Perro", "Vaca"}
	valores := []string{"miau", "guau", "moo"}

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
}

func TestDiccionarioOrdenadoReemplazoDato(t *testing.T) {
	t.Log("Reemplaza datos en claves existentes")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave := "Gato"
	dic.Guardar(clave, "miau")
	dic.Guardar(clave, "miu")
	require.EqualValues(t, "miu", dic.Obtener(clave))
}

func TestDiccionarioOrdenadoBorrar(t *testing.T) {
	t.Log("Borra elementos y verifica comportamiento")
	claves := []string{"Gato", "Perro", "Vaca"}
	valores := []string{"miau", "guau", "moo"}
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	for i := range claves {
		dic.Guardar(claves[i], valores[i])
	}

	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.False(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Borrar(claves[0])
	require.False(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, 1, dic.Cantidad())

	dic.Borrar(claves[1])
	require.EqualValues(t, 0, dic.Cantidad())
}

func TestDiccionarioOrdenadoReutlizacionDeBorrados(t *testing.T) {
	t.Log("Reutiliza una clave borrada")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave := "hola"
	dic.Guardar(clave, "mundo!")
	dic.Borrar(clave)
	dic.Guardar(clave, "mundooo!")
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, "mundooo!", dic.Obtener(clave))
}

func TestDiccionarioOrdenadoConClavesNumericas(t *testing.T) {
	t.Log("Valida claves numéricas")
	cmp := func(a, b int) int { return a - b }
	dic := TDADiccionario.CrearABB[int, string](cmp)
	clave := 10
	dic.Guardar(clave, "Gatito")
	require.EqualValues(t, "Gatito", dic.Borrar(clave))
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
	require.EqualValues(t, 0, dic.Borrar(a1))
}

func TestDiccionarioOrdenadoClaveVacia(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar("", "")
	require.True(t, dic.Pertenece(""))
	require.EqualValues(t, "", dic.Obtener(""))
}

func TestDiccionarioOrdenadoValorNulo(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	dic.Guardar("Pez", nil)
	require.True(t, dic.Pertenece("Pez"))
	require.Nil(t, dic.Obtener("Pez"))
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

func TestDiccionarioOrdenadoIterarVacio(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.Panics(t, func() { iter.VerActual() })
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

func TestIteradorInternoConRango(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
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

func ejecutarPruebaVolumenOrdenado(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	claves := make([]string, n)
	for i := range claves {
		claves[i] = fmt.Sprintf("%08d", i)
		dic.Guardar(claves[i], i)
	}

	for i := 0; i < n; i++ {
		require.EqualValues(b, i, dic.Obtener(claves[i]))
		dic.Borrar(claves[i])
	}
}

func BenchmarkDiccionarioOrdenado(b *testing.B) {
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenOrdenado(b, n)
			}
		})
	}
}
