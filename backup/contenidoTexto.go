package main

type contenidoTexto struct {
	nombre string
	lineas []string
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
