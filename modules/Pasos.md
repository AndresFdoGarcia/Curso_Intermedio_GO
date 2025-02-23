## Modulos

- Son paquetes que se pueden reutilizar en diferentes proyectos
- Se almacenan en el go.mod
- Se identifican con el nombre del modulo
- Se almacenan en el directorio modules

## Comandos

- `go mod init <nombre_modulo>` se recomiendo usar el nombre del modulo de github
- `go mod tidy` para limpiar el go.mod
- `go mod download` para descargar los modulos
- `go mod vendor` para generar el directorio vendor
- `go mod edit -replace <nombre_modulo>=<url_modulo>` para reemplazar un modulo por otro
