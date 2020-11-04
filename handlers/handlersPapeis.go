package handlers

import (
	"html/template"
	"log"
	"net/http"
	mdl "virtus/models"
	sec "virtus/security"
)

func DesignarPapeisHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Designar Papeis Handler")
	if sec.IsAuthenticated(w, r) {
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
		page.Title = "Designar Papéis"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/designarpapeis/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Designar-Papeis", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
