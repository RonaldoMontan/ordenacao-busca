package bubbleSort

import (
	"fmt"
)

func BubbleSort(array []int) {

	// BubbleSort ordena um slice de inteiros utilizando o algoritmo Bubble Sort.
	//
	// O algoritmo percorre o slice múltiplas vezes, comparando elementos adjacentes
	// e trocando-os de lugar sempre que estão fora de ordem. Após cada passagem, o maior 
	// valor "borbulha" para a posição final. 
	//
	// Para melhorar a eficiência, a função interrompe o processo assim que detecta
	// que nenhuma troca foi feita durante uma iteração completa, indicando que o slice já está ordenado.
	//
	// Parâmetros:
	//   - array []int: slice de inteiros que será ordenado in-place (modificado diretamente).
	//
	// Exemplo de uso:
	//   numeros := []int{5, 2, 9, 1}
	//   BubbleSort(numeros)
	//
	// Após a execução, o slice estará ordenado em ordem crescente.

	var wereExchanged bool

	fmt.Println(array)
	for i:=0; i < len(array)-1; i++{
		// o primeiro for é para limitar em até quantas vezes o segundo for devera ser executado.		
		wereExchanged = false
		
		for j:=0; (j < len(array)-i -1); j++{
			
			if (array[j] > array[j+1]){
				fmt.Println("Comparing:", array[j], "and", array[j+1])

				array[j], array[j+1] = array[j+1], array[j]
				
				fmt.Println("\n\t", array)
				wereExchanged = true
			}
		}
		if !wereExchanged {
			break
		}
	}
	fmt.Println("Bubble sort finish !!\n")
	fmt.Println(array)
}