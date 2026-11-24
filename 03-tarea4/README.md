# Tarea 4 – Conversor de almacenamiento en Go

## 1. Propósito

El objetivo de esta tarea es organizar un programa de conversión de unidades de almacenamiento en Go, separando la lógica de conversión del programa principal mediante paquetes, y utilizando funciones y constantes exportadas y no exportadas.

## 2. Estructura del proyecto

```text
03-tarea4/
├── go.mod
├── main.go
├── README.md
└── internal/
    └── almacenamiento/
        └── almacenamiento.go
```

## 3. Instrucciones de uso

El programa permite convertir unidades de almacenamiento B, KB, MB y GB.

Formato:

```bash
go run . convertir <valor> <unidad_origen> <unidad_destino>
```

Ejemplos:

```bash
go run . convertir 2048 KB MB
go run . convertir 1 GB MB
go run . convertir 1500 B KB
go run . convertir 0.5 MB KB
```

## 4. Constantes y funciones

- B: representa un byte.
- KB: representa 1024 bytes.
- MB: representa 1024 KB.
- GB: representa 1024 MB.
- Bytes: permite obtener el valor en bytes de una unidad y comprobar si existe.
- Convertir: realiza la conversión de una unidad de almacenamiento a otra.
- redondear: redondea el resultado a dos decimales. Es una función privada del paquete.

## 5. Pruebas realizadas

Aquí se registran los 8 casos de prueba realizados, los resultados obtenidos y los códigos de salida.
PS C:\Users\USUARIO\Documents\03-tarea4> go run . convertir 2048 KB MB
2048 KB = 2.00 MB
PS C:\Users\USUARIO\Documents\03-tarea4> go run . convertir 1 GB MB
1 GB = 1024.00 MB
PS C:\Users\USUARIO\Documents\03-tarea4> go run . convertir 1500 B KB
1500 B = 1.46 KB
PS C:\Users\USUARIO\Documents\03-tarea4> go run . convertir 0.5 MB KB
0.5 MB = 512.00 KB
PS C:\Users\USUARIO\Documents\03-tarea4> go run . convertir abc KB MB
El valor debe ser un número no negativo
exit status 2
PS C:\Users\USUARIO\Documents\03-tarea4> go run . convertir -5 KB MB
El valor debe ser un número no negativo
exit status 2
PS C:\Users\USUARIO\Documents\03-tarea4> go run . convertir 5 TB MB
Unidad no válida
exit status 3
PS C:\Users\USUARIO\Documents\03-tarea4> go run . convertir 5 KB
Uso: go run . convertir <valor> <unidad_origen> <unidad_destino>
exit status 1
PS C:\Users\USUARIO\Documents\03-tarea4> $LASTEXITCODE
1
PS C:\Users\USUARIO\Documents\03-tarea4> go run . convertir 0.5 MB KB
0.5 MB = 512.00 KB
PS C:\Users\USUARIO\Documents\03-tarea4> $LASTEXITCODE
0
PS C:\Users\USUARIO\Documents\03-tarea4> notepad README.md

## 6. Preguntas

### ¿Por qué Bytes y Convertir empiezan con mayúscula?

Porque en Go los nombres que comienzan con mayúscula son exportados y pueden ser utilizados desde otros paquetes.

### ¿Por qué redondear empieza con minúscula?

Porque es una función no exportada. Solo puede utilizarse dentro de su propio paquete.

### ¿Qué ocurre al dejar redondear dentro del paquete almacenamiento?

La función permanece oculta para otros paquetes, lo que permite mantener la lógica interna protegida.

### ¿Qué demuestra el experimento del compilador?

Demuestra que Go no permite acceder desde main a una función no exportada de otro paquete.
