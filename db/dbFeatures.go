package db

import (
	"log"
	"strconv"
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
}

func updateFeatures() {
	query := " UPDATE features SET author_id = 1, created_at = now()::timestamp WHERE id <= 30"
	log.Println(query)
	db.Exec(query)
}

func createRoleFeatures() {
	for i := 1; i < 5; i++ {
		for j := 1; j <= 30; j++ {
			roleId := strconv.Itoa(i)
			featureId := strconv.Itoa(j)
			query := " INSERT INTO features_roles (role_id, feature_id) VALUES (" + roleId + ", " + featureId + ") "
			db.Exec(query)
		}
	}
}