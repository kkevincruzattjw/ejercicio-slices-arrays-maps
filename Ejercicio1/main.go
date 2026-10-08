package main

import "fmt"

func main() {

	notas := [6][4]float64{
		{8, 9, 7, 10},
		{6, 8, 9, 7},
		{10, 9, 8, 10},
		{7, 6, 8, 9},
		{9, 8, 7, 8},
		{5, 7, 6, 8},
	}
	promedios := []float64{}

	sumaGeneral := 0.0

	for i := 0; i < 6; i++ {

		suma := 0.0
		mayor := notas[i][0]
		menor := notas[i][0]

		for j := 0; j < 4; j++ {

			suma = suma + notas[i][j]

			if notas[i][j] > mayor {
				mayor = notas[i][j]
			}

			if notas[i][j] < menor {
				menor = notas[i][j]
			}
		}
		promedio := suma / 4
		promedios = append(promedios, promedio)
		sumaGeneral = sumaGeneral + promedio

		fmt.Println("Estudiante", i+1)
		fmt.Println("Promedio:", promedio)
		fmt.Println("Nota mayor:", mayor)
		fmt.Println("Nota menor:", menor)
		fmt.Println("/////////////////")
	}

	promedioGeneral := sumaGeneral / 6

	fmt.Println("Promedio general de la clase:", promedioGeneral)
}
