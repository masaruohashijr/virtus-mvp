package handlers

import (
	"log"
	"time"
	mdl "virtus/models"
)

func registrarPesoNotaElemento(produto mdl.ProdutoElemento, currentUser mdl.User) {
	sqlStatement := "UPDATE produtos_elementos SET peso = $1, nota = $2 " +
		" WHERE entidade_id = $3 AND " +
		" ciclo_id = $4 AND " +
		" pilar_id = $5 AND " +
		" componente_id = $6 AND " +
		" elemento_id = $7 "
	log.Println(sqlStatement)
	updtForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	updtForm.Exec(produto.Peso,
		produto.Nota,
		produto.EntidadeId,
		produto.CicloId,
		produto.PilarId,
		produto.ComponenteId,
		produto.ElementoId)

	sqlStatement = "UPDATE produtos_componentes a " +
		" SET peso = $1, nota = (SELECT round(CAST(avg(b.peso) as numeric),2) " +
		" FROM produtos_elementos b " +
		" WHERE b.componente_id = a.componente_id AND " +
		" b.pilar_id = a.pilar_id AND " +
		" b.ciclo_id = a.ciclo_id AND " +
		" b.entidade_id = a.entidade_id " +
		" GROUP BY b.entidade_id, b.ciclo_id, b.pilar_id, b.componente_id) " +
		" WHERE entidade_id = $2 AND " +
		" ciclo_id = $3 AND " +
		" pilar_id = $4 AND " +
		" componente_id = $5 "
	log.Println(sqlStatement)
	/*updtForm, err = Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	updtForm.Exec(produto.Peso,
		produto.Nota,
		produto.EntidadeId,
		produto.CicloId,
		produto.PilarId,
		produto.ComponenteId)*/
}

