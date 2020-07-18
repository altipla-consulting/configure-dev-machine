
# configure-dev-machine

Configura las máquinas de desarrollo desde un script único.


## Instalar las herramientas

```shell
curl -L https://git.io/JJn6D | bash
```

Reinicia la terminal cuando termine el script para aplicar los cambios.

Si es la primera vez que ejecutas el comando y no tenías Docker instalado reinicia la máquina entera directamente para aplicar todos los cambios.


## Herramientas que instala

- **[actools](https://github.com/altipla-consulting/actools)**: Herramientas internas de Altipla Consulting que usamos para ejecutar contenedores en local.
- **[apt](https://packages.ubuntu.com/)**: Paquetes o comandos necesarios en el sistema para funcionar correctamente.
- **[ci](https://github.com/altipla-consulting/ci)**: Ayuda para subir commits a revisión y otras tareas de Git durante el desarrollo.
- **[docker-compose](https://docs.docker.com/compose/)**: Para ejecutar varios contenedores concurrentemente y coordinados en la máquina local.
- **[go](https://golang.org/)**: En la versión concreta que necesitamos. Se instala en `/usr/local/go` como es típico y `GOPATH` apuntará a `~/go`, que es donde se guardarán los programas que se instalan manualmente o la caché de paquetes. Los proyectos se pueden tener donde queramos con los nuevos Go Modules.
- **[mkcert](https://github.com/FiloSottile/mkcert)**: Genera certificados HTTPS en local para el desarrollo, que solo son válidos en la máquina que los genera.
- **[node](https://nodejs.org/en/)**: Ejecuta ficheros escritos en Javascript.
- **[reloader](https://github.com/altipla-consulting/reloader)**: Ayuda a re-ejecutar programas o tests de Go automáticamente cuando cambien los ficheros.
