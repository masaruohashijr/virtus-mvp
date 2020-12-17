package db

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type ElementoESI struct {
	Nome  string
	Itens []string
}

type TipoNotaESI struct {
	Nome      string
	Elementos []ElementoESI
}

type ComponenteESI struct {
	Nome       string
	TiposNotas []TipoNotaESI
}

type PilarESI struct {
	Nome        string
	Componentes []ComponenteESI
}

type CicloESI struct {
	Nome    string
	Pilares []PilarESI
}

func createCicloCompleto() {

	// createTiposNotas()
	var tiposSalvos = make(map[string]int)
	var elementosSalvos = make(map[string]int)
	autor := 1
	tipoMedia := 1
	criadoEm := time.Now().Format("02-Jan-2006 15:04:05")
	statusZero := 0
	idCiclo := 0
	idPilar := 0
	idComponente := 0
	idTipoNota := 0
	idTipoNotaComponente := 0
	idElemento := 0
	idItem := 0
	pesoPadrao := 1
	//var ciclo CicloESI
	cicloESI := montarCiclo()

	stmtElementosComponentes := " INSERT INTO " +
		" elementos_componentes( " +
		" elemento_id, " +
		" componente_id, " +
		" tipo_nota_id, " +
		" peso_padrao, " +
		" author_id, " +
		" criado_em ) %s"

	stmtPilaresCiclos := "INSERT INTO " +
		" pilares_ciclos( " +
		" ciclo_id, " +
		" pilar_id, " +
		" tipo_media, " +
		" peso_padrao, " +
		" author_id, " +
		" criado_em ) %s"

	stmtComponentesPilares := " INSERT INTO " +
		" componentes_pilares( " +
		" pilar_id, " +
		" componente_id, " +
		" tipo_media, " +
		" peso_padrao, " +
		" author_id, " +
		" criado_em ) %s"

	var unsavedPilaresCiclos []string
	var unsavedComponentesPilares []string
	var unsavedElementosComponentes []string
	nome := cicloESI.Nome
	descricao := "Descricao do " + nome
	stmt := " INSERT INTO ciclos(nome, descricao, author_id, criado_em, status_id) " +
		" SELECT $1, $2, $3, $4, $5 WHERE NOT EXISTS (SELECT id FROM ciclos WHERE nome = '" + nome + "' ) RETURNING id"
	/*log.Println(stmt)
	log.Println(nome)
	log.Println(descricao)
	log.Println(autor)
	log.Println(criadoEm)
	log.Println(statusZero)*/
	db.QueryRow(stmt, nome, descricao, autor, criadoEm, statusZero).Scan(&idCiclo)

	if idCiclo == 0 {
		log.Println("SAINDO DO CICLO COMPLETO")
		return
	}

	pesoPadrao = 100
	max := 100
	qtdPilares := len(cicloESI.Pilares)
	for j := 0; j < qtdPilares; j++ {
		nome = cicloESI.Pilares[j].Nome
		stmt := " INSERT INTO pilares(nome, descricao, author_id, criado_em, status_id) " +
			" SELECT $1, $2, $3, $4, $5 WHERE NOT EXISTS (SELECT id FROM pilares WHERE nome = '" + nome + "' ) RETURNING id"
		descricao = "Descricao do " + nome
		db.QueryRow(stmt, nome, descricao, autor, criadoEm, statusZero).Scan(&idPilar)
		log.Println("idPilar: " + strconv.Itoa(idPilar) + " - " + nome)
		pesoPadrao = rand.Intn(max)
		if j < qtdPilares-1 {
			max = max - pesoPadrao
		} else {
			pesoPadrao = max
		}
		stmt = " SELECT " + strconv.Itoa(idCiclo) + ", " +
			strconv.Itoa(idPilar) + ", " +
			strconv.Itoa(tipoMedia) + ", " +
			strconv.Itoa(pesoPadrao) + ", " +
			strconv.Itoa(autor) + ", " +
			" to_timestamp('" + criadoEm + "','DD-Mon-YYYY HH24:MI:SS') " +
			" WHERE NOT EXISTS ( SELECT id FROM pilares_ciclos WHERE ciclo_id = " + strconv.Itoa(idCiclo) + " AND pilar_id = " + strconv.Itoa(idPilar) + " ) "
		unsavedPilaresCiclos = append(unsavedPilaresCiclos, stmt)
		//log.Println(stmt)
		qtdComponentes := len(cicloESI.Pilares[j].Componentes)
		for k := 0; k < qtdComponentes; k++ {
			nome = cicloESI.Pilares[j].Componentes[k].Nome
			idComponente = 0
			idElemento = 0
			stmt := " INSERT INTO componentes(nome, descricao, author_id, criado_em, status_id) " +
				" SELECT $1, $2, $3, $4, $5 WHERE NOT EXISTS (SELECT id FROM componentes WHERE nome = '" + nome + "' ) RETURNING id"
			descricao = "Descricao do " + nome
			db.QueryRow(stmt, nome, descricao, autor, criadoEm, statusZero).Scan(&idComponente)
			log.Println("idComponente: " + strconv.Itoa(idComponente) + " - " + nome)
			pesoPadrao = int(math.Pow(2, float64(k-1)))
			stmt = " SELECT " + strconv.Itoa(idPilar) + ", " +
				strconv.Itoa(idComponente) + ", " +
				strconv.Itoa(tipoMedia) + ", " +
				strconv.Itoa(pesoPadrao) + ", " +
				strconv.Itoa(autor) + ", " +
				" to_timestamp('" + criadoEm + "','DD-Mon-YYYY HH24:MI:SS') " +
				" WHERE NOT EXISTS ( SELECT id FROM componentes_pilares WHERE componente_id = " + strconv.Itoa(idComponente) + " AND pilar_id = " + strconv.Itoa(idPilar) + " ) "
			unsavedComponentesPilares = append(unsavedComponentesPilares, stmt)

			qtdTiposNotas := len(cicloESI.Pilares[j].Componentes[k].TiposNotas)

			for l := 0; l < qtdTiposNotas; l++ {
				idTipoNota = 0
				nome = cicloESI.Pilares[j].Componentes[k].TiposNotas[l].Nome
				letra := nome[0:1]
				descricao = "Descricao do " + nome
				corletra := "C0D0C0"
				stmt := " INSERT INTO tipos_notas ( " +
					" nome, descricao, letra, cor_letra, author_id, criado_em, status_id) " +
					" SELECT  $1, $2, $3, $4, $5, $6, $7 " +
					" WHERE NOT EXISTS (SELECT id FROM tipos_notas WHERE letra = $8) RETURNING id"
				db.QueryRow(stmt, nome, descricao, letra, corletra, autor, criadoEm, statusZero, letra).Scan(&idTipoNota)
				if idTipoNota != 0 {
					tiposSalvos[letra] = idTipoNota
				} else {
					idTipoNota = tiposSalvos[letra]
				}
				log.Println("idTipoNota: " + strconv.Itoa(idTipoNota) + " - " + nome)

				stmt = " INSERT INTO tipos_notas_componentes(componente_id, tipo_nota_id, author_id, criado_em, status_id) " +
					" SELECT $1, $2, $3, $4, $5 WHERE NOT EXISTS (SELECT id FROM tipos_notas_componentes WHERE componente_id = $6 " +
					" AND tipo_nota_id = $7 ) RETURNING id"
				db.QueryRow(stmt, idComponente, idTipoNota, autor, criadoEm, statusZero, idComponente, idTipoNota).Scan(&idTipoNotaComponente)
				log.Println("idTipoNotaComponente: " + strconv.Itoa(idTipoNotaComponente))

				qtdElementos := len(cicloESI.Pilares[j].Componentes[k].TiposNotas[l].Elementos)
				for m := 0; m < qtdElementos; m++ {
					nome = cicloESI.Pilares[j].Componentes[k].TiposNotas[l].Elementos[m].Nome
					stmt := " INSERT INTO elementos(nome, descricao, author_id, criado_em, status_id) " +
						" SELECT $1, $2, $3, $4, $5 RETURNING id"
					descricao = "Descricao do " + nome
					db.QueryRow(stmt, nome, descricao, autor, criadoEm, statusZero).Scan(&idElemento)
					if idElemento != 0 {
						elementosSalvos[nome] = idElemento
					} else {
						idElemento = elementosSalvos[nome]
					}
					log.Println("idElemento: " + strconv.Itoa(idElemento) + " - " + nome)
					pesoPadrao = 1
					stmt = " SELECT " + strconv.Itoa(idElemento) + ", " +
						strconv.Itoa(idComponente) + ", " +
						strconv.Itoa(idTipoNota) + ", " +
						strconv.Itoa(pesoPadrao) + ", " +
						strconv.Itoa(autor) + ", " +
						" to_timestamp('" + criadoEm + "','DD-Mon-YYYY HH24:MI:SS') " +
						" WHERE NOT EXISTS ( SELECT id FROM elementos_componentes WHERE elemento_id = " +
						strconv.Itoa(idElemento) + " AND componente_id = " +
						strconv.Itoa(idComponente) + " AND tipo_nota_id = " + strconv.Itoa(idTipoNota) + " ) "
					//log.Println(stmt)
					unsavedElementosComponentes = append(unsavedElementosComponentes, stmt)
					qtdItens := len(cicloESI.Pilares[j].Componentes[k].TiposNotas[l].Elementos[m].Itens)
					for n := 0; n < qtdItens; n++ {
						nome = cicloESI.Pilares[j].Componentes[k].TiposNotas[l].Elementos[m].Itens[n]
						stmt := " INSERT INTO itens(elemento_id, nome, descricao, author_id, criado_em, status_id) " +
							" SELECT $1, $2, $3, $4, $5, $6 WHERE NOT EXISTS (SELECT id FROM itens WHERE nome = '" + nome + "' and elemento_id = " + strconv.Itoa(idElemento) + " ) RETURNING id"
						descricao = "Descricao do " + nome
						db.QueryRow(stmt, idElemento, nome, descricao, autor, criadoEm, statusZero).Scan(&idItem)
						log.Println("idItem: " + strconv.Itoa(idItem) + " - " + nome)
					}
				}
			}
		}
		BulkInsert(unsavedElementosComponentes, stmtElementosComponentes)
		BulkInsert(unsavedComponentesPilares, stmtComponentesPilares)
		BulkInsert(unsavedPilaresCiclos, stmtPilaresCiclos)
	}
	stmt = " WITH R2 AS " +
		"   (SELECT componente_id, " +
		"           tipo_nota_id, " +
		"           round(round(cast(peso_padrao AS numeric(10, 5))/cast(total AS numeric(10, 5)), 4)*cast(100 AS numeric), 2) AS peso_padrao_percentual " +
		"    FROM " +
		"      (WITH TMP AS " +
		"         (SELECT componente_id, " +
		"                 SUM(peso_padrao) AS total " +
		"          FROM elementos_componentes a " +
		"          WHERE componente_id = a.componente_id " +
		"          GROUP BY componente_id) SELECT a.componente_id, " +
		"                                         a.tipo_nota_id, " +
		"                                         tmp.total AS total, " +
		"                                         sum(a.peso_padrao) AS peso_padrao " +
		"       FROM elementos_componentes a " +
		"       LEFT JOIN TMP ON a.componente_id = TMP.componente_id " +
		"       GROUP BY a.componente_id, " +
		"                a.tipo_nota_id, " +
		"                tmp.total) R1 " +
		"    ORDER BY 1, " +
		"             2) " +
		" UPDATE tipos_notas_componentes " +
		" SET peso_padrao = R2.peso_padrao_percentual " +
		" FROM R2 " +
		" WHERE tipos_notas_componentes.componente_id = R2.componente_id " +
		"   AND tipos_notas_componentes.tipo_nota_id = R2.tipo_nota_id "
	log.Println("stmt UPDATE tipos notas componentes: " + stmt)
	updtForm, err := db.Prepare(stmt)
	if err != nil {
		log.Println(err.Error())
	}
	_, err = updtForm.Exec()
	if err != nil {
		log.Println(err.Error())
	}
}

