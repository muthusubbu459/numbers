package main
import "fmt"

func minimum(x[] int)int{
	value:=x[0]
	for i:=0;i<len(x) ; i++ {
		if(x[i]<=value){
			value=x[i]
		}
	}
	return value
}
func main(){

	x :=[]int{48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17,}
	min:=minimum(x)
	fmt.Println("the min value is ",min)
}
