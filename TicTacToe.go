package main
import "fmt"

func main(){
	var i,j,x,y int
	var a string
	var bingo bool

	bingo = false
	for !bingo{
		fmt.Scan(&x,&y)
		fmt.Scan(&a)
		
		for i = 1; i<=3; i++{
			for j=1;j<=3;j++{
				if i == x && j == y{
					fmt.Print(a)
				}else{
					fmt.Print(" ")
				}
				fmt.Print(" ")
			}
			fmt.Println("")
		}
	}
}

func checkBingo(i,j,x,y,k,d){
	if ((i == 1 && j == 1) && k == "X") {

	}
}