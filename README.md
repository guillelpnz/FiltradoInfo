# TextAnalyzer

## Descripción

El proyecto consiste en una API REST programada en [Go](https://golang.org/) que revisa textos de todo tipo, pudiendo encontrar así redundancias en ellos, o hacer estadísticas sobre el uso del lenguaje en diferentes discursos. (Posiblemente acabe resolviendo algunos, ofreciendo sinónimos). El paquete principal del proyecto está [aquí](src/texto).

## Despliegue correcto y funcionando, con documentación de la conexión entre el repo en GitHub y Netlify/Vercel para despliegue continuo

La conexión entre Vercel y mi GitHub está en [el ejercicio 1 del tema](https://github.com/guillelpnz/Ejercicios/blob/master/Serverless/serverless.md)

Está hecha con un repo de prueba, pero con mi repositorio es igual.

## Integración dentro del proyecto general (es decir, como todo el código deberá tener sus issues y/o HU correspondientes)

En [mis issues](https://github.com/guillelpnz/TextAnalyzer/issues)
se puede ver el avance del código, que va referenciándolos.

## Uso (e integración) de varias plataformas de despliegue

Con Golang, intenté usar Netlify, pero ni el ejemplo de prueba me funcionaba.

También intenté usar Google Cloud Functions. Esta plataforma parecía funcionar,
de hecho hice dos bots, pero no funcionaban bien:

[Link a las imágenes de los bots, así como al código de Google Cloud Functions](docs/bots.md)

Acabé usando Node con Netlify:
El resultado: una función que se supone que debe devolverte los títulos de
los textos de celebridades, que se pasan mediante una variable get.

Como la función que se usa para obtener las variables de get me estaba dando fallo,
hice que al estar indefinido el campo obtenido por get (fallo de dicha función)
muestre los [poemas más famosos de Rafael Alberti](https://guillelpnz-text-analyzer.netlify.app/.netlify/functions/index?autor=%22Rafael%20Alberti%22).

Mientras que al acceder a la ruta 


<!-- ## Elección del contenedor base

Como contenedor base he elegido golang:alpine3.12. He tomado esta decisión porque
en velocidad no había una diferencia significativa golang:latest,
golang:alpine y golang:1.15.3-alpine. En cuanto al espacio,
todos los golang:alpine pesan cerca de 300mb, mientras que golang:latest pesa más
de 800mb. Estas fueron las [pruebas de velocidad que hice](https://github.com/guillelpnz/TextAnalyzer/blob/master/docs/pruebas_velocidad.md)

## Dockerfile correcto

[Dockerfile del proyecto](https://github.com/guillelpnz/TextAnalyzer/blob/master/Dockerfile)

## Uso de GitHub Container Registry

[Paquetes](https://github.com/guillelpnz?tab=packages)

## Docker Hub

[Container](https://hub.docker.com/r/guillelpnz/textanalyzer/tags)

Para que se construya automáticamente, hay que conectar GitHub a Docker Hub mediante un OAUTH. Posteriormente, acceder al apartado Manage Repository/Builds/Configure Automated Builds y ahí activar la pestaña: autobuild. -->

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

## Avance del código

[Carpeta con los fuentes hasta el momento](https://github.com/guillelpnz/TextAnalyzer/tree/master/src/texto)

## Archivo iv.yaml

[Fichero iv.yaml](iv.yaml)

## Issues cerrados

[Issues cerrados](https://github.com/guillelpnz/TextAnalyzer/issues?q=is%3Aissue+is%3Aclosed)

## Issues abiertos

[Issues abiertos](https://github.com/guillelpnz/TextAnalyzer/issues)
