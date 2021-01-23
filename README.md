# Examen Mercadolibre

Proyecto que permite detectar si un humano es un mutante con base en su secuencia de ADN. Este ejercicio hace parte de la prueba técnica realizada por Mercadolibre

## Comenzando 🚀

_Instrucciones para obtener una copia del proyecto en funcionamiento en una máquina local para propósitos de desarrollo y pruebas._


### Pre-requisitos 📋

_Cosas necesarias para instalar el software y como instalarlas_

```
Docker y Docker Compose en sus últimas versiones
```

### Variables de entorno 🔧

_Los siguientes parámetros deben ser añadidos en un archivo .env creado en la raiz del proyecto._
_Parámetros de la API_

```
API_NAME=[Nombre de la API]
API_BASE_DIR=[Espacio de trabajo]
MUTANTS_CRUD_HTTP_PORT=[Puerto de exposición]
MUTANTS_CRUD_PGDB=[Host de la base de datos]
MUTANTS_CRUD_PGPORT=[Puerto de exposición de la base de datos]
```

_Parámetros de la base de datos_

```
POSTGRES_HOST=[Host de la base de datos]
POSTGRES_DB=[Nombre de la base de datos]
POSTGRES_USER=[Usuario de la base de datos]
POSTGRES_PASSWORD=[Contraseña de la base de datos]
POSTGRES_DB_PORT=[Puerto de exposición]
POSTGRES_SCHEMA=[Esquema de la base de datos]
```

### Instalación 👷

_Construir la imagen de Docker Compose_

```
docker-compose build
```

_Subir el Docker Compose_

```
docker-compose up
```

_El servidor se ejecuta en el puerto seleccionado en la varible de entorno MUTANTS_CRUD_HTTP_PORT_
_http://localhost:8085/_

## Construido con 🛠️

_Herramientas_

* [Golang](https://golang.org/) - Lenguaje de programación concurrente y compilado
* [Beego](https://beego.me/) - Framework de código abierto para construir aplicaciones en el lenguaje de Golang
* [PostgreSQL](https://ejs.co/) - Base de datos relacional orientada a objetos y de código abierto
* [Docker](https://www.docker.com/) - Proyecto de código abierto para la automatización de despliegue de aplicaciones por medio de contenedores
* [Docker compose](https://docs.docker.com/compose/) - Herramienta para definir y ejecutar aplicaciones Docker multicontenedor

## Autor ✒️

* Miguel Angel Vega Alonso - [Miavega](https://github.com/Miavega)
