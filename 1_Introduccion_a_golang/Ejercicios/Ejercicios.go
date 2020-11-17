package main

import (
	"fmt"
)

func main() {
	Producto_de_dos_numeros_por_sumas_susesivas()
}

//1_Hacer leer dos numeros A y B, evalue la expresion
//T = 2 + A + A, almacenar el resultado en T

func Evalue_la_exprescio() {

	var A int
	var B int
	fmt.Println("Ingresar numero A: ")
	fmt.Scanln(&A)
	fmt.Println("Ingresar numero B: ")
	fmt.Scanln(&B)
	var T int
	T = 2 + A + B
	fmt.Println("A =", A, "B =", B, "\n", "T = 2 + A + B = ", T)
}

//2_Hacer leer G grados, M mintuos y S segundos comvertir a radianes
//y desplegar el resultado.

func De_grados_min_seg_a_radianes() {

	var G float32
	var M float32
	var S float32
	fmt.Println("ingresar grados: ")
	fmt.Scanln(&G)
	fmt.Println("ingresar minutos: ")
	fmt.Scanln(&M)
	fmt.Println("ingresar segundos: ")
	fmt.Scanln(&S)
	const pi = 3.1416
	var F float32 = pi / 180
	var RAD float32 = F * (G + M/60 + S/3600)
	fmt.Println(RAD)
}

//3_Producto de dos numeros por sumas susesivas.

func Producto_de_dos_numeros_por_sumas_susesivas() {
	var A int
	var B int
	fmt.Println("ingresar A: ")
	fmt.Scanln(&A)
	fmt.Println("ingresar B: ")
	fmt.Scanln(&B)
	var p int
	for i := 0; i < B; i++ {
		p = p + A
	}
	fmt.Println("El producto es: ", p)
}
