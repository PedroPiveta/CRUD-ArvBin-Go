package main

import "fmt"

type No struct {
	num int
	esq *No
	dir *No
}

type ArvBin *No

func comprarIngresso(raiz *ArvBin) {
	var novo No
	fmt.Println("--- BEM VINDO A COMPRA DE INGRESSOS --- ")
	fmt.Println("Digite o número do ingresso")
	fmt.Scanf("%d", &novo.num)
	novo.dir = nil
	novo.esq = nil

	if novo.num <= 0 || novo.num > 101 {
		fmt.Println("\nIngresso indisponível")
		return
	}

	if *raiz == nil {
		*raiz = &novo
	} else {
		atual := *raiz
		var ant *No = nil

		for atual != nil {
			ant = atual

			if novo.num == atual.num {
				fmt.Println("Poltrona já ocupada")
				return
			}
			if novo.num > atual.num {
				atual = atual.dir
			} else {
				atual = atual.esq
			}
		}

		if novo.num > ant.num {
			ant.dir = &novo
		} else {
			ant.esq = &novo
		}
	}
	fmt.Println("Ingresso comprado")
}

func main() {
	var raiz ArvBin
	comprarIngresso(&raiz)
	fmt.Println(raiz)
}
