package main

var a, b int
var c = make(chan bool, 1)

func f() {
	a = 1
	b = 2
	c <- true
}

func g() {
	<-c
	print(b)
	print(a)
}

func main() {
	go f()
	g()

}
