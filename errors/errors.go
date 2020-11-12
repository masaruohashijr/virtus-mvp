package models

import ()

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

const ErroChaveDuplicada = "Erro de duplicidade de chave no banco de dados. Contate o administrador do sistema."
