package db

import (
	"database/sql"
	"log"
	hd "virtus/handlers"
)

var db *sql.DB

func Initialize() {
	db = hd.Db
	createSeq()
	createTable()
	createFeatures()
	createRoles()
	createRoleFeatures()
	createUsers()
	createEscritorios()
	createStatusZERO()
	updateRoles()
	updateFeatures()
	createPKey()
	createFKey()
	createUniqueKey()
}

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
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (29, 'Listar Carteiras', 'listCarteiras')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (30, 'Criar Carteira', 'createCarteira')")
}

func createUniqueKey() {
	db.Exec(" ALTER TABLE ONLY public.actions_status" +
		" ADD CONSTRAINT action_status_unique_key UNIQUE (action_id, origin_status_id, destination_status_id)")

	db.Exec(" ALTER TABLE ONLY public.features_roles" +
		" ADD CONSTRAINT feature_role_unique_key UNIQUE (role_id, feature_id)")

	db.Exec(" ALTER TABLE ONLY public.users" +
		" ADD CONSTRAINT username_unique_key UNIQUE (username)")

	db.Exec(" ALTER TABLE ONLY public.activities_roles" +
		" ADD CONSTRAINT action_role_unique_key UNIQUE (activity_id, role_id)")

	db.Exec(" ALTER TABLE ONLY public.features_activities" +
		" ADD CONSTRAINT features_activities_unique_key UNIQUE (activity_id, feature_id)")
}

