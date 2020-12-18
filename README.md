# TextAnalyzer

## Descripción

El proyecto consiste en una API REST programada en [Go](https://golang.org/) que revisa textos de todo tipo, pudiendo encontrar así redundancias en ellos, o hacer estadísticas sobre el uso del lenguaje en diferentes discursos. (Posiblemente acabe resolviendo algunos, ofreciendo sinónimos). El paquete principal del proyecto está [aquí](src/texto).

## Justificación técnica del framework elegido

He elegido el framework para microservicios [Gin](https://gin-gonic.com/), que se usa de una forma muy
parecida a [Martini](https://github.com/go-martini/martini) (basado en Sinatra), pero puede llegar a ser hasta 40 veces más rápido. Otra alternativa era [Beego](https://beego.me/). Como para el uso que le voy a
dar me servía cualquiera de estas, pero Gin tiene una muy buena documentación y es muy rápido, lo elegí.

También podría no haber elegido ninguno, y quedarme con el paquete http que ya viene en el PATH de Go al
instalarlo, sin embargo, me parece menos legible, al menos si ya has trabajado
con microservicios con otros frameworks para otros lenguajes.

## Diseño en general del API

He implementado 5 HU en forma de distintas acciones que hacer con los endpoints:

- introducirTexto (HU8) -> Mediante una petición POST se envían datos como si de un formulario
se tratara, añadiendo al vector de la clase handler (textos) un nuevo objeto texto

- obtenerRedundantes (HU1) -> Se busca uno de los textos del array y se obtienen las palabras repetidas

- obtenerSinRedundantes (HU5) -> Se busca uno de los textos del array y se obtiene limpiando palabras repetidas

- obtenerPersonas (HU3) -> Se busca uno de los textos del array y se obtienen las alusiones a personas

- obtenerEstadisticas (HU7) -> Se busca uno de los textos del array y se obtienen las frecuencias relativas de aparición de las palabras dentro del texto

Cada una de estas HU tiene una ruta diferente dentro de la app. Están en [main.go](https://github.com/guillelpnz/TextAnalyzer/blob/master/src/main.go)

Se procesan llamando a funciones de mi clase manejadora que está en [textos.go](https://github.com/guillelpnz/TextAnalyzer/blob/master/src/micro/textos.go). Estas hacen uso de funciones de mi api
previamente programadas.

Con respecto a los tests, los ejecuto sin correr el servicio, usando el paquete httptest. Tengo el archivo
[main_test.go](https://github.com/guillelpnz/TextAnalyzer/blob/master/src/main_test.go) y [textos_test.go](https://github.com/guillelpnz/TextAnalyzer/blob/master/src/micro/textos_test.go)

## Buenas prácticas

No he hecho la configuración distribuida, pero sí que tengo un log (gin gonic proporciona uno)
que se puede observar en la ejecución de tests, o si se corre la app y se le hacen peticiones.

También he hecho un middleware, que usa las HU que tenía programadas, pero usando las peticiones http propuestas, las rutas y como almacenamiento un vector. Este es el fichero que contiene la clase, así como
las funciones que se llaman cuando se hace routing: [textos.go](https://github.com/guillelpnz/TextAnalyzer/blob/master/src/micro/textos.go)

## Tests correctos y de acuerdo con las historias de usuario

En mi fichero [textos_test.go](https://github.com/guillelpnz/TextAnalyzer/blob/master/src/micro/textos_test.go), hago tests para ver que funciona de manera correcta el array. Los tests de integración se hacen en el fichero [main_test.go](https://github.com/guillelpnz/TextAnalyzer/blob/master/src/main_test.go). Ahí hago peticiones http y compruebo que la salida es la esperada.

<!-- ## Despliegue correcto y funcionando, con documentación de la conexión entre el repo en GitHub y Netlify/Vercel para despliegue continuo

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
de los pushes, se puede ver el [ejercicio 3 del tema](https://github.com/guillelpnz/Ejercicios/blob/master/Serverless/serverless.md) -->

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
