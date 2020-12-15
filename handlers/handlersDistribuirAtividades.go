package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	mdl "virtus/models"
	sec "virtus/security"
)

type PlanosCfg struct {
	numPlano   string
	cnpb       string
	podeApagar bool
}

func ListDistribuirAtividadesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Distribuir Atividades Handler")
	currentUser := GetUserInCookie(w, r)
	if sec.IsAuthenticated(w, r) && HasPermission(currentUser, "distribuirAtividades") {
		log.Println("--------------")
		msg := r.FormValue("msg")
		log.Println("msg: " + msg)
		errMsg := r.FormValue("errMsg")
		var page mdl.PageEntidadesCiclos
		sql := "SELECT DISTINCT d.codigo, b.entidade_id, d.nome, a.abreviatura " +
			" FROM escritorios a " +
			" LEFT JOIN jurisdicoes b ON a.id = b.escritorio_id " +
			" LEFT JOIN membros c ON c.escritorio_id = b.escritorio_id " +
			" LEFT JOIN entidades d ON d.id = b.entidade_id " +
			" LEFT JOIN users u ON u.id = c.usuario_id " +
			" INNER JOIN ciclos_entidades e ON e.entidade_id = b.entidade_id " +
			" WHERE (c.usuario_id = $1 AND u.role_id in (3,4)) OR (a.chefe_id = $2)"
		log.Println(sql)
		rows, _ := Db.Query(sql, currentUser.Id, currentUser.Id)
		defer rows.Close()
		var entidades []mdl.Entidade
		var entidade mdl.Entidade
		var i = 1
		for rows.Next() {
			rows.Scan(
				&entidade.Codigo,
				&entidade.Id,
				&entidade.Nome,
				&entidade.Escritorio)
			entidade.Order = i
			i++
			entidades = append(entidades, entidade)
		}
		var entidadesCiclos []mdl.Entidade
		for i, entidade := range entidades {
			var ciclosEntidade []mdl.CicloEntidade
			var cicloEntidade mdl.CicloEntidade
			sql = "SELECT b.id, b.nome " +
				" FROM ciclos_entidades a " +
				" LEFT JOIN ciclos b ON a.ciclo_id = b.id " +
				" WHERE a.entidade_id = $1 " +
				" ORDER BY id asc"
			rows, _ = Db.Query(sql, entidade.Id)
			defer rows.Close()
			i = 1
			for rows.Next() {
				rows.Scan(&cicloEntidade.Id, &cicloEntidade.Nome)
				cicloEntidade.Order = i
				i++
				ciclosEntidade = append(ciclosEntidade, cicloEntidade)
			}
			entidade.CiclosEntidade = ciclosEntidade
			entidadesCiclos = append(entidadesCiclos, entidade)
		}
		if errMsg != "" {
			page.ErrMsg = errMsg
		}
		if msg != "" {
			page.Msg = msg
		}
		page.Entidades = entidadesCiclos
		page.AppName = mdl.AppName
		page.Title = "Distribuir Atividades"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/distribuiratividades/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Entidades-Distribuir-Atividades", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateDistribuirAtividadesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Distribuir Atividades Handler")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		r.ParseForm()
		faltouConfigurarPlano := false
		for fieldName, value := range r.Form {
			// log.Println("-------------- fieldName: " + fieldName)
			if strings.HasPrefix(fieldName, "AuditorComponente_") {
				fname := fieldName[7:len(fieldName)]
				//log.Println(fname)
				supervisorId := r.FormValue("SupervisorComponenteId")
				//log.Println(supervisorId)
				entidadeId := r.FormValue("Entidade_" + fname)
				//log.Println(entidadeId)
				cicloId := r.FormValue("Ciclo_" + fname)
				//log.Println(cicloId)
				pilarId := r.FormValue("Pilar_" + fname)
				//log.Println(pilarId)
				planosIds := r.FormValue("Planos_" + fname)
				//log.Println("planosIds: " + planosIds)
				componenteId := r.FormValue("Componente_" + fname)
				//log.Println(fieldName + " - value: " + value[0])
				if value[0] != "" {
					sqlStatement := "UPDATE produtos_componentes SET " +
						" auditor_id=" + value[0] + ", supervisor_id=" + supervisorId +
						" WHERE entidade_id=" + entidadeId +
						" AND ciclo_id=" + cicloId +
						" AND pilar_id=" + pilarId +
						" AND componente_id= " + componenteId
					//log.Println(sqlStatement)
					updtForm, _ := Db.Prepare(sqlStatement)
					_, err := updtForm.Exec()
					if err != nil {
						log.Println(err.Error())
					}
				}
				planos := strings.ReplaceAll(planosIds, "_", ",")
				if len(planos) == 0 {
					faltouConfigurarPlano = true
				}
			} else if strings.HasPrefix(fieldName, "IniciaEmComponente_") {
				fname := fieldName[8:len(fieldName)]
				// log.Println(fname)
				partes := strings.Split(fname, "_")
				entidadeId := partes[1]
				// log.Println(entidadeId)
				cicloId := partes[2]
				// log.Println(cicloId)
				pilarId := partes[3]
				// log.Println(pilarId)
				componenteId := partes[4]
				// log.Println(fieldName + " - value: " + value[0])
				if value[0] != "" {
					sqlStatement := "UPDATE produtos_componentes SET " +
						" inicia_em='" + value[0] + "' " +
						" WHERE entidade_id=" + entidadeId +
						" AND ciclo_id=" + cicloId +
						" AND pilar_id=" + pilarId +
						" AND componente_id= " + componenteId
					log.Println(sqlStatement)
					updtForm, _ := Db.Prepare(sqlStatement)
					_, err := updtForm.Exec()
					if err != nil {
						log.Println(err.Error())
					}
				}
			} else if strings.HasPrefix(fieldName, "TerminaEmComponente_") {
				fname := fieldName[9:len(fieldName)]
				// log.Println(fname)
				partes := strings.Split(fname, "_")
				entidadeId := partes[1]
				// log.Println(entidadeId)
				cicloId := partes[2]
				// log.Println(cicloId)
				pilarId := partes[3]
				// log.Println(pilarId)
				componenteId := partes[4]
				// log.Println(fieldName + " - value: " + value[0])
				if value[0] != "" {
					sqlStatement := "UPDATE produtos_componentes SET " +
						" termina_em='" + value[0] + "' " +
						" WHERE entidade_id=" + entidadeId +
						" AND ciclo_id=" + cicloId +
						" AND pilar_id=" + pilarId +
						" AND componente_id= " + componenteId
					log.Println(sqlStatement)
					updtForm, _ := Db.Prepare(sqlStatement)
					_, err := updtForm.Exec()
					if err != nil {
						log.Println(err.Error())
					}
				}
			}
		}
		if faltouConfigurarPlano {
			http.Redirect(w, r, "/listDistribuirAtividades"+"?errMsg=Faltou configurar quais os planos que serão avaliados.", 301)
		}
		http.Redirect(w, r, "/listDistribuirAtividades"+"?msg=Os demais produtos dos níveis do ciclo foram criados com Sucesso.", 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DistribuirAtividadesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Distribuir Atividades Handler")
	if sec.IsAuthenticated(w, r) {
		entidadeId := r.FormValue("EntidadeId")
		cicloId := r.FormValue("CicloId")
		var page mdl.PageProdutosComponentes
		sql := " SELECT " +
			" a.ciclo_id, c.nome as ciclo_nome, " +
			" a.pilar_id, d.nome as pilar_nome, " +
			" a.componente_id, e.nome as componente_nome, " +
			" coalesce(b.nome,''), a.entidade_id, " +
			" coalesce(R.configurado,'N'), " +
			" coalesce(h.supervisor_id,0) as super_id, coalesce(f.name,'') as supervisor_nome, " +
			" coalesce(a.auditor_id,0) as audit_id, coalesce(g.name,'') as auditor_nome,  " +
			" coalesce(to_char(a.inicia_em,'YYYY-MM-DD'),'') as inicia_em, " +
			" coalesce(to_char(a.termina_em,'YYYY-MM-DD'),'') as termina_em " +
			" FROM produtos_componentes a " +
			" LEFT JOIN entidades b ON a.entidade_id = b.id " +
			" LEFT JOIN ciclos c ON a.ciclo_id = c.id " +
			" LEFT JOIN pilares d ON a.pilar_id = d.id " +
			" LEFT JOIN componentes e ON a.componente_id = e.id " +
			" LEFT JOIN ciclos_entidades h ON (a.entidade_id = h.entidade_id AND a.ciclo_id = h.ciclo_id) " +
			" LEFT JOIN users f ON h.supervisor_id = f.id " +
			" LEFT JOIN users g ON a.auditor_id = g.id " +
			" LEFT JOIN (select a.entidade_id, a.ciclo_id, a.pilar_id, a.componente_id, " +
			" 	CASE WHEN COUNT(i.id)>0 THEN 'S' ELSE 'N' END AS configurado from produtos_componentes a " +
			" 	INNER JOIN produtos_planos i ON (a.entidade_id = i.entidade_id " +
			" 	AND a.ciclo_id = i.ciclo_id " +
			" 	AND a.pilar_id = i.pilar_id " +
			" 	AND a.componente_id = i.componente_id) " +
			" 	GROUP BY 1,2,3,4) R on (a.entidade_id = R.entidade_id " +
			" 	AND a.ciclo_id = R.ciclo_id " +
			" 	AND a.pilar_id = R.pilar_id " +
			" 	AND a.componente_id = R.componente_id) " +
			" WHERE a.entidade_id = " + entidadeId + " AND a.ciclo_id = " + cicloId +
			" ORDER BY d.nome, e.nome "
		log.Println(sql)
		rows, _ := Db.Query(sql)
		defer rows.Close()
		var produtos []mdl.ProdutoComponente
		var produto mdl.ProdutoComponente
		var i = 1
		for rows.Next() {
			rows.Scan(
				&produto.CicloId,
				&produto.CicloNome,
				&produto.PilarId,
				&produto.PilarNome,
				&produto.ComponenteId,
				&produto.ComponenteNome,
				&produto.EntidadeNome,
				&produto.EntidadeId,
				&produto.Configurado,
				&produto.SupervisorId,
				&produto.SupervisorName,
				&produto.AuditorId,
				&produto.AuditorName,
				&produto.IniciaEm,
				&produto.TerminaEm)
			produto.Order = i
			i++
			produtos = append(produtos, produto)
		}
		page.Produtos = produtos
		orderable := ""
		sql = " SELECT " +
			" b.usuario_id, coalesce(c.name,''), UPPER(coalesce(c.name,'')) AS supervisor_nome " +
			" FROM escritorios a " +
			" LEFT JOIN membros b ON a.id = b.escritorio_id " +
			" LEFT JOIN users c ON b.usuario_id = c.id " +
			" WHERE c.role_id = 3 ORDER BY supervisor_nome "
		log.Println(sql)
		rows, _ = Db.Query(sql)
		defer rows.Close()
		var supervisores []mdl.User
		var supervisor mdl.User
		i = 1
		for rows.Next() {
			rows.Scan(&supervisor.Id, &supervisor.Name, &orderable)
			supervisores = append(supervisores, supervisor)
		}
		page.Supervisores = supervisores

		sql = " SELECT " +
			" a.usuario_id, " +
			" coalesce(b.name,''), " +
			" UPPER(coalesce(b.name,'')) AS orderable " +
			" FROM integrantes a " +
			" LEFT JOIN users b " +
			" ON a.usuario_id = b.id " +
			" WHERE " +
			" a.entidade_id = " + entidadeId +
			" AND a.ciclo_id = " + cicloId +
			" AND b.role_id = 4 "
		log.Println(sql)
		rows, _ = Db.Query(sql)
		defer rows.Close()
		var auditores []mdl.User
		var auditor mdl.User
		for rows.Next() {
			rows.Scan(&auditor.Id, &auditor.Name, &orderable)
			auditores = append(auditores, auditor)
		}

		sql = " SELECT id, cnpb, modalidade_id, CASE WHEN recurso_garantidor < 1000000000 THEN recurso_garantidor::numeric::MONEY/1000000||' mi' ELSE recurso_garantidor::numeric::MONEY/1000000000||' bi' END " +
			" FROM planos WHERE entidade_id = $1 ORDER BY recurso_garantidor DESC "
		log.Println(sql)
		rows, _ = Db.Query(sql, entidadeId)
		defer rows.Close()
		var planos []mdl.Plano
		var plano mdl.Plano
		for rows.Next() {
			rows.Scan(&plano.Id, &plano.CNPB, &plano.Modalidade, &plano.RecursoGarantidor)
			planos = append(planos, plano)
		}
		page.Planos = planos
		page.Supervisores = supervisores
		page.Auditores = auditores
		page.AppName = mdl.AppName
		page.Title = "Distribuir Atividades"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/distribuiratividades/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Distribuir-Atividades", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func LoadConfigPlanos(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Config Planos")
	r.ParseForm()
	var entidadeId = r.FormValue("entidadeId")
	var cicloId = r.FormValue("cicloId")
	var pilarId = r.FormValue("pilarId")
	var componenteId = r.FormValue("componenteId")
	configPlanos := ListConfigPlanos(entidadeId, cicloId, pilarId, componenteId)
	jsonConfigPlanos, _ := json.Marshal(configPlanos)
	w.Write([]byte(jsonConfigPlanos))
	log.Println("JSON Config Planos")
}

