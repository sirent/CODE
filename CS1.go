package main

import "fmt"

func main() {
	var n int = 6
	var i, j int
	var angka [10]int
	var a1 int
	var temp1, temp2 int

	fmt.Scan(&a1) // number must not be same as input in a1
	for i := 0; i < n; i++ {
		fmt.Scan(&angka[i])
	}

	// TRY ALONE
	// for i := 0; i < n; i++ {
	// 	if (angka[i+1] > a1 && angka[i] < a1) || angka[i] > a1 && angka[i+1] < a1{
	// 		for x := n; x >= i; x-- {
	// 			angka[x+1] = angka[x]
	// 		}
	// 		angka[i+1] = a1
	// 	}
	// }
	// fmt.Println(angka)

	// FROM BOOK
	i = 0
	for i < n && angka[i] < a1 {
		i++
	}
	temp1 = angka[i]
	angka[i] = a1
	for j = i + 1; j <= n; j++ {
		temp2 = angka[j]
		angka[j] = temp1
		temp1 = temp2
	}
	fmt.Println(angka)
}
