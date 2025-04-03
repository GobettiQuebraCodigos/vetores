package main

import "fmt"

func main(){

	var nomes = []string{"Lucas", "Rafael", "Murilo", "Jubileu", "Michael"}
	rangeOne := nomes[0:2]
	rangeTwo := nomes[3:5]
	rangeThree := nomes[2]
	fmt.Println(nomes)
	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

}