func UpdateConfigPlanos(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Config Planos ===>>> ATUALIZANDO")
	r.ParseForm()
	currentUser := GetUserInCookie(w, r)
	var entidadeId = r.FormValue("entidadeId")
	var cicloId = r.FormValue("cicloId")
	var pilarId = r.FormValue("pilarId")
	var componenteId = r.FormValue("componenteId")
	var planos = r.FormValue("planos")
	var forcar = r.FormValue("forcar")
	planos = strings.TrimSpace(planos)
	array := strings.Split(planos, "_")
	log.Println("planos: " + planos)
	var planosPage []PlanosCfg
	var planoPage PlanosCfg
	sql := " select id, cnpb from planos where entidade_id = " + entidadeId + " order by cnpb "
	log.Println(sql)
	rows, _ := Db.Query(sql)
	defer rows.Close()
	m := make(map[string]string)
	id := 0
	cnpb := ""
	for rows.Next() {
		rows.Scan(&id, &cnpb)
		m[strconv.Itoa(id)] = cnpb
	}
	for i, valor := range array {
		log.Println("i: " + strconv.Itoa(i))
		if strings.TrimSpace(valor) != "" {
			planoPage.numPlano = valor
			planoPage.cnpb = m[planoPage.numPlano]
			planosPage = append(planosPage, planoPage)
		}
	}
	log.Println("Qtd Page: " + strconv.Itoa(len(planosPage)))

	planos = strings.Join(array, ",")
	if len(planos) > 0 {
		planos = planos[:len(planos)-1]
	}
	sql = " select a.plano_id, c.cnpb, " +
		" case when count(b.id) = 0 then true else false end as pode_apagar " +
		" from produtos_elementos a " +
		" left join produtos_elementos_historicos b on " +
		" (a.entidade_id = b.entidade_id and a.ciclo_id = b.ciclo_id " +
		" and a.pilar_id = b.pilar_id and a.componente_id = b.componente_id " +
		" and a.plano_id = b.plano_id) " +
		" inner join planos c on c.id = a.plano_id " +
		" where a.entidade_id = " + entidadeId +
		" and a.ciclo_id = " + cicloId +
		" and a.pilar_id = " + pilarId +
		" and a.componente_id = " + componenteId +
		//" and a.plano_id in (" + planos + ") " +
		" group by 1,2 "
	log.Println(sql)
	rows, _ = Db.Query(sql)
	defer rows.Close()
	var planosBD []PlanosCfg
	var planoBD PlanosCfg
	for rows.Next() {
		rows.Scan(&planoBD.numPlano, &planoBD.cnpb, &planoBD.podeApagar)
		planosBD = append(planosBD, planoBD)
	}
	// Não posso simplesmente apagar eu tenho que testar NOW () no banco PAST (1,2,3,4)
	msgRetorno := ""
	force := false
	log.Println("Qtd BD: " + strconv.Itoa(len(planosBD)))

	if len(planosPage) < len(planosBD) {
		if len(planosPage) == 0 {
			if forcar != "" {
				force = true
			}
			for _, valor := range planosBD {
				log.Println(valor.cnpb)
				if valor.podeApagar || force {
					log.Println("Removendo o " + valor.cnpb)
					deleteProdutoPlano(entidadeId, cicloId, pilarId, componenteId, valor.numPlano)
					msgRetorno += "O plano " + valor.cnpb + " foi removido com Sucesso.\n"
				} else {
					msgRetorno += "O plano " + valor.cnpb + " não pode ser removido por já ter sido avaliado antes.\n"
				}
			}
		} else {
			var diffDB []PlanosCfg = planosBD
			for n := range planosPage {
				if containsPlanoCfg(diffDB, planosPage[n]) {
					diffDB = removePlanoCfg(diffDB, planosPage[n])
				}
			}
			for _, valor := range diffDB {
				if valor.podeApagar || force {
					log.Println("Removendo o " + valor.cnpb)
					deleteProdutoPlano(entidadeId, cicloId, pilarId, componenteId, valor.numPlano)
					msgRetorno += "O plano " + valor.cnpb + " foi removido com Sucesso.\n"
				} else {
					msgRetorno += "O plano " + valor.cnpb + " não pode ser removido por já ter sido avaliado antes.\n"
					log.Println(msgRetorno)
				}
			}
		}
	} else {
		log.Println("Registrar Produtos Planos")
		var diffPage []PlanosCfg = planosPage
		for n := range planosBD {
			log.Println("CNPB: " + planosBD[n].cnpb)
			if containsPlanoCfg(diffPage, planosBD[n]) {
				log.Println("Removendo " + planosBD[n].cnpb)
				diffPage = removePlanoCfg(diffPage, planosBD[n])
			}
		}
		var param mdl.ProdutoPlano
		param.EntidadeId, _ = strconv.ParseInt(entidadeId, 10, 64)
		param.CicloId, _ = strconv.ParseInt(cicloId, 10, 64)
		param.PilarId, _ = strconv.ParseInt(pilarId, 10, 64)
		param.ComponenteId, _ = strconv.ParseInt(componenteId, 10, 64)
		for _, v := range planosPage {
			log.Println("Registrar diffPage como Produtos Planos")
			log.Println("Plano: " + v.numPlano)
			retorno := registrarProdutosPlanos(param, v.numPlano, currentUser)
			if retorno != 0 {
				msgRetorno += "O plano " + v.cnpb + " foi adicionado com Sucesso.\n"
			}
		}
		var diffDB []PlanosCfg = planosBD
		for n := range planosPage {
			if containsPlanoCfg(diffDB, planosPage[n]) {
				diffDB = removePlanoCfg(diffDB, planosPage[n])
			}
		}
		for _, valor := range diffDB {
			if valor.podeApagar || force {
				log.Println("Removendo o " + valor.cnpb)
				deleteProdutoPlano(entidadeId, cicloId, pilarId, componenteId, valor.numPlano)
				msgRetorno += "O plano " + valor.cnpb + " foi removido com Sucesso.\n"
			} else {
				msgRetorno += "O plano " + valor.cnpb + " não pode ser removido por já ter sido avaliado antes.\n"
				log.Println(msgRetorno)
			}
		}
	}
	w.Write([]byte(msgRetorno))
	log.Println("JSON Config Planos")
}

