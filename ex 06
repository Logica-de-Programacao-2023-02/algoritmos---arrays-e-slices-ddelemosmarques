//Crie um Array de inteiros com 10 elementos.
//Em seguida, solicite ao usuário que informe
//um valor e verifique se esse valor está presente no Array.
//Imprima o resultado da busca.

package main

import "fmt"

func main() {

	var arrayInt = [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 99}

	var num int

	//numero informado pelo user
	fmt.Print("Escolha um numero: ")
	fmt.Scan(&num)

	//enquanto o i for menor do que o tamanho do array, i++
	for i := 0; i < len(arrayInt); i++ {

		//caso o numero seja igual ao valor do index i no array
		if num == arrayInt[i] {

			fmt.Printf("O número %d está presente no array\n%d", num, arrayInt)

		}

	}

}
