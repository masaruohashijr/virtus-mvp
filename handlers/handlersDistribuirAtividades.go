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

func ListDistribuirAtividadesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Distribuir Atividades Handler")
	currentUser := GetUserInCookie(w, r)
	if sec.IsAuthenticated(w, r) && HasPermission(currentUser, "distribuirAtividades") {
		log.Println("--------------")
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
		for fieldName, value := range r.Form {
			log.Println("-------------- fieldName: " + fieldName)
			if strings.HasPrefix(fieldName, "AuditorComponente") {
				supervisorId := r.FormValue("SupervisorComponenteId")
				log.Println(supervisorId)
				entidadeId := r.FormValue("Entidade_" + fieldName)
				log.Println(entidadeId)
				cicloId := r.FormValue("Ciclo_" + fieldName)
				log.Println(cicloId)
				pilarId := r.FormValue("Pilar_" + fieldName)
				log.Println(pilarId)
				planosIds := r.FormValue("Planos_" + fieldName)
				log.Println("planosIds: " + planosIds)
				componenteId := r.FormValue("Componente_" + fieldName)
				log.Println(fieldName + " - value: " + value[0])
				if value[0] != "" {
					sqlStatement := "UPDATE produtos_componentes SET " +
						" auditor_id=" + value[0] + ", supervisor_id=" + supervisorId +
						" WHERE entidade_id=" + entidadeId +
						" AND ciclo_id=" + cicloId +
						" AND pilar_id=" + pilarId +
						" AND componente_id= " + componenteId
					log.Println(sqlStatement)
					updtForm, _ := Db.Prepare(sqlStatement)
					_, err := updtForm.Exec()
					if err != nil {
						panic(err.Error())
					}
				}
				planos := strings.ReplaceAll(planosIds, "_", ",")
				if len(planos) >= 1 {
					planos = planos[0 : len(planos)-1]
					var produto mdl.ProdutoPlano
					produto.EntidadeId, _ = strconv.ParseInt(entidadeId, 10, 64)
					produto.CicloId, _ = strconv.ParseInt(cicloId, 10, 64)
					produto.PilarId, _ = strconv.ParseInt(pilarId, 10, 64)
					produto.ComponenteId, _ = strconv.ParseInt(componenteId, 10, 64)
					registrarProdutosPlanos(produto, planos, GetUserInCookie(w, r))
				} else {
					http.Redirect(w, r, "/listDistribuirAtividades"+"?errMsg=Faltou configurar quais os planos que serão avaliados.", 301)
				}
			}
		}
		http.Redirect(w, r, "/listDistribuirAtividades", 301)
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
			" coalesce(h.supervisor_id,0) as super_id, coalesce(f.name,'') as supervisor_nome, " +
			" coalesce(a.auditor_id,0) as audit_id, coalesce(g.name,'') as auditor_nome  " +
			" FROM produtos_componentes a " +
			" LEFT JOIN entidades b ON a.entidade_id = b.id " +
			" LEFT JOIN ciclos c ON a.ciclo_id = c.id " +
			" LEFT JOIN pilares d ON a.pilar_id = d.id " +
			" LEFT JOIN componentes e ON a.componente_id = e.id " +
			" LEFT JOIN ciclos_entidades h ON (a.entidade_id = h.entidade_id AND a.ciclo_id = h.ciclo_id) " +
			" LEFT JOIN users f ON h.supervisor_id = f.id " +
			" LEFT JOIN users g ON a.auditor_id = g.id " +
			" WHERE a.entidade_id = " + entidadeId + " AND a.ciclo_id = " + cicloId +
			" ORDER BY d.nome, e.nome "
		log.Println(sql)
		rows, _ := Db.Query(sql)
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
				&produto.SupervisorId,
				&produto.SupervisorName,
				&produto.AuditorId,
				&produto.AuditorName)
			produto.Order = i
			i++
			// log.Println(produto)
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
		var auditores []mdl.User
		var auditor mdl.User
		for rows.Next() {
			rows.Scan(&auditor.Id, &auditor.Name, &orderable)
			auditores = append(auditores, auditor)
		}

		sql = " SELECT id, cnpb, modalidade_id, recurso_garantidor::NUMERIC::MONEY " +
			" FROM planos WHERE entidade_id = $1 ORDER BY recurso_garantidor DESC "
		log.Println(sql)
		rows, _ = Db.Query(sql, entidadeId)
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
