package a

import "log"

var someVariable string

/*SomeGlobalVariable is some variable that has Global access*/
var SomeGlobalVariable string

func init() {
	log.Println("a.Z initialisation starts")
	someVariable = "Excellent"
	SomeGlobalVariable = "Hello,"
	log.Println("a.Z initialisation ends")
}
