package handlers

import (
	"log"
	"strconv"
	"time"
	mdl "virtus/models"
)

func registrarAuditorComponente(produto mdl.ProdutoElemento) {
	sqlStatement := "UPDATE produtos_componentes SET " +
		" auditor_id=$1, justificativa=$2 " +
		" WHERE entidade_id=$3 " +
		" AND ciclo_id=$4 " +
		" AND pilar_id=$5 " +
		" AND componente_id=$6 "
	log.Println(sqlStatement)
	updtForm, _ := Db.Prepare(sqlStatement)
	_, err := updtForm.Exec(
		produto.AuditorId,
		produto.Motivacao,
		produto.EntidadeId,
		produto.CicloId,
		produto.PilarId,
		produto.ComponenteId)
	if err != nil {
		panic(err.Error())
	}
}

func registrarNotaElemento(produto mdl.ProdutoElemento, currentUser mdl.User) {
	sqlStatement := "UPDATE produtos_elementos a SET nota = " + strconv.Itoa(produto.Nota) + ", " +
		" motivacao_nota = $1 , " +
		" tipo_pontuacao_id = (SELECT case when b.supervisor_id = " + strconv.FormatInt(currentUser.Id, 10) +
		" then 3 else 1 end FROM produtos_componentes b where " +
		" a.entidade_id = b.entidade_id and " +
		" a.ciclo_id = b.ciclo_id and " +
		" a.pilar_id = b.pilar_id and " +
		" a.componente_id = b.componente_id) " +
		" WHERE a.entidade_id = " + strconv.FormatInt(produto.EntidadeId, 10) +
		" AND a.ciclo_id = " + strconv.FormatInt(produto.CicloId, 10) +
		" AND a.pilar_id = " + strconv.FormatInt(produto.PilarId, 10) +
		" AND a.componente_id = " + strconv.FormatInt(produto.ComponenteId, 10) +
		" AND a.elemento_id = " + strconv.FormatInt(produto.ElementoId, 10) +
		" AND a.nota <> " + strconv.Itoa(produto.Nota)
	log.Println(sqlStatement)
	updtForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	updtForm.Exec(produto.Motivacao)
	// Testei e funcionou corretamente
	// PRODUTOS_TIPOS_NOTAS
	sqlStatement = "UPDATE produtos_tipos_notas a " +
		" set nota = (select  " +
		" round(CAST(sum(nota*peso)/sum(peso) as numeric),2) as media " +
		" FROM produtos_elementos b " +
		" WHERE " +
		" a.entidade_id = b.entidade_id " +
		" and a.ciclo_id = b.ciclo_id  " +
		" and a.pilar_id = b.pilar_id " +
		" and a.componente_id = b.componente_id " +
		" and a.tipo_nota_id = b.tipo_nota_id " +
		" GROUP BY b.entidade_id,  " +
		" b.ciclo_id, " +
		" b.pilar_id, " +
		" b.componente_id, " +
		" b.tipo_nota_id " +
		" HAVING sum(peso)>0) " +
		" WHERE a.entidade_id = $1 " +
		" AND a.ciclo_id = $2 "
	log.Println(sqlStatement)
	updtForm, err = Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	updtForm.Exec(produto.EntidadeId, produto.CicloId)
	// PRODUTOS_COMPONENTES
	sqlStatement = "UPDATE produtos_componentes a " +
		" set nota = (select  " +
		" round(CAST(sum(nota*peso)/sum(peso) as numeric),2) as media " +
		" FROM produtos_tipos_notas b " +
		" WHERE " +
		" a.entidade_id = b.entidade_id " +
		" and a.ciclo_id = b.ciclo_id  " +
		" and a.pilar_id = b.pilar_id " +
		" and a.componente_id = b.componente_id " +
		" and a.id_versao_origem is null " +
		" GROUP BY b.entidade_id,  " +
		" b.ciclo_id, " +
		" b.pilar_id, " +
		" b.componente_id " +
		" HAVING sum(peso)>0) " +
		" WHERE a.entidade_id = $1 " +
		" AND a.ciclo_id = $2 "
	log.Println(sqlStatement)
	updtForm, err = Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	updtForm.Exec(produto.EntidadeId, produto.CicloId)
	// PRODUTOS_PILARES
	sqlStatement = "UPDATE produtos_pilares a " +
		" SET nota = (select  " +
		" round(CAST(sum(nota*peso)/sum(peso) as numeric),2) AS media " +
		" FROM produtos_componentes b " +
		" WHERE " +
		" a.entidade_id = b.entidade_id " +
		" AND a.ciclo_id = b.ciclo_id  " +
		" AND a.pilar_id = b.pilar_id " +
		" GROUP BY b.entidade_id,  " +
		" b.ciclo_id, " +
		" b.pilar_id " +
		" HAVING sum(peso)>0) " +
		" WHERE a.entidade_id = $1 " +
		" AND a.ciclo_id = $2 "
	log.Println(sqlStatement)
	updtForm, err = Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	updtForm.Exec(produto.EntidadeId, produto.CicloId)
	// PRODUTOS_CICLOS
	sqlStatement = "UPDATE produtos_ciclos a " +
		" SET nota = (select  " +
		" round(CAST(sum(nota*peso)/sum(peso) as numeric),2) AS media " +
		" FROM produtos_pilares b " +
		" WHERE " +
		" a.entidade_id = b.entidade_id " +
		" AND a.ciclo_id = b.ciclo_id  " +
		" GROUP BY b.entidade_id,  " +
		" b.ciclo_id " +
		" HAVING sum(peso)>0) " +
		" WHERE a.entidade_id = $1 " +
		" AND a.ciclo_id = $2 "
	log.Println(sqlStatement)
	updtForm, err = Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	updtForm.Exec(produto.EntidadeId, produto.CicloId)
	// registrarTiposPontuacao(produto, currentUser)
}

