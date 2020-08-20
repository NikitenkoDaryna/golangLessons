package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"
)

type Person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     string `json:"age"`
	Id      string`json:"id"`
	Email   string `json:"email"`
	Sex     string `json:"sex"`
}
type IWorker interface {
	getRecord() string
}
type Worker struct {
	Person
	job, salary, position string
	yearsOfExperience     int
	backgroundEducation   string
}

type Doctor struct {
	Worker
	numOfPatientsPerDay int
}

type CustomerRepresentative struct {
	Worker
	humanLanguages []string
}
type Programmer struct {
	Worker
	machineLanguages []string
}
type Manager struct {
	Worker
	software []string
}

func Convert(worker IWorker) *Person {
	err := recover()
	if err != nil {
		fmt.Println(err)
	}
	val, ok := worker.(*Person)
	if ok {
		fmt.Println("ok")

	} else
	{
		panic("Not a person")
	}
	return val
}

var workerCounter = Start()
var headCounter = Start()
var goGroup sync.WaitGroup

const postAddr = "http://localhost:3000/postForm"

func main() {

	worker1 := Doctor{Worker{
		Person: Person{
			Name:    "Paul",
			Surname: "Adams",
			Age:     "30",
			Id:      "id1",
			Email:   "paul_adams22@gmail.com",
			Sex:     "male",
		},
		job:                 "surgeon",
		salary:              "30000",
		position:            "head",
		yearsOfExperience:   10,
		backgroundEducation: "univ1",
	},
		10,
	}

	worker2 := Manager{Worker{
		Person: Person{
			Name:    "Avery",
			Surname: "Jackson",
			Age:     "21",
			Id:      "id2",
			Email:   "av_jackson22@gmail.com",
			Sex:     "female",
		},
		job:                 "sales manager",
		salary:              "15000",
		position:            "worker",
		yearsOfExperience:   4,
		backgroundEducation: "univ2",
	},
		[]string{
			"Excel", "World", "Google Forms",
		},
	}
	worker3 := Manager{
		Worker: Worker{
			Person: Person{
				Name:    "George",
				Surname: "Anderson",
				Age:     "24",
				Id:      "id3",
				Email:   "ander25@gmail.com",
				Sex:     "male",
			},
			job:                 "recruiter",
			salary:              "30000",
			position:            "worker",
			yearsOfExperience:   6,
			backgroundEducation: "univ3",
		},

		software: []string{
			"Excel", "World", "Google Forms",
			"PROOFHUB", "Jira",
		},
	}
	Convert(&worker2.Worker.Person)

	worker4 := Doctor{
		Worker{
			Person: Person{
				Name:    "Bailey",
				Surname: "Atkins",
				Age:     "30",
				Id:      "id4",
				Email:   "baily27@gmail.com",
				Sex:     "female",
			},

			job:                 "family doctor",
			salary:              "62000",
			position:            "head",
			yearsOfExperience:   10,
			backgroundEducation: "univ4",
		},
		15,
	}

	worker5 := Programmer{
		Worker{
			Person: Person{
				Name:    "Dave",
				Surname: "Cameron",
				Age:     "45",
				Id:      "id5",
				Email:   "devCam13@yahoo.com",
				Sex:     "male",
			},
			job:                 "C# developer",
			salary:              "62000",
			position:            "head",
			yearsOfExperience:   13,
			backgroundEducation: "univ5",
		},
		[]string{
			"C#", "C++", "Java", "Go",
		},
	}

	workerCounter.initCounter(worker1.Id, worker1.Worker)
	workerCounter.initCounter(worker2.Id, worker2.Worker)
	workerCounter.initCounter(worker3.Id, worker3.Worker)

	headCounter.initCounter(worker4.Id, worker4.Worker)
	headCounter.initCounter(worker5.Id, worker5.Worker)
	//var wg sync.WaitGroup
	//wg.Add(2)

	//getInfo(headCounter.v, "chiefs", &wg)
	//getInfo(workerCounter.v, "average workers", &wg)
	//goGroup.Add(2)
	//getInfoWorkers(workerCounter.v,headCounter.v,&goGroup)
	http.HandleFunc("/page1", mainPage)
	http.HandleFunc("/postForm", postPage)
	http.HandleFunc("/page2", main2Page)
	http.HandleFunc("/getForm", getPage)
	fmt.Printf("Starting server for testing HTTP POST...\n")
	http.ListenAndServe(":3000", nil)

}
func getInfo(m map[string]Worker, pos string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	fmt.Printf("Hello,%s\n", pos)
	for _, v := range m {
		go fmt.Println(*Convert(&v.Person))

	}
	time.Sleep(5 * time.Second)

	fmt.Printf("Bye,%s\n", pos)
}
func mainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/page1" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, "postForm.html")
}
func main2Page(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/page2" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, "getForm.html")
}

func postPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/postForm" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	f, err := os.OpenFile("deck.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer f.Close()


	switch r.Method {
	case "POST":
		workerForm := new(Worker)
		workerForm.Id = GenerateRandomString(GenerateNum())
		workerForm.Name = r.FormValue("name")
		workerForm.Surname = r.FormValue("surname")
		workerForm.Sex = r.FormValue("sex")
		workerForm.Age = r.FormValue("age")
		workerForm.Email = r.FormValue("email")

		worker, err := json.Marshal(workerForm)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		f.Write(worker)
		f.Close()
		w.Write(worker)
		workerCounter.initCounter(workerForm.Id, *workerForm)
		//testing goGroups.Only for output in the console
		//goGroup.Add(1)
		//getInfoWorkers(workerCounter.v, headCounter.v, &goGroup)
	}
}

func getPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/getForm" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		id := r.URL.Query().Get("id")
		st1, _ := workerCounter.v[id]

		_, _ = fmt.Fprintf(w, "Structure was found id:%s, name:%s, surname:%s, age:%s, sex:%s,email:%s",
			st1.Id, st1.Name, st1.Surname, st1.Age, st1.Sex, st1.Email)
		log.Println("structure was found:", st1)
		w.Write([]byte("Received a GET request\n"))
	}

}

func getInfoWorkers(averageWorkers map[string]Worker, heads map[string]Worker, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	for _, v := range averageWorkers {
		go fmt.Println(*Convert(&v.Person))

	}
	time.Sleep(5 * time.Second)
	for _, v := range heads {
		go fmt.Println(*Convert(&v.Person))

	}

}
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(n int) string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	bytes, _ := GenerateRandomBytes(n)

	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes)
}
func GenerateNum() int {
	rand.Seed(time.Now().UnixNano())
	randomNum := random(1, 30)
	return randomNum
}
func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

