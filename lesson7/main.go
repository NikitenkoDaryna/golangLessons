package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type Person struct {
	name, surname string
	age           string
	id            string
	email         string
	sex           string
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

func main() {

	worker1 := Doctor{Worker{
		Person: Person{
			name:    "Paul",
			surname: "Adams",
			age:     "30",
			id:      "id1",
			email:   "paul_adams22@gmail.com",
			sex:     "male",
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
			name:    "Avery",
			surname: "Jackson",
			age:     "21",
			id:      "id2",
			email:   "av_jackson22@gmail.com",
			sex:     "female",
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
				name:    "George",
				surname: "Anderson",
				age:     "24",
				id:      "id3",
				email:   "ander25@gmail.com",
				sex:     "male",
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
				name:    "Bailey",
				surname: "Atkins",
				age:     "30",
				id:      "id4",
				email:   "baily27@gmail.com",
				sex:     "female",
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
				name:    "Dave",
				surname: "Cameron",
				age:     "45",
				id:      "id5",
				email:   "devCam13@yahoo.com",
				sex:     "male",
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

	workerCounter.initCounter(worker1.id, worker1.Worker)
	workerCounter.initCounter(worker2.id, worker2.Worker)
	workerCounter.initCounter(worker3.id, worker3.Worker)

	headCounter.initCounter(worker4.id, worker4.Worker)
	headCounter.initCounter(worker5.id, worker5.Worker)
	//var wg sync.WaitGroup
	//wg.Add(2)

	//getInfo(headCounter.v, "chiefs", &wg)
	//getInfo(workerCounter.v, "average workers", &wg)
	//goGroup.Add(2)
	//getInfoWorkers(workerCounter.v,headCounter.v,&goGroup)
	http.HandleFunc("/page1", mainPage)
	http.HandleFunc("/postForm", postPage)
	http.HandleFunc("/page2",main2Page)
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
func main2Page(w http.ResponseWriter, r *http.Request){
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
	switch r.Method {
	case "Post":
		http.ServeFile(w, r, "postForm.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		w1 := Worker{
			Person: Person{
				name:    r.Form.Get("name"),
				surname: r.Form.Get("surname"),
				age:     r.Form.Get("age"),
				sex:     r.Form.Get("sex"),
				email:   r.Form.Get("email"),
				id:      GenerateRandomString(GenerateNum()),
			},
		}

		_, _ = fmt.Fprintf(w, "Structure is created and uploaded to cache.\nid:%s, name:%s, surname:%s, age:%s, sex:%s,email:%s",
			w1.id, w1.name, w1.surname, w1.age, w1.sex, w1.email)

		workerCounter.initCounter(w1.id, w1)

		goGroup.Add(2)
		getInfoWorkers(workerCounter.v, headCounter.v, &goGroup)
	}

}
func getPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/getForm" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		//http.ServeFile(w, r, "getForm.html")

		id := r.URL.Query().Get("id")
		st1, _ := workerCounter.v[id]

		_, _ = fmt.Fprintf(w, "Structure was found id:%s, name:%s, surname:%s, age:%s, sex:%s,email:%s",
			st1.id, st1.name, st1.surname, st1.age, st1.sex, st1.email)
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
