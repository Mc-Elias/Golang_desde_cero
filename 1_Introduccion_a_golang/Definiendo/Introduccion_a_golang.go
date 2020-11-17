package main

import "fmt"

func main() {

	/*VARIABLES*/
	fmt.Println("---------------------------------------")
	//definiendo variables con su type
	var entero int
	entero = 2 + 6

	//define multiples variables con su type
	var vnombre1, vnombre2, vnombre3 string
	vnombre1 = "A"
	vnombre2 = "B"
	vnombre3 = "C"

	//define variable con un valor inicial
	var Nombre string = "marcos"

	//define multiples variables con valor inicial
	var a, b, c int = 1, 2, 3
	fmt.Println("Entero: ", entero, "cadenas: ",
		vnombre1, vnombre2, vnombre3, "nombre: ", Nombre, "Numbers: ", a, b, c)

	/*CONSTANTES*/
	fmt.Println("---------------------------------------")
	const Pi float32 = 3.1515926
	const i = 10000
	const MaxThread = 10
	const prefix = "astaxie_"

	fmt.Println(Pi, "\n", i, "\n", MaxThread, "\n", prefix)

	/*BOOLEAN*/
	fmt.Println("---------------------------------------")
	var habilitado, dishabilitado = true, false
	fmt.Println(habilitado, " ", dishabilitado)

	/*ARRAY*/
	fmt.Println("---------------------------------------")
	var ar [10]int
	ar[0] = 42
	ar[1] = 13
	fmt.Println("El primer elemento es: ", ar[0])
	fmt.Println("l ultimo elemento es: ", ar[9])
	p := [3]int{1, 2, 3}
	q := [10]int{1, 2, 3}
	r := [...]int{4, 5, 6}
	fmt.Println(p[0], "\n", q[6], "\n", r[2])

	/*REBANADA DE GO*/
	fmt.Println("---------------------------------------")
	slice := []byte{'a', 'b', 'c', 'd'}
	var arr = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

	fmt.Println(slice[1], " ", arr[0])

	/*MAP*/
	fmt.Println("---------------------------------------")
	/*map
	El map se comporta como un diccionario en Python. Usa el formulario map[keyType]valueType para definirlo.
	Veamos un poco de código. Los valores 'set' y 'get' en el mapa son similares a slice , sin embargo el índice en slice sólo puede ser de
	escribe 'int' mientras que el mapa puede usar mucho más que eso: por ejemplo int , string , o lo que quieras. Además, todos son capaces
	para usar == y != para comparar valores*/
	numbers := make(map[string]int)
	numbers["uno"] = 1
	numbers["diez"] = 10
	numbers["tres"] = 3

	fmt.Println("el tercer numero es: ", numbers["tres"])

	/*LEER DATOS DESDE TECLADO*/
	fmt.Println("---------------------------------------")
	var n int //Declarar variable y tipo antes de escanear, esto es obligatorio
	fmt.Scanln(&n)
	fmt.Println(n)
}
