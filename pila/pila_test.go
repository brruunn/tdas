package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	_MENSAJE_PANIC = "La pila esta vacia"
	_CADENA_PRUEBA = "TDA Pila, sobre un arreglo dinámico (A&ED 1C - 2025)"
	_VOL_CHICO     = 100
	_VOL_GRANDE    = 1000000
	_NUM_EULER     = 2.718281
)

// Un pila sin elementos, se comporta como una pila vacía.
func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[rune]()

	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC, func() { pila.VerTope() })
	require.PanicsWithValue(t, _MENSAJE_PANIC, func() { pila.Desapilar() })
}

// Se puede apilar un valor, de cualquier tipo, correctamente.
func TestApilarValor(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()

	pila.Apilar(_CADENA_PRUEBA)
	require.Equal(t, pila.VerTope(), _CADENA_PRUEBA)
}

// Se puede desapilar un valor, de cualquier tipo, correctamente.
func TestDesapilarValor(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float64]()

	pila.Apilar(_NUM_EULER)
	require.Equal(t, pila.Desapilar(), _NUM_EULER)
}

// En una misma posición, puede haber varios elementos.
func TestApilarArreglo(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[[]string]()
	arr := make([]string, _VOL_CHICO)

	pila.Apilar(arr)
	require.Equal(t, pila.VerTope(), arr)
}

// De una misma posición, se pueden sacar varios elementos.
func TestDesapilarArreglo(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[[]float64]()
	arr := make([]float64, _VOL_CHICO)

	pila.Apilar(arr)
	require.Equal(t, pila.Desapilar(), arr)
}

// Una pila de varios elementos, cumple la invariante LIFO.
func TestVariosElems(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[bool]()

	for n := range _VOL_CHICO {
		pila.Apilar(n%2 == 0)
	}

	for !pila.EstaVacia() {
		require.Equal(t, pila.VerTope(), pila.Desapilar())
	}
}

// Una pila de muchos elementos, sigue siendo una pila.
func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	for n := range _VOL_GRANDE {
		pila.Apilar(n + 1)
	}

	for range _VOL_GRANDE {
		pila.Desapilar()
	}
}
