package main
import "fmt"

func main(){
	var num1, num2 int
	Str1:= " Enter the first number "
	fmt.Println(Str1)
	fmt.Scanf("%d\t",&num1)
	Str2:= " Enter the second number "
	fmt.Printf(Str2)
	fmt.Scanf("%d\t",&num2)
	fmt.Println(" the numbers entered are",num1,num2)
	if (num1>num2){
		fmt.Println(" the number 1 is maximum")

	}else if (num1<num2){
		fmt.Println(" the number 2 is maximum")
	}else{
		fmt.Println(" the numbers are maximum")
	}

}

