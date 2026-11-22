package almacenamiento

import (
	"fmt"
	"math"
	"strings"
)

const (
	// B representa un byte.
	B = 1.0

	// KB representa 1024 bytes.
	KB = 1024.0

	// MB representa 1024 KB.
	MB = 1024.0 * KB

	// GB representa 1024 MB.
	GB = 1024.0 * MB
)

// Bytes devuelve los bytes correspondientes a una unidad.
// Si la unidad existe devuelve true.
func Bytes(unidad string) (float64, bool) {
	switch strings.ToUpper(unidad) {
	case "B":
		return B, true
	case "KB":
		return KB, true
	case "MB":
		return MB, true
	case "GB":
		return GB, true
	default:
		return 0, false
	}
}

// Convertir convierte un valor de una unidad a otra.
func Convertir(valor float64, origen, destino string) (float64, error) {
	bytesOrigen, existeOrigen := Bytes(origen)
	bytesDestino, existeDestino := Bytes(destino)

	if !existeOrigen || !existeDestino {
		return 0, fmt.Errorf("unidad no válida")
	}

	resultado := valor * bytesOrigen / bytesDestino

	return redondear(resultado), nil
}

// redondear redondea un número a dos decimales.
func redondear(valor float64) float64 {
	return math.Round(valor*100) / 100
}
