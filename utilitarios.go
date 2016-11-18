package main

func sumaDeNumeros(numeros []int) (sum int) {
	for _, num := range numeros {
		sum += num
	}
	return
}

func largoDeCadenas(cadenas []string) int {
	sumaCaracteres := 0
	for _, cadena := range cadenas {
		sumaCaracteres += len(cadena)
	}
	return sumaCaracteres
}

type a interface{}

func any(vs []iContenido, f func(iContenido) bool) bool {
	for _, v := range vs {
		if v != nil && f(v) {
			return true
		}
	}
	return false
}
