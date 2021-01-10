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

const sqlAvaliarPlanos = " SELECT a.entidade_id, " +
	" 	   coalesce(b.nome,'') as entidade_nome, " +
	"        a.ciclo_id, " +
	" 	   coalesce(c.nome,'') as ciclo_nome, " +
	" 	   coalesce(pdc.nota,0) as ciclo_nota, " +
	"        a.pilar_id,        " +
	" 	   coalesce(d.nome,'') as pilar_nome, " +
	" 	   coalesce(f.peso,0) as pilar_peso, coalesce(f.nota,0) as pilar_nota, " +
	"        a.componente_id, " +
	" 	   coalesce(e.nome,'') as componente_nome, " +
	" 	   coalesce(g.peso,0) as componente_peso, coalesce(g.nota,0) as componente_nota, " +
	" 	   coalesce(g.supervisor_id,0) as super_id, coalesce(h.name,'') as supervisor_nome, " +
	" 	   coalesce(g.auditor_id,0) as super_id, coalesce(i.name,'') as auditor_nome, " +
	" 	   a.tipo_nota_id, m.letra, m.cor_letra, m.nome, " +
	" 	   coalesce(o.peso,0) as tipo_nota_peso, coalesce(o.nota,0) as tipo_nota_nota, " +
	"        a.elemento_id, coalesce(k.nome,'') as elemento_nome, " +
	" 	   coalesce(n.peso,0) as elemento_peso, coalesce(n.nota,0) as elemento_nota, " +
	"	   n.tipo_pontuacao_id, ec.peso_padrao, " +
	" 	   cp.tipo_media, cp.peso_padrao, " +
	" 	   pc.tipo_media, pc.peso_padrao, " +
	" 	   a.item_id, coalesce(l.nome,'') as item_nome, " +
	"      a.plano_id, " +
	"	   j.cnpb, CASE WHEN j.recurso_garantidor < 1000000000 THEN j.recurso_garantidor::numeric::MONEY/1000000||' mi' ELSE j.recurso_garantidor::numeric::MONEY/1000000000||' bi' END, j.modalidade_id, " +
	" 	   coalesce(p.peso,0) as plano_peso, coalesce(p.nota,0) as plano_nota, " +
	"	   to_char(g.inicia_em,'DD/MM/YYYY') as inicia_em, to_char(g.termina_em,'DD/MM/YYYY') as termina_em, " +
	"	   CASE " +
	"	    WHEN now()::TIMESTAMP BETWEEN coalesce(g.inicia_em,to_date('0001-01-01','YYYY-MM-DD')) " +
	"	    AND coalesce(g.termina_em,to_date('9999-12-31','YYYY-MM-DD')) " +
	"	    THEN TRUE " +
	"	    ELSE FALSE " +
	"	   END AS periodo_permitido " +
	" FROM produtos_itens a " +
	" INNER JOIN entidades b ON a.entidade_id = b.id " +
	" INNER JOIN ciclos c ON a.ciclo_id = c.id " +
	" INNER JOIN pilares d ON a.pilar_id = d.id " +
	" INNER JOIN componentes e ON a.componente_id = e.id " +
	" INNER JOIN produtos_pilares f ON " +
	" ( a.pilar_id = f.pilar_id AND  " +
	"   a.ciclo_id = f.ciclo_id AND  " +
	"   a.entidade_id = f.entidade_id ) " +
	" INNER JOIN produtos_componentes g ON  " +
	" ( a.componente_id = g.componente_id AND  " +
	"   a.pilar_id = g.pilar_id AND  " +
	"   a.ciclo_id = g.ciclo_id AND  " +
	"   a.entidade_id = g.entidade_id  " +
	" ) " +
	" LEFT JOIN users h ON g.supervisor_id = h.id " +
	" LEFT JOIN users i ON g.auditor_id = i.id " +
	" INNER JOIN planos j ON a.plano_id = j.id " +
	" INNER JOIN elementos k ON a.elemento_id = k.id   " +
	" INNER JOIN itens l ON a.item_id = l.id   " +
	" INNER JOIN elementos_componentes ec ON ( a.elemento_id = ec.elemento_id AND a.tipo_nota_id = ec.tipo_nota_id AND a.componente_id = ec.componente_id ) " +
	" INNER JOIN componentes_pilares cp ON ( a.componente_id = cp.componente_id AND a.pilar_id = cp.pilar_id ) " +
	" INNER JOIN pilares_ciclos pc ON ( a.pilar_id = pc.pilar_id AND a.ciclo_id = pc.ciclo_id ) " +
	" INNER JOIN tipos_notas m ON a.tipo_nota_id = m.id " +
	" INNER JOIN produtos_elementos n ON  " +
	" 	( a.elemento_id = n.elemento_id AND  " +
	" 	a.tipo_nota_id = n.tipo_nota_id AND  " +
	" 	a.plano_id = n.plano_id AND  " +
	" 	a.componente_id = n.componente_id AND  " +
	" 	a.pilar_id = n.pilar_id AND  " +
	" 	a.ciclo_id = n.ciclo_id AND  " +
	" 	a.entidade_id = n.entidade_id ) " +
	" INNER JOIN produtos_tipos_notas o ON  " +
	" ( a.tipo_nota_id = o.tipo_nota_id AND  " +
	"   a.plano_id = o.plano_id AND " +
	"   a.componente_id = o.componente_id AND " +
	"   a.pilar_id = o.pilar_id AND  " +
	"   a.ciclo_id = o.ciclo_id AND  " +
	"   a.entidade_id = o.entidade_id) " +
	" INNER JOIN produtos_planos p ON  " +
	"  (a.plano_id = p.plano_id AND  " +
	"   a.componente_id = p.componente_id AND  " +
	"   a.pilar_id = p.pilar_id AND  " +
	"   a.ciclo_id = p.ciclo_id AND  " +
	"   a.entidade_id = p.entidade_id)   " +
	" INNER JOIN produtos_ciclos pdc ON " +
	"  (a.ciclo_id = pdc.ciclo_id AND " +
	"   a.entidade_id = pdc.entidade_id) " +
	" WHERE a.entidade_id = $1 " +
	"   AND a.ciclo_id = $2 " +
	" ORDER BY a.ciclo_id, " +
	"          a.pilar_id, " +
	"          a.componente_id, " +
	"          j.recurso_garantidor DESC, " +
	"          a.tipo_nota_id, " +
	"          a.elemento_id, " +
	"          a.item_id "

func ListAvaliarPlanosHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Avaliar Planos Handler")
	currentUser := GetUserInCookie(w, r)
	if sec.IsAuthenticated(w, r) && HasPermission(currentUser, "avaliarPlanos") {
		log.Println("--------------")
		var page mdl.PageEntidadesCiclos
		// Entidades da jurisdição do Escritório ao qual pertenço
		sql := "SELECT DISTINCT d.codigo, b.entidade_id, d.nome, a.abreviatura " +
			" FROM escritorios a " +
			" LEFT JOIN jurisdicoes b ON a.id = b.escritorio_id " +
			" LEFT JOIN membros c ON c.escritorio_id = b.escritorio_id " +
			" LEFT JOIN entidades d ON d.id = b.entidade_id " +
			" LEFT JOIN users u ON u.id = c.usuario_id " +
			" INNER JOIN ciclos_entidades e ON e.entidade_id = b.entidade_id " +
			" INNER JOIN produtos_planos f ON (f.entidade_id = e.entidade_id AND f.ciclo_id = e.ciclo_id) " +
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
		page.Entidades = entidadesCiclos
		page.AppName = mdl.AppName
		page.Title = "Avaliar Planos"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/avaliarplanos/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Entidades-Avaliar-Planos", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func AvaliarPlanosHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Avaliar Planos Handler")
	if sec.IsAuthenticated(w, r) {
		entidadeId := r.FormValue("EntidadeId")
		cicloId := r.FormValue("CicloId")
		var page mdl.PageProdutosItens
		log.Println(sqlAvaliarPlanos)
		rows, _ := Db.Query(sqlAvaliarPlanos, entidadeId, cicloId)
		defer rows.Close()
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
				&produto.ItemId,
				&produto.ItemNome,
				&produto.PlanoId,
				&produto.CNPB,
				&produto.RecursoGarantidor,
				&produto.PlanoModalidade,
				&produto.PlanoPeso,
				&produto.PlanoNota,
				&produto.IniciaEm,
				&produto.TerminaEm,
				&produto.PeriodoPermitido)
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
		defer rows.Close()
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
			" AND b.role_id in (2,3,4) ORDER BY 2 "
		log.Println(sql)
		rows, _ = Db.Query(sql)
		defer rows.Close()
		var auditores []mdl.User
		var auditor mdl.User
		i = 1
		for rows.Next() {
			rows.Scan(&auditor.Id, &auditor.Name)
			//log.Println("Auditor competente: " + auditor.Name)
			auditores = append(auditores, auditor)
		}
		page.Supervisores = supervisores
		page.Auditores = auditores
		page.AppName = mdl.AppName
		page.Title = "Avaliar Planos"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		page.Inc = func(i int) int {
			return i + 1
		}
		var tmpl = template.Must(template.ParseGlob("tiles/avaliarplanos/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Avaliar-Planos", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func AtualizarPlanosHandler(entidadeId string, cicloId string, w http.ResponseWriter, r *http.Request) {
	log.Println("Atualizar Planos Handler")
	var page mdl.PageProdutosItens
	log.Println(sqlAvaliarPlanos)
	rows, _ := Db.Query(sqlAvaliarPlanos, entidadeId, cicloId)
	defer rows.Close()
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
			&produto.ItemId,
			&produto.ItemNome,
			&produto.PlanoId,
			&produto.CNPB,
			&produto.RecursoGarantidor,
			&produto.PlanoModalidade,
			&produto.PlanoPeso,
			&produto.PlanoNota,
			&produto.IniciaEm,
			&produto.TerminaEm,
			&produto.PeriodoPermitido)
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
	defer rows.Close()
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
	defer rows.Close()
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
	page.Title = "Avaliar Planos"
	page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
	page.Inc = func(i int) int {
		return i + 1
	}
	var tmpl = template.Must(template.ParseGlob("tiles/avaliarplanos/*"))
	tmpl.ParseGlob("tiles/*")
	tmpl.ExecuteTemplate(w, "Main-Avaliar-Planos", page)
}

func UpdateAvaliarPlanosHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Avaliar Planos Handler")
	entidadeId := ""
	cicloId := ""
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		currentUser := GetUserInCookie(w, r)
		r.ParseForm()
		var produtoElemento mdl.ProdutoElemento
		motivacaoNota := r.FormValue("MotivacaoNota")
		motivacaoPeso := r.FormValue("MotivacaoPeso")
		motivacaoRemocao := r.FormValue("MotivacaoRemocao")
		log.Println("******************")
		log.Println(motivacaoRemocao)
		log.Println("******************")
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
					produtoElemento.Motivacao = motivacaoRemocao
					//registrarAuditorComponente(produtoElemento)
					//					registrarHistoricoAuditorComponente(produtoElemento, currentUser)
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
				log.Println("TipoNota: " + s[5])
				log.Println("Elemento: " + s[6])
				log.Println("NotaAnterior: " + s[7])
				notaOLD, _ := strconv.Atoi(s[7])
				produtoElemento.EntidadeId, _ = strconv.ParseInt(s[1], 10, 64)
				produtoElemento.CicloId, _ = strconv.ParseInt(s[2], 10, 64)
				produtoElemento.PilarId, _ = strconv.ParseInt(s[3], 10, 64)
				produtoElemento.ComponenteId, _ = strconv.ParseInt(s[4], 10, 64)
				produtoElemento.TipoNotaId, _ = strconv.ParseInt(s[5], 10, 64)
				produtoElemento.ElementoId, _ = strconv.ParseInt(s[6], 10, 64)
				produtoElemento.Nota, _ = strconv.Atoi(value[0])
				if notaOLD != produtoElemento.Nota {
					produtoElemento.Motivacao = motivacaoNota
					registrarNotaElemento(produtoElemento, currentUser)
					registrarHistoricoNotaElemento(produtoElemento, currentUser)
				}
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
				log.Println("PesoAnterior: " + s[7])
				pesoOLD, _ := strconv.ParseFloat(s[7], 10)
				produtoElemento.EntidadeId, _ = strconv.ParseInt(s[1], 10, 64)
				produtoElemento.CicloId, _ = strconv.ParseInt(s[2], 10, 64)
				produtoElemento.PilarId, _ = strconv.ParseInt(s[3], 10, 64)
				produtoElemento.ComponenteId, _ = strconv.ParseInt(s[4], 10, 64)
				produtoElemento.TipoNotaId, _ = strconv.ParseInt(s[5], 10, 64)
				produtoElemento.ElementoId, _ = strconv.ParseInt(s[6], 10, 64)
				produtoElemento.Peso, _ = strconv.ParseFloat(value[0], 64)
				if pesoOLD != produtoElemento.Peso {
					produtoElemento.Motivacao = motivacaoPeso
					registrarPesoElemento(produtoElemento, currentUser)
					registrarHistoricoPesoElemento(produtoElemento, currentUser)
				}
			}
		}
		AtualizarPlanosHandler(entidadeId, cicloId, w, r)
	}
}

