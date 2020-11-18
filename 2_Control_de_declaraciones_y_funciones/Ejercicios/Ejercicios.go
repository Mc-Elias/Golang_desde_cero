package main

import "fmt"

func main() {
	//Valor_absoluto()
	//Desplegar_mayor_de_tres_numeros()
	//Numero_Primo()
	//Factorial_por_sumas()
	//Genarando_n_primeros_primos()
	NumPositivo_cero_numNegativo()
}

//ejecicios con if,for,swithc:
/*1_Hacer leer un numero entero X, hallar su valor absoluto. Desplegar el resultado*/
func Valor_absoluto() {

	var x int
	var y int

	fmt.Println("Ingresar numero x: ")
	fmt.Scanln(&x)

	if x < 0 {
		y = x * (-1)
	} else {
		y = x
	}
	fmt.Println("|", x, "|", " = ", y)
}

/*2_Dado tres numeros A,B,C determinar y desplegar cual es el mayor*/
func Desplegar_mayor_de_tres_numeros() {
	var a, b, c int
	fmt.Println("Ingresar numeros a,b,c: ")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)

	fmt.Println("El mayor es: ")
	if a > b {
		if a > c {
			fmt.Println(a)
		} else {
			fmt.Println(c)
		}
	} else {
		if b > c {
			fmt.Println(b)
		} else {
			fmt.Println(c)
		}
	}
}

/*3_Dado un numero x determinar si es un numero primo, desplegar el numero y el mensaje*/
func Numero_Primo() {
	var x int
	fmt.Println("Leer x")
	fmt.Scanln(&x)

	var cde int = 0
	for i := 1; i <= x; i++ {
		if (x/i)*i == x {
			cde = cde + 1
		}
	}
	if cde == 2 {
		fmt.Println(x, " es primo")
	} else {
		fmt.Println(x, " no es primo")
	}
}

/*4_Calcular el factorial de un numero N por sumas, desplegar resultado.*/
func Factorial_por_sumas() {

	var N, F, SF int
	fmt.Println("intro N: ")
	fmt.Scanln(&N)
	F = 1
	SF = 0

	for i := 1; i <= N; i++ {
		for j := 1; j <= i; j++ {
			SF = SF + F
		}
		F = SF
		SF = 0
	}

	fmt.Println(N, "! = ", F)
}

//Ejecicio usando while en go
/*5_Generar los N primeros numeros primeros y desplegar, con un solo proceso repetitivo*/
func Genarando_n_primeros_primos() {

	var N, CP, SW, P, K int
	fmt.Println("intro N: ")
	fmt.Scanln(&N)
	CP = 0
	SW = 0
	P = 2
	K = 2

	for SW == 0 { // En go la sentencia while se instancia de esta forma
		if P%K == 0 {
			if K == P {
				fmt.Println(P)
				CP = CP + 1
				P = P + 1
				K = 2
				if CP == N {
					SW = 1
				}
			} else {
				P = P + 1
				K = 2
			}
		} else {
			K = K + 1
		}
	}
}

//Ejercicio completo, incluyendo for, if, switch.
/*6_Hacer leer un lote de N numeros. Determinar cuantas veces se presneta la situacion de un numero
positivo seguido de un cero y un negativo. Desplegar el resultado obtenido.*/

/*EJEMPLO: sea N = 8, X = 881, 705, 0, 12, 0, -14, 310, 15.    Salida, CS = 1, tal que 12, 0, -14 cumple
con el planteamiento del problema.*/

func NumPositivo_cero_numNegativo() {

	var N, L, CS, x int
	fmt.Println("INTRO N: ")
	fmt.Scanln(&N)
	L = 1
	CS = 0

	for i := 1; i <= N; i++ {

		fmt.Println("intro X: ")
		fmt.Scanln(&x)

		switch L {
		case 1:
			if x > 0 {
				L = 2
			} else {
				L = 1
			}
		case 2:
			if x == 0 {
				L = 3
			} else {
				if x > 0 {
					L = 2
				} else {
					L = 1
				}
			}
		case 3:
			if x < 0 {
				CS++
				L = 1
			} else {
				if x > 0 {
					L = 2
				} else {
					L = 1
				}
			}
		}
	}
	fmt.Println("CS = ", CS)
}