func registrarTiposPontuacao(produto mdl.ProdutoElemento, currentUser mdl.User) {
	sqlStatement := "UPDATE produtos_tipos_notas a SET " +
		" tipo_pontuacao_id = (SELECT case when b.supervisor_id = $1 " +
		" then 3 else 2 end FROM produtos_componentes b where a.id = b.id) " +
		" WHERE entidade_id = $2 " +
		" AND  ciclo_id = $3 " +
		" AND  pilar_id = $4 " +
		" AND  componente_id = $5 " +
		" AND  tipo_nota_id = $6 "
	log.Println(sqlStatement)
	updtForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	updtForm.Exec(
		currentUser.Id,
		produto.EntidadeId,
		produto.CicloId,
		produto.PilarId,
		produto.ComponenteId,
		produto.TipoNotaId)
	sqlStatement = "UPDATE produtos_componentes a SET " +
		" tipo_pontuacao_id = (SELECT case when b.supervisor_id = $1 " +
		" then 3 else 2 end FROM produtos_componentes b where a.id = b.id) " +
		" WHERE entidade_id = $2 " +
		" AND  ciclo_id = $3 " +
		" AND  pilar_id = $4 " +
		" AND  componente_id = $5 "
	log.Println(sqlStatement)
	updtForm, err = Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	updtForm.Exec(
		currentUser.Id,
		produto.EntidadeId,
		produto.CicloId,
		produto.PilarId,
		produto.ComponenteId)
	sqlStatement = "UPDATE produtos_pilares a SET " +
		" tipo_pontuacao_id = (SELECT case when b.supervisor_id = $1 " +
		" then 3 else 2 end FROM produtos_pilares b where a.id = b.id) " +
		" WHERE entidade_id = $2 " +
		" AND  ciclo_id = $3 " +
		" AND  pilar_id = $4 "
	log.Println(sqlStatement)
	updtForm, err = Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	updtForm.Exec(
		currentUser.Id,
		produto.EntidadeId,
		produto.CicloId,
		produto.PilarId)
	sqlStatement = "UPDATE produtos_ciclos a SET " +
		" tipo_pontuacao_id = (SELECT case when b.supervisor_id = $1 " +
		" then 3 else 2 end FROM produtos_ciclos b where a.id = b.id) " +
		" WHERE entidade_id = $2 " +
		" AND  ciclo_id = $3 "
	log.Println(sqlStatement)
	updtForm, err = Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	updtForm.Exec(
		currentUser.Id,
		produto.EntidadeId,
		produto.CicloId)
}

