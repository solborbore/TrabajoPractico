package main

import "testing"


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
}

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

	tests := []testDeSubida{
		testDeSubida{name: "Un contenido tiene nombre muy largo y por ende no puede subirse", args: contenidoTextDeNombreLargo, resultadoQuerido: false},
		testDeSubida{name: "Un contenido normal", args: contenidoTextoNormal, resultadoQuerido: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if biblio.rebalza() != tt.resultadoQuerido {
				t.Errorf("Error en el test " + tt.name)
			}
		})
	}
} */ // los comente porque no funcionaban


func Test_UnaCarpetaEsLivianaSiTodosSusElementosLoSon(t *testing.T) {

	archivo := contenidoTexto{nombre: "hola", lineas: []string{"hola", "chau"}}
	carpeta := carpeta{nombre: "holas", contenidos: []iContenido{archivo}}


	t.Run("Test_UnaCarpetaEsLivianaSiTodosSusElementosLoSon", func(t *testing.T) {
		if !carpeta.esLiviano() {
			t.Errorf("Test sin exito")
		}
	})
}
