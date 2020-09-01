package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
//making CRUD
type User struct {
	Id    bson.ObjectId `bso:"_id"`
	Name  string        `bson:"name"`
	Age   int           `bson:"age"`
	Email string        `bson:"email"`
}

func main() {
	session, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	usersCollection := session.DB("usersCollection").C("users")
	u1 := &User{Id: bson.NewObjectId(), Name: "Daryna", Age: 18, Email: "daryna_nikitenko@yahoo.com"}
	err = usersCollection.Insert(u1)
	if err != nil {
		fmt.Println(err)
	}
	u2 := &User{Id: bson.NewObjectId(), Name: "Vicky", Age: 22, Email: "vicky_li@gmail.com"}
	u3 := &User{Id: bson.NewObjectId(), Name: "Haryy", Age: 23, Email: "Harold001@gmail.com"}
	err = usersCollection.Insert(u2, u3)
	if err != nil {
		fmt.Println(err)
	}
	result := []User{}
	query1 := bson.M{}
	////query2:=bson.M{"age":bson.M{
	//	"$lt":22,
	//}}

_,	err = usersCollection.UpdateAll(bson.M{"age": 23}, bson.M{"$set": bson.M{"age": 20}})
	if err != nil {
		fmt.Println(err)
	}
	err = usersCollection.Find(query1).All(&result)
	if err != nil {
		fmt.Println(err)
	}
	//for _, v := range result {
	//	fmt.Println(v.Name, v.Id, v.Age, v.Email)
	//}

  _,err=usersCollection.RemoveAll(bson.M{"age":18})
	users2 := []User{}
	usersCollection.Find(bson.M{}).All(&users2)

	for _, u := range users2{

		fmt.Println(u.Email,u.Name,u.Age)
	}
}
