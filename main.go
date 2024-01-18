package main

import (
  "fmt"
  "github.com/PedroPiveta/CRUD-ArvBin-Go/estrutura"
)

func main() {
	var raiz estrutura.ArvBin
	estrutura.ComprarIngresso(&raiz)
	fmt.Println(raiz)
}
