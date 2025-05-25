package cola_prioridad_test

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	TDAColaPrioridad "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"

	"github.com/stretchr/testify/require"
)

var TAMS_VOLUMEN = []int{12500, 25000, 50000, 100000, 200000, 400000}

// TESTS

func TestHeapVacio(t *testing.T) {
	t.Log("Comprueba que el heap vacío no tiene elementos")
	heap := TDAColaPrioridad.CrearHeap(func(a, b int) int { return a - b })
	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestEncolarYDesencolar(t *testing.T) {
	t.Log("Encolar y desencolar algunos elementos, y verificar el máximo en cada paso")
	heap := TDAColaPrioridad.CrearHeap(func(a, b int) int { return a - b })

	heap.Encolar(5)
	require.Equal(t, 5, heap.VerMax())
	require.Equal(t, 1, heap.Cantidad())

	heap.Encolar(10)
	require.Equal(t, 10, heap.VerMax())
	require.Equal(t, 2, heap.Cantidad())

	heap.Encolar(3)
	require.Equal(t, 10, heap.VerMax())
	require.Equal(t, 3, heap.Cantidad())

	require.Equal(t, 10, heap.Desencolar())
	require.Equal(t, 5, heap.VerMax())
	require.Equal(t, 2, heap.Cantidad())

	require.Equal(t, 5, heap.Desencolar())
	require.Equal(t, 3, heap.VerMax())
	require.Equal(t, 1, heap.Cantidad())

	require.Equal(t, 3, heap.Desencolar())
	require.Equal(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
}

func TestEncolarDesencolarAlternado(t *testing.T) {
	t.Log("Encolar y desencolar alternadamente, verificando el máximo en cada paso")
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

	for i := range 10 {
		heap.Encolar(i)
		require.Equal(t, heap.VerMax(), heap.Desencolar())
	}
}

func TestStringsPorLargo(t *testing.T) {
	t.Log("Heap con strings ordenados de mayor a menor largo")
	heap := TDAColaPrioridad.CrearHeap(func(a, b string) int { return len(a) - len(b) })

	heap.Encolar("a")
	heap.Encolar("abc")
	heap.Encolar("ab")

	require.Equal(t, "abc", heap.VerMax())
	require.Equal(t, "abc", heap.Desencolar())
	require.Equal(t, "ab", heap.VerMax())
}

func TestStringsCompare(t *testing.T) {
	t.Log("Heap con strings ordenados por criterio lexicográfico")
	heap := TDAColaPrioridad.CrearHeap(strings.Compare)

	heap.Encolar("Elefante")
	heap.Encolar("Abeja")
	heap.Encolar("Burro")
	heap.Encolar("Aguila")
	require.Equal(t, "Elefante", heap.Desencolar())
	require.Equal(t, "Burro", heap.Desencolar())
	require.Equal(t, "Aguila", heap.VerMax())

	heap.Encolar("Dromedario")
	heap.Encolar("Gato")
	heap.Encolar("Orangutan")
	require.Equal(t, "Orangutan", heap.Desencolar())
	require.Equal(t, "Gato", heap.Desencolar())

	require.Equal(t, "Dromedario", heap.Desencolar())
	require.Equal(t, "Aguila", heap.VerMax())
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
	require.Equal(t, "Juan", heap.VerMax().nombre)
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
	require.Equal(t, 1, heap.Cantidad())
}

func TestHeapDesdeArregloVacio(t *testing.T) {
	t.Log("Crear heap desde arreglo vacío y verificar propiedad del heap")
	var arr []int
	heap := TDAColaPrioridad.CrearHeapArr(arr, func(a, b int) int { return a - b })

	heap.Encolar(5)
	heap.Encolar(3)
	require.Equal(t, 5, heap.VerMax())
	require.Equal(t, 2, heap.Cantidad())

	heap.Encolar(8)
	heap.Encolar(15)
	require.Equal(t, 15, heap.VerMax())
	require.Equal(t, 4, heap.Cantidad())

	heap.Desencolar()
	heap.Desencolar()
	require.Equal(t, 5, heap.VerMax())
	heap.Encolar(20)
	require.Equal(t, 20, heap.VerMax())
}

func TestHeapDesdeArreglo(t *testing.T) {
	t.Log("Crear heap desde arreglo y verificar propiedad de heap")
	arr := []int{15, 3, 8, 20, 5}
	heap := TDAColaPrioridad.CrearHeapArr(arr, func(a, b int) int { return a - b })

	require.Equal(t, 20, heap.VerMax())
	require.Equal(t, len(arr), heap.Cantidad())
	require.Equal(t, 20, heap.Desencolar())
	require.Equal(t, 15, heap.VerMax())
	require.Equal(t, len(arr)-1, heap.Cantidad())

	require.Equal(t, []int{15, 3, 8, 20, 5}, arr) // El arreglo original no se vió modificado
}

func TestHeapSort(t *testing.T) {
	t.Log("Ordenar un arreglo usando HeapSort, de menor a mayor, y luego, de mayor a menor")
	elementos := []int{9, 3, 7, 1, 5, 10, 2, 8, 6, 4}
	menorAMayor := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mayorAMenor := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	TDAColaPrioridad.HeapSort(elementos, func(a, b int) int { return a - b })
	require.Equal(t, menorAMayor, elementos)

	TDAColaPrioridad.HeapSort(elementos, func(a, b int) int { return b - a })
	require.Equal(t, mayorAMenor, elementos)
}

// BENCHMARKS

func ejecutarPruebaVolumenHeap(b *testing.B, n int) {
	/* Heap de máximos */
	cmpMax := func(a, b int) int { return a - b }
	heapMax := TDAColaPrioridad.CrearHeap(cmpMax)
	dicHeapMax := TDADiccionario.CrearABB[int, int](cmpMax)

	heapMax.Encolar(500000)
	dicHeapMax.Guardar(500000, 500000)

	/*
		Un diccionario, nos permite verificar más rápido si
		un número se repite o no; además, dada la raíz y el
		rango de posibles números, no podría desbalancearse
	*/

	cantMax := 1
	for cantMax < n {
		valor := rand.Intn(1000000)
		if !dicHeapMax.Pertenece(valor) {
			heapMax.Encolar(valor)
			dicHeapMax.Guardar(valor, valor)
			cantMax++
		}
	}

	require.EqualValues(b, n, cantMax, "La cantidad de elementos es incorrecta")
	require.EqualValues(b, n, dicHeapMax.Cantidad(), "La cantidad de elementos es incorrecta")
	require.EqualValues(b, n, heapMax.Cantidad(), "Encolar muchos elementos no funciona correctamente")

	okMax := true
	var anteriorMax int
	for i := 0; i < n; i++ {
		maxActual := heapMax.VerMax()
		elemento := heapMax.Desencolar()

		// Verificar coherencia entre VerMax() y Desencolar()
		if maxActual != elemento {
			okMax = false
			break
		}

		// Validar orden no creciente (cada elemento debe ser <= al anterior)
		if i > 0 && elemento > anteriorMax {
			okMax = false
			break
		}
		anteriorMax = elemento
	}

	require.True(b, okMax, "Desencolar muchos elementos no funciona correctamente en heap de máximos")
	require.EqualValues(b, 0, heapMax.Cantidad())

	/* Heap de mínimos */
	cmpMin := func(a, b int) int { return b - a }
	heapMin := TDAColaPrioridad.CrearHeap(cmpMin)
	dicHeapMin := TDADiccionario.CrearABB[int, int](cmpMax) // Misma comparación para el ABB

	heapMin.Encolar(500000)
	dicHeapMin.Guardar(500000, 500000)

	cantMin := 1
	for cantMin < n {
		valor := rand.Intn(1000000)
		if !dicHeapMin.Pertenece(valor) {
			heapMin.Encolar(valor)
			dicHeapMin.Guardar(valor, valor)
			cantMin++
		}
	}

	require.EqualValues(b, n, cantMin, "La cantidad de elementos es incorrecta")
	require.EqualValues(b, n, dicHeapMin.Cantidad(), "La cantidad de elementos es incorrecta")
	require.EqualValues(b, n, heapMin.Cantidad(), "Encolar muchos elementos no funciona correctamente")

	okMin := true
	var anteriorMin int
	for i := 0; i < n; i++ {
		minActual := heapMin.VerMax()
		elemento := heapMin.Desencolar()

		// Verificar coherencia entre VerMax() y Desencolar()
		if minActual != elemento {
			okMin = false
			break
		}

		// Validar orden no decreciente (cada elemento debe ser >= al anterior)
		if i > 0 && elemento < anteriorMin {
			okMin = false
			break
		}
		anteriorMin = elemento
	}

	require.True(b, okMin, "Desencolar muchos elementos no funciona correctamente en heap de mínimos")
	require.EqualValues(b, 0, heapMin.Cantidad())
}

func BenchmarkHeap(b *testing.B) {
	b.Log("Prueba de stress del heap. Se encolan muchos números aleatorios, de un rango grande, verificando" +
		"con un diccionario que no se repitan; al desencolar valor a valor, se devuelve el de mayor prioridad." +
		"Funciona tanto para un heap de máximos, como de mínimos.")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenHeap(b, n)
			}
		})
	}
}

