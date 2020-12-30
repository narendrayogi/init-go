package z

import "log"

var someVariable string

/*SomeGlobalVariable is some variable that has Global access*/
var SomeGlobalVariable string

func init() {
	log.Println("z.Z initialisation starts")
	someVariable = "Excellent"
	SomeGlobalVariable = "World!!"
	log.Println("z.Z initialisation ends")
}
