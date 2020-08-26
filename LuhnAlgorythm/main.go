package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	myScanner := bufio.NewScanner(os.Stdin)
	myScanner.Scan()
	number := myScanner.Text()
	fmt.Println("before Luhn algorythm:", number)
	isGenuine,res,_:=Luhn(number)
	fmt.Println("after Luhn algorythm:",res,isGenuine)
	nextNum:=getNextNum(number)
	fmt.Println("next number:",nextNum)
}
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func Luhn(number string) (bool,string,int){
	number=Reverse(number)
	length := len(number)
	numArray := strings.Split(number, "")
	sum:=0
	str:=""
	fmt.Println(numArray)
	for i := 1; i < length; i += 2 {
		num1, err := strconv.Atoi(numArray[i])
		if err != nil {
			fmt.Println(err.Error())

		}
		num1 *= 2
		if num1/10 >= 1 {
			a := num1 / 10
			b := num1 % 10
			num1 = a + b
		}
		numArray[i] = strconv.Itoa(num1)

	}
	str=strings.Join(numArray,"")
	for j:=0;j<len(numArray);j++{
		c,err:=strconv.Atoi(numArray[j])
		if err!=nil{
			fmt.Println(err.Error())
		}
		sum+=c
	}
	str=Reverse(str)
	if sum%10==0{
		return true,str,sum
	}
	return false,str,sum
}
func getNextNum(number string) int{
	numArray := strings.Split(number, "")
	arrOfNums:=append(numArray,"0")
	strNum:=""
	for i:=0;i<len(arrOfNums);i++{
	strNum+=arrOfNums[i]
	}
	//fmt.Println(arrOfNums)
     isGenuine,_,sum:=Luhn(strNum)
     if isGenuine{
     	return 0
	 }
	 return 10-sum%10

}