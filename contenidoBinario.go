package main

type contenidoBinario struct {
	nombre            string
	secuencia         []int64 //Se presupone que cada uno de los int representa a un byte
	compresion        iTipoCompresion
	fechaModificacion t.Time
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

func (c *contenidoBinario) actualizarFechaModificacion() {
	c.fechaModificacion = t.Now()
}

func (c *contenidoBinario) renombrar(nuevoNombre string) {
	c.nombre = nuevoNombre
	c.actualizarFechaModificacion()
}
