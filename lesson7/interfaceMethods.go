package main

import "fmt"
//Worker
func(w *Worker)getRecord() string{
	return fmt.Sprintf("name:%s, surname:%s, sex:%s, email:%s\njob:%s, salary:%s, position:%s\n yearsOfExperience:%v, backGround:%s",w.name,w.surname,w.sex,w.email,w.job,w.salary,w.position,w.yearsOfExperience,w.backgroundEducation)
}
//Person
func(p *Person)getRecord() string{
	return fmt.Sprintf("name:%s, surname:%s\nage:%v, sex:%s\nid:%s, email:%s,",p.name,p.surname,p.age,p.sex,p.id,p.email)
}
//Doctor
func(d *Doctor)getRecord() string{
	return fmt.Sprintf("job:%s, salary:%s, position:%s\n patiensPerDay:%v",d.job,d.salary,d.position,d.numOfPatientsPerDay)
}

//CustomerRepresentative
func(ct *CustomerRepresentative) getRecord() string{
	return fmt.Sprintf("job:%s, salary:%s, position:%s\n languages:%s",ct.job,ct.salary,ct.position,ct.humanLanguages)
}

//Programmer
func(p *Programmer)getRecord() string{
	return fmt.Sprintf("job:%s, salary:%s, position:%s\n languages:%s",p.job,p.salary,p.position,p.machineLanguages)
}

//Manager

func(m *Manager)getRecord() string{
	return fmt.Sprintf("job:%s, salary:%s, position:%s\nsoftware:%s",m.job,m.salary,m.position,m.software)
}