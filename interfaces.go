package main

import "time"

type iContenido interface {
	tamanio() int
	getNombre() string
	nombrePermitido() bool
	esLiviano() bool
	actualizarFechaModificacion()
	fechaDeModificacion() time.Time
	renombrar(nuevoNombre string)
	estaRoto(b iBiblioteca) bool
}

type iBiblioteca interface {
	puedeSubirContenido(iContenido) bool
	eliminarContenido(iContenido)
	eliminarContenidosRotos()
	existeEnBiblioteca(contenido iContenido) bool
}

type iTipoCompresion interface {
	tamanioReducido(int) int
}
