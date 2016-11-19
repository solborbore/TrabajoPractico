package main

type iContenido interface {
	tamanio() int
	getNombre() string
	nombrePermitido() bool
	esLiviano() bool
	actualizarFechaModificacion()
	renombrar(nuevoNombre string)
	estaRoto(b biblioteca) bool
}

type iBiblioteca interface {
	puedeSubirContenido(iContenido) bool
	eliminarContenido(iContenido)
	eliminarContenidosRotos()
}

type iTipoCompresion interface {
	tamanioReducido(int) int
}
