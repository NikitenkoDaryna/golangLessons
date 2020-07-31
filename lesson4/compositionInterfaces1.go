package main

import "fmt"

type greeter interface{
	greet(string) string
}
type american struct{

}
type ukrainian struct{

}
func(a *american) greet(name string) (message string){
	return fmt.Sprintf("Hello %s",name)
}
func(a *ukrainian) greet(name string) (message string){
	return fmt.Sprintf("Привіт %s",name)
}
func sayHello(g greeter,name string){
	fmt.Print(g.greet(name))
}

//composition interfaces

type animal interface{
	walker
	runner
}
type walker interface {
	walk()
}
type runner interface {
	run()
}
type cat struct {

}
type bird interface {
	walker
	flier
}
type flier interface{
	fly()
}
type eagle struct{
	walker
	flier
}
type dog struct {

}
func (cat *cat) walk(){
	fmt.Println("cat is walking")
}
func (cat *cat) run(){
	fmt.Println("cat is running")
}
func (dog *dog) walk(){
	fmt.Println("dog is walking")
}
func (dog *dog) run(){
	fmt.Println("dog is running")
}
func (eagle *eagle) walk(){
	fmt.Println("eagle is walking")
}
func (eagle *eagle) fly(){
	fmt.Println("eagle is flying")
}

func walk(w walker){
	w.walk()
}
