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
	createRoleAdmin()
	createRoleFeatures()
	createStatusZERO()
	createAdmin()
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
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (19, 'Listar Matrizes', 'listMatrizes')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (20, 'Criar Matriz', 'createMatriz')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (21, 'Listar Ciclos', 'listCiclos')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (22, 'Criar Ciclo', 'createCiclo')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (23, 'Listar Entidades', 'listEntidades')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (24, 'Criar Entidade', 'createEntidade')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (25, 'Listar Planos', 'listPlanos')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (26, 'Criar Plano', 'createPlano')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (27, 'Listar Equipes', 'listEquipes')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (28, 'Criar Equipe', 'createEquipe')")
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

	// MATRIZES
	db.Exec("ALTER TABLE ONLY public.matrizes" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES public.users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY public.matrizes" +
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

	// EQUIPES
	db.Exec("ALTER TABLE ONLY public.equipes" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES public.users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY public.equipes" +
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
	db.Exec("ALTER TABLE ONLY public.matrizes ADD CONSTRAINT matrizes_pkey PRIMARY KEY (id)")
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
	query := "INSERT INTO status (id, name, stereotype)" +
		" SELECT 0, '-', '' " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE id = 0)"
	log.Println(query)
	db.Exec(query)
}

func createAdmin() {
	query := "INSERT INTO users (id, username, password, email, mobile, name, role_id)" +
		" SELECT 1, 'masaru', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'masaru@vindixit.com', '61 984385415', 'Masaru Ohashi Jr', 1" +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'masaru')"
	log.Println(query)
	db.Exec(query)
}

func createRoleAdmin() {
	query := " INSERT INTO roles (id, name) " +
		" SELECT 1, 'Admin' " +
		" WHERE NOT EXISTS (SELECT id FROM roles WHERE name = 'Admin')"
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
	// Sequence MATRIZES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.matrizes_id_seq " +
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
	// Sequence EQUIPES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.equipes_id_seq " +
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

	// Table WORKFLOWS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.workflows  (" +
			" id integer DEFAULT nextval('public.workflows_id_seq'::regclass) NOT NULL," +
			" name character varying(255) NOT NULL," +
			" entity_type character varying(50)," +
			" start_at timestamp without time zone," +
			" end_at timestamp without time zone)")

	// Table ACTIONS
	db.Exec(" CREATE TABLE IF NOT EXISTS public.actions (" +
		"id integer DEFAULT nextval('public.actions_id_seq'::regclass) NOT NULL, " +
		"name character varying(255) NOT NULL, " +
		"origin_status_id integer, " +
		"destination_status_id integer, " +
		"other_than boolean)")

	// Table STATUS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.status  (" +
			" id integer DEFAULT nextval('public.status_id_seq'::regclass) NOT NULL," +
			" name character varying(255) NOT NULL," +
			" stereotype character varying(255) NULL)")

	// Table ROLES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.roles  (" +
			" id integer DEFAULT nextval('public.roles_id_seq'::regclass) NOT NULL," +
			" name character varying(255) NOT NULL)")
	// Table FEATURES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.features  (" +
			" id integer DEFAULT nextval('public.features_id_seq'::regclass) NOT NULL," +
			" name character varying(255) NOT NULL," +
			" code character varying(255) NOT NULL)")
	// Table USERS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.users (" +
			" id integer DEFAULT nextval('public.users_id_seq'::regclass) NOT NULL," +
			" username character varying(255) NOT NULL," +
			" password character varying(255) NOT NULL," +
			" email character varying(255) NOT NULL," +
			" mobile character varying(255) NOT NULL," +
			" role_id integer NOT NULL," +
			" name character varying(255))")

	// Table NOTAS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.notas (" +
			" id integer DEFAULT nextval('public.notas_id_seq'::regclass) NOT NULL," +
			" elemento_id integer," +
			" tipo_nota_id integer," +
			" nota double precision," +
			" author_id integer," +
			" data_criacao timestamp without time zone )")

	// Table ITENS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.itens (" +
			" id integer DEFAULT nextval('public.itens_id_seq'::regclass) NOT NULL," +
			" elemento_id integer," +
			" titulo character varying(255) NOT NULL," +
			" descricao character varying(4000)," +
			" avaliacao character varying(4000)," +
			" data_criacao timestamp without time zone," +
			" author_id integer," +
			" status_id integer)")

	// Table ACTIONS_STATUS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.actions_status (" +
			" id integer DEFAULT nextval('actions_status_id_seq'::regclass)," +
			" action_id integer," +
			" origin_status_id integer," +
			" destination_status_id integer)")

	// Table ELEMENTOS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.elementos (" +
			" id integer DEFAULT nextval('public.elementos_id_seq'::regclass) NOT NULL," +
			" titulo character varying(255) NOT NULL," +
			" descricao character varying(4000)," +
			" peso integer DEFAULT 1 NOT NULL," +
			" author_id integer," +
			" data_criacao timestamp without time zone," +
			" status_id integer)")

	// Table COMPONENTES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.componentes (" +
			" id integer DEFAULT nextval('public.componentes_id_seq'::regclass) NOT NULL," +
			" titulo character varying(255) NOT NULL," +
			" tipo_media character varying(20)," +
			" peso integer," +
			" author_id integer," +
			" data_criacao timestamp without time zone," +
			" status_id integer)")

	// Table MATRIZES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.matrizes (" +
			" id integer DEFAULT nextval('public.matrizes_id_seq'::regclass) NOT NULL," +
			" titulo character varying(255) NOT NULL," +
			" tipo_media character varying(20)," +
			" peso integer," +
			" author_id integer," +
			" data_criacao timestamp without time zone," +
			" status_id integer)")

	// Table CICLOS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.ciclos (" +
			" id integer DEFAULT nextval('public.ciclos_id_seq'::regclass) NOT NULL," +
			" titulo character varying(255) NOT NULL," +
			" tipo_media character varying(20)," +
			" peso integer," +
			" author_id integer," +
			" data_criacao timestamp without time zone," +
			" status_id integer)")

	// Table ENTIDADES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.entidades (" +
			" id integer DEFAULT nextval('public.entidades_id_seq'::regclass) NOT NULL," +
			" titulo character varying(255) NOT NULL," +
			" tipo_media character varying(20)," +
			" peso integer," +
			" author_id integer," +
			" data_criacao timestamp without time zone," +
			" status_id integer)")

	// Table PLANOS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.planos (" +
			" id integer DEFAULT nextval('public.planos_id_seq'::regclass) NOT NULL," +
			" titulo character varying(255) NOT NULL," +
			" tipo_media character varying(20)," +
			" peso integer," +
			" author_id integer," +
			" data_criacao timestamp without time zone," +
			" status_id integer)")

	// Table EQUIPES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.equipes (" +
			" id integer DEFAULT nextval('public.equipes_id_seq'::regclass) NOT NULL," +
			" titulo character varying(255) NOT NULL," +
			" tipo_media character varying(20)," +
			" peso integer," +
			" author_id integer," +
			" data_criacao timestamp without time zone," +
			" status_id integer)")

	// Table CARTEIRAS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.carteiras (" +
			" id integer DEFAULT nextval('public.carteiras_id_seq'::regclass) NOT NULL," +
			" titulo character varying(255) NOT NULL," +
			" tipo_media character varying(20)," +
			" peso integer," +
			" author_id integer," +
			" data_criacao timestamp without time zone," +
			" status_id integer)")
}
