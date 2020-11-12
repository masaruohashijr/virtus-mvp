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

func ListMatrizesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Matrizes Handler")
	currentUser := GetUserInCookie(w, r)
	if sec.IsAuthenticated(w, r) && HasPermission(currentUser, "viewMatriz") {
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
		page.Title = "Matriz de Trabalho"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/matrizes/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Entidades-Matrizes", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ExecutarMatrizHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Matriz Handler")
	if sec.IsAuthenticated(w, r) {
		entidadeId := r.FormValue("EntidadeId")
		cicloId := r.FormValue("CicloId")
		var page mdl.PageMatriz
		sql := " SELECT  " +
			" a.ciclo_id, c.nome as ciclo_nome, " +
			" a.pilar_id, d.nome as pilar_nome, " +
			" a.componente_id, e.nome as componente_nome, " +
			" coalesce(j.peso,0) as componente_peso, coalesce(j.nota,0) as componente_nota, " +
			" (SELECT count(1) FROM (SELECT tipo_nota_id FROM elementos_componentes WHERE componente_id = a.componente_id GROUP BY tipo_nota_id) R) as qtdTiposNotas, " +
			" coalesce(k.peso,0) as pilar_peso, coalesce(k.nota,0) as pilar_nota, " +
			" (SELECT count(1) FROM (SELECT componente_id FROM componentes_pilares WHERE pilar_id = a.pilar_id GROUP BY componente_id) R) as qtdComponentes, " +
			" coalesce(l.nota,0) as ciclo_nota, " +
			" (SELECT count(1) FROM (SELECT pilar_id FROM pilares_ciclos WHERE ciclo_id = a.ciclo_id GROUP BY pilar_id) R) as qtdPilares, " +
			" a.elemento_id, f.nome as elemento_nome, " +
			" coalesce(n.peso,0) as elemento_peso, coalesce(n.nota,0) as elemento_nota, " +
			" coalesce(o.peso,0) as tipo_nota_peso, coalesce(o.nota,0) as tipo_nota_nota, " +
			" ec.tipo_nota_id, m.letra, m.cor_letra, ec.peso_padrao, " +
			" cp.tipo_media, cp.peso_padrao, " +
			" pc.tipo_media, pc.peso_padrao, " +
			" coalesce(ce.tipo_media,0), " +
			" coalesce(to_char(ce.inicia_em,'DD/MM/YYYY')) as inicia_em, " +
			" coalesce(to_char(ce.termina_em,'DD/MM/YYYY')) as termina_em, " +
			" a.item_id, g.nome as item_nome, " +
			" coalesce(b.nome,''), a.entidade_id, " +
			" coalesce(j.supervisor_id,0) as super_id, coalesce(h.name,'') as supervisor_nome, " +
			" coalesce(j.auditor_id,0) as audit_id, coalesce(i.name,'') as auditor_nome " +
			" FROM produtos_itens a " +
			" LEFT JOIN entidades b ON a.entidade_id = b.id " +
			" LEFT JOIN ciclos c ON a.ciclo_id = c.id " +
			" LEFT JOIN pilares d ON a.pilar_id = d.id " +
			" LEFT JOIN componentes e ON a.componente_id = e.id " +
			" LEFT JOIN produtos_elementos n ON " +
			" ( a.elemento_id = n.elemento_id AND " +
			" a.componente_id = n.componente_id AND " +
			" a.pilar_id = n.pilar_id AND " +
			" a.ciclo_id = n.ciclo_id AND " +
			" a.entidade_id = n.entidade_id )" +
			" LEFT JOIN produtos_tipos_notas o ON " +
			" ( a.tipo_nota_id = o.tipo_nota_id AND " +
			" a.componente_id = o.componente_id AND " +
			" a.pilar_id = o.pilar_id AND " +
			" a.ciclo_id = o.ciclo_id AND " +
			" a.entidade_id = o.entidade_id )" +
			" LEFT JOIN produtos_componentes j ON " +
			" ( a.componente_id = j.componente_id AND " +
			" a.pilar_id = j.pilar_id AND " +
			" a.ciclo_id = j.ciclo_id AND " +
			" a.entidade_id = j.entidade_id )" +
			" LEFT JOIN produtos_pilares k ON " +
			" ( a.pilar_id = k.pilar_id AND " +
			"   a.ciclo_id = k.ciclo_id AND " +
			"   a.entidade_id = k.entidade_id ) " +
			" LEFT JOIN produtos_ciclos l ON " +
			" ( a.ciclo_id = l.ciclo_id AND " +
			"   a.entidade_id = l.entidade_id ) " +
			" LEFT JOIN elementos_componentes ec ON " +
			" ( a.elemento_id = ec.elemento_id AND a.componente_id = ec.componente_id ) " +
			" LEFT JOIN componentes_pilares cp ON " +
			" ( a.componente_id = cp.componente_id AND a.pilar_id = cp.pilar_id ) " +
			" LEFT JOIN pilares_ciclos pc ON " +
			" ( a.pilar_id = pc.pilar_id AND a.ciclo_id = pc.ciclo_id ) " +
			" LEFT JOIN ciclos_entidades ce ON a.componente_id = ce.id  " +
			" LEFT JOIN elementos f ON a.elemento_id = f.id  " +
			" LEFT JOIN itens g ON a.item_id = g.id  " +
			" LEFT JOIN users h ON j.supervisor_id = h.id  " +
			" LEFT JOIN users i ON j.auditor_id = i.id  " +
			" LEFT JOIN tipos_notas m ON ec.tipo_nota_id = m.id " +
			" WHERE a.entidade_id = " + entidadeId + " AND a.ciclo_id = " + cicloId +
			" ORDER BY a.ciclo_id, " +
			" a.pilar_id,  " +
			" a.componente_id, " +
			" a.tipo_nota_id, " +
			" a.elemento_id, " +
			" a.item_id "
		log.Println(sql)
		rows, _ := Db.Query(sql)
		var elementosMatriz []mdl.ElementoDaMatriz
		var elementoMatriz mdl.ElementoDaMatriz
		var i = 1
		cicloColspan := 0
		pilarColspan := 0
		auxPilarNome := ""
		componenteColspan := 0
		auxComponenteNome := ""
		for rows.Next() {
			rows.Scan(
				&elementoMatriz.CicloId,
				&elementoMatriz.CicloNome,
				&elementoMatriz.PilarId,
				&elementoMatriz.PilarNome,
				&elementoMatriz.ComponenteId,
				&elementoMatriz.ComponenteNome,
				&elementoMatriz.ComponentePeso,
				&elementoMatriz.ComponenteNota,
				&elementoMatriz.ComponenteQtdTiposNotas,
				&elementoMatriz.PilarPeso,
				&elementoMatriz.PilarNota,
				&elementoMatriz.PilarQtdComponentes,
				&elementoMatriz.CicloNota,
				&elementoMatriz.CicloQtdPilares,
				&elementoMatriz.ElementoId,
				&elementoMatriz.ElementoNome,
				&elementoMatriz.ElementoPeso,
				&elementoMatriz.ElementoNota,
				&elementoMatriz.TipoNotaPeso,
				&elementoMatriz.TipoNotaNota,
				&elementoMatriz.TipoNotaId,
				&elementoMatriz.TipoNotaLetra,
				&elementoMatriz.TipoNotaCorLetra,
				&elementoMatriz.PesoPadraoEC,
				&elementoMatriz.TipoMediaCPId,
				&elementoMatriz.PesoPadraoCP,
				&elementoMatriz.TipoMediaPCId,
				&elementoMatriz.PesoPadraoPC,
				&elementoMatriz.TipoMediaCEId,
				&elementoMatriz.IniciaEm,
				&elementoMatriz.TerminaEm,
				&elementoMatriz.ItemId,
				&elementoMatriz.ItemNome,
				&elementoMatriz.EntidadeNome,
				&elementoMatriz.EntidadeId,
				&elementoMatriz.SupervisorId,
				&elementoMatriz.SupervisorName,
				&elementoMatriz.AuditorId,
				&elementoMatriz.AuditorName)
			elementoMatriz.Order = i
			i++
			log.Println(elementoMatriz)
			if auxPilarNome != elementoMatriz.PilarNome {
				log.Println("Calculando Pilar colspan")
				pilarColspan = calcularColspan("pilar", elementoMatriz.PilarId)
				auxPilarNome = elementoMatriz.PilarNome
			}
			if auxComponenteNome != elementoMatriz.ComponenteNome {
				log.Println("Calculando Componente colspan")
				componenteColspan = calcularColspan("componente", elementoMatriz.ComponenteId)
				auxComponenteNome = elementoMatriz.ComponenteNome
			}
			if cicloColspan == 0 {
				log.Println("Calculando Ciclo colspan")
				cicloColspan = calcularColspan("ciclo", elementoMatriz.CicloId)
			}
			elementoMatriz.ComponenteColSpan = componenteColspan
			elementoMatriz.PilarColSpan = pilarColspan
			elementoMatriz.CicloColSpan = cicloColspan
			elementosMatriz = append(elementosMatriz, elementoMatriz)
		}
		page.ElementosDaMatriz = elementosMatriz

		sql = " SELECT " +
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
			" AND b.role_id = 4 "
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
		page.Title = "Matriz"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/matrizes/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Matriz", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func calcularColspan(tipo string, identificador int64) int {
	sql := ""
	if tipo == "ciclo" {
		tipo = "pc.ciclo"
	} else if tipo == "pilar" {
		tipo = "pc.pilar"
	} else if tipo == "componente" {
		tipo = "cp.componente"
	}
	sql = " SELECT COUNT(1) FROM ( " +
		" SELECT ciclo_id, pc.pilar_id, cp.componente_id, tnc.tipo_nota_id " +
		" FROM tipos_notas_componentes tnc " +
		" LEFT JOIN componentes c ON tnc.componente_id = c.id " +
		" LEFT JOIN componentes_pilares cp ON c.id = cp.componente_id " +
		" LEFT JOIN pilares p ON p.id = cp.pilar_id " +
		" LEFT JOIN pilares_ciclos pc ON p.id = pc.pilar_id " +
		" WHERE " + tipo + "_id = $1) R "
	log.Println(sql)
	rows, _ := Db.Query(sql, identificador)
	resultado := 0
	if rows.Next() {
		rows.Scan(&resultado)
		return resultado
	}
	return 0
}

