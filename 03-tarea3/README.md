# Tarea 3 - Conversor de almacenamiento en Go

## 1. Descripción

Este programa permite convertir unidades de almacenamiento como B, KB, MB y GB. Se ejecuta desde la línea de comandos, recibe el valor y las unidades de origen y destino, valida los datos ingresados y muestra el resultado con dos decimales.

Las unidades se convierten utilizando múltiplos de 1024. El programa también informa errores cuando el valor no es numérico, es negativo o se utiliza una unidad no válida.

## 2. Cómo ejecutar

Para ejecutar el programa se utiliza el siguiente formato:

```bash
go run . <valor> <unidad_origen> <unidad_destino>
```

Ejemplos:

```bash
go run . 2048 KB MB
go run . 1 GB MB
go run . 1500 B KB
go run . 0.5 MB KB
```

Códigos de salida:

- 0: ejecución correcta.
- 1: cantidad incorrecta de argumentos.
- 2: valor inválido o negativo.
- 3: unidad no válida.

## 3. Pruebas realizadas

| Comando | Resultado esperado | Código |
|---|---|---|
| go run . 2048 KB MB | 2048 KB = 2.00 MB | 0 |
| go run . 1 GB MB | 1 GB = 1024.00 MB | 0 |
| go run . 1500 B KB | 1500 B = 1.46 KB | 0 |
| go run . 0.5 MB KB | 0.5 MB = 512.00 KB | 0 |
| go run . abc KB MB | Error de valor | 1 |
| go run . -5 KB MB | Error de valor | 1 |
| go run . 5 TB MB | Error de unidad | 1|
| go run . 5 KB | Mensaje de uso | 1 |

## 4. Preguntas

### a) ¿Por qué el valor está en os.Args[1] y no en os.Args[0]?

Porque os.Args[0] contiene el nombre o la ruta del programa ejecutado. Los argumentos que escribe el usuario empiezan en la posición 1, por eso el valor se encuentra en os.Args[1].

### b) ¿Por qué los mensajes de error se escriben en stderr y no con fmt.Println?

Porque stderr es el canal destinado a los mensajes de error. Esto permite separar los errores de la salida normal del programa.

### c) ¿Qué pasaría si escribieras valor, _ := strconv.ParseFloat(...) e ignoraras el error al ejecutar convertir abc KB MB?

El programa ignoraría el error que indica que abc no es un número. Por eso no podría validar correctamente el dato ingresado y no mostraría el error correspondiente con el código de salida 2.

### d) ¿Qué ventaja tiene declarar los códigos de salida con const en lugar de escribir 1, 2 y 3 dentro de os.Exit?

Permite identificar qué significa cada código, hace el programa más fácil de leer y evita tener números sin explicación repartidos en el código.
