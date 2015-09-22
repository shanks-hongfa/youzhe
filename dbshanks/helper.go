package dbshanks
import (
	"gopkg.in/mgo.v2"
/*"gopkg.in/mgo.v2/bson"
"log"
"fmt"*/
	"gopkg.in/mgo.v2/bson"
	"time"
)


var session *mgo.Session

var user_c *mgo.Collection


func Destroy() {
	session.Close()
}
func queryBags() {
	result := &[]Bag{}
	user_c.Find(bson.M{}).All(result);
}


func sqlhelp(num int) {
	//var err error
	session_,_:=mgo.Dial("127.0.0.1:27017")
	/*if err != nil {
		panic(err)
	}*/
	session_.SetMode(mgo.Monotonic, true)
	user_c=session_.DB("sell").C("bag")
	//user_c.Update(bson.M{})
	time.Sleep(time.Second*time.Duration(num))
	println(time.Second*time.Duration(num))
	//session.Copy()
	session_.Close()
	end <- 1
}
var end chan int32 = make(chan int32, 10);

func Init() {
	count:=200
	//session, _= mgo.Dial("127.0.0.1:27017")
	for i := 0; i<count; i++ {
		go sqlhelp((i+1))
	}
	var j int32
	for j=0; j<int32(count); j=j+<-end {
		println(j)
	}
	//session.Close()
	//defer session.Close()

	// Optional. Switch the session to a monotonic behavior.

	/*	c := session.DB("sell").C("bag")

		err = c.Insert(&Bag{"Ale", "+55 53 8116 9639"},
			&Person{"Cla", "+55 53 8402 8510"})
		if err != nil {
			log.Fatal(err)
		}

		result := Person{}
		err = c.Find(bson.M{"name": "Ale"}).One(&result)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Phone:", result.Phone)

		var people []Person
		c.Find(bson.M{"name":"Ale"}).All(&people)

		for _,value:= range people{
			fmt.Println("Phone:", value)
		}*/

}