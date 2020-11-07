package db

import (
	"log"
	"strconv"
	"strings"
)

func createFeatures() {
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 1, 'Listar Workflows', 'listWorkflows' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 1)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 2, 'Criar Workflow', 'createWorkflow' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 2)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 3, 'Listar Elementos', 'listElementos' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 3)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 4, 'Criar Elemento', 'createElemento' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 4)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 5, 'Listar Usuários', 'listUsers' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 5)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 6, 'Criar Usuário', 'createUser' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 6)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 7, 'Listar Papéis', 'listRoles' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 7)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 8, 'Criar Papel', 'createRole' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 8)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 9, 'Listar Status', 'listStatus' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 9)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 10, 'Criar Status', 'createStatus' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 10)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 11, 'Listar Funcionalidades', 'listFeatures' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 11)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 12, 'Criar Funcionalidade', 'createFeature' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 12)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 13, 'Listar Ações', 'listActions' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 13)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 14, 'Criar Ação', 'createAction' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 14)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 15, 'Criar Item', 'createItem' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 15)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 16, 'Listar Itens', 'listItens' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 16)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 17, 'Listar Componentes', 'listComponentes' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 17)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 18, 'Criar Componente', 'createComponente' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 18)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 19, 'Listar Pilares', 'listPilares' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 19)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 20, 'Criar Pilar', 'createPilar' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 20)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 21, 'Listar Ciclos', 'listCiclos' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 21)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 22, 'Criar Ciclo', 'createCiclo' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 22)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 23, 'Listar Entidades', 'listEntidades' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 23)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 24, 'Criar Entidade', 'createEntidade' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 24)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 25, 'Listar Planos', 'listPlanos' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 25)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 26, 'Criar Plano', 'createPlano' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 26)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 27, 'Listar Escritórios', 'listEscritorios' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 27)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 28, 'Criar Escritório', 'createEscritorio' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 28)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 29, 'Listar Tipos de Notas', 'listTiposNotas' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 29)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 30, 'Criar Tipo de Nota', 'createTipoNota' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 30)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 31, 'Designar Equipes', 'designarEquipes' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 31)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 32, 'Distribuir Papéis', 'distribuirPapeis' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 32)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 33, 'Avaliar Papéis', 'avaliarPapeis' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 33)")
	db.Exec("INSERT INTO public.features (id, name, code) SELECT 34, 'Visualizar Matriz', 'viewMatriz' WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 34)")
}

func updateFeatures() {
	query := " UPDATE features SET author_id = 1 WHERE id <= 34"
	log.Println(query)
	db.Exec(query)
}

func createRoleFeatures() {
	stmt1 := " INSERT INTO features_roles (role_id, feature_id) "
	stmt2 := ""
	for i := 1; i < 5; i++ {
		for j := 1; j <= 34; j++ {
			roleId := strconv.Itoa(i)
			featureId := strconv.Itoa(j)
			stmt2 = stmt2 + " SELECT " + roleId + ", " + featureId + " WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE feature_id = " + featureId + " AND role_id = " + roleId + ") UNION "
		}
	}
	pos := strings.LastIndex(stmt2, "UNION")
	stmt2 = stmt2[:pos]
	log.Println(stmt1 + stmt2)
	db.Exec(stmt1 + stmt2)
	stmt1 = " INSERT INTO features_roles (role_id, feature_id) " +
		" SELECT 5, a.id FROM features a " +
		" WHERE NOT EXISTS ( " +
		" SELECT 1  " +
		" FROM features_roles b " +
		" WHERE b.role_id = 5 AND b.feature_id = a.id) " +
		" AND SUBSTRING(a.code,1,4) = 'list' "
	log.Println(stmt1)
	db.Exec(stmt1)
}
