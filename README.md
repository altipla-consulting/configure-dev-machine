
# configure-dev-machine

Configura las máquinas de desarrollo desde un script único.


## Instalar las herramientas

```shell
curl -L https://git.io/JJn6D | bash
```

Reinicia la terminal cuando termine el script para aplicar los cambios.

Si es la primera vez que ejecutas el comando y no tenías Docker instalado reinicia la máquina entera directamente para aplicar todos los cambios.


## Herramientas que instala

- **[az](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest)**: La línea de comandos de Azure para administrar los servicios de los clientes que lo usan.
- **[actools](https://github.com/altipla-consulting/actools)**: Herramientas internas de Altipla Consulting que usamos para ejecutar contenedores en local.
- **[apt](https://packages.ubuntu.com/)**: Paquetes o comandos necesarios en el sistema para funcionar correctamente.
- **[ci](https://github.com/altipla-consulting/ci)**: Ayuda para subir commits a revisión y otras tareas de Git durante el desarrollo.
- **[docker-compose](https://docs.docker.com/compose/)**: Para ejecutar varios contenedores concurrentemente y coordinados en la máquina local.
- **[enospc](https://stackoverflow.com/questions/22475849/node-js-what-is-enospc-error-and-how-to-solve)**: Configura el kernel para evitar el error típico al observar cambios en los ficheros locales.
- **[gcloud](https://cloud.google.com/sdk)**: Incluyendo programas como `kubectl` y su alias a `k`.
- **[go](https://golang.org/)**: En la versión concreta que necesitamos. Se instala en `/usr/local/go` como es típico y `GOPATH` apuntará a `~/go`, que es donde se guardarán los programas que se instalan manualmente o la caché de paquetes. Los proyectos se pueden tener donde queramos con los nuevos Go Modules.
- **[ipv4-forwarding](https://stackoverflow.com/questions/41453263/docker-networking-disabled-warning-ipv4-forwarding-is-disabled-networking-wil)**: Soluciona un fallo de Docker que evita que los contenedores puedan conectarse a la red con normalidad.
- **[jnet](https://github.com/altipla-consulting/jnet)**: Compilador de JSONNET a JSON para las configuraciones de Kubernetes.
- **[mkcert](https://github.com/FiloSottile/mkcert)**: Genera certificados HTTPS en local para el desarrollo, que solo son válidos en la máquina que los genera.
- **[node](https://nodejs.org/en/)**: Ejecuta ficheros escritos en Javascript.
- **[npmpackages](https://www.npmjs.com/package/lerna)**: Instala el paquete de lerna que usamos para administrar algunos proyectos con varios frontends. También actualiza npm a la última versión.
- **[reloader](https://github.com/altipla-consulting/reloader)**: Ayuda a re-ejecutar programas o tests de Go automáticamente cuando cambien los ficheros.
- **[stern](https://github.com/wercker/stern)**: Emite los logs de una app de Kubernetes y se queda escuchando a las nuevas líneas que van surgiendo.
