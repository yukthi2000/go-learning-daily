package main 

import (
	"fmt"
	"reflect"
)


func main(){

	//Integer types
	var num int =34
	fmt.Println("Integer:", num,"Type:",reflect.TypeOf(num))

	//Float point
	var pi float64=3.1421
	fmt.Println("Float:",pi,"Type:",reflect.TypeOf(pi))

	//Boolean
	var isGO bool = true
	fmt.Println("Boolean :",isGO,"Type",reflect.TypeOf(isGO))

	//String
	var message string="hello world"
	fmt.Println("String:", message, "Type:", reflect.TypeOf(message))

	//Byte and Rune(characters)
	var letter byte ='A'
	var unicode rune ='✓'
	fmt.Println("Byte (ASCII):", letter, "Rune (Unicode):", unicode)

	// Complex Numbers
	var complexNum complex128 = complex(5, 7) 
	fmt.Println("Complex Number:", complexNum)

	//Type conversion 
	var a int =4
	var b float64 = float64(a)
	fmt.Println("converted value :",b)

	var c float64 = 3.432
	var d int = int(c)
	fmt.Println("converted value :",d)

	//String to integer 
	var strNum string="100"
	var intNum int 
	fmt.Sscanf(strNum,"%d",&intNum)
	fmt.Println("String to Int:", intNum)


}