func createFKey() {

	// ELEMENTOS
	db.Exec("ALTER TABLE ONLY public.elementos" +
		" ADD CONSTRAINT users_fkey FOREIGN KEY (author_id)" +
		" REFERENCES public.users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY public.elementos" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES public.status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// COMPONENTES
	db.Exec("ALTER TABLE ONLY public.componentes" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES public.users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY public.componentes" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES public.status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// PILARES
	db.Exec("ALTER TABLE ONLY public.pilares" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES public.users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY public.pilares" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES public.status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// CICLOS
	db.Exec("ALTER TABLE ONLY public.ciclos" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES public.users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY public.ciclos" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES public.status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// ENTIDADES
	db.Exec("ALTER TABLE ONLY public.entidades" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES public.users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY public.entidades" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES public.status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// PLANOS
	db.Exec("ALTER TABLE ONLY public.planos" +
		" ADD CONSTRAINT entidades_fkey FOREIGN KEY (entidade_id)" +
		" REFERENCES public.entidades (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY public.planos" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES public.users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY public.planos" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES public.status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// CICLOS ENTIDADES
	db.Exec("ALTER TABLE ONLY public.ciclos_entidades" +
		" ADD CONSTRAINT entidades_fkey FOREIGN KEY (entidade_id)" +
		" REFERENCES public.entidades (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// EQUIPES
	db.Exec("ALTER TABLE ONLY public.escritorios" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES public.users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY public.escritorios" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES public.status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// CARTEIRAS
	db.Exec("ALTER TABLE ONLY public.carteiras" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES public.users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY public.carteiras" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES public.status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// ACTIVITIES
	db.Exec("ALTER TABLE public.activities ADD CONSTRAINT action_fkey FOREIGN KEY (action_id)" +
		" REFERENCES public.actions (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE public.activities ADD CONSTRAINT expiration_action_fkey FOREIGN KEY (expiration_action_id)" +
		" REFERENCES public.actions (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE public.activities ADD CONSTRAINT workflow_fkey FOREIGN KEY (workflow_id)" +
		" REFERENCES public.workflows (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.actions " +
		" ADD CONSTRAINT destination_status_fkey FOREIGN KEY (destination_status_id)" +
		" REFERENCES public.status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.actions " +
		" ADD CONSTRAINT origin_status_fkey FOREIGN KEY (origin_status_id)" +
		" REFERENCES public.status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.actions " +
		" ADD CONSTRAINT workflows_fkey FOREIGN KEY (workflow_id)" +
		" REFERENCES public.workflows (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.features_roles " +
		" ADD CONSTRAINT features_fkey FOREIGN KEY (feature_id)" +
		" REFERENCES public.features (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.features_roles " +
		" ADD CONSTRAINT roles_fkey FOREIGN KEY (role_id)" +
		" REFERENCES public.roles (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.users " +
		" ADD CONSTRAINT roles_fkey FOREIGN KEY (role_id)" +
		" REFERENCES public.roles (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.activities_roles " +
		" ADD CONSTRAINT activities_fkey FOREIGN KEY (activity_id)" +
		" REFERENCES public.activities (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.activities_roles " +
		" ADD CONSTRAINT roles_fkey FOREIGN KEY (role_id)" +
		" REFERENCES public.roles (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.features_activities " +
		" ADD CONSTRAINT activities_fkey FOREIGN KEY (activity_id)" +
		" REFERENCES public.activities (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.features_activities " +
		" ADD CONSTRAINT features_fkey FOREIGN KEY (feature_id)" +
		" REFERENCES public.features (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.actions_status " +
		" ADD CONSTRAINT actions_fkey FOREIGN KEY (action_id)" +
		" REFERENCES public.actions (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.actions_status " +
		" ADD CONSTRAINT origin_status_fkey FOREIGN KEY (origin_status_id)" +
		" REFERENCES public.status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.actions_status " +
		" ADD CONSTRAINT destination_status_fkey FOREIGN KEY (destination_status_id)" +
		" REFERENCES public.status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")
}

func createPKey() {
	db.Exec("ALTER TABLE ONLY public.activities ADD CONSTRAINT activities_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.roles ADD CONSTRAINT roles_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.entidades ADD CONSTRAINT entidades_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.planos ADD CONSTRAINT planos_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.ciclos ADD CONSTRAINT ciclos_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.pilares ADD CONSTRAINT pilares_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.componentes ADD CONSTRAINT componentes_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.elementos ADD CONSTRAINT elementos_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.features ADD CONSTRAINT features_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.itens ADD CONSTRAINT itens_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.users ADD CONSTRAINT users_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.workflows ADD CONSTRAINT workflows_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.status ADD CONSTRAINT status_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.actions ADD CONSTRAINT actions_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.actions_status ADD CONSTRAINT actions_status_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.activities_roles ADD CONSTRAINT activities_roles_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.features_activities ADD CONSTRAINT features_activities_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.features_roles ADD CONSTRAINT features_roles_pkey PRIMARY KEY (id)")
}

func createStatusZERO() {
	query := "INSERT INTO status (id, name, stereotype, description, author_id, created_at)" +
		" SELECT 0, '-', '', '', 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM status WHERE id = 0)"
	log.Println(query)
	db.Exec(query)
}

func createEscritorios() {
	sql := "INSERT INTO public.escritorios( " +
		" id, nome, descricao, author_id, criado_em) " +
		" SELECT 1, 'Escritório de Representação - Pernambuco', 'Escritório de Representação - Pernambuco', 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM escritorios WHERE id = 1)"
	db.Exec(sql)
	sql = "INSERT INTO public.escritorios( " +
		" id, nome, descricao, author_id, criado_em) " +
		" SELECT 2, 'Escritório de Representação - São Paulo', 'Escritório de Representação - São Paulo', 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM escritorios WHERE id = 2)"
	db.Exec(sql)
	sql = "INSERT INTO public.escritorios( " +
		" id, nome, descricao, author_id, criado_em) " +
		" SELECT 3, 'Escritório de Representação - Minas Gerais', 'Escritório de Representação - Minas Gerais', 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM escritorios WHERE id = 3)"
	db.Exec(sql)
	sql = "INSERT INTO public.escritorios( " +
		" id, nome, descricao, author_id, criado_em) " +
		" SELECT 4, 'Escritório de Representação - Paraná', 'Escritório de Representação - Paraná', 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM escritorios WHERE id = 4)"
	db.Exec(sql)
	sql = "INSERT INTO public.escritorios( " +
		" id, nome, descricao, author_id, criado_em) " +
		" SELECT 5, 'Escritório de Representação - Alagoas', 'Escritório de Representação - Alagoas', 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM escritorios WHERE id = 5)"
	db.Exec(sql)
}
func createUsers() {
	sql := "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 1, 'aria', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'aria@gmail.com', '61 984385415', 'Ária Ohashi', 1, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'aria')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 2, 'masaru', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'masaru@gmail.com', '61 984385415', 'Masaru Ohashi Jr', 1, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'masaru')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 3, 'arnaldo', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'arnaldo@gmail.com', '61 984385415', 'Arnaldo Burle', 1, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'arnaldo')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 4, 'ana', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'ana@gmail.com', '61 984385415', 'Ana Carolina Baasch', 2, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'ana')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 5, 'annette', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'annette@gmail.com', '61 984385415', 'Annette Lopes Pinto ', 2, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'annette')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 6, 'carlos', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'carlos@gmail.com', '61 984385415', 'Carlos Marne Dias Alves', 2, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'carlos')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 7, 'christian', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'christian@gmail.com', '61 984385415', 'Christian Aggensteiner Catunda', 2, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'christian')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 8, 'dagomar', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'dagomar@gmail.com', '61 984385415', 'Dagomar Alécio Anhê', 2, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'dagomar')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 9, 'david', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'david@gmail.com', '61 984385415', 'David Prates Coutinho', 3, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'david')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 10, 'elthon', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'elthon@gmail.com', '61 984385415', 'Elthon Baier Nunes', 3, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'elthon')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 11, 'fabio', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'fábio@gmail.com', '61 984385415', 'Fábio Lucas de Albuquerque Lima', 3, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'fabio')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 12, 'fabricio', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'fabricio@gmail.com', '61 984385415', 'Fabricio Cardoso de Meneses', 3, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'fabricio')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 13, 'felipe', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'felipe@gmail.com', '61 984385415', 'Felipe Spolavori Martins', 3, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'felipe')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 14, 'fernando', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'fernando@gmail.com', '61 984385415', 'Fernando Duarte Folle', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'fernando')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 15, 'hilton', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'hilton@gmail.com', '61 984385415', 'Hilton de Enzo Mitsunaga', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'hilton')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 16, 'chedeak', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'josé@gmail.com', '61 984385415', 'José Carlos Sampaio Chedeak', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'chedeak')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 17, 'jose', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'josé@gmail.com', '61 984385415', 'José de Arimatéria Pinheiro Torres', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'jose')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 18, 'luciano', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'luciano@gmail.com', '61 984385415', 'Luciano Draghetti', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'luciano')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 19, 'lucio', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'lucio@gmail.com', '61 984385415', 'Lucio Rodrigues Capelletto', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'lucio')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 20, 'luis', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'luis@gmail.com', '61 984385415', 'Luis Ronaldo Martins Angoti', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'luis')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 21, 'manoel', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'manoel@gmail.com', '61 984385415', 'Manoel Robson Aguiar', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'manoel')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 22, 'mauricio', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'maurício@gmail.com', '61 984385415', 'Maurício de Aguirre Nakata', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'mauricio')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 23, 'milton', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'milton@gmail.com', '61 984385415', 'Milton Santos', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'milton')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 24, 'otavio', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'otávio@gmail.com', '61 984385415', 'Otávio Lima Reis', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'otavio')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 25, 'paulo', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'paulo@gmail.com', '61 984385415', 'Paulo Roberto Pereira De Macedo', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'paulo')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 26, 'peterson', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'peterson@gmail.com', '61 984385415', 'Peterson Gonçalves', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'peterson')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 27, 'rita', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'rita@gmail.com', '61 984385415', 'Rita de Cassia Correa da Silva', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'rita')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 28, 'sergio', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'sergio@gmail.com', '61 984385415', 'Sergio Djundi Taniguchi', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'sergio')"
	db.Exec(sql)
}

func createRoles() {
	query := " INSERT INTO roles (id, name, description, created_at) " +
		" SELECT 1, 'Admin', 'Admin' , now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM roles WHERE name = 'Admin')"
	db.Exec(query)
	query = " INSERT INTO roles (id, name, description, created_at) " +
		" SELECT 2, 'Chefe', 'Chefe' , now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM roles WHERE name = 'Chefe')"
	db.Exec(query)
	query = " INSERT INTO roles (id, name, description, created_at) " +
		" SELECT 3, 'Supervisor', 'Supervisor' , now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM roles WHERE name = 'Supervisor')"
	db.Exec(query)
	query = " INSERT INTO roles (id, name, description, created_at) " +
		" SELECT 4, 'Auditor', 'Auditor' , now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM roles WHERE name = 'Auditor')"
	db.Exec(query)
}

func updateRoles() {
	query := " UPDATE roles SET author_id = 1 WHERE name = 'Admin' AND (SELECT author_id FROM roles WHERE name = 'Admin') IS NULL "
	log.Println(query)
	db.Exec(query)
}

func updateFeatures() {
	query := " UPDATE features SET author_id = 1, created_at = now()::timestamp WHERE id <= 30"
	log.Println(query)
	db.Exec(query)
}

func createRoleFeatures() {
	query := " INSERT INTO features_roles (role_id, feature_id) VALUES (1, 1) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1, 2) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1, 3) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1, 4) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1, 5) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1, 6) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1, 7) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1, 8) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1, 9) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,10) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,11) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,12) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,13) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,14) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,15) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,16) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,17) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,18) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,19) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,20) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,21) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,22) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,23) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,24) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,25) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,26) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,27) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,28) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,29) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,30) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2, 1) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2, 2) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2, 3) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2, 4) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2, 5) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2, 6) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2, 7) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2, 8) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2, 9) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2,10) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2,11) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2,12) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2,13) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2,14) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2,15) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2,16) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2,17) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2,18) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2,19) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2,20) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2,21) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2,22) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2,23) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2,24) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2,25) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2,26) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2,27) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2,28) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2,29) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (2,30) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3, 1) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3, 2) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3, 3) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3, 4) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3, 5) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3, 6) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3, 7) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3, 8) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3, 9) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3,10) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3,11) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3,12) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3,13) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3,14) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3,15) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3,16) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3,17) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3,18) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3,19) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3,20) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3,21) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3,22) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3,23) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3,24) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3,25) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3,26) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3,27) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3,28) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3,29) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (3,30) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4, 1) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4, 2) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4, 3) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4, 4) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4, 5) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4, 6) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4, 7) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4, 8) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4, 9) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4,10) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4,11) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4,12) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4,13) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4,14) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4,15) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4,16) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4,17) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4,18) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4,19) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4,20) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4,21) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4,22) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4,23) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4,24) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4,25) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4,26) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4,27) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4,28) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4,29) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (4,30) "
	//log.Println(query)
	db.Exec(query)
}

func createSeq() {
	// Sequence ACTIONS_STATUS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.activities_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ACTIONS_STATUS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.actions_status_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ACTIVITIES_ROLES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.activities_roles_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence FEATURES_ACTIVITIES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.features_activities_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence FEATURES_ROLES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.features_roles_id_seq " +
		" START WITH 2" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence WORKFLOWS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.workflows_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence STATUS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.status_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ACTIONS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.actions_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ROLES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.roles_id_seq " +
		" START WITH 2" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence FEATURES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.features_id_seq " +
		" START WITH 18" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence USERS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.users_id_seq " +
		" START WITH 2" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ITENS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.itens_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ELEMENTOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.elementos_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence COMPONENTES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.componentes_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence PILARES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.pilares_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence PILARES_CICLOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.pilares_ciclos_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence CICLOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.ciclos_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence CICLOS_ENTIDADES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.ciclos_entidades_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ENTIDADES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.entidades_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence PLANOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.planos_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ESCRITORIOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.escritorios_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence CARTEIRAS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.carteiras_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
}

func createTable() {
	// Table ACTIONS
	db.Exec(" CREATE TABLE IF NOT EXISTS public.actions (" +
		" id integer DEFAULT nextval('public.actions_id_seq'::regclass) NOT NULL, " +
		" name character varying(255) NOT NULL, " +
		" origin_status_id integer, " +
		" destination_status_id integer, " +
		" other_than boolean, " +
		" description character varying(4000)," +
		" author_id integer," +
		" created_at timestamp without time zone," +
		" id_versao_origem integer," +
		" status_id integer)")

	// Table ACTIONS_STATUS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.actions_status (" +
			" id integer DEFAULT nextval('actions_status_id_seq'::regclass)," +
			" action_id integer," +
			" origin_status_id integer," +
			" destination_status_id integer)")

	// Table ACTIVITIES
	db.Exec(" CREATE TABLE public.activities (" +
		" id integer NOT NULL DEFAULT nextval('activities_id_seq'::regclass)," +
		" workflow_id integer," +
		" action_id integer," +
		" expiration_action_id integer," +
		" expiration_time_days integer," +
		" start_at timestamp without time zone," +
		" end_at timestamp without time zone)")

	// Table ACTIVITIES_ROLES
	db.Exec(
		" CREATE TABLE public.activities_roles (" +
			" id integer DEFAULT nextval('activities_roles_id_seq'::regclass)," +
			" activity_id integer," +
			" role_id integer)")

	// Table CICLOS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.ciclos (" +
			" id integer DEFAULT nextval('public.ciclos_id_seq'::regclass) NOT NULL," +
			" nome character varying(255) NOT NULL," +
			" descricao character varying(4000)," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table CICLOS_ENTIDADES
	db.Exec(
		" CREATE TABLE public.ciclos_entidades (" +
			" id integer DEFAULT nextval('ciclos_entidades_id_seq'::regclass)," +
			" ciclo_id integer," +
			" entidade_id integer," +
			" tipo_media integer," +
			" nota double precision," +
			" inicia_em timestamp without time zone," +
			" termina_em timestamp without time zone," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table COMPONENTES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.componentes (" +
			" id integer DEFAULT nextval('public.componentes_id_seq'::regclass) NOT NULL," +
			" nome character varying(255) NOT NULL," +
			" descricao character varying(4000)," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table ELEMENTOS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.elementos (" +
			" id integer DEFAULT nextval('public.elementos_id_seq'::regclass) NOT NULL," +
			" nome character varying(255) NOT NULL," +
			" descricao character varying(4000)," +
			" peso integer DEFAULT 1 NOT NULL," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" status_id integer)")

	// Table ENTIDADES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.entidades (" +
			" id integer DEFAULT nextval('public.entidades_id_seq'::regclass) NOT NULL," +
			" nome character varying(255) NOT NULL," +
			" descricao character varying(4000)," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table ESCRITORIOS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.escritorios (" +
			" id integer DEFAULT nextval('public.escritorios_id_seq'::regclass) NOT NULL," +
			" nome character varying(255) NOT NULL," +
			" descricao character varying(4000)," +
			" chefe_id integer," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table FEATURES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.features  (" +
			" id integer DEFAULT nextval('public.features_id_seq'::regclass) NOT NULL," +
			" name character varying(255) NOT NULL," +
			" code character varying(255) NOT NULL," +
			" description character varying(4000)," +
			" author_id integer," +
			" created_at timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table FEATURES_ROLES
	db.Exec(
		" CREATE TABLE public.features_roles (" +
			" id integer DEFAULT nextval('features_roles_id_seq'::regclass)," +
			" feature_id integer," +
			" role_id integer)")

	// Table FEATURES_ACTIVITIES
	db.Exec(
		" CREATE TABLE public.features_activities (" +
			" id integer DEFAULT nextval('features_activities_id_seq'::regclass)," +
			" feature_id integer," +
			" activity_id integer)")

	// Table ITENS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.itens (" +
			" id integer DEFAULT nextval('public.itens_id_seq'::regclass) NOT NULL," +
			" elemento_id integer," +
			" nome character varying(255) NOT NULL," +
			" descricao character varying(4000)," +
			" avaliacao character varying(4000)," +
			" criado_em timestamp without time zone," +
			" author_id integer," +
			" status_id integer)")

	// Table NOTAS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.notas (" +
			" id integer DEFAULT nextval('public.notas_id_seq'::regclass) NOT NULL," +
			" elemento_id integer," +
			" tipo_nota_id integer," +
			" nota double precision," +
			" author_id integer," +
			" criado_em timestamp without time zone )")

	// Table PILARES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.pilares (" +
			" id integer DEFAULT nextval('public.pilares_id_seq'::regclass) NOT NULL," +
			" nome character varying(255) NOT NULL," +
			" descricao character varying(4000)," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table PILARES_CICLOS
	db.Exec(
		" CREATE TABLE public.pilares_ciclos (" +
			" id integer DEFAULT nextval('pilares_ciclos_id_seq'::regclass)," +
			" pilar_id integer," +
			" ciclo_id integer," +
			" tipo_media integer," +
			" peso_padrao integer," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table PLANOS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.planos (" +
			" id integer DEFAULT nextval('public.planos_id_seq'::regclass) NOT NULL," +
			" entidade_id integer," +
			" nome character varying(255) NOT NULL," +
			" descricao character varying(4000)," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table ROLES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.roles  (" +
			" id integer DEFAULT nextval('public.roles_id_seq'::regclass) NOT NULL," +
			" name character varying(255) NOT NULL," +
			" description character varying(4000)," +
			" author_id integer," +
			" created_at timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table STATUS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.status  (" +
			" id integer DEFAULT nextval('public.status_id_seq'::regclass) NOT NULL," +
			" name character varying(255) NOT NULL," +
			" description character varying(4000)," +
			" author_id integer," +
			" created_at timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer," +
			" stereotype character varying(255))")

	// Table USERS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.users (" +
			" id integer DEFAULT nextval('public.users_id_seq'::regclass) NOT NULL," +
			" name character varying(255)," +
			" username character varying(255) NOT NULL," +
			" password character varying(255) NOT NULL," +
			" email character varying(255)," +
			" mobile character varying(255)," +
			" role_id integer," +
			" escritorio_id integer," +
			" author_id integer," +
			" criado_em timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")

	// Table WORKFLOWS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.workflows  (" +
			" id integer DEFAULT nextval('public.workflows_id_seq'::regclass) NOT NULL," +
			" name character varying(255) NOT NULL," +
			" description character varying(4000)," +
			" entity_type character varying(50)," +
			" start_at timestamp without time zone," +
			" end_at timestamp without time zone," +
			" author_id integer," +
			" created_at timestamp without time zone," +
			" id_versao_origem integer," +
			" status_id integer)")
}
