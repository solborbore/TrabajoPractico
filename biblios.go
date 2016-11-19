package main

import (
	"fmt"
	"reflect"
)

func main() {
	biblio := biblioteca{contenidos: make([]iContenido, 8), limiteIndividual: 40000}

	contenidoTextoNormal := contenidoTexto{nombre: "normal", lineas: []string{"", ""}}
	otroContenidoTextoNormal := contenidoTexto{nombre: "normal 2", lineas: []string{"", ""}}

	biblio.subirContenido(contenidoTextoNormal)
	biblio.subirContenido(otroContenidoTextoNormal)

	fmt.Println(reflect.DeepEqual([]iContenido{}, biblio.buscarPorNombre("asd")))
	fmt.Println(reflect.DeepEqual([]iContenido{contenidoTextoNormal, otroContenidoTextoNormal}, biblio.buscarPorNombre("normal")))
}
