
# configure-dev-machine

Configura las máquinas de desarrollo desde un script único.


## Instalar las herramientas

```shell
curl -L https://git.io/Jvihi | bash
```

Reinicia la terminal cuando termine el script para aplicar los cambios.

Si es la primera vez que ejecutas el comando y no tenías Docker instalado reinicia la máquina entera directamente para aplicar todos los cambios.


## Herramientas que instala

- [Go](https://golang.org/): En la versión concreta que necesitamos. Se instala en `/usr/local/go` como es típico y `GOPATH` apuntará a `~/go`, que es donde se guardarán los programas que se instalan manualmente o la caché de paquetes.
- [Docker Compose](https://docs.docker.com/compose/)
- [Actools](https://github.com/altipla-consulting/actools): Herramientas internas de Altipla Consulting que usamos para ejecutar contenedores en local.
- [Reloader](https://github.com/altipla-consulting/reloader): La última versión disponible. Ayuda a re-ejecutar programas o tests de Go automáticamente cuando cambien los ficheros.
