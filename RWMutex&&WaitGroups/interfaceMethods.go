package main

import "fmt"

//Doctor
func (d *Doctor) getSalary() string {
	return fmt.Sprintf("doctor's salary-%s:", d.salary)
}
func (d *Doctor) getJob() string {
	return fmt.Sprintf("doctor's job-%s:", d.job)
}
func (d *Doctor) getPosition() string {
	return fmt.Sprintf("doctor's position-%s:", d.position)
}
func(d *Doctor) getName() string{
	return fmt.Sprintf("doctor's name:%s",d.name)
}
func(d *Doctor) getSurname() string{
	return fmt.Sprintf("doctor's surname:%s",d.surname)
}
func(d *Doctor) getAge() int{
	return d.age
}
//CustomerRepresentative
func (c *CustomerRepresentative) getSalary() string {
	return fmt.Sprintf("CustomerRepresentative's salary-%s:", c.salary)
}
func (c *CustomerRepresentative) getJob() string {
	return fmt.Sprintf("CustomerRepresentative's job-%s:", c.job)
}
func (c *CustomerRepresentative) getPosition() string {
	return fmt.Sprintf("CustomerRepresentative's position-%s:", c.position)
}
func(c *CustomerRepresentative) getName() string{
	return fmt.Sprintf("CustomerRepresentative's name:%s",c.name)
}
func(c *CustomerRepresentative) getSurname() string{
	return fmt.Sprintf("CustomerRepresentative's surname:%s",c.surname)
}
func(c *CustomerRepresentative) getAge() int{
	return c.age
}
//Programmer
func (p *Programmer) getSalary() string {
	return fmt.Sprintf("programmer's salary-%s:", p.salary)
}
func (p *Programmer) getJob() string {
	return fmt.Sprintf("programmer's job-%s:", p.job)
}
func (p *Programmer) getPosition() string {
	return fmt.Sprintf("programmer's position-%s:", p.position)
}
func(p *Programmer) getName() string{
	return fmt.Sprintf("programmer's name:%s",p.name)
}
func(p *Programmer) getSurname() string{
	return fmt.Sprintf("programmer's surname:%s",p.surname)
}
func(p *Programmer) getAge() int{
	return p.age
}
//Manager
func (m *Manager) getSalary() string {
	return fmt.Sprintf("manager's salary-%s:", m.salary)
}
func (m *Manager) getJob() string {
	return fmt.Sprintf("manager's job-%s:", m.job)
}
func (m *Manager) getPosition() string {
	return fmt.Sprintf("manager's position-%s:", m.position)
}
func(m *Manager) getName() string{
	return fmt.Sprintf("manager's name:%s",m.name)
}
func(m *Manager) getSurname() string{
	return fmt.Sprintf("manager's surname:%s",m.surname)
}
func(m *Manager) getAge() int{
	return m.age
}
