package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	mdl "virtus/models"
	sec "virtus/security"
)

const sqlPapeis = " SELECT  " +
	" a.entidade_id, coalesce(b.nome,''),  " +
	" a.ciclo_id, c.nome as ciclo_nome, " +
	" coalesce(l.nota,0) as ciclo_nota, " +
	" a.pilar_id, d.nome as pilar_nome, " +
	" coalesce(k.peso,0) as pilar_peso, coalesce(k.nota,0) as pilar_nota, " +
	" a.componente_id, e.nome as componente_nome, " +
	" coalesce(j.peso,0) as componente_peso, coalesce(j.nota,0) as componente_nota, " +
	" coalesce(j.supervisor_id,0) as super_id, coalesce(h.name,'') as supervisor_nome, " +
	" coalesce(j.auditor_id,0) as audit_id, coalesce(i.name,'') as auditor_nome, " +
	" ec.tipo_nota_id, m.letra, m.cor_letra, m.nome, " +
	" coalesce(o.peso,0) as tipo_nota_peso, coalesce(o.nota,0) as tipo_nota_nota, " +
	" a.elemento_id, f.nome as elemento_nome, " +
	" coalesce(n.peso,0) as elemento_peso, coalesce(n.nota,0) as elemento_nota, " +
	" n.tipo_pontuacao_id, " +
	" ec.peso_padrao, " +
	" cp.tipo_media, cp.peso_padrao, " +
	" pc.tipo_media, pc.peso_padrao, " +
	" ce.tipo_media, ce.inicia_em, ce.termina_em, " +
	" a.item_id, g.nome as item_nome " +
	" FROM produtos_itens a " +
	" INNER JOIN entidades b ON a.entidade_id = b.id " +
	" INNER JOIN ciclos c ON a.ciclo_id = c.id " +
	" INNER JOIN pilares d ON a.pilar_id = d.id " +
	" INNER JOIN componentes e ON a.componente_id = e.id " +
	" INNER JOIN produtos_tipos_notas o ON " +
	" ( a.tipo_nota_id = o.tipo_nota_id AND " +
	" a.componente_id = o.componente_id AND " +
	" a.pilar_id = o.pilar_id AND " +
	" a.ciclo_id = o.ciclo_id AND " +
	" a.entidade_id = o.entidade_id )" +
	" INNER JOIN produtos_elementos n ON " +
	" ( a.elemento_id = n.elemento_id AND " +
	" a.componente_id = n.componente_id AND " +
	" a.pilar_id = n.pilar_id AND " +
	" a.ciclo_id = n.ciclo_id AND " +
	" a.entidade_id = n.entidade_id )" +
	" INNER JOIN produtos_componentes j ON " +
	" ( a.componente_id = j.componente_id AND " +
	" a.pilar_id = j.pilar_id AND " +
	" a.ciclo_id = j.ciclo_id AND " +
	" a.entidade_id = j.entidade_id )" +
	" INNER JOIN produtos_pilares k ON " +
	" ( a.pilar_id = k.pilar_id AND " +
	"   a.ciclo_id = k.ciclo_id AND " +
	"   a.entidade_id = k.entidade_id ) " +
	" INNER JOIN produtos_ciclos l ON " +
	" ( a.ciclo_id = l.ciclo_id AND " +
	"   a.entidade_id = l.entidade_id ) " +
	" INNER JOIN elementos_componentes ec ON " +
	" ( a.elemento_id = ec.elemento_id AND a.componente_id = ec.componente_id ) " +
	" INNER JOIN componentes_pilares cp ON " +
	" ( a.componente_id = cp.componente_id AND a.pilar_id = cp.pilar_id ) " +
	" INNER JOIN pilares_ciclos pc ON " +
	" ( a.pilar_id = pc.pilar_id AND a.ciclo_id = pc.ciclo_id ) " +
	" INNER JOIN ciclos_entidades ce ON (a.ciclo_id = ce.ciclo_id and a.entidade_id = ce.entidade_id)  " +
	" INNER JOIN elementos f ON a.elemento_id = f.id  " +
	" INNER JOIN itens g ON a.item_id = g.id  " +
	" INNER JOIN users h ON j.supervisor_id = h.id  " +
	" INNER JOIN users i ON j.auditor_id = i.id  " +
	" INNER JOIN tipos_notas m ON ec.tipo_nota_id = m.id " +
	" WHERE a.entidade_id = $1 AND a.ciclo_id = $2 " +
	" ORDER BY a.ciclo_id, " +
	" a.pilar_id,  " +
	" a.componente_id, " +
	" a.tipo_nota_id, " +
	" a.elemento_id, " +
	" a.item_id "

func ListAvaliarPapeisHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Avaliar Papeis Handler")
	currentUser := GetUserInCookie(w, r)
	if sec.IsAuthenticated(w, r) && HasPermission(currentUser, "avaliarPapeis") {
		log.Println("--------------")
		var page mdl.PageEntidadesCiclos
		// Entidades da jurisdição do Escritório ao qual pertenço
		sql := "SELECT a.entidade_id, b.nome FROM jurisdicoes a " +
			" INNER JOIN ciclos_entidades d ON d.entidade_id = a.entidade_id " +
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

func AvaliarPapeisHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Avaliar Papéis Handler")
	if sec.IsAuthenticated(w, r) {
		entidadeId := r.FormValue("EntidadeId")
		cicloId := r.FormValue("CicloId")
		var page mdl.PageProdutosItens
		log.Println(sqlPapeis)
		rows, _ := Db.Query(sqlPapeis, entidadeId, cicloId)
		var produtos []mdl.ProdutoItem
		var produto mdl.ProdutoItem
		var i = 1
		for rows.Next() {
			rows.Scan(
				&produto.EntidadeId,
				&produto.EntidadeNome,
				&produto.CicloId,
				&produto.CicloNome,
				&produto.CicloNota,
				&produto.PilarId,
				&produto.PilarNome,
				&produto.PilarPeso,
				&produto.PilarNota,
				&produto.ComponenteId,
				&produto.ComponenteNome,
				&produto.ComponentePeso,
				&produto.ComponenteNota,
				&produto.SupervisorId,
				&produto.SupervisorName,
				&produto.AuditorId,
				&produto.AuditorName,
				&produto.TipoNotaId,
				&produto.TipoNotaLetra,
				&produto.TipoNotaCorLetra,
				&produto.TipoNotaNome,
				&produto.TipoNotaPeso,
				&produto.TipoNotaNota,
				&produto.ElementoId,
				&produto.ElementoNome,
				&produto.ElementoPeso,
				&produto.ElementoNota,
				&produto.TipoPontuacaoId,
				&produto.PesoPadraoEC,
				&produto.TipoMediaCPId,
				&produto.PesoPadraoCP,
				&produto.TipoMediaPCId,
				&produto.PesoPadraoPC,
				&produto.TipoMediaCEId,
				&produto.IniciaEm,
				&produto.TerminaEm,
				&produto.ItemId,
				&produto.ItemNome)
			produto.Order = i
			i++
			//log.Println(produto)
			produtos = append(produtos, produto)
		}
		page.Produtos = produtos

		sql := " SELECT " +
			" a.usuario_id, " +
			" coalesce(b.name,'') " +
			" FROM integrantes a " +
			" LEFT JOIN users b " +
			" ON a.usuario_id = b.id " +
			" WHERE " +
			" a.entidade_id = " + entidadeId +
			" AND a.ciclo_id = " + cicloId +
			" AND b.role_id = 3 "
		log.Println(sql)
		rows, _ = Db.Query(sql)
		var supervisores []mdl.User
		var supervisor mdl.User
		i = 1
		for rows.Next() {
			rows.Scan(&supervisor.Id, &supervisor.Name)
			supervisores = append(supervisores, supervisor)
		}
		page.Supervisores = supervisores

		sql = " SELECT " +
			" a.usuario_id, " +
			" b.name " +
			" FROM integrantes a " +
			" LEFT JOIN users b " +
			" ON a.usuario_id = b.id " +
			" WHERE " +
			" a.entidade_id = " + entidadeId +
			" AND a.ciclo_id = " + cicloId +
			" AND b.role_id in (2,3,4) "
		log.Println(sql)
		rows, _ = Db.Query(sql)
		var auditores []mdl.User
		var auditor mdl.User
		i = 1
		for rows.Next() {
			rows.Scan(&auditor.Id, &auditor.Name)
			log.Println("Auditor competente: " + auditor.Name)
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

func AtualizarPapeisHandler(entidadeId string, cicloId string, w http.ResponseWriter, r *http.Request) {
	log.Println("Atualizar Papéis Handler")
	var page mdl.PageProdutosItens
	log.Println(sqlPapeis)
	rows, _ := Db.Query(sqlPapeis, entidadeId, cicloId)
	var produtos []mdl.ProdutoItem
	var produto mdl.ProdutoItem
	var i = 1
	for rows.Next() {
		rows.Scan(
			&produto.EntidadeId,
			&produto.EntidadeNome,
			&produto.CicloId,
			&produto.CicloNome,
			&produto.CicloNota,
			&produto.PilarId,
			&produto.PilarNome,
			&produto.PilarPeso,
			&produto.PilarNota,
			&produto.ComponenteId,
			&produto.ComponenteNome,
			&produto.ComponentePeso,
			&produto.ComponenteNota,
			&produto.SupervisorId,
			&produto.SupervisorName,
			&produto.AuditorId,
			&produto.AuditorName,
			&produto.TipoNotaId,
			&produto.TipoNotaLetra,
			&produto.TipoNotaCorLetra,
			&produto.TipoNotaNome,
			&produto.TipoNotaPeso,
			&produto.TipoNotaNota,
			&produto.ElementoId,
			&produto.ElementoNome,
			&produto.ElementoPeso,
			&produto.ElementoNota,
			&produto.TipoPontuacaoId,
			&produto.PesoPadraoEC,
			&produto.TipoMediaCPId,
			&produto.PesoPadraoCP,
			&produto.TipoMediaPCId,
			&produto.PesoPadraoPC,
			&produto.TipoMediaCEId,
			&produto.IniciaEm,
			&produto.TerminaEm,
			&produto.ItemId,
			&produto.ItemNome)
		produto.Order = i
		i++
		// log.Println(produto)
		produtos = append(produtos, produto)
	}
	page.Produtos = produtos

	sql := " SELECT " +
		" a.usuario_id, " +
		" coalesce(b.name,'') " +
		" FROM integrantes a " +
		" LEFT JOIN users b " +
		" ON a.usuario_id = b.id " +
		" WHERE " +
		" a.entidade_id = " + entidadeId +
		" AND a.ciclo_id = " + cicloId +
		" AND b.role_id = 3 "
	log.Println(sql)
	rows, _ = Db.Query(sql)
	var supervisores []mdl.User
	var supervisor mdl.User
	i = 1
	for rows.Next() {
		rows.Scan(&supervisor.Id, &supervisor.Name)
		supervisores = append(supervisores, supervisor)
	}
	page.Supervisores = supervisores

	sql = " SELECT " +
		" a.usuario_id, " +
		" b.name " +
		" FROM integrantes a " +
		" LEFT JOIN users b " +
		" ON a.usuario_id = b.id " +
		" WHERE " +
		" a.entidade_id = " + entidadeId +
		" AND a.ciclo_id = " + cicloId +
		" AND b.role_id in (2,3,4) "
	log.Println(sql)
	rows, _ = Db.Query(sql)
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
}

func UpdateAvaliarPapeisHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Avaliar Papéis Handler")
	entidadeId := ""
	cicloId := ""
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		r.ParseForm()
		var produtoElemento mdl.ProdutoElemento
		for key, value := range r.Form {
			if strings.HasPrefix(key, "AuditorComponente_") {
				s := strings.Split(key, "_")
				log.Println("AuditorId_NEW: " + value[0])
				log.Println("AuditorId_OLD: " + s[5])
				log.Println("Entidade: " + s[1])
				entidadeId = s[1]
				log.Println("Ciclo: " + s[2])
				cicloId = s[2]
				log.Println("Pilar: " + s[3])
				log.Println("Componente: " + s[4])
				produtoElemento.EntidadeId, _ = strconv.ParseInt(s[1], 10, 64)
				produtoElemento.CicloId, _ = strconv.ParseInt(s[2], 10, 64)
				produtoElemento.PilarId, _ = strconv.ParseInt(s[3], 10, 64)
				produtoElemento.ComponenteId, _ = strconv.ParseInt(s[4], 10, 64)
				auditorIdOLD, _ := strconv.ParseInt(s[5], 10, 64)
				produtoElemento.AuditorId, _ = strconv.ParseInt(value[0], 10, 64)
				if auditorIdOLD != produtoElemento.AuditorId {
					registrarHistoricoAuditorComponente(produtoElemento, GetUserInCookie(w, r))
					registrarAuditorComponente(produtoElemento)
				}
			}
			if strings.HasPrefix(key, "ElementoNota") {
				log.Println("Nota: " + value[0])
				s := strings.Split(key, "_")
				log.Println("Entidade: " + s[1])
				entidadeId = s[1]
				log.Println("Ciclo: " + s[2])
				cicloId = s[2]
				log.Println("Pilar: " + s[3])
				log.Println("Componente: " + s[4])
				log.Println("Elemento: " + s[5])
				produtoElemento.EntidadeId, _ = strconv.ParseInt(s[1], 10, 64)
				produtoElemento.CicloId, _ = strconv.ParseInt(s[2], 10, 64)
				produtoElemento.PilarId, _ = strconv.ParseInt(s[3], 10, 64)
				produtoElemento.ComponenteId, _ = strconv.ParseInt(s[4], 10, 64)
				produtoElemento.ElementoId, _ = strconv.ParseInt(s[5], 10, 64)
				produtoElemento.Nota, _ = strconv.ParseFloat(value[0], 64)
				registrarNotaElemento(produtoElemento, GetUserInCookie(w, r))
			}
			if strings.HasPrefix(key, "ElementoPeso") {
				log.Println("Peso: " + value[0])
				s := strings.Split(key, "_")
				log.Println("Entidade: " + s[1])
				entidadeId = s[1]
				log.Println("Ciclo: " + s[2])
				cicloId = s[2]
				log.Println("Pilar: " + s[3])
				log.Println("Componente: " + s[4])
				log.Println("Elemento: " + s[5])
				produtoElemento.EntidadeId, _ = strconv.ParseInt(s[1], 10, 64)
				produtoElemento.CicloId, _ = strconv.ParseInt(s[2], 10, 64)
				produtoElemento.PilarId, _ = strconv.ParseInt(s[3], 10, 64)
				produtoElemento.ComponenteId, _ = strconv.ParseInt(s[4], 10, 64)
				produtoElemento.TipoNotaId, _ = strconv.ParseInt(s[5], 10, 64)
				produtoElemento.ElementoId, _ = strconv.ParseInt(s[6], 10, 64)
				produtoElemento.Peso, _ = strconv.ParseFloat(value[0], 64)
				registrarPesoElemento(produtoElemento, GetUserInCookie(w, r))
			}
		}
		AtualizarPapeisHandler(entidadeId, cicloId, w, r)
	}
}
