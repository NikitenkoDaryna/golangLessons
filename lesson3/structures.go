package main

import (
	"math/rand"
	"strconv"
	"time"
)

type human struct {
	name    string
	surname string
	age     int
}

func (h human) Info() (res string) {
	res = h.name + " " + h.surname + "\nage:" + strconv.Itoa(h.age)
	return
}
func (h human) HappyBirthday() (res string) {
	rand.Seed(time.Now().UnixNano())
	min := h.age
	max := 100-h.age
	res = "Congratulations!You turn " + strconv.Itoa(rand.Intn(max-min+1)+min)
	return
}

type worker struct {
	basicHuman *human
}

func (w worker) New(name string, surname string, age int) {
	w.basicHuman = &human{name, surname, age}
}

type artist struct {
	w                     *worker
	magnetismAndCharisma  bool
	publicSpeakingAbility bool
	specialization        string
}
type chef struct {
	w                     *worker
	bakingAbility         string
	ordinaryCuisineSkills string
	senseOfTaste          string
	senseOfSmell          string
	specialization        string
}
type doctor struct {
	w           *worker
	trustworthy bool
	openMinded  bool
}
type teacher struct {
	w             *worker
	openMinded    bool
	sincerity     bool
	emotionality  bool
	likesChildren bool
}
type programmer struct {
	w                    *worker
	hasProbabilitySkills bool
	persistence          bool
	hasPatience          bool
}
