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
	} else {
		lista.ultimo.siguiente = nodoNuevo
	}

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
	nuevoNodo := &nodoLista[T]{dato: elem}

	if iter.actual.siguiente == nil {
		iter.lista.ultimo.siguiente = nuevoNodo
		iter.lista.ultimo = nuevoNodo
		iter.actual = nuevoNodo
	} else if iter.lista.primero == iter.actual {
		nuevoNodo.siguiente = iter.actual
		iter.lista.primero = nuevoNodo
		iter.actual = nuevoNodo
	} else {
		nuevoNodo.siguiente = iter.actual
		punteroAnterior := iter.lista.primero
		for punteroAnterior.siguiente != iter.actual {
			punteroAnterior = punteroAnterior.siguiente
		}
		punteroAnterior.siguiente = nuevoNodo
		iter.actual = nuevoNodo
	}

	iter.lista.largo++
}
func (iter *iteradorLista[T]) Borrar() T {
	iter.chequearIterador()

	dato := iter.actual.dato

	if iter.actual == nil {
		return dato
	}

	if iter.lista.largo == 1 { //Estoy borrando el elemento de una lista de 1 elemento
		iter.lista.ultimo = nil
		iter.lista.primero = nil
		iter.actual = nil
	} else if iter.actual.siguiente == nil { //Estoy borrando un elemento al final
		punteroAnteUltimo := iter.lista.primero
		for punteroAnteUltimo.siguiente != iter.lista.ultimo {
			punteroAnteUltimo = punteroAnteUltimo.siguiente
		}
		iter.lista.ultimo = punteroAnteUltimo
		iter.lista.ultimo.siguiente = nil
		iter.actual = punteroAnteUltimo
		iter.actual.siguiente = nil
	} else if iter.lista.primero == iter.actual { //Estoy borrando al principio
		iter.lista.primero = iter.lista.primero.siguiente
		iter.actual = iter.actual.siguiente
	} else { //Estoy borrando en el medio
		punteroAnterior := iter.lista.primero
		for iter.actual != punteroAnterior.siguiente {
			punteroAnterior = punteroAnterior.siguiente
		}
		punteroAnterior.siguiente = iter.actual.siguiente
		iter.actual = iter.actual.siguiente
	}

	iter.lista.largo--

	return dato
}
