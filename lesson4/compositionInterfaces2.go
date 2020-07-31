package main

import "fmt"

type person struct {
	name, surname string
	age           int
}
type worker interface {

	regularWorker
	nonRegularWorker
}
type regularWorker interface {
	attendClasses()
	doHomework()
	work()
	hangOutWithSb(name string)
}
type nonRegularWorker interface {
	work()
	hangOutWithSb(name string)
	cook()
}
type student struct {
	person
}

func (st *student) work() {
fmt.Println("working for 4 hours")
}

func (st *student) hangOutWithSb(name string) {
	fmt.Println("hanging out in the meadow with ",name)
}

func (nst *nonStudent) work() {
fmt.Println("working for 8 hours")
}

func (nst *nonStudent) hangOutWithSb(name string) {
	fmt.Println("hang out in bar with ",name)
}


type nonStudent struct {
	person
}



func (st *student) attendClasses() {
	fmt.Println("feel tired,but heading to class...")

}
func (st *student) doHomework() {
	fmt.Println("doing labs,tasks(relevant/non-relevant)")

}
func (nst *nonStudent) cook() {
	fmt.Println("practicing cooking")
}


