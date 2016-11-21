package main

import (
	"reflect"
	s "strings"
	"testing"
	"time"
)

func enTamanioDeseado(contenido iContenido, tamanioDeseado int) bool {
	return contenido.tamanio() == tamanioDeseado
}

func Test_elTamanioDeUnArchivoDeTextoEsElLargoDeSusCadenasPor16(t *testing.T) {

	archivo := contenidoTexto{nombre: "hola", lineas: []string{"hola", "chau"}}

	valorDeseado := 128

	t.Run("Test_elTamanioDeUnArchivoDeTextoEsElLargoDeSusCadenasPor16", func(t *testing.T) {
		if !enTamanioDeseado(archivo, valorDeseado) {
			t.Errorf("Test sin exito")
		}
	})
}

func Test_elTamanioDeUnArchivoBinarioSinCompresionEsElTamanioDeSuSecuencia(t *testing.T) {

	comp := sinCompresion{}

	archivo := contenidoBinario{nombre: "hola", secuencia: []int64{1001, 1100}, compresion: comp}

	valorDeseado := 2

	t.Run("Test_elTamanioDeUnArchivoBinarioSinCompresionEsElTamanioDeSuSecuencia", func(t *testing.T) {
		if !enTamanioDeseado(archivo, valorDeseado) {
			t.Errorf("Test sin exito")
		}
	})
}

/*func Test_elTamanioDeUnArchivoBinarioConReflateEsElTamanioDeSuSecuenciaMenosUn20(t *testing.T) {

	comp := reflate{}

	archivo := contenidoBinario{nombre: "hola", secuencia: []int64{1001, 1100}, compresion: comp}

	valorDeseado := 2

	t.Run("Test_elTamanioDeUnArchivoBinarioConReflateEsElTamanioDeSuSecuenciaMenosUn20", func(t *testing.T) {
		if !enTamanioDeseado(archivo, valorDeseado) {
			t.Errorf("Test sin exito")
		}
	})
}

func Test_elTamanioDeUnArchivoBinarioConBsip2EsElTamanioDeSuSecuenciaMenosUn40SiElOriginalEsPesado(t *testing.T) {

	comp := bsip2{}

	archivo := contenidoBinario{nombre: "hola", secuencia: []int64{1001, 1100}, compresion: comp}

	valorDeseado := 2

	t.Run("Test_elTamanioDeUnArchivoBinarioConBsip2EsElTamanioDeSuSecuenciaMenosUn40SiElOriginalEsPesado", func(t *testing.T) {
		if !enTamanioDeseado(archivo, valorDeseado) {
			t.Errorf("Test sin exito")
		}
	})
}

func Test_elTamanioDeUnArchivoBinarioConBsip2EsElTamanioDeSuSecuenciaMenosUn10SiElOriginalNoEsPesado(t *testing.T) {

	comp := bsip2{}

	archivo := contenidoBinario{nombre: "hola", secuencia: []int64{1001, 1100}, compresion: comp}

	valorDeseado := 2

	t.Run("Test_elTamanioDeUnArchivoBinarioConBsip2EsElTamanioDeSuSecuenciaMenosUn10SiElOriginalNoEsPesado", func(t *testing.T) {
		if !enTamanioDeseado(archivo, valorDeseado) {
			t.Errorf("Test sin exito")
		}
	})
}*/

func cadenaMuyLarga() string {
	return s.Repeat("a", 1024*151)
}

func Test_elTamanioDeUnArchivoEsMayorA150MBYNoEsLiviano(t *testing.T) {

	archivo := contenidoTexto{nombre: "hola", lineas: []string{cadenaMuyLarga()}}

	t.Run("Test_elTamanioDeUnArchivoEsMayorA150MBYNoEsLiviano", func(t *testing.T) {
		if archivo.esLiviano() {
			t.Errorf("Test sin exito")
		}
	})
}