func AtualizarMatrizHandler(entidadeId string, cicloId string, w http.ResponseWriter, r *http.Request) {
	log.Println("Atualizar Papéis Handler")
	var page mdl.PageProdutosItens
	sql := " SELECT  " +
		" a.ciclo_id, c.nome as ciclo_nome, " +
		" a.pilar_id, d.nome as pilar_nome, " +
		" a.componente_id, e.nome as componente_nome, " +
		" coalesce(j.peso,0) as componente_peso, coalesce(j.nota,0) as componente_nota, " +
		" coalesce(k.peso,0) as pilar_peso, coalesce(k.nota,0) as pilar_nota, " +
		" coalesce(l.nota,0) as ciclo_nota, " +
		" a.elemento_id, f.nome as elemento_nome, " +
		" coalesce(n.peso,0) as elemento_peso, coalesce(n.nota,0) as elemento_nota, " +
		" ec.tipo_nota_id, m.letra, m.cor_letra, ec.peso_padrao, " +
		" cp.tipo_media, cp.peso_padrao, " +
		" pc.tipo_media, pc.peso_padrao, " +
		" ce.tipo_media, ce.inicia_em, ce.termina_em, " +
		" a.item_id, g.nome as item_nome, " +
		" coalesce(b.nome,''), a.entidade_id, " +
		" coalesce(j.supervisor_id,0) as super_id, coalesce(h.name,'') as supervisor_nome, " +
		" coalesce(j.auditor_id,0) as audit_id, coalesce(i.name,'') as auditor_nome " +
		" FROM produtos_itens a " +
		" LEFT JOIN entidades b ON a.entidade_id = b.id " +
		" LEFT JOIN ciclos c ON a.ciclo_id = c.id " +
		" LEFT JOIN pilares d ON a.pilar_id = d.id " +
		" LEFT JOIN componentes e ON a.componente_id = e.id " +
		" LEFT JOIN produtos_elementos n ON " +
		" ( a.elemento_id = n.elemento_id AND " +
		" a.componente_id = n.componente_id AND " +
		" a.pilar_id = n.pilar_id AND " +
		" a.ciclo_id = n.ciclo_id AND " +
		" a.entidade_id = n.entidade_id )" +
		" LEFT JOIN produtos_componentes j ON " +
		" ( a.componente_id = j.componente_id AND " +
		" a.pilar_id = j.pilar_id AND " +
		" a.ciclo_id = j.ciclo_id AND " +
		" a.entidade_id = j.entidade_id )" +
		" LEFT JOIN produtos_pilares k ON " +
		" ( a.pilar_id = k.pilar_id AND " +
		"   a.ciclo_id = k.ciclo_id AND " +
		"   a.entidade_id = k.entidade_id ) " +
		" LEFT JOIN produtos_ciclos l ON " +
		" ( a.ciclo_id = l.ciclo_id AND " +
		"   a.entidade_id = l.entidade_id ) " +
		" LEFT JOIN elementos_componentes ec ON " +
		" ( a.elemento_id = ec.elemento_id AND a.componente_id = ec.componente_id ) " +
		" LEFT JOIN componentes_pilares cp ON " +
		" ( a.componente_id = cp.componente_id AND a.pilar_id = cp.pilar_id ) " +
		" LEFT JOIN pilares_ciclos pc ON " +
		" ( a.pilar_id = pc.pilar_id AND a.ciclo_id = pc.ciclo_id ) " +
		" LEFT JOIN ciclos_entidades ce ON a.componente_id = ce.id  " +
		" LEFT JOIN elementos f ON a.elemento_id = f.id  " +
		" LEFT JOIN itens g ON a.item_id = g.id  " +
		" LEFT JOIN users h ON j.supervisor_id = h.id  " +
		" LEFT JOIN users i ON j.auditor_id = i.id  " +
		" LEFT JOIN tipos_notas m ON ec.tipo_nota_id = m.id " +
		" WHERE a.entidade_id = " + entidadeId + " AND a.ciclo_id = " + cicloId +
		" ORDER BY a.ciclo_id, " +
		" a.pilar_id,  " +
		" a.componente_id, " +
		" a.elemento_id, " +
		" a.item_id "
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
			&produto.ComponentePeso,
			&produto.ComponenteNota,
			&produto.PilarPeso,
			&produto.PilarNota,
			&produto.CicloNota,
			&produto.ElementoId,
			&produto.ElementoNome,
			&produto.ElementoPeso,
			&produto.ElementoNota,
			&produto.TipoNotaId,
			&produto.TipoNotaLetra,
			&produto.TipoNotaCorLetra,
			&produto.PesoPadraoEC,
			&produto.TipoMediaCPId,
			&produto.PesoPadraoCP,
			&produto.TipoMediaPCId,
			&produto.PesoPadraoPC,
			&produto.TipoMediaCEId,
			&produto.IniciaEm,
			&produto.TerminaEm,
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
		// log.Println(produto)
		produtos = append(produtos, produto)
	}
	page.Produtos = produtos

	sql = " SELECT " +
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
		" AND b.role_id = 4 "
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
	page.Title = "Matriz"
	page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
	var tmpl = template.Must(template.ParseGlob("tiles/matrizes/*"))
	tmpl.ParseGlob("tiles/*")
	tmpl.ExecuteTemplate(w, "Main-Matrizes", page)
}

func UpdateMatrizHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Matriz Handler")
	entidadeId := ""
	cicloId := ""
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		r.ParseForm()
		var produtoElemento mdl.ProdutoElemento
		for key, value := range r.Form {
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
				log.Println("TipoNota: " + s[5])
				log.Println("Elemento: " + s[6])
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
