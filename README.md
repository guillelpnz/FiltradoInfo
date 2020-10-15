# TextAnalyzer

<img src="docs/imagenes/logo.jpg" alt="logo" width=400/>

## Descripción

El proyecto consiste en una API REST programada en [Go](https://golang.org/) que revisa textos de todo tipo, pudiendo encontrar así redundancias en ellos, o hacer estadísticas sobre el uso del lenguaje en diferentes discursos. (Posiblemente acabe resolviendo algunos, ofreciendo sinónimos). El paquete principal del proyecto está [aquí](src/texto).

## Motivación

Durante el trayecto que llevo recorrido de carrera, he tenido que hacer numerosas documentaciones, exposiciones, explicaciones, etc. Esto hizo que me diera cuenta de que paso bastante tiempo revisando si uso palabras de manera redundante. Por lo que se me ocurrió esta pequeña API que facilita el trabajo de analizar textos.

## Uso de la aplicación

1. Para poder hacer uso de TextAnalyzer debes previamente
[instalar](https://golang.org/dl/) y [configurar](https://golang.org/doc/install)
Go correctamente en tu sistema.

2. Debes descargarte este repositorio.

- Para poder testear la aplicación, debes estar en el directorio raíz del
proyecto y ejecutar el comando:

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; `make test`

- Para ver las cabeceras de los métodos del paquete texto debes estar en el
directorio raíz del proyecto y ejecutar el comando:

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; `make doc`

## Herramientas utilizadas

El lenguaje que se va a utilizar es Go. Para conocer el resto de herramientas [haz click aquí](docs/herramientas.md).

## Configuración de Git

Para conocer la configuración previa de git [haz click aquí](docs/documentacion.md).

## Pasos seguidos para desarrollar el proyecto

Para conocer los pasos seguidos para la realización del proyecto [haz click aquí](docs/pasos.md).

## Interfaz inicial

[Acceder al archivo](src/texto.go)

## Testeo de la interfaz

[Acceder al archivo](src/texto_test.go)

## Funciones auxiliares

[Acceder al archivo](src/funcs.go)

## Testeo de funciones auxiliares

[Acceder al archivo](src/funcs_test.go)

## Archivo iv.yaml

[Acceder al archivo](iv.yaml)

## Issues cerrados

[Acceder al archivo](https://github.com/guillelpnz/TextAnalyzer/issues?q=is%3Aissue+is%3Aclosed)

## Issues abiertos

[Acceder al archivo](https://github.com/guillelpnz/TextAnalyzer/issues)