func registrarProdutosCiclos(currentUser mdl.User, entidadeId string, cicloId string) {
	sqlStatement := "INSERT INTO public.produtos_ciclos ( " +
		" entidade_id, " +
		" ciclo_id, " +
		" tipo_pontuacao_id, " +
		" author_id, " +
		" criado_em ) " +
		" SELECT " +
		entidadeId + ", " +
		cicloId + ", " +
		" $1, " +
		" $2, " +
		" $3 " +
		" FROM ciclos_entidades a " +
		" WHERE NOT EXISTS " +
		"  (SELECT 1 " +
		"   FROM produtos_ciclos b " +
		"   WHERE b.entidade_id = a.entidade_id " +
		"     AND b.ciclo_id = a.ciclo_id) RETURNING id "
	log.Println(sqlStatement)
	produtoCicloId := 0
	err := Db.QueryRow(
		sqlStatement,
		mdl.Calculada,
		currentUser.Id,
		time.Now()).Scan(&produtoCicloId)
	if err != nil {
		log.Println(err)
	}
	sqlStatement = "INSERT INTO produtos_pilares " +
		" (entidade_id, ciclo_id, pilar_id, peso, tipo_pontuacao_id, author_id, criado_em) " +
		" SELECT " +
		entidadeId + ", " +
		cicloId + ", " +
		" a.pilar_id, " +
		" a.peso_padrao, " +
		" $1, " +
		" $2, " +
		" $3 " +
		" FROM pilares_ciclos a " +
		" WHERE NOT EXISTS " +
		"  (SELECT 1 " +
		"   FROM produtos_pilares b " +
		"   WHERE b.entidade_id = " + entidadeId +
		"     AND b.ciclo_id = a.ciclo_id " +
		"     AND b.pilar_id = a.pilar_id) RETURNING id"
	log.Println(sqlStatement)
	produtoPilarId := 0
	err = Db.QueryRow(
		sqlStatement,
		mdl.Manual,
		currentUser.Id,
		time.Now()).Scan(&produtoPilarId)
	if err != nil {
		log.Println(err)
	}
	sqlStatement = "INSERT INTO public.produtos_componentes ( " +
		" entidade_id, " +
		" ciclo_id, " +
		" pilar_id, " +
		" componente_id, " +
		" peso, " +
		" tipo_pontuacao_id, " +
		" author_id, " +
		" criado_em ) " +
		" SELECT " + entidadeId + ", " + cicloId + ", a.pilar_id, b.componente_id, b.peso_padrao, $1, $2, $3 FROM " +
		" pilares_ciclos a LEFT JOIN " +
		" componentes_pilares b ON a.pilar_id = b.pilar_id " +
		" WHERE NOT EXISTS " +
		"  (SELECT 1 " +
		"   FROM produtos_componentes c " +
		"   WHERE c.entidade_id = " + entidadeId +
		"     AND c.ciclo_id = a.ciclo_id " +
		"     AND c.pilar_id = a.pilar_id " +
		"     AND c.componente_id = b.componente_id) RETURNING id"
	log.Println(sqlStatement)
	produtoComponenteId := 0
	err = Db.QueryRow(
		sqlStatement,
		mdl.Manual,
		currentUser.Id,
		time.Now()).Scan(&produtoComponenteId)
	if err != nil {
		log.Println(err)
	}
	sqlStatement = "INSERT INTO public.produtos_elementos ( " +
		" entidade_id, " +
		" ciclo_id, " +
		" pilar_id, " +
		" componente_id, " +
		" elemento_id, " +
		" tipo_nota_id, " +
		" peso," +
		" tipo_pontuacao_id, " +
		" author_id, " +
		" criado_em ) " +
		" SELECT " + entidadeId + ", " + cicloId + ", a.pilar_id, b.componente_id, " +
		" c.elemento_id, c.tipo_nota_id, c.peso_padrao, $1, $2, $3 " +
		" FROM " +
		" pilares_ciclos a " +
		" LEFT JOIN " +
		" componentes_pilares b ON a.pilar_id = b.pilar_id " +
		" LEFT JOIN " +
		" elementos_componentes c ON b.componente_id = c.componente_id " +
		" WHERE NOT EXISTS " +
		"  (SELECT 1 " +
		"   FROM produtos_elementos d " +
		"   WHERE d.entidade_id = " + entidadeId +
		"     AND d.ciclo_id = a.ciclo_id " +
		"     AND d.pilar_id = a.pilar_id " +
		"     AND d.componente_id = b.componente_id " +
		"     AND d.elemento_id = c.elemento_id) RETURNING id"
	log.Println(sqlStatement)
	produtoElementoId := 0
	err = Db.QueryRow(
		sqlStatement,
		mdl.Manual,
		currentUser.Id,
		time.Now()).Scan(&produtoElementoId)
	if err != nil {
		log.Println(err)
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
		" SELECT " + entidadeId + ", " + cicloId + ", " +
		" a.pilar_id, b.componente_id, c.elemento_id, d.id, $1, $2 " +
		" FROM pilares_ciclos a " +
		" LEFT JOIN componentes_pilares b ON a.pilar_id = b.pilar_id " +
		" LEFT JOIN elementos_componentes c ON b.componente_id = c.componente_id " +
		" LEFT JOIN itens d ON c.elemento_id = d.elemento_id " +
		" WHERE NOT EXISTS " +
		"    (SELECT 1 " +
		"     FROM produtos_itens e " +
		"     WHERE e.entidade_id = " + entidadeId +
		"       AND e.ciclo_id = a.ciclo_id " +
		"       AND e.pilar_id = a.pilar_id " +
		"       AND e.componente_id = b.componente_id " +
		"	   AND e.elemento_id = c.elemento_id " +
		"	   AND e.item_id = d.id) "
	log.Println(sqlStatement)
	produtoItemId := 0
	err = Db.QueryRow(
		sqlStatement,
		currentUser.Id,
		time.Now()).Scan(&produtoItemId)
	if err != nil {
		log.Println(err)
	}
}
