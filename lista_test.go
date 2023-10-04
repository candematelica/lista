package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const GRANTAMANIO = 10000

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo(), "El largo de una lista vacia es 0")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() }, "Hay un panic al querer borrar el primer elemento de una lista vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "Hay un panic al querer ver el primer elemento de una lista vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() }, "Hay un panic al querer ver el ultimo elemento de una lista vacia")
}

func TestInsertarYBorrarVariosElementosNumericos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(3)
	require.EqualValues(t, 3, lista.VerPrimero(), "En una lista vacia, los elementos numericos se insertan correctamente al principio")
	require.EqualValues(t, 3, lista.VerUltimo(), "En una lista vacia, los elementos numericos se insertan correctamente al principio")
	require.EqualValues(t, 1, lista.Largo(), "Luego de insertar un elemento, el largo aumenta")
	lista.InsertarUltimo(-1)
	require.EqualValues(t, 3, lista.VerPrimero())
	require.EqualValues(t, -1, lista.VerUltimo(), "Los elementos numericos se insertan correctamente al final")
	require.EqualValues(t, 2, lista.Largo())
	lista.InsertarPrimero(0)
	require.EqualValues(t, 0, lista.VerPrimero())
	require.EqualValues(t, -1, lista.VerUltimo())
	require.EqualValues(t, 3, lista.Largo())

	require.False(t, lista.EstaVacia())

	lista.BorrarPrimero()
	require.EqualValues(t, 3, lista.VerPrimero(), "Los elementos numericos se borran correctamente")
	require.EqualValues(t, -1, lista.VerUltimo(), "Los elementos numericos se borran correctamente")
	require.EqualValues(t, 2, lista.Largo(), "Al borrar un elemento, el largo disminuye")
	lista.BorrarPrimero()
	require.EqualValues(t, -1, lista.VerPrimero())
	require.EqualValues(t, -1, lista.VerUltimo())
	require.EqualValues(t, 1, lista.Largo())
}

func TestInsertarYBorrarVariosElementosDeTipoString(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()

	lista.InsertarPrimero("elem1")
	require.EqualValues(t, "elem1", lista.VerPrimero(), "En una lista vacia, los elementos de tipo string se insertan correctamente al principio")
	require.EqualValues(t, "elem1", lista.VerUltimo(), "En una lista vacia, los elementos de tipo string se insertan correctamente al principio")
	require.EqualValues(t, 1, lista.Largo(), "Luego de insertar un elemento, el largo aumenta")
	lista.InsertarUltimo("elem2")
	require.EqualValues(t, "elem1", lista.VerPrimero())
	require.EqualValues(t, "elem2", lista.VerUltimo(), "Los elementos de tipo string se insertan correctamente al final")
	require.EqualValues(t, 2, lista.Largo())
	lista.InsertarPrimero("elem3")
	require.EqualValues(t, "elem3", lista.VerPrimero())
	require.EqualValues(t, "elem2", lista.VerUltimo())
	require.EqualValues(t, 3, lista.Largo())

	require.False(t, lista.EstaVacia())

	lista.BorrarPrimero()
	require.EqualValues(t, "elem1", lista.VerPrimero(), "Los elementos de tipo string se borran correctamente")
	require.EqualValues(t, "elem2", lista.VerUltimo(), "Los elementos de tipo string se borran correctamente")
	require.EqualValues(t, 2, lista.Largo(), "Al borrar un elemento, el largo disminuye")
	lista.BorrarPrimero()
	require.EqualValues(t, "elem2", lista.VerPrimero())
	require.EqualValues(t, "elem2", lista.VerUltimo())
	require.EqualValues(t, 1, lista.Largo())
}

func TestListaConUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(7)
	require.EqualValues(t, 7, lista.VerPrimero(), "En una lista vacia, los elementos se insertan correctamente al final")
	require.EqualValues(t, 7, lista.VerUltimo(), "En una lista vacia, los elementos se insertan correctamente al final")
	require.EqualValues(t, 1, lista.Largo())

	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia(), "Al borrar un elemento de una lista de largo 1, esta se comporta como vacia")
	require.EqualValues(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
}

func TestIteradorExterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(3)
	lista.InsertarUltimo(-1)
	lista.InsertarPrimero(0)
	lista.InsertarUltimo(9)
	require.EqualValues(t, 4, lista.Largo())
	require.False(t, lista.EstaVacia())

	var primero_insertado bool
	var medio_insertado bool

	for iter1 := lista.Iterador(); iter1.HaySiguiente(); iter1.Siguiente() {
		if !primero_insertado {
			iter1.Insertar(5)
			primero_insertado = true
		} else if iter1.VerActual() == 3 && !medio_insertado {
			iter1.Insertar(2)
			medio_insertado = true
		} else if iter1.VerActual() == 9 {
			iter1.Siguiente()
			iter1.Insertar(-7)
			break
		}
	}

	var i int
	for iter2 := lista.Iterador(); iter2.HaySiguiente(); iter2.Siguiente() {
		switch i {
		case 0:
			require.EqualValues(t, 5, iter2.VerActual(), "Al insertar un elemento en la posicion en la que se crea el iterador, efectivamente se inserta al principio")
		case 2:
			require.EqualValues(t, 2, iter2.VerActual(), "Insertar un elemento en el medio se hace en la posicion correcta")
		case 6:
			require.EqualValues(t, -7, iter2.VerActual(), "Insertar un elemento cuando el iterador esta al final efectivamente es equivalente a insertar al final")
		}
		i++
	}

	var primero_eliminado bool
	for iter3 := lista.Iterador(); iter3.HaySiguiente(); iter3.Siguiente() {
		if !primero_eliminado {
			iter3.Borrar()
			primero_eliminado = true
		} else if iter3.VerActual() == 3 {
			iter3.Borrar()
		} else if iter3.VerActual() == -7 {
			iter3.Borrar()
			break
		}
	}

	var j int
	for iter4 := lista.Iterador(); iter4.HaySiguiente(); iter4.Siguiente() {
		switch j {
		case 0:
			require.EqualValues(t, 0, iter4.VerActual(), "Al remover el elemento cuando se crea el iterador, cambia el primer elemento de la lista")
		case 1:
			require.EqualValues(t, 2, iter4.VerActual(), "Al remover un elemento del medio, este no esta")
		case 2:
			require.EqualValues(t, -1, iter4.VerActual(), "Al remover un elemento del medio, este no esta")
		case 3:
			require.EqualValues(t, 9, iter4.VerActual(), "Remover el ultimo elemento con el iterador cambia el ultimo de la lista")
		}
		j++
	}

}

func TestIteradorInternoConCondicionDeCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(3)
	lista.InsertarUltimo(-1)
	lista.InsertarPrimero(0)
	lista.InsertarUltimo(9)
	require.False(t, lista.EstaVacia())

	var suma int
	lista.Iterar(func(dato int) bool {
		if dato == -1 {
			return false
		}
		suma = suma + dato
		return true
	})
	require.EqualValues(t, 3, suma, "El iterador interno funciona")
}

func TestIteradorInternoListaCompleta(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(3)
	lista.InsertarUltimo(-1)
	lista.InsertarPrimero(0)
	lista.InsertarUltimo(9)
	require.False(t, lista.EstaVacia())

	var suma int
	lista.Iterar(func(dato int) bool {
		suma = suma + dato
		return true
	})
	require.EqualValues(t, 11, suma, "Se puede iterar toda la lista correctamente con el iterador interno")
}

func TestVolumenPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 0; i < GRANTAMANIO; i++ {
		lista.InsertarPrimero(i)
	}
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 9999, lista.VerPrimero())
	require.EqualValues(t, 0, lista.VerUltimo())
	require.EqualValues(t, GRANTAMANIO, lista.Largo())
	for i := GRANTAMANIO - 1; i >= 0; i-- {
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, 0, lista.VerUltimo())
		require.EqualValues(t, i+1, lista.Largo())
		lista.BorrarPrimero()
	}
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
}
func TestVolumenUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 0; i < GRANTAMANIO; i++ {
		lista.InsertarUltimo(i)
	}
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.VerPrimero())
	require.EqualValues(t, 9999, lista.VerUltimo())
	require.EqualValues(t, GRANTAMANIO, lista.Largo())
	for i := 0; i < GRANTAMANIO; i++ {
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, 9999, lista.VerUltimo())
		require.EqualValues(t, GRANTAMANIO-i, lista.Largo())
		lista.BorrarPrimero()
	}
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
}
