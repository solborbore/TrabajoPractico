package main

import t "time"

type contenidoTexto struct {
	nombre            string
	lineas            []string
	fechaModificacion t.Time
}

func (c contenidoTexto) nombrePermitido() bool {
	return len(c.nombre) < 200
}

func (c contenidoTexto) getNombre() string {
	return c.nombre
}

func (c contenidoTexto) tamanio() int {
	return 16 * largoDeCadenas(c.lineas)
}

func (c contenidoTexto) esLiviano() bool {
	return c.tamanio()/1024 < 150
}

func (c *contenidoTexto) actualizarFechaModificacion() {
	c.fechaModificacion = t.Now()
}

func (c *contenidoTexto) renombrar(nuevoNombre string) {
	c.nombre = nuevoNombre
	c.actualizarFechaModificacion()
}
