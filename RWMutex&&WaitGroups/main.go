package main

import (
	"fmt"
	"sync"
	"time"
)

type Person interface {
	getName() string
	getSurname() string
	getAge() int
}

type Worker interface {
	//Person
	getSalary() string
	getJob() string
	getPosition() string
}

type Doctor struct {
	//Worker
	//Person
	name, surname                 string
	age                           int
	job, salary, position, status string
}
type CustomerRepresentative struct {
	//Worker
	//Person
	name, surname                 string
	age                           int
	job, salary, position, status string
}
type Programmer struct {
	//Worker
	//Person
	name, surname                 string
	age                           int
	job, salary, position, status string
}
type Manager struct {
	//Worker
	//Person
	name, surname                 string
	age                           int
	job, salary, position, status string
}

func Convert(worker interface{}) {
	err := recover()
	if err != nil {
		fmt.Println(err)
	}
	val, ok := worker.(Person)
	if ok {
		//fmt.Println("ok")
		fmt.Println(val.getName(), val.getSurname(), val.getAge())
	} else
	{
		panic("Not a worker")
	}

}
func main() {

	var worker1 Worker = &Doctor{
		name:     "Paul",
		surname:  "Adams",
		age:      30,
		job:      "surgeon",
		salary:   "52000",
		position: "worker",
	}
	Convert(worker1)

	var worker2 Worker = &Manager{
		name:     "Avery",
		surname:  "Jackson",
		age:      21,
		job:      "sales manager",
		salary:   "15000",
		position: "worker",
	}
	var worker3 Worker = &Manager{
		name:     "George",
		surname:  "Anderson",
		age:      24,
		job:      "recruiter",
		salary:   "30000",
		position: "worker",
	}
	Convert(worker2)
	var worker4 Worker = &Doctor{
		name:     "Bailey",
		surname:  "Atkins",
		age:      30,
		job:      "family doctor",
		salary:   "62000",
		position: "head",
	}
	var worker5 Worker = &Programmer{
		name:     "Dave",
		surname:  "Cameron",
		age:      45,
		job:      "C# developer",
		salary:   "62000",
		position: "head",
	}
	var workerCounter = Start()
	var headCounter = Start()
	workerCounter.initCounter(worker1)
	workerCounter.initCounter(worker2)
	workerCounter.initCounter(worker3)

	headCounter.initCounter(worker4)
	headCounter.initCounter(worker5)
	var wg sync.WaitGroup
	wg.Add(2)
	
	getInfo(headCounter.v, "chiefs", &wg)
	getInfo(workerCounter.v, "average workers", &wg)
}
func getInfo(m map[Worker]int, pos string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	fmt.Printf("Hello,%s\n", pos)
	for w := range m {
		go Convert(w)
	}
	time.Sleep(5 * time.Second)

	fmt.Printf("Bye,%s\n", pos)
}
