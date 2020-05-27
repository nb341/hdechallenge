package main

import (
	"fmt"
)

func main()  {
	var t int

	fmt.Scanf("%d", &t)
	outerLoop(t)

}
func outerLoop(i int) int {
	if i>0{
		var j int
		fmt.Scanf("%d",&j)
		fmt.Println(innerLoop(j))
		outerLoop(i-1)

	}
	return 0
}

func innerLoop(j int) int {
	var ans int
	if j>0 {
		var a int
		fmt.Scanf("%d",&a)
		ans = sumS(a)+innerLoop(j-1)
	} else
	{
		ans=0
	}

    return ans


}
func sumS(a int) int  {

	 if a>0{
	 	return a*a
	 }
	 return 0
}
