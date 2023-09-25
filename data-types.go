/* var identifier type

var a int
var b bool
var f float
var str string
*/

package main

import "fmt"

func main() {

	var name string = "Krishna"
	fmt.Printf("My name is: %s\n", name)

	var mobile int = 4321
	fmt.Printf("Contact: %d\n", mobile)

	var married bool = false
	fmt.Printf("Marriage status: %v \n", married)

	var mydata float32 = 3.142
	fmt.Printf("The floating point is %g\n", mydata)

	// Print binary
	var mybinary int = 7
	fmt.Printf("The binary value of mydata is %b\n", mybinary)

	fmt.Printf("%T %T %T %T %T\n", name, mobile, married, mydata, mybinary)
	// Prints type of the variable

	fmt.Printf("%v %v %v %v %v\n", name, mobile, married, mydata, mybinary)
	// Prints value of the variable

}
