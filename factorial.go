package main
import "fmt"
func factorial(num int) int{
	if (num ==0){
		return 1
	} else if (num ==1){
		return 1
	}else {
		num = num * factorial (num-1)
		return num
	}

}
func main(){
	var num int
	fmt.Println(" get the number for which you want to find factorial for")
	fmt.Scanf("%d",&num)
	numerical:= factorial(num)
	fmt.Println(" the factorial value is ",numerical)
}
