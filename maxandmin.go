package main
import (
"fmt"

)
func sort( num[20]int,n int)(int,int){
min :=num[0]
max :=num[0]
//var new[] int
for i:=0;i<n;i++ {
if (num[i]<=min) {
min=num[i]
}else if (num[i]>=max) {
max= num[i]
}
}
return min,max
}



func main() {
var n int
var num[20] int
fmt.Println("enter the number of integers to be sorted")
fmt.Scanf("%d\t", &n)
for i := 0; i<n; i++ {
fmt.Println(" enter the  integer %d\t")
fmt.Scanf("%d\t",&num[i])

}
min,max:=sort(num,n)
fmt.Println("minimum value is ", min)
fmt.Println("maximum value is ", max)
}

