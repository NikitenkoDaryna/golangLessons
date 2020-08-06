package main

import (
	"fmt"
)

type ErrNegativeSqrt float64
type SquareNegativeSide float64
type Person struct{
	age int
}
type AgeError struct{
	age int
	what string
}
func(er *AgeError) Error() (s string){
	s =fmt.Sprintf("age: %v\n message:%v",er.age,er.what)
	return s
}
func(per *Person) getAge() (int,error){
	var myErr AgeError
	var e error
	myErr.age=per.age
	myErr.what="age can't be <0 or >100"
	if per.age<0 || per.age>100{
		e=&myErr
		myErr.age=per.age
		myErr.what="age can't be <0 or >100"
		return per.age,e
	}
	return per.age,nil
}
func Sqrt(x float64) (float64, error) {
	var n ErrNegativeSqrt
	var e error
	if x < 0 {
		e = n

		return x, e
	}

	return x, nil
}
func Area(x float64) (float64, error) {
	var y SquareNegativeSide
	var e error
	if x < 0 {

		e = y
		return x, e
	}
	return x * x, nil
}
func (e SquareNegativeSide) Error() (s string) {
	s = fmt.Sprintf("error,side can't be negative!!you input:%v", float64(e*e))
	return
}
func (e ErrNegativeSqrt) Error() (s string) {
	s = fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
	return
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(3))
p:=Person{
	25,
}

	fmt.Println(Area(-9))
	fmt.Println(p.getAge())
}
