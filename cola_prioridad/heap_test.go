package cola_prioridad_test

import (
	"testing"

	TDAColaPrioridad "tdas/cola_prioridad"

	"github.com/stretchr/testify/require"
)

func TestHeapVacio(t *testing.T) {
	t.Log("Comprueba que el heap vacio no tiene elementos")
	heap := TDAColaPrioridad.CrearHeap(func(a, b int) int { return a - b })
	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestEncolarYDesencolar(t *testing.T) {
	t.Log("Encolar algunos elementos y verificar el máximo en cada paso")
	heap := TDAColaPrioridad.CrearHeap(func(a, b int) int { return a - b })

	heap.Encolar(5)
	require.Equal(t, 5, heap.VerMax())

	heap.Encolar(10)
	require.Equal(t, 10, heap.VerMax())

	heap.Encolar(3)
	require.Equal(t, 10, heap.VerMax())

	require.Equal(t, 10, heap.Desencolar())
	require.Equal(t, 5, heap.VerMax())
}

func TestEncolarDesencolarAlternado(t *testing.T) {
	t.Log("Encolar y desencolar alternadamente, verificando el maximo en cada paso")
	heap := TDAColaPrioridad.CrearHeap(func(a, b int) int { return a - b })

	heap.Encolar(10)
	require.Equal(t, 10, heap.VerMax())
	require.Equal(t, 10, heap.Desencolar())
	require.True(t, heap.EstaVacia())

	heap.Encolar(20)
	heap.Encolar(15)
	require.Equal(t, 20, heap.VerMax())
	require.Equal(t, 20, heap.Desencolar())
	require.Equal(t, 15, heap.VerMax())
	require.Equal(t, 15, heap.Desencolar())
	require.True(t, heap.EstaVacia())

	heap.Encolar(5)
	heap.Encolar(10)
	heap.Encolar(3)
	heap.Encolar(8)
	require.Equal(t, 10, heap.VerMax())
	require.Equal(t, 10, heap.Desencolar())
	require.Equal(t, 8, heap.VerMax())
	require.Equal(t, 8, heap.Desencolar())
	require.Equal(t, 5, heap.VerMax())
	require.Equal(t, 5, heap.Desencolar())
	require.Equal(t, 3, heap.VerMax())
	require.Equal(t, 3, heap.Desencolar())
	require.True(t, heap.EstaVacia())
}

func TestHeapDesdeArreglo(t *testing.T) {
	t.Log("Crear heap desde arreglo y verificar propiedad de heap")
	arr := []int{15, 3, 8, 20, 5}
	heap := TDAColaPrioridad.CrearHeapArr(arr, func(a, b int) int { return a - b })

	require.Equal(t, 20, heap.VerMax())
	require.Equal(t, 20, heap.Desencolar())
	require.Equal(t, 15, heap.VerMax())
	require.Equal(t, 4, heap.Cantidad())
}

func TestPruebaDeVolumen(t *testing.T) {
	t.Log("Prueba de volumen con muchos elementos")
	heap := TDAColaPrioridad.CrearHeap(func(a, b int) int { return a - b })
	n := 10000

	for i := 0; i < n; i++ {
		heap.Encolar(i)
		require.Equal(t, i, heap.VerMax())
	}

	require.Equal(t, n, heap.Cantidad())

	for i := n - 1; i >= 0; i-- {
		require.Equal(t, i, heap.Desencolar())
	}

	require.True(t, heap.EstaVacia())
}

func TestHeapSort(t *testing.T) {
	t.Log("Ordenar un arreglo usando HeapSort")
	elementos := []int{9, 3, 7, 1, 5, 10, 2, 8, 6, 4}
	esperado := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	TDAColaPrioridad.HeapSort(elementos, func(a, b int) int { return a - b })
	require.Equal(t, esperado, elementos)
}

func TestStrings(t *testing.T) {
	t.Log("Heap con strings ordenado por longitud")
	heap := TDAColaPrioridad.CrearHeap(func(a, b string) int { return len(a) - len(b) })

	heap.Encolar("a")
	heap.Encolar("abc")
	heap.Encolar("ab")

	require.Equal(t, "abc", heap.VerMax())
	require.Equal(t, "abc", heap.Desencolar())
	require.Equal(t, "ab", heap.VerMax())
}

func TestStructs(t *testing.T) {
	t.Log("Heap con estructuras personalizadas")
	type persona struct {
		nombre string
		edad   int
	}

	heap := TDAColaPrioridad.CrearHeap(func(a, b persona) int { return a.edad - b.edad })

	heap.Encolar(persona{"Juan", 30})
	heap.Encolar(persona{"Ana", 25})
	heap.Encolar(persona{"Pedro", 40})

	require.Equal(t, 40, heap.VerMax().edad)
	require.Equal(t, "Pedro", heap.Desencolar().nombre)
	require.Equal(t, 30, heap.VerMax().edad)
}

func TestHeapConElementosIguales(t *testing.T) {
	t.Log("Heap con elementos de igual prioridad")
	heap := TDAColaPrioridad.CrearHeap(func(a, b int) int { return 0 }) // Todos iguales

	heap.Encolar(5)
	heap.Encolar(5)
	heap.Encolar(5)

	require.Equal(t, 3, heap.Cantidad())
	require.Equal(t, 5, heap.Desencolar())
	require.Equal(t, 5, heap.Desencolar())
}

func TestPanics(t *testing.T) {
	t.Log("Verificar que panic funciona correctamente")
	heap := TDAColaPrioridad.CrearHeap(func(a, b int) int { return a - b })

	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })

	heap.Encolar(1)
	heap.Desencolar()

	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
}
