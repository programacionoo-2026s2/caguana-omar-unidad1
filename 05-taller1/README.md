
# 05-taller1 - Estadísticas de notas

## Descripción

Este programa utiliza una función variádica llamada estadisticas para calcular la nota menor, la nota mayor y el promedio de un grupo de calificaciones. También utiliza resultados con nombre y un error centinela para controlar cuando no se reciben datos.

## Cómo ejecutar

Desde la carpeta 05-taller1, ejecutar el siguiente comando:

```bash
go run .
```

## Resultado

La ejecución del programa muestra:

```text
4 9.75 7.0625 <nil>
min=4.00 max=9.75 prom=7.06
min=7.00 max=10.00 prom=8.17
error: no se recibieron datos
```

Los resultados de las notas 8.5, 6, 9.75 y 4 son:
- Menor: 4
- Mayor: 9.75
- Promedio: 7.0625

Para el slice de notas 10, 7 y 7.5:
- Menor: 7
- Mayor: 10
- Promedio: 8.17

## Preguntas

**a) ¿Por qué los resultados se llaman menor y mayor y no min y max? ¿Qué pasaría con la llamada min(menor, n) si el resultado se llamara min?**

Se llaman menor y mayor porque son los nombres de los resultados que devuelve la función. Si uno de ellos se llamara min, podría ocultar la función integrada min de Go dentro de ese bloque y causar un conflicto al intentar utilizarla.

**b) ¿Qué hacen los tres puntos en estadisticas(notas...)? ¿Qué mensaje da Go si los quitas y por qué?**

Los tres puntos permiten pasar todos los elementos del slice como valores individuales a la función variádica. Si se quitan, Go produce un error porque se estaría intentando pasar un slice de tipo []float64 como si fuera un único valor float64.
