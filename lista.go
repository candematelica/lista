package lista

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
	//'VerActual' devuelve el elemento sobre el cual estamos parados en la lista, en caso de que
	//se llame a esta funcion cuando el iterador ya haya terminado de iterar, devolvera un panic
	VerActual() T
	//'HaySiguiente' devuelve true si se puede seguir iterando (es decir si hay elementos despues
	//del cual estemos parados). En caso contrario devuelve false
	HaySiguiente() bool
	//'Siguiente' devuelve el siguiente elemento a la posicion en la que estamos parados en la lista, en caso
	//de que se llame a esta funcion cuando el iterador ya haya terminado de iterar, devolvera un panic
	Siguiente() T
	//'Insertar' inserta un elemento pasado por parametro en la posicion en la que estamos parados en la lista
	Insertar(T)
	//'Borrar' borra y devuelve el elemento sobre el cual estemos parados en la lista, en caso de que
	//se llame a esta funcion cuando el iterador ya haya terminado de iterar, devolvera un panic
	Borrar() T
}
