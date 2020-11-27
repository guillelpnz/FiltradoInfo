# TextAnalyzer

## Descripción

El proyecto consiste en una API REST programada en [Go](https://golang.org/) que revisa textos de todo tipo, pudiendo encontrar así redundancias en ellos, o hacer estadísticas sobre el uso del lenguaje en diferentes discursos. (Posiblemente acabe resolviendo algunos, ofreciendo sinónimos). El paquete principal del proyecto está [aquí](src/texto).

## Despliegue correcto y funcionando, con documentación de la conexión entre el repo en GitHub y Netlify/Vercel para despliegue continuo

La conexión entre Vercel y mi GitHub está en [el ejercicio 1 del tema](https://github.com/guillelpnz/Ejercicios/blob/master/Serverless/serverless.md)

Está hecha con un repo de prueba, pero con mi repositorio se conecta de la misma forma.

Una vez que se ha conectado GitHub con mi repositorio, nos debería de aparecer
en la pestaña Overview de Vercel. El endpoint que se nos asignará, contendrá
un token aleatorio, haciendo la url menos legible, por lo que decidí cambiar
el dominio en Settings > Domain.

La función que desplegué en esta plataforma es la correspondiente a mi HU5.
La función trata de eliminar palabras repetidas de un texto. Está implementada
en [este fichero](https://github.com/guillelpnz/TextAnalyzer/blob/master/src/texto/texto.go).
Se llama ObtenerSinRedundantes(). Básicamente, sobre un objeto tipo Texto (mi clase), modifica
su atributo contenido, eliminando palabras repetidas.

Esta es la única operación que hice de configuración de Vercel, además de
conectarlo con mi repo de GitHub. Los builds automáticos se configuran
automáticamente.

Para poder desplegarla tuve que hacer cambios:

Como vercel coge las funciones del directorio /api, tuve que crear ese directorio y colocar ahí
un fichero .go con mi función. [Éste es el fichero obtener-sin-redundantes.go](https://github.com/guillelpnz/TextAnalyzer/blob/master/api/obtener-sin-redundantes.go).

Lo que cambia es que ahora el texto a modificar nos lo pasan por query string,
por lo que hacemos un saneamiento de la cadena que nos pasan antes de llamar
al método de mi clase, y tras inicializar un objeto con el contenido de lo que
nos han pasado por GET, ya llamamos a ObtenerSinRedundantes.

Otro cambio es que en mi código, esta función devuelve un slice de string.
En cambio, al tener que devolver un string tuve que hacer una pequeña modificación,
recorriendo este slice y concatenando sus palabras en un string, separando por espacios.

## Integración dentro del proyecto general (es decir, como todo el código deberá tener sus issues y/o HU correspondientes)

En [mis issues](https://github.com/guillelpnz/TextAnalyzer/issues)
se puede ver el avance del código, donde voy referenciándolos.

## Uso (e integración) de varias plataformas de despliegue

Con Golang, intenté usar Netlify, pero ni el ejemplo de prueba me funcionaba.

También intenté usar Google Cloud Functions. Esta plataforma parecía funcionar,
de hecho hice dos bots, pero no funcionaban bien:

[Link a las imágenes de los bots, así como al código de Google Cloud Functions](docs/bots.md)

Acabé programando un bot Node que usa un endpoint de Netlify.
El bot devuelve los títulos de los textos de celebridades. (HU4)
En mi caso he recopilado algunos de Rafael Alberti y Federico García Lorca.

Para más información sobre el funcionamiento del bot, así como la automatización
de los pushes, se puede ver el [ejercicio 3 del tema](https://github.com/guillelpnz/Ejercicios/blob/master/Serverless/serverless.md)

## Originalidad

He cambiado de lenguaje a Node para la anterior rúbrica, y he hecho un bot.

He probado 3 sistemas de FaaS.

Hice 2 bots en Google Cloud Functions, que aunque fueron fallidos, me sirvieron
para entender más en profundidad el tema y dicha plataforma

Programé una función correspondiente a una nueva HU, así como su test.

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
