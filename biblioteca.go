package main

import "errors"

type biblioteca struct {
	contenidos       []iContenido
	limiteIndividual float64
}

func (b biblioteca) existeEnBiblioteca(contenido iContenido) bool {
	return any(b.contenidos, func(conte iContenido) bool { return conte == contenido })
}

func (b biblioteca) cumpleTamanioRequerido(contenido iContenido) bool {
	return float64(contenido.tamanio()) <= b.limiteIndividual
}

func (b biblioteca) permiteNombre(contenido iContenido) bool {
	return contenido.nombrePermitido()
}

func (b biblioteca) tamanio() (sum float64) {
	if len(b.contenidos) != 0 {
		for _, contenido := range b.contenidos {
			sum += float64(contenido.tamanio())
		}
	}
	return
}

func (b biblioteca) rebalza() bool {
	return b.tamanio() > 500*b.limiteIndividual
}

func (b biblioteca) puedeSubirContenido(contenido iContenido) bool {
	condicionesDeContenido := []func(iContenido) bool{
		b.existeEnBiblioteca, b.cumpleTamanioRequerido, b.permiteNombre}
	for _, condicion := range condicionesDeContenido {
		if !condicion(contenido) {
			return false
		}
	}
	if b.rebalza() {
		return false
	}

	return true
}

func (b biblioteca) concretarSubida(contenido iContenido) {
	b.contenidos = append(b.contenidos, contenido)
}

func (b biblioteca) subirContenido(contenido iContenido) error {
	if b.puedeSubirContenido(contenido) {
		b.concretarSubida(contenido)
	} else {
		return errors.New("No se puede subir el contenido")
	}

	return nil
}