func SalvarPesoElemento(w http.ResponseWriter, r *http.Request) {
	log.Println("Salvar Peso Elemento")
	r.ParseForm()
	var entidadeId = r.FormValue("entidadeId")
	var cicloId = r.FormValue("cicloId")
	var pilarId = r.FormValue("pilarId")
	var planoId = r.FormValue("planoId")
	var tipoNotaId = r.FormValue("tipoNotaId")
	var componenteId = r.FormValue("componenteId")
	var elementoId = r.FormValue("elementoId")
	motivacaoPeso := r.FormValue("motivacao")
	peso := r.FormValue("peso")
	var produtoElemento mdl.ProdutoElemento
	produtoElemento.EntidadeId, _ = strconv.ParseInt(entidadeId, 10, 64)
	produtoElemento.CicloId, _ = strconv.ParseInt(cicloId, 10, 64)
	produtoElemento.PilarId, _ = strconv.ParseInt(pilarId, 10, 64)
	produtoElemento.PlanoId, _ = strconv.ParseInt(planoId, 10, 64)
	produtoElemento.TipoNotaId, _ = strconv.ParseInt(tipoNotaId, 10, 64)
	produtoElemento.ComponenteId, _ = strconv.ParseInt(componenteId, 10, 64)
	produtoElemento.ElementoId, _ = strconv.ParseInt(elementoId, 10, 64)
	produtoElemento.Peso, _ = strconv.ParseFloat(peso, 64)
	produtoElemento.Motivacao = motivacaoPeso
	currentUser := GetUserInCookie(w, r)
	pesosAtuais := registrarPesoElemento(produtoElemento, currentUser)
	registrarHistoricoPesoElemento(produtoElemento, currentUser)
	jsonPesosAtuais, _ := json.Marshal(pesosAtuais)
	w.Write([]byte(jsonPesosAtuais))
	log.Println("----------")
}

