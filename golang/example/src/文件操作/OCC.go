package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// open files r and w
	r, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	w, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer w.Close()

	// do the actual work
	n, err := io.Copy(w, r)
	if err != nil {
		panic(err)
	}
	log.Printf("Copied %v bytes\n", n)
}