func ejecutarPruebaVolumenHeapArr(b *testing.B, n int) {

	/* Heap de máximos */

	cmpMax := func(a, b int) int { return a - b }
	var arr []int
	dic := TDADiccionario.CrearABB[int, int](cmpMax)

	arr = append(arr, 500000)
	dic.Guardar(500000, 500000)

	cantidad := 1
	for cantidad < n {
		valor := rand.Intn(1000000)
		if !dic.Pertenece(valor) {
			arr = append(arr, valor)
			dic.Guardar(valor, valor)
			cantidad++
		}
	}

	require.EqualValues(b, n, cantidad, "La cantidad de elementos es incorrecta")
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")
	require.EqualValues(b, n, len(arr), "La cantidad de elementos es incorrecta")

	heapMax := TDAColaPrioridad.CrearHeapArr(arr, cmpMax)

	okMax := true
	for range n {
		okMax = heapMax.VerMax() == heapMax.Desencolar()
		if !okMax {
			break
		}
	}

	require.True(b, okMax, "Desencolar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, heapMax.Cantidad())

	/* Heap de mínimos */

	cmpMin := func(a, b int) int { return b - a }
	heapMin := TDAColaPrioridad.CrearHeapArr(arr, cmpMin) // Como el arreglo original no cambia, podemos reutilizarlo

	okMin := true
	for range n {
		okMin = heapMin.VerMax() == heapMin.Desencolar()
		if !okMin {
			break
		}
	}

	require.True(b, okMin, "Desencolar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, heapMin.Cantidad())
}

func BenchmarkHeapArr(b *testing.B) {
	b.Log("Prueba de stress del heap, según un arreglo. En un arreglo, se guardan muchos números aleatorios, de un rango grande," +
		"verificando con un diccionario que no se repitan. Se lo usa tanto para un heap de máximos, como para un heap de mínimos." +
		"Al desencolar, se devuelve el valor de mayor prioridad.")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenHeapArr(b, n)
			}
		})
	}
}

