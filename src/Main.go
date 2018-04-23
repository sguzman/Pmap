package main

import "fmt"

const (
    cores = 4
)

func pmap(f func(int), length int) {
    endSignal := make(chan bool, cores)
    threadFunc := func(start int) {
        fmt.Printf("Starting thread %d\n", start)
        for i := start; i < length; i += cores {
            f(i)
        }
        fmt.Printf("Done on thread %d\n", start)
        endSignal <- true
    }

    for i := 0; i < cores; i += 1 {
        go threadFunc(i)
    }

    for i := 0; i < cores; i += 1 {
        fmt.Printf("Waiting on thread %d...\n", i)
        <- endSignal
    }
}

func main() {
    pmap(func (i int) { fmt.Println(i) }, 100000)
}
