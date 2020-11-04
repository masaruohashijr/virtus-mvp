package handlers

import (
	"log"
	"time"
	mdl "virtus/models"
)

func registrarProdutosCiclos(currentUser mdl.User, entidadeId string, cicloId string) {
	sqlStatement := "INSERT INTO public.produtos_ciclos ( " +
		" entidade_id, " +
		" ciclo_id, " +
		" tipo_pontuacao_id, " +
		" author_id, " +
		" criado_em ) " +
		" values ($1, $2, $3, $4, $5) RETURNING id"
	log.Println(sqlStatement)
	produtoCicloId := 0
	err := Db.QueryRow(
		sqlStatement,
		entidadeId,
		cicloId,
		mdl.Manual,
		currentUser.Id,
		time.Now()).Scan(&produtoCicloId)
	if err != nil {
		panic(err)
	}
	sqlStatement = "INSERT INTO public.produtos_pilares ( " +
		" entidade_id, " +
		" ciclo_id, " +
		" pilar_id, " +
		" tipo_pontuacao_id, " +
		" author_id, " +
		" criado_em ) " +
		" SELECT $1, $2, pilar_id, $3, $4, $5 FROM " +
		" pilares_ciclos WHERE ciclo_id = $6 RETURNING id"
	log.Println(sqlStatement)
	produtoPilarId := 0
	err = Db.QueryRow(
		sqlStatement,
		entidadeId,
		cicloId,
		mdl.Manual,
		currentUser.Id,
		time.Now(), cicloId).Scan(&produtoPilarId)
	if err != nil {
		panic(err)
	}
	sqlStatement = "INSERT INTO public.produtos_componentes ( " +
		" entidade_id, " +
		" ciclo_id, " +
		" pilar_id, " +
		" componente_id, " +
		" tipo_pontuacao_id, " +
		" author_id, " +
		" criado_em ) " +
		" SELECT $1, $2, a.pilar_id, b.componente_id, $3, $4, $5 FROM " +
		" pilares_ciclos a LEFT JOIN " +
		" componentes_pilares b ON a.pilar_id = b.pilar_id " +
		" WHERE a.ciclo_id = $6 RETURNING id"
	log.Println(sqlStatement)
	produtoComponenteId := 0
	err = Db.QueryRow(
		sqlStatement,
		entidadeId,
		cicloId,
		mdl.Manual,
		currentUser.Id,
		time.Now(),
		cicloId).Scan(&produtoComponenteId)
	if err != nil {
		panic(err)
	}
	sqlStatement = "INSERT INTO public.produtos_elementos ( " +
		" entidade_id, " +
		" ciclo_id, " +
		" pilar_id, " +
		" componente_id, " +
		" elemento_id, " +
		" tipo_nota_id, " +
		" tipo_pontuacao_id, " +
		" author_id, " +
		" criado_em ) " +
		" SELECT $1, $2, a.pilar_id, b.componente_id, c.elemento_id, c.tipo_nota_id, $3, $4, $5 " +
		" FROM " +
		" pilares_ciclos a " +
		" LEFT JOIN " +
		" componentes_pilares b ON a.pilar_id = b.pilar_id " +
		" LEFT JOIN " +
		" elementos_componentes c ON b.componente_id = c.componente_id " +
		" WHERE a.ciclo_id = $6 RETURNING id"
	log.Println(sqlStatement)
	produtoElementoId := 0
	err = Db.QueryRow(
		sqlStatement,
		entidadeId,
		cicloId,
		mdl.Manual,
		currentUser.Id,
		time.Now(),
		cicloId).Scan(&produtoElementoId)
	if err != nil {
		panic(err)
	}
	sqlStatement = "INSERT INTO public.produtos_itens ( " +
		" entidade_id, " +
		" ciclo_id, " +
		" pilar_id, " +
		" componente_id, " +
		" elemento_id, " +
		" item_id, " +
		" author_id, " +
		" criado_em ) " +
		" SELECT $1, $2, a.pilar_id, b.componente_id, c.elemento_id, d.id, $3, $4 " +
		" FROM " +
		" pilares_ciclos a " +
		" LEFT JOIN " +
		" componentes_pilares b ON a.pilar_id = b.pilar_id " +
		" LEFT JOIN " +
		" elementos_componentes c ON b.componente_id = c.componente_id " +
		" LEFT JOIN " +
		" itens d ON d.elemento_id = c.elemento_id " +
		" WHERE a.ciclo_id = $5 RETURNING id"
	log.Println(sqlStatement)
	produtoItemId := 0
	err = Db.QueryRow(
		sqlStatement,
		entidadeId,
		cicloId,
		currentUser.Id,
		time.Now(),
		cicloId).Scan(&produtoItemId)
	if err != nil {
		panic(err)
	}
}
