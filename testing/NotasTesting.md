# Testing 🧪

- Testing es una practica que nos permite verificar que el codigo funcione correctamente ✅
- Se puede hacer manual 👋 o automatizado 🤖
- Se puede hacer en un solo archivo 📄 o en varios 📚
- Se puede hacer en una sola funcion 🔨 o en varias ⚙️


## Comandos 💻

- `go test` para ejecutar los tests ▶️
- `go test -v` para ejecutar los tests con mas detalle 🔍
- `go test -cover` para ejecutar los tests con cobertura 🎯
- `go test -coverprofile=coverage.out` para generar un archivo de cobertura 📊
- `go test -cpuprofile=cpu.out` para generar un archivo de cobertura de CPU 📈


## Coverage 📈

- `go tool cover -html=coverage.out` para generar un archivo html de cobertura 📱
- `go tool cover -func=coverage.out` para generar un archivo de cobertura en formato de texto ✅


## Profiling 📈

- `go test -cpuprofile=cpu.out` para generar un archivo de cobertura de CPU 📈
En caso de querer ver el archivo de cobertura de CPU o memoria, se puede usar el comando:

- `go tool pprof cpu.out` para ver el archivo de cobertura de CPU 📈
- Se usa el comando `top` para ver el top 10 de funciones que consumen mas CPU
- Se usa el comando `list <funcion>` para ver el codigo de la funcion e indicar que linea ocupa mas tiempo
- Se usa el comando `web` para ver el codigo de la funcion en formato html.
- Se usa el comando `pdf` para ver el codigo de la funcion en formato pdf.


### ***Requerimientos***

- Tener instalado `graphviz` para poder usar el comando `web`
- Configurar el navegador para que se pueda abrir el archivo html, en caso de usar Ubuntu se pueden usar los comandos:
```
$ sudo apt install wslu
```
```
$ echo "export BROWSER=wslview" >> ~/.zshrc
```
```
$ source ~/.zshrc
```