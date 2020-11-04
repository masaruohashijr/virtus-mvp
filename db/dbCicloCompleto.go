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

func createCicloCompleto(qtdCiclos int, qtdPilares int, qtdComponentes int, qtdElementos int, qtdItens int) {
	//createTiposNotas()
	// Ciclos - código: 1
	autor := 1
	tipoMedia := 1
	criadoEm := time.Now().Format("02-Jan-2006 15:04:05")
	log.Println(criadoEm)
	statusZero := 0
	idCiclo := 0
	idPilar := 0
	idItem := 0
	pesoPadrao := 2
	tipoNotaId := 1

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
		" public.componentes_pilares( " +
		" pilar_id, " +
		" componente_id, " +
		" tipo_media, " +
		" peso_padrao, " +
		" author_id, " +
		" criado_em ) %s"

	for i := 1; i <= qtdCiclos; i++ {
		var unsavedPilaresCiclos []string
		var unsavedComponentesPilares []string
		var unsavedElementosComponentes []string
		nome := "Ciclo " + strconv.Itoa(i)
		log.Println(nome)
		descricao := "Descricao do " + nome
		stmt := " INSERT INTO ciclos(nome, descricao, author_id, criado_em, status_id) " +
			" SELECT $1, $2, $3, $4, $5 WHERE NOT EXISTS (SELECT id FROM ciclos WHERE nome = '" + nome + "' ) RETURNING id"
		db.QueryRow(stmt, nome, descricao, autor, criadoEm, statusZero).Scan(&idCiclo)
		// log.Println("idCiclo: " + strconv.Itoa(idCiclo))
		// Pilares - código: 11
		pesoPadrao = 100
		max := 100
		for j := 1; j <= qtdPilares; j++ {
			nome = "Pilar " + strconv.Itoa(i) + strconv.Itoa(j)
			stmt := " INSERT INTO pilares(nome, descricao, author_id, criado_em, status_id) " +
				" SELECT $1, $2, $3, $4, $5 WHERE NOT EXISTS (SELECT id FROM pilares WHERE nome = '" + nome + "' ) RETURNING id"
			//log.Println(nome)
			descricao = "Descricao do " + nome
			db.QueryRow(stmt, nome, descricao, autor, criadoEm, statusZero).Scan(&idPilar)
			//log.Println(pesoPadrao)
			pesoPadrao = rand.Intn(max)
			if j < qtdPilares {
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
			// Componentes - código: 111
			for k := 1; k <= qtdComponentes; k++ {
				nome = "Componente " + strconv.Itoa(i) +
					strconv.Itoa(j) +
					strconv.Itoa(k)
					//log.Println(nome)
				idComponente := 0
				idElemento := 0
				stmt := " INSERT INTO componentes(nome, descricao, author_id, criado_em, status_id) " +
					" SELECT $1, $2, $3, $4, $5 WHERE NOT EXISTS (SELECT id FROM componentes WHERE nome = '" + nome + "' ) RETURNING id"
				descricao = "Descricao do " + nome
				db.QueryRow(stmt, nome, descricao, autor, criadoEm, statusZero).Scan(&idComponente)
				pesoPadrao = int(math.Pow(2, float64(k-1)))
				stmt = " SELECT " + strconv.Itoa(idPilar) + ", " +
					strconv.Itoa(idComponente) + ", " +
					strconv.Itoa(tipoMedia) + ", " +
					strconv.Itoa(pesoPadrao) + ", " +
					strconv.Itoa(autor) + ", " +
					" to_timestamp('" + criadoEm + "','DD-Mon-YYYY HH24:MI:SS') " +
					" WHERE NOT EXISTS ( SELECT id FROM componentes_pilares WHERE componente_id = " + strconv.Itoa(idComponente) + " AND pilar_id = " + strconv.Itoa(idPilar) + " ) "
				unsavedComponentesPilares = append(unsavedComponentesPilares, stmt)
				// Componentes - código: 111
				for l := 1; l <= qtdComponentes; l++ {
					nome = "Elemento " + strconv.Itoa(i) +
						strconv.Itoa(j) + strconv.Itoa(k) +
						strconv.Itoa(l)
					//log.Println(nome)
					stmt := " INSERT INTO elementos(nome, descricao, author_id, criado_em, status_id) " +
						" SELECT $1, $2, $3, $4, $5 WHERE NOT EXISTS (SELECT id FROM elementos WHERE nome = '" + nome + "' ) RETURNING id"
					descricao = "Descricao do " + nome
					db.QueryRow(stmt, nome, descricao, autor, criadoEm, statusZero).Scan(&idElemento)
					pesoPadrao = int(math.Pow(2, float64(l-1)))
					stmt = " SELECT " + strconv.Itoa(idElemento) + ", " +
						strconv.Itoa(idComponente) + ", " +
						strconv.Itoa(tipoNotaId) + ", " +
						strconv.Itoa(pesoPadrao) + ", " +
						strconv.Itoa(autor) + ", " +
						" to_timestamp('" + criadoEm + "','DD-Mon-YYYY HH24:MI:SS') " +
						" WHERE NOT EXISTS ( SELECT id FROM elementos_componentes WHERE elemento_id = " + strconv.Itoa(idElemento) + " AND componente_id = " + strconv.Itoa(idComponente) + " ) "
					//	log.Println(stmt)
					unsavedElementosComponentes = append(unsavedElementosComponentes, stmt)
					if tipoNotaId == 1 {
						tipoNotaId = 2
					} else {
						tipoNotaId = 1
					}
					// Componentes - código: 111
					for m := 1; m <= qtdComponentes; m++ {
						nome = "Item " + strconv.Itoa(i) +
							strconv.Itoa(j) + strconv.Itoa(k) +
							strconv.Itoa(l) + strconv.Itoa(m)
						//log.Println(nome)
						stmt := " INSERT INTO itens(elemento_id, nome, descricao, author_id, criado_em, status_id) " +
							" SELECT $1, $2, $3, $4, $5, $6 WHERE NOT EXISTS (SELECT id FROM itens WHERE nome = '" + nome + "' ) RETURNING id"
						descricao = "Descricao do " + nome
						db.QueryRow(stmt, idElemento, nome, descricao, autor, criadoEm, statusZero).Scan(&idItem)
					}
				}
			}
		}
		BulkInsert(unsavedElementosComponentes, stmtElementosComponentes)
		BulkInsert(unsavedComponentesPilares, stmtComponentesPilares)
		BulkInsert(unsavedPilaresCiclos, stmtPilaresCiclos)
	}
}

func BulkInsert(unsaved []string, pStmt string) {
	stmt := fmt.Sprintf(pStmt, strings.Join(unsaved, " UNION "))
	log.Println(stmt)
	db.Exec(stmt)
}
