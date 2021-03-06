# Chapter9: Leveraging Go concurrency
One of the important feature of Go is concurrency. This chapter shows what concurrent execution is, and how modeled/designed in Go. We also walk through the two main utilities, goroutine and channel.

## 9.1: Cocurrency and Parallelism
> Concurrent execution is handling many tasks in the same time. Parallel execution is doing many tasks in the same time. - Rob Pike (Co-developer of Go)

Concurrent execution handles multiple tasks, enables them to interact with each other. These tasks are treated as done concurrently (not ordered).

Parallel execution is similar to concurrent execution but different. In parallel execution, tasks begin and be executed at the same time. Parallel execution often used in order to reduce the total operation time. Parallel execution often assumes independent computational resources (like CPU), concurrent execution shares the same resources.

Go can implement (seemingly) parallel execution, but concurrency is actually implemented. Concurrent programming in Go is implemented by mainly two utilities, **goroutine** and **channels**.

## 9.2: Goroutine

Goroutine is a function that multiple operations concurrently. This seems to be similar to *thread* (indeed goroutine uses thread), but not. Goroutine can do more operations than thread because goroutine is light. Goroutine initially starts a operation with a light stack then gradually increases/reduces stack size if nesessary.

Former versions of Go (~v1.4) uses only single CPU by default. Since v1.5, Go changes to use CPUs as many as possible.

Goroutine requires launching time through it is light. In the test in `goroutine_benchmark_1` shows, the cost using goroutine is larger than that of actual operation. In `goroutine_benchmark_2`, `GoPrint2` is much faster than `print2`. This is because those operations require a few moment and goroutine gets meaningfull. Scheduling and execution on multiple CPUs require cost. Execution time depends on the balance of the cost and the effect of concurrency.

WaitGroup is a utility to make all goroutines finished before the following task. WaitGroup is included in `sync` package.

## 9.3: Channel

Channel can be seen as a kind of box. Goroutines can interact only with this box. Goroutine put what wants to pass in this box so that other goroutines can get it. Channel is a typed value, enables goroutines to interact with. Channel is assigned by `make`, resulting value is a pointer to the actual data.

- `ch := make(chan int)`: int channel
- `ch := make(chan int, 10)`: int channel with size 10 buffer
- `ch <- 1`: put int 1 into ch(annel)
- `i := <- ch`: get value from ch and assign it to i
- `ch := make(chan <- string)`: send-only string channel
- `ch := make(<-chan string)`: receive-only string channel

`channel_interact/channel.go` shows that the interacted numbers via channel are ordered, the number does not increment unless the value are pull from the channel (varing printed Caught/Thres order is not probramatic).

Channel with buffer allows processes to continue until the channel is empty. This is useful when you want to limit the amount of process per time.

## 9.4: Web application and concurrent execution

Concurrent execution introduced above can be used in web applications.
In this chapter, we are going to create a mosaic photo generation application.
This application simply make user to upload an image, then generate mosaic photo based on it.
We assume that tile image is prepared and its size is modified.

See [https://github.com/mushahiroyuki/gowebprog/tree/master/ch09](https://github.com/mushahiroyuki/gowebprog/tree/master/ch09) for the static files.

- Build: `go build` under `ch9/mosaic` directory
- Execute: `GOMAXPROCS=1 ./mosaic` under `ch9/mosaic` with built file

## Summary
- Go web server itself has concurrency. Each received requests are processed by each goroutine
- Concurrency and parallelism are complement but different. Concurrency inplies multiple processes to be begun in the same time, executed, end, and enables them to interact with. Parallelism just say multiple execution in the same time.
- Go provides namely two functions - goroutine and channel - to implement concurrency. Go does not support parallel operation directly.
- Goroutine is used to create concurrent execution. Channel enables goroutines to interact with.
- Non-buffered channel is synchronous, stop to write messages while containing some data. Buffered channel allows asynchronous writing while the buffer has space.
- `select` can be used to select channel (firstly prepared one) from multiple ones
- `WaitGroup` also can be used to synchronize channels
- Concurrent web application runs much faster than sequential one, even on single CPU
- Concurrent-type web application possibly benefits from concurrent execution automatically