func registrarPesoElemento(produto mdl.ProdutoElemento, currentUser mdl.User) {
	sqlStatement := "UPDATE produtos_elementos SET peso = $1, motivacao_peso = $2 " +
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
		produto.Motivacao,
		produto.EntidadeId,
		produto.CicloId,
		produto.PilarId,
		produto.ComponenteId,
		produto.ElementoId)
	// Testei e funcionou corretamente
	sqlStatement = "UPDATE produtos_tipos_notas as p SET peso = R1.peso " +
		"FROM " +
		"(WITH TMP AS " +
		"     (SELECT entidade_id, " +
		"             ciclo_id, " +
		"             pilar_id, " +
		"             componente_id, " +
		"             round(CAST(sum(peso) AS numeric), 2) AS TOTAL " +
		"      FROM produtos_elementos " +
		"      WHERE componente_id = " + strconv.FormatInt(produto.ComponenteId, 10) +
		"        AND pilar_id = " + strconv.FormatInt(produto.PilarId, 10) +
		"        AND ciclo_id = " + strconv.FormatInt(produto.CicloId, 10) +
		"        AND entidade_id = " + strconv.FormatInt(produto.EntidadeId, 10) +
		"      GROUP BY entidade_id, " +
		"               ciclo_id, " +
		"               pilar_id, " +
		"               componente_id)  " +
		"   SELECT r.entidade_id, r.ciclo_id, r.pilar_id, r.componente_id, r.tipo_nota_id, round(CAST((sum(r.peso)/(sum(t.TOTAL)/count(1)))*100 AS numeric), 2) AS peso " +
		"   FROM " +
		"     (SELECT b.entidade_id, " +
		"             b.ciclo_id, " +
		"             b.pilar_id, " +
		"             b.componente_id, " +
		"             b.tipo_nota_id, " +
		"             b.peso " +
		"      FROM produtos_elementos b " +
		"      WHERE componente_id = " + strconv.FormatInt(produto.ComponenteId, 10) +
		"        AND pilar_id = " + strconv.FormatInt(produto.PilarId, 10) +
		"        AND ciclo_id = " + strconv.FormatInt(produto.CicloId, 10) +
		"        AND entidade_id = " + strconv.FormatInt(produto.EntidadeId, 10) +
		"   ) r " +
		"   LEFT JOIN TMP t ON r.entidade_id = t.entidade_id " +
		"   AND r.ciclo_id = t.ciclo_id " +
		"   AND r.pilar_id = t.pilar_id " +
		"   AND r.componente_id = t.componente_id " +
		"   GROUP BY r.entidade_id, " +
		"            r.ciclo_id, " +
		"            r.pilar_id, " +
		"            r.componente_id, " +
		"            r.tipo_nota_id) AS R1 " +
		"	WHERE " +
		"		 p.tipo_nota_id = R1.tipo_nota_id " +
		"	 	 AND p.componente_id = R1.componente_id " +
		"        AND p.pilar_id = R1.pilar_id " +
		"        AND p.ciclo_id = R1.ciclo_id " +
		"        AND p.entidade_id = R1.entidade_id "
	log.Println(sqlStatement)
	updtForm, err = Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	res, _ := updtForm.Exec()
	log.Println(res.RowsAffected())
	// Testei e funcionou corretamente
	sqlStatement = "UPDATE produtos_componentes a " +
		" SET peso = (SELECT round(CAST(avg(b.peso) as numeric),2) " +
		" FROM produtos_elementos b " +
		" WHERE b.componente_id = a.componente_id AND " +
		" b.pilar_id = a.pilar_id AND " +
		" b.ciclo_id = a.ciclo_id AND " +
		" b.entidade_id = a.entidade_id " +
		" GROUP BY b.entidade_id, b.ciclo_id, b.pilar_id, b.componente_id) " +
		" WHERE entidade_id = $1 AND " +
		" ciclo_id = $2 AND " +
		" pilar_id = $3 AND " +
		" componente_id = $4 "
	log.Println(sqlStatement)
	updtForm, err = Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	updtForm.Exec(
		produto.EntidadeId,
		produto.CicloId,
		produto.PilarId,
		produto.ComponenteId)
}