func Test_elTamanioDeUnArchivoEsMenorA150MBYEsLiviano(t *testing.T) {

	archivo := contenidoTexto{nombre: "hola", lineas: []string{"hola", "chau"}}

	t.Run("Test_elTamanioDeUnArchivoEsMenorA150MBYEsLiviano", func(t *testing.T) {
		if !archivo.esLiviano() {
			t.Errorf("Test sin exito")
		}
	})
}

type testDeSubida struct {
	name             string
	args             iContenido
	resultadoQuerido bool
}

func Test_bloqueosEnSubida(t *testing.T) {
	biblio := biblioteca{contenidos: make([]iContenido, 8), limiteIndividual: 40000}

	contenidoTextDeNombreLargo := contenidoTexto{nombre: cadenaMuyLarga(), lineas: []string{"", ""}}
	contenidoTextoNormal := contenidoTexto{nombre: "", lineas: []string{"", ""}}
	contenidoTextoQueYaExiste := contenidoTexto{nombre: "ya existo", lineas: []string{"", ""}}
	contenidoDemasiadoGrande := contenidoTexto{nombre: "soy grande", lineas: []string{cadenaMuyLarga(), cadenaMuyLarga()}}

	biblio.contenidos = append(biblio.contenidos, contenidoTextoQueYaExiste)

	tests := []testDeSubida{
		testDeSubida{name: "Un contenido tiene nombre muy largo y por ende no puede subirse", args: contenidoTextDeNombreLargo, resultadoQuerido: false},
		testDeSubida{name: "Un contenido cumple todo y puede subirse", args: contenidoTextoNormal, resultadoQuerido: true},
		testDeSubida{name: "Un contenido ya existe en la biblioteca y por ende no puede subirse", args: contenidoTextoQueYaExiste, resultadoQuerido: false},
		testDeSubida{name: "Un contenido es demasiado grande y por ende no puede subirse", args: contenidoDemasiadoGrande, resultadoQuerido: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if biblio.puedeSubirContenido(tt.args) != tt.resultadoQuerido {
				t.Errorf("Error en el test " + tt.name)
			}
		})
	}
}

func Test_fallasAlSubir(t *testing.T) {
	biblio := biblioteca{contenidos: make([]iContenido, 8), limiteIndividual: 40000}

	contenidoTextDeNombreLargo := contenidoTexto{nombre: cadenaMuyLarga(), lineas: []string{"", ""}}
	contenidoTextoNormal := contenidoTexto{nombre: "normal", lineas: []string{"", ""}}
	contenidoTextoNormal2 := contenidoTexto{nombre: "normal", lineas: []string{"", ""}}
	contenidoDemasiadoGrande := contenidoTexto{nombre: "soy grande", lineas: []string{cadenaMuyLarga(), cadenaMuyLarga()}}

	biblio.subirContenido(contenidoTextoNormal2)
	tests := []testDeSubida{
		testDeSubida{name: "Un contenido tiene nombre muy largo y da error al querer subirse", args: contenidoTextDeNombreLargo, resultadoQuerido: true},
		testDeSubida{name: "Un contenido cumple todo y se sube sin error", args: contenidoTextoNormal, resultadoQuerido: false},
		testDeSubida{name: "Un contenido ya existe en la biblioteca y da error al querer subirse", args: contenidoTextoNormal2, resultadoQuerido: true},
		testDeSubida{name: "Un contenido es demasiado grande y da error al querer subirse", args: contenidoDemasiadoGrande, resultadoQuerido: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if e := biblio.subirContenido(tt.args); (e != nil) != tt.resultadoQuerido {
				t.Errorf("Error en el test " + tt.name)
			}
		})
	}
}

type testDeBusqueda struct {
	name             string
	args             string
	resultadoQuerido []iContenido
}

