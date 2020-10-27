package db

import (
	"log"
	"strconv"
)

func createFeatures() {
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (1, 'Listar Workflows', 'listWorkflows')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (2, 'Criar Workflow', 'createWorkflow')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (3, 'Listar Elementos', 'listElementos')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (4, 'Criar Elemento', 'createElemento')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (5, 'Listar Usuários', 'listUsers')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (6, 'Criar Usuário', 'createUser')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (7, 'Listar Papéis', 'listRoles')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (8, 'Criar Papel', 'createRole')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (9, 'Listar Status', 'listStatus')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (10, 'Criar Status', 'createStatus')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (11, 'Listar Funcionalidades', 'listFeatures')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (12, 'Criar Funcionalidade', 'createFeature')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (13, 'Listar Ações', 'listActions')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (14, 'Criar Ação', 'createAction')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (15, 'Criar Item', 'createItem')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (16, 'Listar Itens', 'listItens')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (17, 'Listar Componentes', 'listComponentes')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (18, 'Criar Componente', 'createComponente')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (19, 'Listar Pilares', 'listPilares')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (20, 'Criar Pilar', 'createPilar')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (21, 'Listar Ciclos', 'listCiclos')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (22, 'Criar Ciclo', 'createCiclo')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (23, 'Listar Entidades', 'listEntidades')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (24, 'Criar Entidade', 'createEntidade')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (25, 'Listar Planos', 'listPlanos')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (26, 'Criar Plano', 'createPlano')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (27, 'Listar Escritórios', 'listEscritorios')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (28, 'Criar Escritório', 'createEscritorio')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (29, 'Listar Tipos de Notas', 'listTiposNotas')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (30, 'Criar Tipo de Nota', 'createTipoNota')")
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
