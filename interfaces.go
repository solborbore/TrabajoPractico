package main

import "time"

type iContenido interface {
	tamanio() int
	getNombre() string
	nombrePermitido() bool
	esLiviano() bool
	fechaDeModificacion() time.Time
	estaRoto(b iBiblioteca) bool
	actualizarFechaModificacion()
	renombrar(nuevoNombre string)
}

type iBiblioteca interface {
	existeEnBiblioteca(contenido iContenido) bool
	puedeSubirContenido(iContenido) bool

	rebalza() bool
	noExisteEnBiblioteca(contenido iContenido) bool
	cumpleTamanioRequerido(contenido iContenido) bool
	permiteNombre(contenido iContenido) bool
	tamanio() float64
	subirContenido(contenido iContenido) error
	buscarPorNombre(nombre string) (contenidos []iContenido)
	cantContenidos() int
	ordenarPorFecha() []iContenido
	ultimos5Contenidos() [5]iContenido
}

type iBibliotecaP *interface {
	concretarSubida(contenido iContenido)
	eliminarContenido(iContenido)
	eliminarContenidosRotos()
}

type iTipoCompresion interface {
	tamanioReducido(int) int
}
