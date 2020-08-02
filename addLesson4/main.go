package main

import (
	"fmt"
	"reflect"
)

type worker interface {
	getName(name string) string
	getJob(job string) string
	getPosition(pos string) string
}
type teacher struct {
	worker
}
type doctor struct {
	worker
}
type programmer struct {
	worker
}

func (t *teacher) getName(name string) string {
	return "name:" + name
}
func (t *teacher) getJob(job string) string {
	return "job:" + job
}
func (t *teacher) getPosition(pos string) string {
	return "position:" + pos
}

//doctor
func (d *doctor) getName(name string) string {
	return "name:" + name
}
func (d *doctor) getJob(job string) string {
	return "job:" + job
}
func (d *doctor) getPosition(pos string) string {
	return "position:" + pos
}

//programmer
func (p *programmer) getName(name string) string {
	return "name:" + name
}
func (p *programmer) getJob(job string) string {
	return "job:" + job
}
func (p *programmer) getPosition(pos string) string {
	return "position:" + pos
}
func getType(w []worker) interface{}{
	for  i:=0;i<len(w);i++{
		fmt.Println(reflect.TypeOf(w[i]))
	}
	return " "
}
func main() {
	var teacher, doctor, programmer worker = &teacher{}, &doctor{}, &programmer{}

	var workers = map[worker]string{}
	workers[teacher] = teacher.getName("Michael") +"\n"+ teacher.getJob("teacher") +"\n"+ teacher.getPosition("math teacher")
	workers[doctor] = doctor.getName("Daniel") +"\n"+ doctor.getJob("doctor") +"\n"+ doctor.getPosition("plastic surgeon")
	workers[programmer] = programmer.getName("Emily") +"\n"+ programmer.getJob("programmer") + "\n"+programmer.getPosition("junior go developer")

	fmt.Println(workers[teacher]+"\n")
	fmt.Println(workers[doctor]+"\n")
	fmt.Println(workers[programmer]+"\n")

	array:=[]worker{teacher,doctor,programmer}
	fmt.Println(getType(array))


}
