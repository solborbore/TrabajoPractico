package main

import "errors"
import s "strings"

type superBiblioteca struct {
	contenidos       []iContenido
}

func (sb superBiblioteca) existeEnBiblioteca(contenido iContenido) bool {
	return any(sb.contenidos, func(conte iContenido) bool { return conte.getNombre() == contenido.getNombre() })
}

func (sb superBiblioteca) noExisteEnBiblioteca(contenido iContenido) bool {
	return !sb.existeEnBiblioteca(contenido)
}

func (sb superBiblioteca) permiteNombre(contenido iContenido) bool {
	return contenido.nombrePermitido()
}

func (sb superBiblioteca) tamanio() float64 {
	sum := 0.0
	for _, contenido := range sb.contenidos {
		if contenido != nil {
			sum += float64(contenido.tamanio())
		}
	}
	return sum
}

func (sb superBiblioteca) puedeSubirContenido(contenido iContenido) bool {
	condicionesDeContenido := []func(iContenido) bool{
		sb.noExisteEnBiblioteca, sb.permiteNombre}
	for _, condicion := range condicionesDeContenido {
		if !condicion(contenido) {
			return false
		}
	}

	if !contenido.estaRoto(sb) {
		return false
	}

	return true
}

func (sb superBiblioteca) concretarSubida(contenido iContenido) {
	contenido.actualizarFechaModificacion()
	(&sb).contenidos = append(sb.contenidos, contenido)
}

func (sb superBiblioteca) subirContenido(contenido iContenido) error {
	if sb.puedeSubirContenido(contenido) {
		sb.concretarSubida(contenido)
	} else {
		return errors.New("No se puede subir el contenido")
	}

	return nil
}

func (sb superBiblioteca) buscarPorNombre(nombre string) (contenidos []iContenido) {
	contenidos = filter(sb.contenidos, func(conte iContenido) bool { return s.Contains(conte.getNombre(), nombre) })
	return
}

func (sb superBiblioteca) eliminarContenido(contenido iContenido) {

	(&sb).contenidos = filter(sb.contenidos, func(conte iContenido) bool { return conte.getNombre() != contenido.getNombre() })

}

func (sb superBiblioteca) eliminarContenidosRotos() {
	(&sb).contenidos = filter(sb.contenidos, func(conte iContenido) bool { return !conte.estaRoto(sb) })
}

func (sb superBiblioteca) cantContenidos() int {
	return len(sb.contenidos)
}

func (sb superBiblioteca) ordenarPorFecha() []iContenido {
	var aux iContenido
	var contenidosOrdenados = sb.contenidos
	for i := 0 ; i< sb.cantContenidos() ; i++ {
		for j := 0 ; j < sb.cantContenidos() - i ; j++ {
			if contenidosOrdenados[j].fechaDeModificacion().After(contenidosOrdenados[j+1].fechaDeModificacion()){
				aux = contenidosOrdenados[j]
				contenidosOrdenados[j] = contenidosOrdenados[j+1]
				contenidosOrdenados[j+1] = aux
			}
		}
	}
	return contenidosOrdenados
}

func (sb superBiblioteca) ultimos5Contenidos() [5]iContenido {
	var ultimos5 [5]iContenido
	for i:=0 ; i< 5 ; i++ {
		ultimos5[i] = sb.contenidos[i] 
	}
	return ultimos5
}