package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	const (
		maxWorkers   = 4
		maxQueueSize = 20
		port         = 8081
	)

	jobQueue := make(chan Job, maxQueueSize)
	dispatcher := NewDispatcher(jobQueue, maxWorkers)
	dispatcher.Run()

	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, r, jobQueue)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

// Job representa una tarea para calcular un número de Fibonacci
// Name: identificador de la tarea
// Delay: tiempo de espera artificial
// Number: número para calcular su Fibonacci
type Job struct {
	Name   string
	Delay  time.Duration
	Number int
}

// Worker representa un trabajador que procesa Jobs
// Id: identificador único del worker
// JobQue: canal para recibir tareas específicas del worker
// WorkerPool: canal compartido entre todos los workers para distribuir trabajo
// Quit: canal para señalizar la terminación del worker
type Worker struct {
	Id         int
	JobQue     chan Job
	WorkerPool chan chan Job
	Quit       chan bool
}

// NewWorker crea y retorna un nuevo worker
// id: identificador único para el worker
// workerPool: canal compartido para la distribución de trabajo
func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{
		Id:         id,
		JobQue:     make(chan Job),  // Canal de tareas del worker
		WorkerPool: workerPool,      //Canal de canales de tareas, este canal se comparte entre todos los workers
		Quit:       make(chan bool), // Canal para parar el worker
	}
}

// Start inicia el worker en una goroutine separada
// El worker se registra en el WorkerPool y espera por nuevos trabajos
func (w Worker) Start() {
	//Se inicia de manera concurrente un ciclo sin fin
	go func() {
		for {
			//Al worker pool se manda el canal de worker, este se manda cada vez iteracion, es decir cuando el worker termino de hacer un job
			w.WorkerPool <- w.JobQue
			select {
			case job := <-w.JobQue:
				//Si se recibe un job en el canal de salida se para el worker (lo sca del ciclo)
				fmt.Printf("Worker with id %d started\n", w.Id)
				fib := Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("Worker with id %d finished whit result %d\n", w.Id, fib)

			case <-w.Quit:
				//Si se recibe un job en el canal de salida se para el worker (lo sca del ciclo)
				fmt.Printf("Worker with id %d Stopped\n", w.Id)
				return
			}
		}
	}()
}

// Stop envía una señal para detener el worker
func (w Worker) Stop() {
	go func() {
		w.Quit <- true
	}()
}

// Fibonacci calcula el número de Fibonacci de forma recursiva
// n: número para calcular su valor en la secuencia de Fibonacci
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// Dispatcher coordina la distribución de trabajos entre los workers
// WorkerPool: pool de canales de workers disponibles
// MaxWorkers: número máximo de workers permitidos
// JobQueue: cola global de trabajos pendientes
type Dispatcher struct {
	WorkerPool chan chan Job //Canal de canales de tareas, este se les pasa a cada worker nuevo
	MaxWorkers int           //Cantidad máxima de workers
	JobQueue   chan Job      //Canal de tareas, se puede ver como un canal global de tareas que despues se reparten entre workers
}

// NewDispatcher crea y retorna un nuevo dispatcher
// jobQueue: canal para recibir nuevos trabajos
// maxWorkers: número máximo de workers a crear
func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	worker := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		JobQueue:   jobQueue,
		MaxWorkers: maxWorkers,
		WorkerPool: worker,
	}
}

// Dispatch inicia el ciclo principal de distribución de trabajos
// Toma trabajos de JobQueue y los asigna a workers disponibles
func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			go func() {
				workerJobQueue := <-d.WorkerPool
				workerJobQueue <- job
			}()
		}
	}
}

// Run inicializa el dispatcher creando los workers
// y comenzando el proceso de distribución
func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start()
	}
	go d.Dispatch()
}

// RequestHandler maneja las peticiones HTTP entrantes para crear nuevos trabajos
// w: respuesta HTTP
// r: petición HTTP
// jobQueue: cola global de trabajos donde se agregarán las nuevas tareas
func RequestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Invalid delay", http.StatusBadRequest)
		return
	}

	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Invalid value", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Invalid name", http.StatusBadRequest)
		return
	}

	job := Job{Name: name, Delay: delay, Number: value}
	jobQueue <- job
	w.WriteHeader(http.StatusCreated)
}
