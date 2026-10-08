package ejercicio2

import "fmt"

func actividadGanadora(votos map[string]int) string {

    actividad := ""
    mayor := 0

	  for nombre, cantidad := range votos {

        if cantidad > mayor {
            mayor = cantidad
            actividad = nombre
        }
	}