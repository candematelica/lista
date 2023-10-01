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
	lista  *listaEnlazada[T]
	actual *nodoLista[T]
}

func CrearListaEnlazada[T any]() Lista[T] {

	return &listaEnlazada[T]{primero: nil, ultimo: nil, largo: 0}
}

func crearNodo[T any](dato T) *nodoLista[T] {
	return &nodoLista[T]{dato: dato}
}

func (lista listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil
}

func (lista *listaEnlazada[T]) InsertarPrimero(elem T) {
	nodoNuevo := crearNodo[T](elem)
	if lista.EstaVacia() {
		lista.ultimo = nodoNuevo
	} else {
		nodoNuevo.siguiente = lista.primero

	}
	lista.primero = nodoNuevo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(elem T) {
	nodoNuevo := crearNodo[T](elem)

	if lista.EstaVacia() {
		lista.primero = nodoNuevo
	}

	lista.ultimo.siguiente = lista.ultimo
	lista.ultimo = nodoNuevo
	lista.largo++
}

func chequearListaVacia[T any](lista *listaEnlazada[T]) {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {

	chequearListaVacia(lista)

	dato := lista.primero.dato

	lista.primero = lista.primero.siguiente

	lista.largo--

	if lista.largo == 0 {
		lista.ultimo = nil
	}

	return dato

}

func (lista *listaEnlazada[T]) VerPrimero() T {
	chequearListaVacia[T](lista)
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	chequearListaVacia[T](lista)
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	nodoActual := lista.primero

	for nodoActual != nil {
		valor := nodoActual.dato
		visitar(valor)
		nodoActual = nodoActual.siguiente
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {

	return &iteradorLista[T]{lista: lista, actual: lista.primero}
}

func (iter *iteradorLista[T]) chequearIterador() {
	if !iter.HaySiguiente() {
		panic("El iterador ya itero")
	}
}

func (iter *iteradorLista[T]) VerActual() T {
	iter.chequearIterador()

	return iter.actual.dato
}

func (iter *iteradorLista[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iteradorLista[T]) Siguiente() {
	iter.chequearIterador()

	iter.actual = iter.actual.siguiente

}

func (iter *iteradorLista[T]) Insertar(elem T) {
	if !iter.HaySiguiente() {
		puntero_ante_ultimo := iter.actual
		iter.actual.dato = elem
		iter.actual.siguiente = nil
		puntero_ante_ultimo.siguiente = iter.actual
	} else {
		iter.actual.siguiente = iter.actual
		iter.actual.dato = elem

	}
}

func (iter *iteradorLista[T]) Borrar() T {
	iter.chequearIterador()

	dato := iter.actual.dato
	if !iter.HaySiguiente() {
		iter.actual = iter.lista.ultimo
	} else {
		iter.actual.dato = iter.actual.siguiente.dato
		iter.actual.siguiente = iter.actual.siguiente.siguiente
	}

	return dato
}
