package main

import "time"

type link struct {
	nombre            string
	referencia        iContenido
	fechaModificacion time.Time
}

func (c link) getNombre() string {
	return c.nombre
}

func (c link) tamanio() int {
	return 0
}

func (c link) esLiviano() bool {
	return true
}

func (c link) estaRoto(b biblioteca) bool {
	return c.referencia.estaRoto(b)
}

func (c *link) actualizarFechaModificacion() {
	c.fechaModificacion = time.Now()
}

func (c *link) renombrar(nuevoNombre string) {
	c.nombre = nuevoNombre
	c.actualizarFechaModificacion()
}

func (c link) nombrePermitido() bool {
	return len(c.nombre) < 200
}

type carpeta struct {
	nombre            string
	contenidos        []iContenido
	fechaModificacion time.Time
}

func (c carpeta) nombrePermitido() bool {
	return len(c.nombre) < 200
}

func (c *carpeta) actualizarFechaModificacion() {
	c.fechaModificacion = time.Now()
}

func (c *carpeta) renombrar(nuevoNombre string) {
	c.nombre = nuevoNombre
	c.actualizarFechaModificacion()
}

func (c carpeta) getNombre() string {
	return c.nombre
}

func (c carpeta) tamanio() int {
	return 0
}

func (c carpeta) esLiviano() bool {
	return all(c.contenidos, func(conte iContenido) bool { return conte.esLiviano() })
}

func (c carpeta) estaRoto(b biblioteca) bool {
	return any(c.contenidos, func(conte iContenido) bool { return conte.estaRoto(b) })
}
