package main
import "fmt"

func sample(){  
var i int
for i = 1; i <= 5; i++ {
	fmt.Println(i)
		}
		fmt.Println("done wiht for loop function ")
}


func display (a int) {

	fmt.Println(a)
}


func main () {

	defer sample()
	defer display(1)
	defer display(2)
	defer display(3)
	defer display(4)

	 display(1)
	 display(2)
	 display(3)
	 display(4)
	fmt.Println("In the main looop")
}