func ejecutarPruebaVolumenHeapSort(b *testing.B, n int) {

	/* Ordenar de menor a mayor */

	cmpMax := func(a, b int) int { return a - b }
	var arrMinMay []int
	dicMinMay := TDADiccionario.CrearABB[int, int](cmpMax)

	arrMinMay = append(arrMinMay, 500000)
	dicMinMay.Guardar(500000, 500000)

	cantMinMay := 1
	for cantMinMay < n {
		valor := rand.Intn(1000000)
		if !dicMinMay.Pertenece(valor) {
			arrMinMay = append(arrMinMay, valor)
			dicMinMay.Guardar(valor, valor)
			cantMinMay++
		}
	}

	require.EqualValues(b, n, cantMinMay, "La cantidad de elementos es incorrecta")
	require.EqualValues(b, n, dicMinMay.Cantidad(), "La cantidad de elementos es incorrecta")
	require.EqualValues(b, n, len(arrMinMay), "La cantidad de elementos es incorrecta")

	TDAColaPrioridad.HeapSort(arrMinMay, cmpMax)

	okMinMay := true
	for i := 1; i < len(arrMinMay); i++ {
		okMinMay = cmpMax(arrMinMay[i], arrMinMay[i-1]) > 0 // Mi elemento debe ser mayor a su anterior
		if !okMinMay {
			break
		}
	}

	require.True(b, okMinMay, "No se ordenaron los elementos correctamente")

	/* Ordenar de mayor a menor */

	cmpMin := func(a, b int) int { return b - a }
	var arrMayMin []int
	dicMayMin := TDADiccionario.CrearABB[int, int](cmpMax)

	arrMayMin = append(arrMayMin, 500000)
	dicMayMin.Guardar(500000, 500000)

	cantMayMin := 1
	for cantMayMin < n {
		valor := rand.Intn(1000000)
		if !dicMayMin.Pertenece(valor) {
			arrMayMin = append(arrMayMin, valor)
			dicMayMin.Guardar(valor, valor)
			cantMayMin++
		}
	}

	require.EqualValues(b, n, cantMayMin, "La cantidad de elementos es incorrecta")
	require.EqualValues(b, n, dicMayMin.Cantidad(), "La cantidad de elementos es incorrecta")
	require.EqualValues(b, n, len(arrMayMin), "La cantidad de elementos es incorrecta")

	TDAColaPrioridad.HeapSort(arrMayMin, cmpMin)

	okMayMin := true
	for i := 1; i < len(arrMayMin); i++ {
		okMayMin = cmpMax(arrMayMin[i], arrMayMin[i-1]) < 0 // Mi elemento debe ser menor a su anterior
		if !okMayMin {
			break
		}
	}

	require.True(b, okMayMin, "No se ordenaron los elementos correctamente")
}

func BenchmarkHeapSort(b *testing.B) {
	b.Log("Prueba de stress del HeapSort. Ordena dos arreglos, de muchos números aleatorios desordenados;" +
		"al primero, de menor a mayor, y al segundo, de mayor a menor. En el primer caso, verifica que todo" +
		"elemento sea mayor a su anterior; en el segundo caso, que todo elemento sea menor a su anterior.")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenHeapSort(b, n)
			}
		})
	}
}
