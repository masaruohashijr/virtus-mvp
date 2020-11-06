package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	mdl "virtus/models"
	sec "virtus/security"
)

func ListDistribuirPapeisHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Distribuir Papeis Handler")
	if sec.IsAuthenticated(w, r) {
		log.Println("--------------")
		currentUser := GetUserInCookie(w, r)
		var page mdl.PageEntidadesCiclos
		sql := "SELECT b.entidade_id, c.nome " +
			"FROM escritorios a " +
			"LEFT JOIN jurisdicoes b ON a.id = b.escritorio_id " +
			"LEFT JOIN entidades c ON c.id = b.entidade_id " +
			"WHERE a.chefe_id = $1"
		log.Println(sql)
		rows, _ := Db.Query(sql, currentUser.Id)
		var entidades []mdl.Entidade
		var entidade mdl.Entidade
		var i = 1
		for rows.Next() {
			rows.Scan(
				&entidade.Id,
				&entidade.Nome)
			entidade.Order = i
			i++
			sql = "SELECT b.id, b.nome " +
				" FROM ciclos_entidades a " +
				" LEFT JOIN ciclos b ON a.ciclo_id = b.id " +
				" WHERE a.entidade_id = $1 " +
				" ORDER BY id asc"
			rows, _ = Db.Query(sql, entidade.Id)
			var ciclosEntidade []mdl.CicloEntidade
			var cicloEntidade mdl.CicloEntidade
			i = 1
			for rows.Next() {
				rows.Scan(&cicloEntidade.Id, &cicloEntidade.Nome)
				cicloEntidade.Order = i
				i++
				ciclosEntidade = append(ciclosEntidade, cicloEntidade)
			}
			entidade.CiclosEntidade = ciclosEntidade
			entidades = append(entidades, entidade)
		}
		page.Entidades = entidades
		page.AppName = mdl.AppName
		page.Title = "Distribuir Papéis"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/distribuirpapeis/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Entidades-Distribuir-Papeis", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListAvaliarPapeisHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Avaliar Papeis Handler")
	if sec.IsAuthenticated(w, r) {
		log.Println("--------------")
		currentUser := GetUserInCookie(w, r)
		var page mdl.PageEntidadesCiclos
		// Entidades da jurisdição do Escritório ao qual pertenço
		sql := "SELECT a.entidade_id, b.nome FROM jurisdicoes a " +
			" LEFT JOIN entidades b ON a.entidade_id = b.id " +
			" LEFT JOIN membros c ON a.escritorio_id = c.escritorio_id " +
			" WHERE c.usuario_id = $1"
		log.Println(sql)
		rows, _ := Db.Query(sql, currentUser.Id)
		var entidades []mdl.Entidade
		var entidade mdl.Entidade
		var i = 1
		for rows.Next() {
			rows.Scan(
				&entidade.Id,
				&entidade.Nome)
			entidade.Order = i
			i++
			sql = "SELECT b.id, b.nome " +
				" FROM ciclos_entidades a " +
				" LEFT JOIN ciclos b ON a.ciclo_id = b.id " +
				" WHERE a.entidade_id = $1 " +
				" ORDER BY id asc"
			rows, _ = Db.Query(sql, entidade.Id)
			var ciclosEntidade []mdl.CicloEntidade
			var cicloEntidade mdl.CicloEntidade
			i = 1
			for rows.Next() {
				rows.Scan(&cicloEntidade.Id, &cicloEntidade.Nome)
				cicloEntidade.Order = i
				i++
				ciclosEntidade = append(ciclosEntidade, cicloEntidade)
			}
			entidade.CiclosEntidade = ciclosEntidade
			entidades = append(entidades, entidade)
		}
		page.Entidades = entidades
		page.AppName = mdl.AppName
		page.Title = "Avaliar Papéis"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/avaliarpapeis/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Entidades-Avaliar-Papeis", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateDistribuirPapeisHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Distribuir Papeis Handler")
	log.Println("--------------")
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		r.ParseForm()
		for key, value := range r.Form {
			if strings.HasPrefix(key, "SupervisorComponente") {
				log.Println(key + "- value: " + value[0])
				if value[0] != "" {
					sqlStatement := "UPDATE produtos_componentes SET " +
						"supervisor_id=$1 WHERE id=$2"
					log.Println(sqlStatement)
					updtForm, _ := Db.Prepare(sqlStatement)
					_, err := updtForm.Exec(value[0], key[20:len(key)])
					if err != nil {
						panic(err.Error())
					}
					log.Println("Statement: " + sqlStatement)
				}
			}
			if strings.HasPrefix(key, "AuditorComponente") {
				log.Println(key + "- value: " + value[0])
				if value[0] != "" {
					sqlStatement := "UPDATE produtos_componentes SET " +
						"auditor_id=$1 WHERE id=$2"
					log.Println(sqlStatement)
					updtForm, _ := Db.Prepare(sqlStatement)
					_, err := updtForm.Exec(value[0], key[17:len(key)])
					if err != nil {
						panic(err.Error())
					}
					log.Println("Statement: " + sqlStatement)
				}
			}
		}
		http.Redirect(w, r, "/listDistribuirPapeis", 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DistribuirPapeisHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Distribuir Papéis Handler")
	if sec.IsAuthenticated(w, r) {
		currentUser := GetUserInCookie(w, r)
		entidadeId := r.FormValue("EntidadeId")
		cicloId := r.FormValue("CicloId")
		var page mdl.PageProdutosComponentes
		sql := " SELECT " +
			" a.ciclo_id, c.nome as ciclo_nome, " +
			" a.pilar_id, d.nome as pilar_nome, " +
			" a.componente_id, e.nome as componente_nome, " +
			" coalesce(b.nome,''), a.entidade_id, " +
			" coalesce(a.supervisor_id,0) as super_id, coalesce(f.name,'') as supervisor_nome, " +
			" coalesce(a.auditor_id,0) as audit_id, coalesce(g.name,'') as auditor_nome  " +
			" FROM produtos_componentes a " +
			" LEFT JOIN entidades b ON a.entidade_id = b.id " +
			" LEFT JOIN ciclos c ON a.ciclo_id = c.id " +
			" LEFT JOIN pilares d ON a.pilar_id = d.id " +
			" LEFT JOIN componentes e ON a.componente_id = e.id " +
			" LEFT JOIN users f ON a.supervisor_id = f.id " +
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
			//			log.Println(produto)
			produtos = append(produtos, produto)
		}
		page.Produtos = produtos

		sql = " SELECT " +
			" b.usuario_id, coalesce(c.name,'') " +
			" FROM escritorios a " +
			" LEFT JOIN membros b ON a.id = b.escritorio_id " +
			" LEFT JOIN users c ON b.usuario_id = c.id " +
			" WHERE a.chefe_id = $1 AND c.role_id = 3 "
		log.Println(sql)
		rows, _ = Db.Query(sql, currentUser.Id)
		var supervisores []mdl.User
		var supervisor mdl.User
		i = 1
		for rows.Next() {
			rows.Scan(&supervisor.Id, &supervisor.Name)
			supervisores = append(supervisores, supervisor)
		}
		page.Supervisores = supervisores

		sql = " SELECT " +
			" b.usuario_id, coalesce(c.name,'') " +
			" FROM escritorios a " +
			" LEFT JOIN membros b ON a.id = b.escritorio_id " +
			" LEFT JOIN users c ON b.usuario_id = c.id " +
			" WHERE a.chefe_id = $1 AND c.role_id = 4 "
		log.Println(sql)
		rows, _ = Db.Query(sql, currentUser.Id)
		var auditores []mdl.User
		var auditor mdl.User
		i = 1
		for rows.Next() {
			rows.Scan(&auditor.Id, &auditor.Name)
			auditores = append(auditores, auditor)
		}
		page.Supervisores = supervisores
		page.Auditores = auditores
		page.AppName = mdl.AppName
		page.Title = "Distribuir Papéis"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/distribuirpapeis/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Distribuir-Papeis", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func AvaliarPapeisHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Avaliar Papéis Handler")
	if sec.IsAuthenticated(w, r) {
		currentUser := GetUserInCookie(w, r)
		entidadeId := r.FormValue("EntidadeId")
		cicloId := r.FormValue("CicloId")
		var page mdl.PageProdutosItens
		sql := " SELECT  " +
			" a.ciclo_id, c.nome as ciclo_nome,  " +
			" a.pilar_id, d.nome as pilar_nome,  " +
			" a.componente_id, e.nome as componente_nome,  " +
			" a.elemento_id, f.nome as elemento_nome,  " +
			" a.item_id, g.nome as item_nome,  " +
			" coalesce(b.nome,''), a.entidade_id,  " +
			" coalesce(j.supervisor_id,0) as super_id, coalesce(h.name,'') as supervisor_nome,  " +
			" coalesce(j.auditor_id,0) as audit_id, coalesce(i.name,'') as auditor_nome   " +
			" FROM produtos_itens a  " +
			" LEFT JOIN entidades b ON a.entidade_id = b.id  " +
			" LEFT JOIN ciclos c ON a.ciclo_id = c.id  " +
			" LEFT JOIN pilares d ON a.pilar_id = d.id  " +
			" LEFT JOIN componentes e ON a.componente_id = e.id  " +
			" LEFT JOIN produtos_componentes j ON a.componente_id = j.id  " +
			" LEFT JOIN elementos f ON a.elemento_id = f.id  " +
			" LEFT JOIN itens g ON a.item_id = g.id  " +
			" LEFT JOIN users h ON j.supervisor_id = h.id  " +
			" LEFT JOIN users i ON j.auditor_id = i.id  " +
			" WHERE a.entidade_id = " + entidadeId + " AND a.ciclo_id = " + cicloId +
			" ORDER BY d.nome, e.nome "
		log.Println(sql)
		rows, _ := Db.Query(sql)
		var produtos []mdl.ProdutoItem
		var produto mdl.ProdutoItem
		var i = 1
		for rows.Next() {
			rows.Scan(
				&produto.CicloId,
				&produto.CicloNome,
				&produto.PilarId,
				&produto.PilarNome,
				&produto.ComponenteId,
				&produto.ComponenteNome,
				&produto.ElementoId,
				&produto.ElementoNome,
				&produto.ItemId,
				&produto.ItemNome,
				&produto.EntidadeNome,
				&produto.EntidadeId,
				&produto.SupervisorId,
				&produto.SupervisorName,
				&produto.AuditorId,
				&produto.AuditorName)
			produto.Order = i
			i++
			//			log.Println(produto)
			produtos = append(produtos, produto)
		}
		page.Produtos = produtos

		sql = " SELECT " +
			" b.usuario_id, coalesce(c.name,'') " +
			" FROM escritorios a " +
			" LEFT JOIN membros b ON a.id = b.escritorio_id " +
			" LEFT JOIN users c ON b.usuario_id = c.id " +
			" WHERE a.chefe_id = $1 AND c.role_id = 3 "
		log.Println(sql)
		rows, _ = Db.Query(sql, currentUser.Id)
		var supervisores []mdl.User
		var supervisor mdl.User
		i = 1
		for rows.Next() {
			rows.Scan(&supervisor.Id, &supervisor.Name)
			supervisores = append(supervisores, supervisor)
		}
		page.Supervisores = supervisores

		sql = " SELECT " +
			" b.usuario_id, coalesce(c.name,'') " +
			" FROM escritorios a " +
			" LEFT JOIN membros b ON a.id = b.escritorio_id " +
			" LEFT JOIN users c ON b.usuario_id = c.id " +
			" WHERE a.chefe_id = $1 AND c.role_id = 4 "
		log.Println(sql)
		rows, _ = Db.Query(sql, currentUser.Id)
		var auditores []mdl.User
		var auditor mdl.User
		i = 1
		for rows.Next() {
			rows.Scan(&auditor.Id, &auditor.Name)
			auditores = append(auditores, auditor)
		}
		page.Supervisores = supervisores
		page.Auditores = auditores
		page.AppName = mdl.AppName
		page.Title = "Avaliar Papéis"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/avaliarpapeis/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Avaliar-Papeis", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
