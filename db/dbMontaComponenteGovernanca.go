package db

import ()

func initGovernanca() ComponenteESI {
	var componente ComponenteESI
	var tiposNotas []TipoNotaESI
	var tipoNota TipoNotaESI
	var elementos []ElementoESI

	var itens = []string{"Governança"}
	var elemento ElementoESI
	elemento.Itens = itens
	elemento.Nome = "Governança"
	elementos = append(elementos, elemento)

	tipoNota.Nome = "Avaliação"
	tipoNota.Elementos = elementos
	tiposNotas = append(tiposNotas, tipoNota)
	componente.Nome = "Governança"
	componente.TiposNotas = tiposNotas

	return componente
}
