package main

import "errors"
import s "strings"

type biblioteca struct {
	contenidos       []iContenido
	limiteIndividual float64
}

func (b biblioteca) existeEnBiblioteca(contenido iContenido) bool {
	return any(b.contenidos, func(conte iContenido) bool { return conte.getNombre() == contenido.getNombre() })
}

func (b biblioteca) noExisteEnBiblioteca(contenido iContenido) bool {
	return !b.existeEnBiblioteca(contenido)
}

func (b biblioteca) cumpleTamanioRequerido(contenido iContenido) bool {
	return float64(contenido.tamanio()) <= b.limiteIndividual
}

func (b biblioteca) permiteNombre(contenido iContenido) bool {
	return contenido.nombrePermitido()
}

func (b biblioteca) tamanio() float64 {
	sum := 0.0
	for _, contenido := range b.contenidos {
		if contenido != nil {
			sum += float64(contenido.tamanio())
		}
	}
	return sum
}

func (b biblioteca) rebalza() bool {
	return b.tamanio() > 500*b.limiteIndividual
}

func (b biblioteca) puedeSubirContenido(contenido iContenido) bool {
	condicionesDeContenido := []func(iContenido) bool{
		b.noExisteEnBiblioteca, b.cumpleTamanioRequerido, b.permiteNombre}
	for _, condicion := range condicionesDeContenido {
		if !condicion(contenido) {
			return false
		}
	}

	if !contenido.estaRoto(b) {
		return false
	}

	if b.rebalza() {
		return false
	}

	return true
}

func (b biblioteca) concretarSubida(contenido iContenido) {
	contenido.actualizarFechaModificacion()
	(&b).contenidos = append(b.contenidos, contenido)
}

func (b biblioteca) subirContenido(contenido iContenido) error {
	if b.puedeSubirContenido(contenido) {
		b.concretarSubida(contenido)
	} else {
		return errors.New("No se puede subir el contenido")
	}

	return nil
}

func (b biblioteca) buscarPorNombre(nombre string) (contenidos []iContenido) {
	contenidos = filter(b.contenidos, func(conte iContenido) bool { return s.Contains(conte.getNombre(), nombre) })
	return
}

func (b biblioteca) eliminarContenido(contenido iContenido) {

	(&b).contenidos = filter(b.contenidos, func(conte iContenido) bool { return conte.getNombre() != contenido.getNombre() })

}

func (b biblioteca) eliminarContenidosRotos() {
	(&b).contenidos = filter(b.contenidos, func(conte iContenido) bool { return !conte.estaRoto(b) })
}
