package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

const LimiteSubtitulo = 42

const (
	ErrorUso    = 1
	ErrorVacio  = 2
	ErrorLimite = 3
)

func main() {
	// Validar que exista exactamente un argumento.
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Uso: go run . \"texto\"")
		os.Exit(ErrorUso)
	}

	texto := os.Args[1]

	// Rechazar textos vacíos o que solo tengan espacios.
	if strings.TrimSpace(texto) == "" {
		fmt.Fprintln(os.Stderr, "Error: el texto no puede estar vacío.")
		os.Exit(ErrorVacio)
	}

	var letras, digitos, mayusculas, multibyte int
	var palabras int
	enPalabra := false

	// Recorrer el texto rune por rune.
	for _, r := range texto {
		if unicode.IsLetter(r) {
			letras++
		}

		if unicode.IsDigit(r) {
			digitos++
		}

		if unicode.IsUpper(r) {
			mayusculas++
		}

		if utf8.RuneLen(r) > 1 {
			multibyte++
		}

		// Contar palabras separadas por espacios.
		if unicode.IsSpace(r) {
			enPalabra = false
		} else if !enPalabra {
			palabras++
			enPalabra = true
		}
	}

	bytes := len(texto)
	runas := utf8.RuneCountInString(texto)

	// Mostrar el informe.
	fmt.Printf("texto: %q\n", texto)
	fmt.Printf("bytes: %d\n", bytes)
	fmt.Printf("runas: %d\n", runas)
	fmt.Printf("palabras: %d\n", palabras)
	fmt.Printf("letras: %d\n", letras)
	fmt.Printf("dígitos: %d\n", digitos)
	fmt.Printf("mayúsculas: %d\n", mayusculas)
	fmt.Printf("runas de más de 1 byte: %d\n", multibyte)

	// Comprobar el límite de subtítulos.
	if runas > LimiteSubtitulo {
		fmt.Printf("límite de 42 runas: excedido\n")
		fmt.Fprintln(os.Stderr, "Error: el texto supera el límite de 42 runas.")
		os.Exit(ErrorLimite)
	}

	fmt.Println("límite de 42 runas: dentro")
}
