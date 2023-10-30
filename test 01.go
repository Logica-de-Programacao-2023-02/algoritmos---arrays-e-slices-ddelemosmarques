//Crie um Array de inteiros com 5 elementos
//e imprima cada elemento em uma linha separada.

package main

import "fmt"

func main() {

	var numInteiros [5]int

	numInteiros = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(numInteiros); i++ {
		fmt.Println(numInteiros[i])
	}
}
