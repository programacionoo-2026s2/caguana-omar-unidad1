Tarea 03 - Ámbito y analizador Shadow

1. ¿Qué imprime el programa?

El programa original imprime:

total: 0

Esto ocurre porque dentro del ciclo for se declara otra variable llamada total, utilizando el operador :=.

2. ¿Por qué ocurre? (Ámbito)

Esto ocurre por el ámbito de las variables. En Go, una variable declarada dentro de un bloque solamente existe dentro de ese bloque.

En este caso, existen dos variables llamadas total: una fuera del ciclo y otra dentro del ciclo for.

La variable interior oculta a la variable exterior. Por eso, cuando termina el ciclo, la función devuelve el valor de la variable exterior, que sigue siendo 0.

Este comportamiento se conoce como shadowing u ocultamiento de variables.

3. Corrección cambiando un solo carácter

Para corregir el programa se cambia el operador := por = en la siguiente línea:

Antes:

total := total + i

Después:

total = total + i

De esta manera se utiliza la variable exterior y se acumulan correctamente los valores.

El resultado corregido es:

total: 6

4. Investigación del analizador Shadow

Shadow es un analizador de código de Go que permite detectar posibles problemas relacionados con variables que tienen el mismo nombre en diferentes ámbitos.

Su función es identificar cuando una variable declarada dentro de un bloque oculta a otra variable declarada en un ámbito exterior.

Esto puede provocar errores porque el programador puede pensar que está modificando una variable cuando realmente está utilizando otra diferente.

Sus principales funciones son:

Detectar posibles variables ocultas.

Identificar errores relacionados con el ámbito.

Ayudar a mejorar la calidad del código.

Evitar confusiones al utilizar variables con nombres repetidos.

5. Ejecución de go vet ./...

El comando go vet ./... analiza el código del proyecto para encontrar posibles errores y construcciones sospechosas.

Sin embargo, el analizador Shadow no viene habilitado de forma predeterminada en go vet.

Para detectar este problema se debe instalar el analizador Shadow mediante:

go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest

Después se puede ejecutar con:

go vet -vettool="$env:USERPROFILE\go\bin\shadow.exe" ./...

De esta manera se pueden detectar posibles problemas de ocultamiento de variables.

6. Conclusión

Con esta actividad aprendí que el ámbito determina dónde puede utilizarse una variable dentro de un programa.

También aprendí que utilizar := en lugar de = puede crear una nueva variable y ocultar otra que ya existe.

Finalmente, comprendí que el analizador Shadow ayuda a detectar este tipo de situaciones y permite escribir programas más claros y evitar errores.