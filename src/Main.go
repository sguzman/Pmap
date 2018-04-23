package main

import (
    "fmt"
    "time"
)

const (
    cores = 8
)

func pmap(f func(int), length int) {
    endSignal := make(chan bool, cores)
    threadFunc := func(start int) {
        for i := start; i < length; i += cores {
            f(i)
        }
        endSignal <- true
    }

    for i := 0; i < cores; i += 1 {
        go threadFunc(i)
    }

    for i := 0; i < cores; i += 1 {
        <- endSignal
    }
}

func main() {
    pmap(func (i int) {
        time.Sleep(100000)
        fmt.Println(i)
    }, 100000)
}
