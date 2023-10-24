package main

import "fmt"

func main() {

	//declaaracao d matrizes
	var matrizInteiros [3][2]int

	//eclaracao variaveis
	var matrizValores int

	//loop para as linhas

	for i := 0; i < 3; i++ {
		//loop para colunas (dnto do loop d linahs)

		for j := 0; j < 2; j++ {

			//coletade valores
			fmt.Printf("escolha o valor da posição [#{i}] [#{j}]: ")
			fmt.Scan(&matrizValores)

			// aribuicao dos valores a matriz
			matrizInteiros[i][j] = matrizValores
		}
	}
	//print da matriz
	//fmt.Printf("\nSua matriz ficou: \n %d", matrizInteiros)
	for _, ints := range matrizInteiros {
		fmt.Println(ints)
	}

}
