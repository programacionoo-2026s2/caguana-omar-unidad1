# Taller 1 - FizzBuzz en Go

## 1. Descripción

Este programa utiliza un bucle for y un switch sin condición para recorrer los números del 1 al 30.

Cuando un número es múltiplo de 3 imprime Fizz, cuando es múltiplo de 5 imprime Buzz y cuando es múltiplo de ambos imprime FizzBuzz. Si no cumple ninguna condición, imprime el número.

También utiliza tres contadores para registrar cuántas veces aparece cada palabra.

## 2. Cómo ejecutar

Para ejecutar el programa se utiliza el siguiente comando:

```bash
go run .
```

## 3. Resultado

### Resultado con el orden correcto

```text
1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz 16 17 Fizz 19 Buzz Fizz 22 23 Fizz Buzz 26 Fizz 28 29 FizzBuzz
Fizz=8 Buzz=4 FizzBuzz=2
```

### Experimento cambiando el orden

Al colocar el case de 3 antes del case de 15, el resultado de los contadores fue:

```text
Fizz=10 Buzz=4 FizzBuzz=0
```

## 4. Preguntas

### a) ¿Qué salida y qué conteos obtuviste al poner el case de 3 antes que el de 15, y por qué FizzBuzz dejó de aparecer?

Al colocar primero el caso de 3, los números 15 y 30 se imprimen como Fizz. Esto sucede porque ambos son múltiplos de 3 y de 15, pero Go ejecuta únicamente el primer case verdadero. Por eso FizzBuzz deja de aparecer y sus conteos quedan en cero.

### b) Cuando i vale 15, ¿cuántos case son verdaderos y cuál ejecuta Go? ¿Por qué no hace falta escribir break al final de cada case?

Cuando i vale 15, tres condiciones son verdaderas: la de 15, la de 3 y la de 5. Go ejecuta únicamente el primer case verdadero. No hace falta escribir break porque en Go el switch no continúa automáticamente con el siguiente case.
