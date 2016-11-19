package main

import "time"

type contenidoBinario struct {
	nombre            string
	secuencia         []int64 //Se presupone que cada uno de los int representa a un byte
	compresion        iTipoCompresion
	fechaModificacion time.Time
}

func (c contenidoBinario) tamanio() int {
	return c.compresion.tamanioReducido(len(c.secuencia))
}

func (c contenidoBinario) getNombre() string {
	return c.nombre
}

func (c contenidoBinario) nombrePermitido() bool {
	return len(c.nombre) < 200
}

func (c contenidoBinario) esLiviano() bool {
	return c.tamanio()/1024 < 150
}

func (c contenidoBinario) actualizarFechaModificacion() {
	c.fechaModificacion = time.Now()
} /* tuvimos que cambiar este metodo/funcion porque si le pasabamos el puntero y no el contenido, este no era implementado por la interfaz,
no se podia pasar por parametro en funciones que esperaban un type iContenido */

func (c contenidoBinario) renombrar(nuevoNombre string) {
	c.nombre = nuevoNombre
	c.actualizarFechaModificacion()
}

func (c contenidoBinario) estaRoto(b biblioteca) bool {
	return !b.existeEnBiblioteca(c)
}