func BulkInsert(unsaved []string, pStmt string) {
	stmt := fmt.Sprintf(pStmt, strings.Join(unsaved, " UNION "))
	log.Println(stmt)
	db.Exec(stmt)
}

func montarCiclo() CicloESI {
	var cicloESI CicloESI
	var pilarRC PilarESI
	var pilarG PilarESI
	var pilarEF PilarESI
	var pilares []PilarESI
	var componentes []ComponenteESI

	pilarRC.Nome = "Riscos e Controles"
	componenteRiscoDeCredito := initRiscoDeCredito()
	componenteRiscoDeMercado := initRiscoDeMercado()
	componenteRiscoDeLiquidez := initRiscoDeLiquidez()
	componenteRiscoAtuarial := initRiscoAtuarial()
	componentes = append(componentes, componenteRiscoDeCredito, componenteRiscoDeMercado, componenteRiscoDeLiquidez, componenteRiscoAtuarial)
	pilarRC.Componentes = componentes

	pilarG.Nome = "Governança"
	componente := initGovernanca()
	componentes = make([]ComponenteESI, 0)
	componentes = append(componentes, componente)
	pilarG.Componentes = componentes

	pilarEF.Nome = "Econômico-Financeiro"
	componenteSolvencia := initEconomicoFinanceiro("Solvência")
	componenteInvestimentosAtivos := initEconomicoFinanceiro("Investimentos/Ativos")
	componenteAtuarial := initEconomicoFinanceiro("Atuarial")
	componenteResultados := initEconomicoFinanceiro("Resultados")
	componenteEficienciaOperacional := initEconomicoFinanceiro("Eficiência Operacional")
	componentes = make([]ComponenteESI, 0)
	componentes = append(componentes, componenteSolvencia, componenteInvestimentosAtivos, componenteAtuarial, componenteResultados, componenteEficienciaOperacional)
	pilarEF.Componentes = componentes
	pilares = append(pilares, pilarRC, pilarG, pilarEF)
	cicloESI.Nome = "Ciclo ESI"
	cicloESI.Pilares = pilares
	return cicloESI
}
