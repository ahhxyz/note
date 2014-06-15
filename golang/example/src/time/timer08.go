package main

import "time"

func main() {
        ch := make(chan bool)
        end := make(chan bool)
        go func() {
                defer func() { end <- true }()
                select {
                case <-ch:
                        print("OK\n")
                case <-time.After(time.Second * 2):
                }
        }()
        go func() {
                time.Sleep(time.Second * 3)
                ch <- true
        }()
        <-end
}