func registrarProdutosCiclos(currentUser mdl.User, entidadeId string, cicloId string) {
	sqlStatement := "INSERT INTO produtos_ciclos ( " +
		" entidade_id, " +
		" ciclo_id, " +
		" nota, " +
		" tipo_pontuacao_id, " +
		" author_id, " +
		" criado_em ) " +
		" SELECT " +
		entidadeId + ", " +
		cicloId + ", " +
		" 1, " +
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
		" (entidade_id, ciclo_id, pilar_id, peso, nota, tipo_pontuacao_id, author_id, criado_em) " +
		" SELECT " +
		entidadeId + ", " +
		cicloId + ", " +
		" a.pilar_id, " +
		" a.peso_padrao, " +
		" 1, " +
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
		mdl.Calculada,
		currentUser.Id,
		time.Now()).Scan(&produtoPilarId)
	if err != nil {
		log.Println(err)
	}

	sqlStatement = "INSERT INTO produtos_elementos ( " +
		" entidade_id, " +
		" ciclo_id, " +
		" pilar_id, " +
		" componente_id, " +
		" elemento_id, " +
		" tipo_nota_id, " +
		" peso," +
		" nota," +
		" tipo_pontuacao_id, " +
		" author_id, " +
		" criado_em ) " +
		" SELECT " + entidadeId + ", " + cicloId + ", a.pilar_id, b.componente_id, " +
		" c.elemento_id, c.tipo_nota_id, c.peso_padrao, 1, $1, $2, $3 " +
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
		mdl.Calculada,
		currentUser.Id,
		time.Now()).Scan(&produtoElementoId)
	if err != nil {
		log.Println(err)
	}

	sqlStatement = "INSERT INTO public.produtos_componentes ( " +
		" entidade_id, " +
		" ciclo_id, " +
		" pilar_id, " +
		" componente_id, " +
		" peso, " +
		" nota, " +
		" tipo_pontuacao_id, " +
		" author_id, " +
		" criado_em ) " +
		" SELECT " + entidadeId + ", " + cicloId + ", a.pilar_id, b.componente_id, " +
		" round(CAST(avg(c.peso_padrao) AS numeric),2), 1, " +
		" $1, $2, $3 " +
		" FROM " +
		" PILARES_CICLOS a " +
		" LEFT JOIN COMPONENTES_PILARES b ON (a.pilar_id = b.pilar_id) " +
		" LEFT JOIN ELEMENTOS_COMPONENTES c ON (b.componente_id = c.componente_id) " +
		" WHERE  " +
		" a.ciclo_id = " + cicloId +
		" AND NOT EXISTS " +
		"  (SELECT 1 " +
		"   FROM produtos_componentes c " +
		"   WHERE c.entidade_id = " + entidadeId +
		"     AND c.ciclo_id = a.ciclo_id " +
		"     AND c.pilar_id = a.pilar_id " +
		"     AND c.componente_id = b.componente_id) " +
		" GROUP BY 1,2,3,4 ORDER BY 1,2,3,4" +
		" RETURNING id"
	log.Println(sqlStatement)
	produtoComponenteId := 0
	err = Db.QueryRow(
		sqlStatement,
		mdl.Calculada,
		currentUser.Id,
		time.Now()).Scan(&produtoComponenteId)
	if err != nil {
		log.Println(err)
	}

	sqlStatement = "INSERT INTO public.produtos_tipos_notas ( " +
		" entidade_id, " +
		" ciclo_id, " +
		" pilar_id, " +
		" componente_id, " +
		" tipo_nota_id, " +
		" peso, " +
		" nota, " +
		" tipo_pontuacao_id, " +
		" author_id, " +
		" criado_em ) " +
		" SELECT " + entidadeId + ", " + cicloId + ", a.pilar_id, b.componente_id, " +
		" d.tipo_nota_id, " +
		" round(CAST(avg(d.peso_padrao) AS numeric),2), 1, " +
		" $1, $2, $3 " +
		" FROM " +
		" PILARES_CICLOS a " +
		" LEFT JOIN COMPONENTES_PILARES b ON (a.pilar_id = b.pilar_id) " +
		" LEFT JOIN ELEMENTOS_COMPONENTES c ON (b.componente_id = c.componente_id) " +
		" LEFT JOIN TIPOS_NOTAS_COMPONENTES d ON (b.componente_id = d.componente_id AND c.tipo_nota_id = d.tipo_nota_id) " +
		" WHERE  " +
		" a.ciclo_id = " + cicloId +
		" AND NOT EXISTS " +
		"  (SELECT 1 " +
		"   FROM produtos_tipos_notas e " +
		"   WHERE e.entidade_id = " + entidadeId +
		"     AND e.ciclo_id = a.ciclo_id " +
		"     AND e.pilar_id = a.pilar_id " +
		"     AND e.tipo_nota_id = d.tipo_nota_id " +
		"     AND e.componente_id = b.componente_id) " +
		" GROUP BY 1,2,3,4,5 ORDER BY 1,2,3,4,5" +
		" RETURNING id"
	log.Println(sqlStatement)
	produtoTipoNotaId := 0
	err = Db.QueryRow(
		sqlStatement,
		mdl.Calculada,
		currentUser.Id,
		time.Now()).Scan(&produtoTipoNotaId)
	if err != nil {
		log.Println(err)
	}

	sqlStatement = "INSERT INTO produtos_itens ( " +
		" entidade_id, " +
		" ciclo_id, " +
		" pilar_id, " +
		" componente_id, " +
		" tipo_nota_id, " +
		" elemento_id, " +
		" item_id, " +
		" author_id, " +
		" criado_em ) " +
		" SELECT " + entidadeId + ", " + cicloId + ", " +
		" a.pilar_id, b.componente_id, c.tipo_nota_id, c.elemento_id, d.id, $1, $2 " +
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

	log.Println("INICIANDO CICLO --  UPDATE NOTA")
	// Testei e funcionou corretamente
	sqlStatement = "UPDATE produtos_componentes a " +
		" set nota = (select  " +
		" sum(nota*peso)/sum(peso) as media " +
		" FROM produtos_elementos b " +
		" WHERE " +
		" a.entidade_id = b.entidade_id " +
		" and a.ciclo_id = b.ciclo_id  " +
		" and a.pilar_id = b.pilar_id " +
		" and a.componente_id = b.componente_id " +
		" GROUP BY b.entidade_id,  " +
		" b.ciclo_id, " +
		" b.pilar_id, " +
		" b.componente_id " +
		" HAVING sum(peso)>0) " +
		" WHERE a.entidade_id = $1 " +
		" AND a.ciclo_id = $2 "
	log.Println(sqlStatement)
	updtForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	updtForm.Exec(entidadeId, cicloId)

	// PRODUTOS TIPOS NOTAS
	sqlStatement = "UPDATE produtos_tipos_notas a " +
		" set nota = (select  " +
		" sum(nota*peso)/sum(peso) as media " +
		" FROM produtos_elementos b " +
		" WHERE " +
		" a.entidade_id = b.entidade_id " +
		" and a.ciclo_id = b.ciclo_id  " +
		" and a.pilar_id = b.pilar_id " +
		" and a.componente_id = b.componente_id " +
		" and a.tipo_nota_id = b.tipo_nota_id " +
		" GROUP BY b.entidade_id,  " +
		" b.ciclo_id, " +
		" b.pilar_id, " +
		" b.componente_id, " +
		" b.tipo_nota_id " +
		" HAVING sum(peso)>0) " +
		" WHERE a.entidade_id = $1 " +
		" AND a.ciclo_id = $2 "
	log.Println(sqlStatement)
	updtForm, err = Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	updtForm.Exec(entidadeId, cicloId)
}
