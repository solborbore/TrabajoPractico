package main


type link struct{
  nombre string
  referencia iContenido

  }

func (c link) getNombre() string {
  	return c.nombre
  }

func (c link) tamanio() int {
  	return 0
  }

func (c link) esLiviano() bool {
  	return true
  }

type carpeta struct{
    nombre string
    contenidos  []iContenido
    }


func (c carpeta) getNombre() string {
    	return c.nombre
    }

func (c carpeta) tamanio() int {
    	return 0
    }

func (c carpeta) esLiviano() bool {
    	return all(c.contenidos, func(conte iContenido) bool { return conte.esLiviano() })
    }
