type Lista[T any] interface {
	//Esta vacia devuelve verdadero si la lista no contiene ningun elemento, falso en caso contario.
	EstaVacia() bool
	// insertar primero, inserta el elemento pasado por parametro al inicio de la lista
	InsertarPrimero(T)
	//insertar ultimo, inserta el elemento pasado por parametro al final de la lista
	InsertarUltimo(T)
	//borrar primero, borra el primer elemento de la lista, devuelve un panic si esta vacia
	BorrarPrimero() T
	//ver primero, muestra el primer elemento de la lista, devuelve un panic si esta vacia
	VerPrimero() T
	//ver ultimo, muestra el primer elemento de la lista, devuelve un panic si esta vacia
	VerUltimo() T
	//ver largo devuelve la cantidad de elementos que tiene la lista
	Largo() int
	//iterador interno
	Iterar(visitar func(T) bool)
	//iterador externo
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
	VerActual() T
	HaySiguiente() bool
	Siguiente()
	Insertar(T)
	Borrar() T
}