func SalvarNotaElemento(w http.ResponseWriter, r *http.Request) {
	log.Println("Salvar Nota Elemento")
	r.ParseForm()
	var entidadeId = r.FormValue("entidadeId")
	var cicloId = r.FormValue("cicloId")
	var pilarId = r.FormValue("pilarId")
	var planoId = r.FormValue("planoId")
	var tipoNotaId = r.FormValue("tipoNotaId")
	var componenteId = r.FormValue("componenteId")
	var elementoId = r.FormValue("elementoId")
	motivacaoNota := r.FormValue("motivacao")
	nota := r.FormValue("nota")
	var produtoElemento mdl.ProdutoElemento
	produtoElemento.EntidadeId, _ = strconv.ParseInt(entidadeId, 10, 64)
	produtoElemento.CicloId, _ = strconv.ParseInt(cicloId, 10, 64)
	produtoElemento.PilarId, _ = strconv.ParseInt(pilarId, 10, 64)
	produtoElemento.PlanoId, _ = strconv.ParseInt(planoId, 10, 64)
	produtoElemento.TipoNotaId, _ = strconv.ParseInt(tipoNotaId, 10, 64)
	produtoElemento.ComponenteId, _ = strconv.ParseInt(componenteId, 10, 64)
	produtoElemento.ElementoId, _ = strconv.ParseInt(elementoId, 10, 64)
	produtoElemento.Nota, _ = strconv.Atoi(nota)
	produtoElemento.Motivacao = motivacaoNota
	currentUser := GetUserInCookie(w, r)
	notasAtuais := registrarNotaElemento(produtoElemento, currentUser)
	registrarHistoricoNotaElemento(produtoElemento, currentUser)
	//log.Println(notasAtuais)
	jsonNotasAtuais, _ := json.Marshal(notasAtuais)
	w.Write([]byte(jsonNotasAtuais))
	log.Println("----------")

}

