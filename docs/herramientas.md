# Herramientas que se van a usar y justificación de las mismas

## Lenguaje: Go

He elegido este lenguaje porque está bastante bien documentado, es simple,
ha sido desarrollado por Google y creo que puede serme útil en un futuro
para otros proyectos.

## Framework para logging: Logrus

Esta librería es la más utilizada por programadores Go para el logging de su
aplicación. He decidido usarlo porque aunque Go tiene su propia biblioteca,
Logrus hace uso de ella para incorporar mejoras, como el manejo de JSON.
El logueo de una aplicación en Go es un problema que ya está resuelto
y no es mi intención volver a resolverlo.

## Posible framework para manejar tests y mocks: Testify

Por el momento tan solo estoy usando la biblioteca de tests propia de Go,
pero puede que cuando tenga que incorporar mocks me sea más fácil hacer
uso de esta herramienta. No es algo seguro.
