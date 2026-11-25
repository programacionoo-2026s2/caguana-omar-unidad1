# Tarea 3 – Analizador de texto en Go

## Descripción

Este programa está desarrollado en Go y permite analizar un texto ingresado por el usuario mediante argumentos de línea de comandos.

El programa realiza el conteo de bytes, runas, palabras, letras, dígitos, mayúsculas y runas que ocupan más de un byte en UTF-8.

También valida que el texto no esté vacío y que no supere el límite de 42 runas.

## Funcionalidades

- Recepción de texto mediante argumentos.
- Validación de argumentos.
- Validación de texto vacío.
- Conteo de bytes.
- Conteo de runas Unicode.
- Conteo de palabras.
- Conteo de letras.
- Conteo de dígitos.
- Conteo de letras mayúsculas.
- Conteo de runas multibyte.
- Validación del límite de 42 runas.
- Manejo de códigos de salida.

## Tecnologías utilizadas

- Go
- Visual Studio Code
- Git
- GitHub

## Ejecución

Para ejecutar el programa:

```bash
go run . "Canción de cuna 2026"
```

## Ejemplo

```text
texto: "Canción de cuna 2026"
bytes: 21
runas: 20
palabras: 4
letras: 13
dígitos: 4
mayúsculas: 1
runas de más de 1 byte: 1
límite de 42 runas: dentro
```

## Autor

Omar Caguana

## Tareas

- [03-tarea3 - Conversor de almacenamiento](./03-tarea3/)

## 04-taller2 - FizzBuzz en Go

Programa que utiliza for, switch sin condición y contadores para resolver el ejercicio FizzBuzz del 1 al 30.

[Ver taller 04-taller1](./04-taller1/)
