package main

type sinCompresion struct{}
type reflate struct{}
type bsip2 struct{}

func (c sinCompresion) tamanioReducido(original int) int {
	return original
}

func (c reflate) tamanioReducido(original int) int {
	return int(float32(original) * 0.8)
}

func (c bsip2) tamanioReducido(original int) int {
	if original > 100 {
		return int(float32(original) * 0.6)
	}
	return int(float32(original) * 0.9)
}