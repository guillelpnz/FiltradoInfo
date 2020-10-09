# Documentación del hito 0

El hito 0 trata sobre la puesta en marcha de las herramientas necesarias para el desarrollo de un proyecto. Se usa Git y Github. Hay que seguir los siguientes pasos:

## Configurar el nombre y correo en GitHub y en Git

Git ya configurado con nombre y correo:

![](imagenes/capt.png)

## Añadir una clave pública a GitHub

Hay que crear un par de claves pública/privada y añadir la privada al agente ssh y la pública a GitHub:

![](imagenes/creacion_publica.png)

Nos aparecerán dos archivos en el directorio .ssh, en mi caso:

![](imagenes/claves_creadas.png)

Añadimos al agente ssh la clave privada:

![](imagenes/agregar_clave_privada_agente.png)

Copiamos la clave pública

![](imagenes/copiar_clave_publica.png)

Y la pegamos en GitHub:

![](imagenes/clave_publica_copiada.png)