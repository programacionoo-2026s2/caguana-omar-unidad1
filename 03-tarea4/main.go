package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"tarea4/internal/almacenamiento"
)

const (
	ErrUso    = 1
	ErrValor  = 2
	ErrUnidad = 3
)

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintln(os.Stderr, "Uso: go run . convertir <valor> <unidad_origen> <unidad_destino>")
		os.Exit(ErrUso)
	}

	if os.Args[1] != "convertir" {
		fmt.Fprintln(os.Stderr, "Comando inválido")
		os.Exit(ErrUso)
	}

	valor, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil || valor < 0 {
		fmt.Fprintln(os.Stderr, "El valor debe ser un número no negativo")
		os.Exit(ErrValor)
	}

	origen := strings.ToUpper(os.Args[3])
	destino := strings.ToUpper(os.Args[4])

	_, existeOrigen := almacenamiento.Bytes(origen)
	_, existeDestino := almacenamiento.Bytes(destino)

	if !existeOrigen || !existeDestino {
		fmt.Fprintln(os.Stderr, "Unidad no válida")
		os.Exit(ErrUnidad)
	}

	resultado, err := almacenamiento.Convertir(valor, origen, destino)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(ErrUnidad)
	}

	fmt.Printf("%g %s = %.2f %s\n", valor, origen, resultado, destino)
}