func SalvarAuditorComponente(w http.ResponseWriter, r *http.Request) {
	log.Println("Salvar Auditor Componente")
	r.ParseForm()
	var entidadeId = r.FormValue("entidadeId")
	var cicloId = r.FormValue("cicloId")
	var pilarId = r.FormValue("pilarId")
	var componenteId = r.FormValue("componenteId")
	motivacao := r.FormValue("motivacao")
	auditorAnterior := r.FormValue("auditorAnterior")
	auditorNovo := r.FormValue("auditorNovo")
	var produtoComponente mdl.ProdutoComponente
	produtoComponente.EntidadeId, _ = strconv.ParseInt(entidadeId, 10, 64)
	produtoComponente.CicloId, _ = strconv.ParseInt(cicloId, 10, 64)
	produtoComponente.PilarId, _ = strconv.ParseInt(pilarId, 10, 64)
	produtoComponente.ComponenteId, _ = strconv.ParseInt(componenteId, 10, 64)
	produtoComponente.AuditorId, _ = strconv.ParseInt(auditorNovo, 10, 64)
	produtoComponente.AuditorAnteriorId, _ = strconv.ParseInt(auditorAnterior, 10, 64)
	produtoComponente.Motivacao = motivacao
	currentUser := GetUserInCookie(w, r)
	registrarAuditorComponente(produtoComponente, currentUser)
	registrarHistoricoAuditorComponente(produtoComponente, currentUser)
	jsonOK, _ := json.Marshal("OK")
	w.Write(jsonOK)
	log.Println("----------")

}

func SalvarPesoPilar(w http.ResponseWriter, r *http.Request) {
	log.Println("Salvar Peso Pilar")
	r.ParseForm()
	var entidadeId = r.FormValue("entidadeId")
	var cicloId = r.FormValue("cicloId")
	var pilarId = r.FormValue("pilarId")
	motivacaoPeso := r.FormValue("motivacao")
	peso := r.FormValue("peso")
	var produtoPilar mdl.ProdutoPilar
	produtoPilar.EntidadeId, _ = strconv.ParseInt(entidadeId, 10, 64)
	produtoPilar.CicloId, _ = strconv.ParseInt(cicloId, 10, 64)
	produtoPilar.PilarId, _ = strconv.ParseInt(pilarId, 10, 64)
	produtoPilar.Peso, _ = strconv.ParseFloat(peso, 64)
	produtoPilar.Motivacao = motivacaoPeso
	currentUser := GetUserInCookie(w, r)
	registrarPesoPilar(produtoPilar)
	registrarHistoricoPesoPilar(produtoPilar, currentUser)
	w.Write([]byte("OK"))
	log.Println("----------")
}

func LoadAnalise(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Analise")
	r.ParseForm()
	var rota = r.FormValue("btn")
	analise := getAnalise(rota)
	w.Write([]byte(analise))
	log.Println("Fim Load Analise")
}

func SalvarAnalise(w http.ResponseWriter, r *http.Request) {
	log.Println("Salvar Analise")
	r.ParseForm()
	var rota = r.FormValue("acionadoPor")
	var analise = r.FormValue("analise")
	retorno := setAnalise(rota, analise)
	w.Write([]byte(retorno))
	log.Println("----------")
}

func LoadDescricao(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Descrição")
	r.ParseForm()
	var rota = r.FormValue("btn")
	descricao := getDescricao(rota)
	jsonDescricao, _ := json.Marshal(descricao)
	w.Write([]byte(jsonDescricao))
	log.Println("JSON Descrição")
}
