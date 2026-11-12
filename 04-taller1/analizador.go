
package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	texto := "¡Programación en Go: rápido, simple y divertido! Año 2026."

	// 1. Contar bytes y runas
	bytes := len(texto)
	runas := utf8.RuneCountInString(texto)

	// 2. Contar palabras
	palabras := len(strings.Fields(texto))

	// Variables para los contadores
	mayusculas := 0
	digitos := 0
	tildes := 0

	// 3 y 4. Recorrer el texto y contar caracteres
	for _, caracter := range texto {

		if unicode.IsUpper(caracter) {
			mayusculas++
		}

		if unicode.IsDigit(caracter) {
			digitos++
		}

		if strings.ContainsRune("áéíóúÁÉÍÓÚñÑ", caracter) {
			tildes++
		}
	}

	// Mostrar resultados
	fmt.Println("Bytes:", bytes)
	fmt.Println("Runas:", runas)
	fmt.Println("Palabras:", palabras)
	fmt.Println("Mayúsculas:", mayusculas)
	fmt.Println("Dígitos:", digitos)
	fmt.Println("Con tilde/ñ:", tildes)
}