func deleteProdutoPlano(entidadeId string, cicloId string, pilarId string, componenteId string, planoId string) {
	sqlStatement := "DELETE FROM produtos_itens WHERE " +
		" entidade_id = " + entidadeId +
		" and ciclo_id = " + cicloId +
		" and pilar_id = " + pilarId +
		" and componente_id = " + componenteId +
		" and plano_id = " + planoId
	log.Println(sqlStatement)
	updtForm, _ := Db.Prepare(sqlStatement)
	_, err := updtForm.Exec()
	if err != nil {
		log.Println(err.Error())
	}
	sqlStatement = "DELETE FROM produtos_elementos_historicos WHERE " +
		" entidade_id = " + entidadeId +
		" and ciclo_id = " + cicloId +
		" and pilar_id = " + pilarId +
		" and componente_id = " + componenteId +
		" and plano_id = " + planoId
	log.Println(sqlStatement)
	updtForm, _ = Db.Prepare(sqlStatement)
	_, err = updtForm.Exec()
	if err != nil {
		log.Println(err.Error())
	}
	sqlStatement = "DELETE FROM produtos_elementos WHERE " +
		" entidade_id = " + entidadeId +
		" and ciclo_id = " + cicloId +
		" and pilar_id = " + pilarId +
		" and componente_id = " + componenteId +
		" and plano_id = " + planoId
	log.Println(sqlStatement)
	updtForm, _ = Db.Prepare(sqlStatement)
	_, err = updtForm.Exec()
	if err != nil {
		log.Println(err.Error())
	}
	sqlStatement = "DELETE FROM produtos_tipos_notas WHERE " +
		" entidade_id = " + entidadeId +
		" and ciclo_id = " + cicloId +
		" and pilar_id = " + pilarId +
		" and componente_id = " + componenteId +
		" and plano_id = " + planoId
	log.Println(sqlStatement)
	updtForm, _ = Db.Prepare(sqlStatement)
	_, err = updtForm.Exec()
	if err != nil {
		log.Println(err.Error())
	}
	sqlStatement = "DELETE FROM produtos_planos WHERE " +
		" entidade_id = " + entidadeId +
		" and ciclo_id = " + cicloId +
		" and pilar_id = " + pilarId +
		" and componente_id = " + componenteId +
		" and plano_id = " + planoId
	log.Println(sqlStatement)
	updtForm, _ = Db.Prepare(sqlStatement)
	_, err = updtForm.Exec()
	if err != nil {
		log.Println(err.Error())
	}
}

func removePlanoCfg(planos []PlanosCfg, planoCfgToBeRemoved PlanosCfg) []PlanosCfg {
	var newPlanosCfg []PlanosCfg
	for i := range planos {
		if planos[i].numPlano != planoCfgToBeRemoved.numPlano {
			newPlanosCfg = append(newPlanosCfg, planos[i])
		}
	}
	return newPlanosCfg
}

func containsPlanoCfg(planosCfg []PlanosCfg, planoCfgCompared PlanosCfg) bool {
	for n := range planosCfg {
		log.Println(planosCfg[n].numPlano)
		log.Println(planoCfgCompared.numPlano)
		if planosCfg[n].numPlano == planoCfgCompared.numPlano {
			return true
		}
	}
	return false
}
