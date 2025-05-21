package insertionSort

import (
	"fmt"
)

func InsertionSort(array []int) {
    // InsertionSort ordena um slice de inteiros utilizando o algoritmo Insertion Sort.
    //
    // O algoritmo percorre o slice uma vez, inserindo cada elemento na posição correta
    // dentro da parte já ordenada (à esquerda). Para isso, ele desloca elementos maiores
    // para a direita até encontrar a posição adequada para o elemento atual.
    //
    // Diferente do Bubble Sort, o Insertion Sort é eficiente para conjuntos pequenos ou
    // quase ordenados, com complexidade O(n) no melhor caso e O(n²) no pior caso.
    //
    // Parâmetros:
    //   - array []int: slice de inteiros que será ordenado in-place (modificado diretamente).
    //
    // Exemplo de uso:
    //   numeros := []int{5, 2, 9, 1}
    //   InsertionSort(numeros)
    //
    // Após a execução, o slice estará ordenado em ordem crescente:
    //   [1, 2, 5, 9]
    //
	// Após a execução, o slice estará ordenado em ordem crescente.

    for i := 0; i < len(array); i++ {
        key := array[i]
        j := i - 1
        
        for j >= 0 && array[j] > key {
            array[j+1] = array[j]
            j--
        }
        array[j+1] = key
        fmt.Println("\n\t", array) 
    }
    fmt.Println("Insertion sort finish !!\n")
    fmt.Println(array)
}