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

type iteradorLista[T any] struct {
	lista *listaEnlazada[T]
	actual *nodoLista[T]
}

func CrearListaEnlazada[T any]() Lista[T]{
	return &listaEnlazada{}
}

func crearNodo[T any](dato T) *nodoLista[T]{
	return &nodoLIsta[T]{dato: dato}
}

func (lista listaEnlazada[T])[T any]EstaVacia() bool {
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

	//Acá arriba no habría que agregar:
	//lista.ultimo := ante_ultimo
	//ante_ultimo.siguiente = nodoNuevo
	//??
	lista.ultimo = nodoNuevo
	lista.largo ++
}

func chequearListaVacia[T any](lista listaEnlazada[T]) {
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

func (lista listaEnlazada[T])VerPrimero[T any]() T{
	chequearListaVacia[T](lista)
	return lista.primero.dato
}

func (lista listaEnlazada[T])VerUlitmo[T any]() T{
	chequearListaVacia[T](lista)
	return lista.ultimo.dato
}

func (lista listaEnlazada[T])Largo[T any]() int{
	return lista.largo
}

func (lista listaEnlazada[T])Iterar[T any](visitar func(T) bool){
	nodoActual := lista.primero

	for nodoActual != nil{
		valor := nodoActual.dato
		visitar(valor)
		nodoActual = nodoActual.siguiente
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorLista[T]{lista, lista.VerPrimero()}
}

func (iter iteradorLista[T]) chequearIterador() {
	if !iter.HaySiguiente() {
		panic("El iterador ya itero")
	}
}

func (iter iteradorLista[T]) VerActual() T {
	iter.chequearIterador()

	return iter.actual.dato
}

func (iter iteradorLista[T]) HaySiguiente() bool {
	return iter.actual.siguiente != nil
}

func (iter *iteradorLista[T]) Siguiente() T {
	iter.chequearIterador()
	
	iter.actual = iter.actual.siguiente
	return iter.actual.dato
}

func (iter *iteradorLista[T]) Insertar(elem T) {
	if !iter.HaySiguiente() {
		puntero_ante_ultimo := iter.actual
		iter.actual.dato = elem
		iter.actual.siguiente = nil
		puntero_ante_ultimo.siguiente = iter.actual
	} else {
		puntero_siguiente := iter.actual
		iter.actual.dato = elem
		iter.actual.siguiente = puntero_siguiente
	}
}

func (iter *iteradorLista[T]) Borrar() T {
	iter.chequearIterador()

	dato := iter.actual.dato
	if !iter.HaySiguiente() {
		iter.actual.dato = nil
		iter.actual = iter.lista.ultimo
	} else {
		iter.actual.dato = iter.actual.siguiente.dato
		iter.actual.siguiente = iter.actual.siguiente.siguiente
	}

	return dato
}
