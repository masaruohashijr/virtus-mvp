package db

import ()

func initEconomicoFinanceiro(nome string) ComponenteESI {
	var componente ComponenteESI
	var tiposNotas []TipoNotaESI
	var tipoNota TipoNotaESI
	var elementoA ElementoESI
	var elementoB ElementoESI
	var elementoC ElementoESI
	var elementos []ElementoESI

	var itensA = []string{"A1", "A2"}
	elementoA.Itens = itensA
	elementoA.Nome = "A"
	elementos = append(elementos, elementoA)
	var itensB = []string{"B1", "B2"}
	elementoB.Itens = itensB
	elementoB.Nome = "B"
	elementos = append(elementos, elementoB)
	var itensC = []string{"C1", "C2"}
	elementoC.Itens = itensC
	elementoC.Nome = "C"
	elementos = append(elementos, elementoC)

	tipoNota.Nome = "Avaliação"
	tipoNota.Elementos = elementos
	tiposNotas = append(tiposNotas, tipoNota)

	componente.Nome = nome
	componente.TiposNotas = tiposNotas
	return componente
}
