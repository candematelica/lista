package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}


func CrearListaEnlazada[T any]() Lista[T]{
	return &listaEnlazada{ primero:nil, ultimo:nil, largo:0}
}

func crearNodo[T any](dato T) *nodoLista[T]{

	return &nodoLIsta[T]{dato: dato, prox: nil}
}

func (lista *listaEnlazada[T])[T any]EstaVacia() bool {
	return lista.primero == nil 
}

func (lista *listaEnlazada[T])InsertarPrimero[T any](elem T){
	nodoNuevo := crearNodo[T](elem)
	if lista.EstaVacia(){
		lista.ultimo = nodoNuevo
	} else {
		nodoNuevo.siguiente = lista.primero

	}
	lista.primero = nodoNuevo
	lista.largo++
}

func (lista *listaEnlazada[T])InsertarUltimo[T any](elem T){
	nodoNuevo := crearNodo[T](elem)

	if lista.EstaVacia(){
		lista.primero = nodoNuevo
	}

	lista.ultimo = nodoNuevo
	lista.largo ++
}

func chequearListaVacia[T any](lista *listaEnlazada[T]) {
	if lista.EstaVacia() {

		panic("La lista esta vacia")
	}

}

func (lista *listaEnlazada[T])BorrarPrimero[T any]() T{

	chequearListaVacia[T](lista)

	dato := lista.primero.dato

	lista.primero = lista.primero.siguiente

	lista.largo --

	if lista.largo == 0{
		lista.ultimo == nil
	}

	return dato

}

func (lista *listaEnlazada[T])VerPrimero[T any]() T{
	return lista.primero.dato
}

func (lista *listaEnlazada[T])VerUlitmo[T any]() T{
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T])Largo[T any]() int{
	return lista.largo
}

func (lista *listaEnlazada[T])Iterador[T any](visitar func(T) bool){
	nodoActual := lista.primero

	for nodoActual != nil{
		valor := nodoActual.dato
		visitar(valor)
		nodoActual = nodoActual.siguiente
	}
}

func (lista listaEnlazada[T]) Iterar(visitar func(T) bool) {
	
}
