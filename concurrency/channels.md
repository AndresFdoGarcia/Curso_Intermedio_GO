# 📡 Channels

Los canales son una forma de comunicación entre goRoutins. Existen dos tipos de canales:

- `unbuffered`: No tienen capacidad para almacenar mensajes. Por lo tanto, la goRoutine que envía el mensaje debe esperar a que la goRoutine que recibe el mensaje esté lista para recibirlo. Estará <strong>bloqueada</strong> hasta que se reciba el mensaje.
- `buffered`: Tienen una capacidad para almacenar mensajes. Se puede especificar la capacidad al crear el canal, esto como segundo argumento en la función `make`.


## 📝 Declaración

```go
c := make(chan int) // unbuffered
c := make(chan int, 10) // buffered
``` 

## 📤 Enviar y recibir mensajes

```go
c <- 1 // enviar

fmt.Println(<-c) // recibir 
```

## 🔒 Buffered channels como Semáforos

Los canales buffered se pueden usar como semáforos.

```go
c := make(chan int, 1)
``` 

# ⏳ WaitGroup

Sirve para esperar a que todas las goRoutins finalicen.

-  Se importa el paquete `sync`.

```go
import "sync"
```

- Se declara una variable de tipo `WaitGroup`.

```go
var wg sync.WaitGroup
```

- Se agrega una goRoutin a la espera y se le indica que debe esperar a que se finalice.

```go
wg.Add(1)
```

- Se ejecuta la goRoutin y se indica que debe esperar a que se finalice.

```go
go func() {
	defer wg.Done()

}
```

- Se espera a que se finalice la goRoutin.

```go
wg.Wait()
```

# 📦 Pipelines

Es una forma de comunicación entre goRoutins que se basa en la idea de que una goRoutine produce datos y otra los consume.

Es importante que el canal no se cierre, ya que si lo hace, la goRoutine que lo está consumiendo se quedará esperando a que se le envíen más datos. Además la goRutin que consume debe cerrar el canal de salida y no asignarle un nuevo valor al canal de entrada.

Para evitar el problema se puede definir un canal de salida y un canal de entrada.

- Canal de salida (Escritura): `chan<-` 
- Canal de entrada (Lectura): `<-chan` 


En ***Generator*** se declara el canal `c` como solo escritura `chan<-`.
```go
func Generator(c chan<- int) {
	for i := 1; i <= 10; i++ {
		c <- i
	}
}
```

En ***Double*** se declara el canal `in` como solo lectura `<-chan` y el canal `out` como solo escritura `chan<-`.
```go
func Double(in <-chan int, out chan<- int) {
	for value := range in {
		out <- value * 2
	}
}
```

