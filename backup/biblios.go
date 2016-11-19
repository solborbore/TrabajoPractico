package main

import "fmt"

func main() {
	archivo := contenidoTexto{nombre: "hola", lineas: []string{"hola", "chau"}}
	fmt.Println(archivo.tamanio())
}
