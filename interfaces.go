package main

type iContenido interface {
	tamanio() int
	getNombre() string
	nombrePermitido() bool
	esLiviano() bool
	actualizarFechaModificacion()
	renombrar(nuevoNombre string)
}

type iBiblioteca interface {
	puedeSubirContenido(iContenido) bool
}

type iTipoCompresion interface {
	tamanioReducido(int) int
}