func Test_buscarNombres(t *testing.T) {
	biblio := biblioteca{contenidos: make([]iContenido, 8), limiteIndividual: 40000}

	contenidoTextoNormal := contenidoTexto{nombre: "normal", lineas: []string{"", ""}}
	otroContenidoTextoNormal := contenidoTexto{nombre: "normal 2", lineas: []string{"", ""}}

	biblio.subirContenido(contenidoTextoNormal)
	biblio.subirContenido(otroContenidoTextoNormal)

	tests := []testDeBusqueda{
		testDeBusqueda{name: "Busco un nombre que no esta en la biblioteca, y la lista de resultado esta vacia", args: "asd", resultadoQuerido: []iContenido{}},
		testDeBusqueda{name: "Busco un nombre que existe en biblioteca, y recibo una lista de resultados", args: "normal", resultadoQuerido: []iContenido{contenidoTextoNormal, otroContenidoTextoNormal}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.resultadoQuerido, biblio.buscarPorNombre(tt.args)) {
				t.Errorf("Error en el test " + tt.name)
			}
		})
	}
}

/*type testDeFecha struct {
	name             string
	args             iContenido
	fAplicado        func(iContenido)
	resultadoQuerido t.TimeS
}

func Test_fechaModificacion(t *testing.T) {
	biblio := biblioteca{contenidos: make([]iContenido, 8), limiteIndividual: 40000}

	contenidoTextoNormal := contenidoTexto{nombre: "normal", lineas: []string{"", ""}}

	biblio.subirContenido(contenidoTextoNormal)
	biblio.subirContenido(otroContenidoTextoNormal)

	tests := []testDeBusqueda{
		testDeFecha{name: "Un contenido se sube a la biblioteca y su fecha de modificacion es la actual", args: contenidoTextoNormal, resultadoQuerido: []iContenido{}},
		testDeFecha{name: "Busco un nombre que existe en biblioteca, y recibo una lista de resultados", args: "normal", resultadoQuerido: []iContenido{contenidoTextoNormal, otroContenidoTextoNormal}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.resultadoQuerido, biblio.buscarPorNombre(tt.args)) {
				t.Errorf("Error en el test " + tt.name)
			}
		})
	}
}*/

func Test_UnaCarpetaEsLivianaSiTodosSusElementosLoSon(t *testing.T) {

	archivo := contenidoTexto{nombre: "hola", lineas: []string{"hola", "chau"}}
	carpeta := carpeta{nombre: "holas", contenidos: []iContenido{archivo}}

	t.Run("Test_UnaCarpetaEsLivianaSiTodosSusElementosLoSon", func(t *testing.T) {
		if !carpeta.esLiviano() {
			t.Errorf("Test sin exito")
		}
	})
}

func Test_RemoverUnElementoDeLaBiblioteca(t *testing.T) {

	archivo1 := contenidoTexto{nombre: "sol", lineas: []string{"holazzz", "chauuuu"}}
	archivo2 := contenidoTexto{nombre: "juan", lineas: []string{"freza", "talvezzzz"}}

	biblio := biblioteca{contenidos: []iContenido{archivo1, archivo2}, limiteIndividual: 40000}

	cantidadContenidos := len(biblio.contenidos)

	biblio.eliminarContenido(archivo1)

	t.Run("Test_RemoverUnElementoDeLaBiblioteca", func(t *testing.T) {
		if !(len(biblio.contenidos) == cantidadContenidos-1) {
			t.Errorf("Test sin exito")
		}
	})
} //no funciona este test porq no pudimos hacer el remove porq go es una kk

func Test_LosContenidosReferenciadosNoFueronSubidosYNoSeSube(t *testing.T) {
	unArchivo := contenidoTexto{nombre: "tom", lineas: []string{"Buenas", "Tardes"}}
	unLink := link{nombre: "link a tom", referencia: unArchivo , fechaModificacion: time.Now()}

	biblio := biblioteca{contenidos: []iContenido{}, limiteIndividual: 40000}

	biblio.subirContenido(unLink)

	t.Run("LosContenidosReferenciadosNoFueronSubidosYNoSeSube", func(t *testing.T) {
		if (len(biblio.contenidos) == 1) {
			t.Errorf("Test sin exito")
}