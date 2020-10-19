# Herramientas que se van a usar y justificación de las mismas

## Lenguaje: Go

He elegido este lenguaje porque está bastante bien documentado, es simple,
ha sido desarrollado por Google y creo que puede serme útil en un futuro
para otros proyectos.

## Herramienta para automatización de tests: Make

Esta herramienta permite al usuario interactuar con la aplicación de una manera
muy clara y sencilla, y me permite a mí como programador tener una lista
de acciones que puede realizar el usuario con la aplicación en un mismo
archivo. Usaré esta herramienta en lugar de el task runner de Go porque me ahorro tener que escribir los argumentos del comando "go" al ejecutar, testear o compilar la aplicación una y otra vez. Además es algo muy estandarizado dentro de la estructura de los proyectos programados en Go.

## Framework para logging: Logrus

Esta librería es la más utilizada por programadores Go para el logging de su
aplicación. He decidido usarlo porque aunque Go tiene su propia biblioteca,
Logrus hace uso de ella para incorporar mejoras, como el manejo de JSON.
El logueo de una aplicación en Go es un problema que ya está resuelto
y no es mi intención volver a resolverlo.

## Posible framework para manejar tests y mocks: Testify

Por el momento tan solo estoy usando la biblioteca de tests propia de Go,
porque creo que para las distintas tareas que tengo que realizar en mi
aplicación hasta ahora, para lo único que me ayudarían las bibliotecas
de test sería para generar un código más legible, y no aprovecharía su
funcionalidad a penas, ocupando un espacio de manera no eficiente.
Dicho esto, puede que cuando tenga que incorporar mocks me sea más fácil hacer uso de esta herramienta. No es algo seguro.
