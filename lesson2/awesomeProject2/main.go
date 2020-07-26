package main

import (
	"encoding/json"
	"fmt"
)

type worker struct {
	WorkingHours    string   `json:"workingHours"`
	MonthlySalary   []salary `json:"monthlySalary"`
	IsStudent       bool     `json:"isStudent"`
	Responsibility  bool     `json:"responsibility"`
	Punctuality     bool     `json:"punctuality"`
	Mindfulness     string   `json:"mindfulness"`
	Communicability bool     `json:"communicability"`
	_human          human
	Languages       []language
}
type salary struct {
	Basic, HRA, TA float64
}

//creating array
var badQualities = []string{"dishonest", "selfish", "cruel", "hypocritical", "greedy", "short-tempered"}

type language struct {
	English bool `json:"english"`
	Japanese bool `json:"japanese"`
	Chinese bool `json:"chinese"`
	French bool `json:"french"`
	Spanish bool `json:"spanish"`
	Turkish bool `json:"turkish"`
	German bool `json:"German"`

}
//to work with JSon,the fields should start with uppercase letter
type artist struct {
	Worker                worker
	MagnetismAndCharisma  bool   `json:"magnetismAndCharisma"`
	PublicSpeakingAbility bool   `json:"publicSpeakingAbility"`
	Specialization        string `json:"specialization"`
}
type chef struct {
	Worker               worker
	BakingAbility         bool   `json:"bakingAbility"`
	OrdinaryCuisineSkills bool   `json:"ordinaryCuisine"`
	SenseOfTaste          bool   `json:"taste"`
	SenseOfSmell          bool   `json:"smell"`
	Specialization        string `json:"specialization"`
}
type doctor struct {
	Worker     worker
	Trustworthy bool `json:"trustworthy"`
	OpenMinded  bool `json:"openMinded"`
}
type teacher struct {
	Worker       worker
	OpenMinded    bool `json:"openMinded"`
	Sincerity     bool `json:"sincerity"`
	Emotionality  bool `json:"emotionality"`
	LikesChildren bool `json:"lovesChildren"`
}
type programmer struct {
	Worker        worker
	HasProbabilitySkills bool `json:"probabilitySkills"`
	Persistence    bool `json:"persistence"`
	HasPatience    bool `json:"patience"`
}
type human struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
	Age           int `json:"age"`
}

var worker2 *worker

func main() {
	fmt.Println("It's second lesson")
	worker1 := worker{WorkingHours: "8:30-19:00", MonthlySalary: []salary{
		{Basic: 1000, HRA: 1200, TA: 1400},
		{Basic: 1200, HRA: 1300, TA: 1500},
		{Basic: 1300, HRA: 1500, TA: 1600},
	}, IsStudent: true, _human: human{Name: "George", Surname: "Waskin", Age: 19}}
	fmt.Println(worker1)
	worker2 := &worker{"7:30-18:00", []salary{
		{1400, 1600, 2000},
		{1600, 1800, 2300},
	}, false, true, true, "70%", true, human{"Vasya", "Nat", 24}, []language{{English: true, German: true, Spanish: true}}}
	fmt.Println("---------------------")
	fmt.Println(worker2)
	fmt.Println("----------------------")
	fmt.Println("Type and Value")
	fmt.Printf("%v,%T", worker1, worker1)
	fmt.Println()
	fmt.Printf("%v,%T\n", worker2, worker2)
	fmt.Println("Section of specific workers:")

	artist1 := artist{
		Worker:                worker{WorkingHours: "6:00-17:00", MonthlySalary: []salary{{1000, 1300, 1800}}, Responsibility: true, Punctuality: true, Mindfulness: "50%", Communicability: true, _human: human{"Mary", "Johnson", 21}, Languages: []language{{English: true, French: true}}},
		MagnetismAndCharisma:  true,
		PublicSpeakingAbility: true,
		Specialization:        "architect",
	}
	fmt.Println("Working with json")
//using marshal function working with json
	bytearray, err := json.MarshalIndent(artist1,"","  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytearray))
//compare length and capacity
	fmt.Println("Playing with slices:")
	s := badQualities[:4]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	s = badQualities[2:5]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	s = badQualities[4:]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	//playing with pointers and dereferencing
	x := 25
	fmt.Println(&x)
	fmt.Println(&*&x, x)
	mutateNum(&x)
	fmt.Println(&x, x)
}

func mutateNum(n *int) {
	*n = 1090
}
