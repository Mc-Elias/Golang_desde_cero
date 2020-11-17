package main

import "fmt"

func main() {
	/*
	   Estructuras de control.
	   Las estructuras de control de Go están relacionadas a las de C pero
	   difieren en importantes maneras. No hay bucles do o while, solo un ligeramente
	   generalizado for; switch es más flexible; if y switch aceptan una declaración de inicio
	   opcional como la del for; las declaraciones break y continue toman una etiqueta opcional
	   para identificar qué interrumpir o continuar; y hay nuevas estructuras de control
	   incluyendo un tipo switch y un multiplexor de comunicaciones multivía, select.
	   La sintaxis también es ligeramente diferente: no hay paréntesis y los cuerpos
	   siempre tienen que estar delimitados por llaves.
	*/

	//if
	var x int = 10

	if x > 0 {
		fmt.Println(x)
	}

	//For
	/*El for en GO esta unificada for y while, y no existe do-while*/

	/*
		// Como un for C
		for inicio; condición; incremento {
		}

		// Como un while C
		for condición {
		}

		// Como un for(;;) C
		for {
		}
	*/

	for i := 0; i < x; i++ {
		fmt.Println(i)
	}

}

//Switch
/*
	El switch de Go es más general que el de C. Las expresiones no es
	necesario que sean constantes o incluso enteros, los casos se evalúan
	e arriba hacia abajo hasta encontrar una coincidencia y si el switch no tiene
	una expresión este cambia a true. Por lo tanto es posible —e idiomático— escribir
	una cadena de if-else-if-else como un switch.
*/

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

//Funciones
func sigue(s string) string {
	fmt.Println("entrando a:", s)
	return s
}

func abandona(s string) {
	fmt.Println("dejando:", s)
}

func a() {
	defer abandona(sigue("a"))
	fmt.Println("en a")
}

func b() {
	defer abandona(sigue("b"))
	fmt.Println("en b")
	a()
}
