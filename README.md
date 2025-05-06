# Algoritmos de Busca e Ordenação em Go

[![Go Version](https://img.shields.io/badge/go-1.20+-blue.svg)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](https://opensource.org/licenses/MIT)

Este repositório contém implementações de algoritmos clássicos de **busca** e **ordenação** desenvolvidos como complemento aos estudos da disciplina _"Algoritmos e Programação de Computadores II"_ da **UNIVESP**.

## 🎯 Objetivo

Implementar algoritmos fundamentais da Ciência da Computação em **Go (Golang)**, com:
- Código limpo e documentado
- Exemplos práticos
- Explicações detalhadas
- Análise de complexidade

## 📚 Algoritmos Implementados

### Ordenação
- [x] [Bubble Sort](bubbleSort/) (com otimização para parada antecipada)
- [ ] Insertion Sort
- [ ] Selection Sort
- [ ] Merge Sort
- [ ] Quick Sort

### Busca
- [ ] Busca Linear
- [ ] Busca Binária
- [ ] Busca por Hash

## 🧮 Complexidade dos Algoritmos

| Algoritmo      | Melhor Caso | Caso Médio | Pior Caso | Espaço |
|----------------|------------|------------|----------|--------|
| Bubble Sort    | O(n)       | O(n²)      | O(n²)    | O(1)   |
| Insertion Sort | O(n)       | O(n²)      | O(n²)    | O(1)   |
| Busca Binária  | O(1)       | O(log n)   | O(log n) | O(1)   |

## 🗂️ Estrutura do Projeto

```plaintext
ordenacao-busca/
├── main.go
├── bubbleSort/
│ └── bubble-sort.go
├── go.mod
└── README.md
```


## ▶️ Como Executar

### Pré-requisitos
- Go 1.20 ou superior

### Executando o projeto
```bash
# Clonar o repositório (opcional)
git clone https://github.com/seu-usuario/ordenacao-busca.git
cd ordenacao-busca
``

# Executar o programa principal
go run main.go
