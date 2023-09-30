package lista_test

import (
	"lista"
	"testing"
  
	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	
}

func TestInsertarYBorrarVariosElementosNumericos(t *testing.T) {
	//En este la idea seria probar insertar y borrar sin que la lista llegue a estar vacía
  	//Probar acá funciones como el Largo, VerPrimero y VerUltimo
}

func TestInsertarYBorrarVariosElementosDeTipoString(t *testing.T) {
	//En este la idea seria probar insertar y borrar sin que la lista llegue a estar vacía
  	//Probar acá funciones como el Largo, VerPrimero y VerUltimo
}

func TestListaConUnElemento(t *testing.T) {
  	//Aca probamos que al borrar una lista con un elemento esta se comporte como vacia	
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
	//Usar los iteradores?
  	//Hacer uno usando BorrarPrimero y InsertarPrimero y otro usando BorrarPrimero y InsertarUltimo
}
