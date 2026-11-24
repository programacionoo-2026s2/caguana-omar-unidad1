package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ErrUso    = 1
	ErrValor  = 2
	ErrUnidad = 3
)

func bytesUnidad(unidad string) (float64, bool) {
	switch strings.ToUpper(unidad) {
	case "B":
		return 1, true
	case "KB":
		return 1024, true
	case "MB":
		return 1024 * 1024, true
	case "GB":
		return 1024 * 1024 * 1024, true
	default:
		return 0, false
	}
}

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintln(os.Stderr, "Uso: go run . <valor> <unidad_origen> <unidad_destino>")
		os.Exit(ErrUso)
	}

	valor, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil || valor < 0 {
		fmt.Fprintln(os.Stderr, "El valor debe ser un número no negativo")
		os.Exit(ErrValor)
	}

	origen := strings.ToUpper(os.Args[2])
	destino := strings.ToUpper(os.Args[3])

	bytesOrigen, existeOrigen := bytesUnidad(origen)
	bytesDestino, existeDestino := bytesUnidad(destino)

	if !existeOrigen || !existeDestino {
		fmt.Fprintln(os.Stderr, "Unidad no válida")
		os.Exit(ErrUnidad)
	}

	resultado := valor * bytesOrigen / bytesDestino

	fmt.Printf("%g %s = %.2f %s\n", valor, origen, resultado, destino)
}
