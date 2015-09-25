
package main
import( "fmt"
)


func fibo(num int) int{
	if (num ==0){
		return 0
	} else if (num ==1){
		return 1
	}else {

			sum := fibo(num-1)+fibo(num-2)
			return sum




	}

}
func main() {
	var num, i int
	fmt.Println(" get the number for which you want to see fibonacci for")
	fmt.Scanf("%d", &num)
	fmt.Println("the fibonacci series will be")
	for i= 0; i<num; i++ {
		numerical := fibo(i)
		fmt.Printf("%d\n ", numerical)
	}
}

