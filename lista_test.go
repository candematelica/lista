package lista_test

import (
	"lista"
	"testing"
  
	"github.com/stretchr/testify/require"
)

const int granTamanio = 10000

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo(), "El largo de una pila vacia es 0")
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

	//Al insertar un elemento en la posicion en la que se crea el iterador, efectivamente se inserta al principio
	//Insertar un elemento cuando el iterador esta al final efectivamente es equivalente a insertar al final
	//Insertar un elemento en el medio se hace en la posicion correcta
	//Al remover el elemento cuando se crea el iterador, cambia el primer elemento de la lista
	//Remover el ultimo elemento con el iterador cambia el ultimo de la lista
	//Verificar que al remover un elemento del medio, este no esta
}

func TestIteradorInterno(t *testing.T) {
	//Probar que funcione
	//Se puede usar el iterador, por ejemplo, para calcular una suma de todos los elementos de la lista
	//La prueba NO debe depender de imprimir dentro de la funcion visitar
	//Probar el caso de iteracion sin condicion de corte (iterar toda la lista)
	//Probar iteracion con condicion de corte (que en un momento determinado la funcion visitar de false)
}

func TestVolumen(t *testing.T) {
	//Ver, no se si esta bien este o.o
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 0; i < granTamanio; i++ {
		lista.InsertarPrimero(i)
	}
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 9999, lista.VerPrimero())
	require.EqualValues(t, 0, lista.VerUltimo())
	require.EqualValues(t, granTamanio, lista.Largo())
	for i := granTamanio - 1; i > 0; i-- {
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

	for i := 0; i < granTamanio; i++ {
		lista.InsertarUltimo(i)
	}
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.VerPrimero())
	require.EqualValues(t, 9999, lista.VerUltimo())
	require.EqualValues(t, granTamanio, lista.Largo())
	for i := 0; i < granTamanio; i++ {
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, 9999, lista.VerUltimo())
		require.EqualValues(t, granTamanio-i, lista.Largo())
		lista.BorrarPrimero()
	}
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
}
