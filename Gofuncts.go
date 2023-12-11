/*Write a function in Go language that takes two 
integers as parameters and returns their sum*/
package main
import "fmt"
func sum (num1 int, num2 int) int {
	return num1+num2
}
func main() {
	fmt.Println(sum)
}


/*Create a function that accepts a variable number of strings and
concatenates them into a single string.*/
package main
import ("fmt")

func main() {
  var village = "Amechi"
  var town = "Awkunanaw"
  var state = "Enugu"
  var country = "Nigeria"
  
  fmt.Print(village, " ", town, " ", state, " ", country)
}
