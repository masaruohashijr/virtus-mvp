package db

import ()

func initEconomicoFinanceiro(nome string) ComponenteESI {
	var componente ComponenteESI
	var tiposNotas []TipoNotaESI
	var tipoNota TipoNotaESI
	var elemento ElementoESI
	var elementos []ElementoESI

	var itens = []string{"Avaliação"}
	elemento.Itens = itens
	elemento.Nome = "Avaliação"
	elementos = append(elementos, elemento)

	tipoNota.Nome = "Avaliação"
	tipoNota.Elementos = elementos
	tiposNotas = append(tiposNotas, tipoNota)

	componente.Nome = nome
	componente.TiposNotas = tiposNotas
	return componente
}
