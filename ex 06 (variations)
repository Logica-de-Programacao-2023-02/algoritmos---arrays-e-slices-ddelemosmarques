//Crie um Array de strings com 10 elementos.
//Em seguida, solicite ao usuário que informe
//um valor e verifique se esse valor está presente no Array.
//Imprima o resultado da busca.

package main

import "fmt"

func main() {
	nomes := [10]string{"João", "Maria", "Ana", "José", "Andre", "Caio", "Ian", "Pedro", "Felipe", "Kauã"}

	var nome string

	fmt.Println("escreva um nome")
	fmt.Scan(&nome)

	encontrado := false

	for i := 0; i < len(nomes); i++ {
		if nome == nomes[i] {
			encontrado = true
			break
		}
	}

	if encontrado {
		fmt.Printf(" o nome %s está presente no array\n,", nome)
	} else {
		fmt.Printf("o nome %s não foi encontrado no array.\n", nome)
	}

}
