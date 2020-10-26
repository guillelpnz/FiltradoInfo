# TextAnalyzer

## Descripción

El proyecto consiste en una API REST programada en [Go](https://golang.org/) que revisa textos de todo tipo, pudiendo encontrar así redundancias en ellos, o hacer estadísticas sobre el uso del lenguaje en diferentes discursos. (Posiblemente acabe resolviendo algunos, ofreciendo sinónimos). El paquete principal del proyecto está [aquí](src/texto).

## Elección del contenedor base

Como contenedor base he elegido golang:alpine3.12. He tomado esta decisión porque
en velocidad no había una diferencia significativa golang:latest,
golang:alpine y golang:1.15.3-alpine. En cuanto al espacio,
todos los golang:alpine pesan cerca de 300mb, mientras que golang:latest pesa más
de 800mb. Estas fueron las [pruebas de velocidad que hice](https://github.com/guillelpnz/TextAnalyzer/blob/master/docs/pruebas_velocidad.md)

## Dockerfile correcto

[Dockerfile del proyecto](https://github.com/guillelpnz/TextAnalyzer/blob/master/Dockerfile)

## Uso de GitHub Container Registry

[Página de mi paquete](https://github.com/users/guillelpnz/packages/container/package/textanalyzer%2Fgolang-alpine)

## Docker Hub

[Página de mi container](https://hub.docker.com/r/guillelpnz/textanalyzer/tags)

Para que se construya automáticamente, hay que conectar GitHub a Docker Hub mediante un OAUTH. Posteriormente, acceder al apartado Manage Repository/Builds/Configure Automated Builds y ahí activar la pestaña: autobuild.

## Prueba actualizacion en GitHub Container Registry

<!-- ## Motivación

Durante el trayecto que llevo recorrido de carrera, he tenido que hacer numerosas documentaciones, exposiciones, explicaciones, etc. Esto hizo que me diera cuenta de que paso bastante tiempo revisando si uso palabras de manera redundante. Por lo que se me ocurrió esta pequeña API que facilita el trabajo de analizar textos. -->

<!-- ## Uso de la aplicación

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

El lenguaje que se va a utilizar es Go. Estas son el resto de [herramientas](docs/herramientas.md). -->
<!-- 

## Interfaz inicial

[Interfaz inicial](src/texto.go)

## Testeo de la interfaz

[Testeo de la interfaz](src/texto_test.go)

## Funciones auxiliares

[Funciones auxiliares](src/funcs.go)

## Testeo de funciones auxiliares

[Testeo de funciones auxiliares](src/funcs_test.go)

## Archivo iv.yaml

[Fichero iv.yaml](iv.yaml) -->

## Issues cerrados

[Issues cerrados](https://github.com/guillelpnz/TextAnalyzer/issues?q=is%3Aissue+is%3Aclosed)

## Issues abiertos

[Issues abiertos](https://github.com/guillelpnz/TextAnalyzer/issues)
