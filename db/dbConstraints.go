package db

import ()

func createUniqueKey() {
	db.Exec(" ALTER TABLE ONLY actions_status" +
		" ADD CONSTRAINT action_status_unique_key UNIQUE (action_id, origin_status_id, destination_status_id)")
	db.Exec(" ALTER TABLE ONLY features_roles" +
		" ADD CONSTRAINT feature_role_unique_key UNIQUE (role_id, feature_id)")
	db.Exec(" ALTER TABLE ONLY users" +
		" ADD CONSTRAINT username_unique_key UNIQUE (username)")
	db.Exec(" ALTER TABLE ONLY activities_roles" +
		" ADD CONSTRAINT action_role_unique_key UNIQUE (activity_id, role_id)")
	db.Exec(" ALTER TABLE ONLY features_activities" +
		" ADD CONSTRAINT features_activities_unique_key UNIQUE (activity_id, feature_id)")
	db.Exec(" ALTER TABLE ONLY jurisdicoes" +
		" ADD CONSTRAINT jurisdicoes_unique_key UNIQUE (escritorio_id, entidade_id)")
	db.Exec(" ALTER TABLE ONLY membros" +
		" ADD CONSTRAINT membros_unique_key UNIQUE (escritorio_id, usuario_id)")
	db.Exec(" ALTER TABLE ONLY ciclos_entidades" +
		" ADD CONSTRAINT ciclos_entidades_unique_key UNIQUE (entidade_id, ciclo_id)")
	db.Exec(" ALTER TABLE ONLY pilares_ciclos" +
		" ADD CONSTRAINT pilares_ciclos_unique_key UNIQUE (ciclo_id, pilar_id)")
	db.Exec(" ALTER TABLE ONLY componentes_pilares" +
		" ADD CONSTRAINT componentes_pilares_unique_key UNIQUE (pilar_id, componente_id)")
	db.Exec(" ALTER TABLE ONLY elementos_componentes" +
		" ADD CONSTRAINT elementos_componentes_unique_key UNIQUE (componente_id, elemento_id)")
	db.Exec(" ALTER TABLE ONLY tipos_notas_componentes" +
		" ADD CONSTRAINT tipos_notas_componentes_unique_key UNIQUE (componente_id, tipo_nota_id)")
	db.Exec(" ALTER TABLE ONLY produtos_ciclos" +
		" ADD CONSTRAINT produtos_ciclos_unique_key UNIQUE (entidade_id, ciclo_id)")
	db.Exec(" ALTER TABLE ONLY produtos_pilares" +
		" ADD CONSTRAINT produtos_pilares_unique_key UNIQUE (entidade_id, ciclo_id, pilar_id)")
	db.Exec(" ALTER TABLE ONLY produtos_componentes" +
		" ADD CONSTRAINT produtos_componentes_unique_key UNIQUE (entidade_id, ciclo_id, pilar_id, componente_id)")
	db.Exec(" ALTER TABLE ONLY produtos_tipos_notas" +
		" ADD CONSTRAINT produtos_tipos_notas_unique_key UNIQUE (entidade_id, ciclo_id, pilar_id, componente_id, tipo_nota_id)")
	db.Exec(" ALTER TABLE ONLY produtos_elementos" +
		" ADD CONSTRAINT produtos_elementos_unique_key UNIQUE (entidade_id, ciclo_id, pilar_id, componente_id, tipo_nota_id, elemento_id)")
	db.Exec(" ALTER TABLE ONLY produtos_itens" +
		" ADD CONSTRAINT produtos_itens_unique_key UNIQUE (entidade_id, ciclo_id, pilar_id, componente_id, elemento_id, item_id)")
}

func createFKey() {
	// ACTIONS
	db.Exec("ALTER TABLE ONLY actions " +
		" ADD CONSTRAINT destination_status_fkey FOREIGN KEY (destination_status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY actions " +
		" ADD CONSTRAINT origin_status_fkey FOREIGN KEY (origin_status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY actions " +
		" ADD CONSTRAINT workflows_fkey FOREIGN KEY (workflow_id)" +
		" REFERENCES workflows (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	//  ACTIONS_STATUS
	db.Exec("ALTER TABLE ONLY actions_status " +
		" ADD CONSTRAINT actions_fkey FOREIGN KEY (action_id)" +
		" REFERENCES actions (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY actions_status " +
		" ADD CONSTRAINT origin_status_fkey FOREIGN KEY (origin_status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY actions_status " +
		" ADD CONSTRAINT destination_status_fkey FOREIGN KEY (destination_status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	// ACTIVITIES
	db.Exec("ALTER TABLE activities ADD CONSTRAINT action_fkey FOREIGN KEY (action_id)" +
		" REFERENCES actions (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE activities ADD CONSTRAINT expiration_action_fkey FOREIGN KEY (expiration_action_id)" +
		" REFERENCES actions (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE activities ADD CONSTRAINT workflow_fkey FOREIGN KEY (workflow_id)" +
		" REFERENCES workflows (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	// ACTIVITIES_ROLES
	db.Exec("ALTER TABLE ONLY activities_roles " +
		" ADD CONSTRAINT activities_fkey FOREIGN KEY (activity_id)" +
		" REFERENCES activities (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY activities_roles " +
		" ADD CONSTRAINT roles_fkey FOREIGN KEY (role_id)" +
		" REFERENCES roles (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	// CICLOS
	db.Exec("ALTER TABLE ONLY ciclos" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY ciclos" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// CICLOS ENTIDADES
	db.Exec("ALTER TABLE ONLY ciclos_entidades" +
		" ADD CONSTRAINT entidades_fkey FOREIGN KEY (entidade_id)" +
		" REFERENCES entidades (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY ciclos_entidades" +
		" ADD CONSTRAINT ciclos_fkey FOREIGN KEY (ciclo_id)" +
		" REFERENCES ciclos (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY ciclos_entidades" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY ciclos_entidades" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// COMPONENTES
	db.Exec("ALTER TABLE ONLY componentes" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY componentes" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// COMPONENTES PILARES
	db.Exec("ALTER TABLE ONLY componentes_pilares" +
		" ADD CONSTRAINT componentes_fkey FOREIGN KEY (componente_id)" +
		" REFERENCES componentes (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY componentes_pilares" +
		" ADD CONSTRAINT pilares_fkey FOREIGN KEY (pilar_id)" +
		" REFERENCES pilares (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY componentes_pilares" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY componentes_pilares" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// ELEMENTOS
	db.Exec("ALTER TABLE ONLY elementos" +
		" ADD CONSTRAINT users_fkey FOREIGN KEY (author_id)" +
		" REFERENCES users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY elementos" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// ELEMENTOS_COMPONENTES
	db.Exec("ALTER TABLE ONLY elementos_componentes" +
		" ADD CONSTRAINT tipos_notas_fkey FOREIGN KEY (tipo_nota_id)" +
		" REFERENCES tipos_notas (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY elementos_componentes" +
		" ADD CONSTRAINT elementos_fkey FOREIGN KEY (elemento_id)" +
		" REFERENCES elementos (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY elementos_componentes" +
		" ADD CONSTRAINT componentes_fkey FOREIGN KEY (componente_id)" +
		" REFERENCES componentes (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY elementos_componentes" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY elementos_componentes" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// ENTIDADES
	db.Exec("ALTER TABLE ONLY entidades" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY entidades" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// ESCRITÓRIOS
	db.Exec("ALTER TABLE ONLY escritorios" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY escritorios" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// FEATURES_ACTIVITIES
	db.Exec("ALTER TABLE ONLY features_activities " +
		" ADD CONSTRAINT activities_fkey FOREIGN KEY (activity_id)" +
		" REFERENCES activities (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY features_activities " +
		" ADD CONSTRAINT features_fkey FOREIGN KEY (feature_id)" +
		" REFERENCES features (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	// FEATURES_ROLES
	db.Exec("ALTER TABLE ONLY features_roles " +
		" ADD CONSTRAINT features_fkey FOREIGN KEY (feature_id)" +
		" REFERENCES features (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY features_roles " +
		" ADD CONSTRAINT roles_fkey FOREIGN KEY (role_id)" +
		" REFERENCES roles (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	// ITENS
	db.Exec("ALTER TABLE ONLY itens" +
		" ADD CONSTRAINT elementos_fkey FOREIGN KEY (elemento_id)" +
		" REFERENCES elementos (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	// PILARES
	db.Exec("ALTER TABLE ONLY pilares" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY pilares" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// PILARES_CICLOS
	db.Exec("ALTER TABLE ONLY pilares_ciclos" +
		" ADD CONSTRAINT pilares_fkey FOREIGN KEY (pilar_id)" +
		" REFERENCES pilares (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY pilares_ciclos" +
		" ADD CONSTRAINT ciclos_fkey FOREIGN KEY (ciclo_id)" +
		" REFERENCES ciclos (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY pilares_ciclos" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY pilares_ciclos" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// PLANOS
	db.Exec("ALTER TABLE ONLY planos" +
		" ADD CONSTRAINT entidades_fkey FOREIGN KEY (entidade_id)" +
		" REFERENCES entidades (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY planos" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY planos" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// PRODUTOS_CICLOS
	db.Exec("ALTER TABLE ONLY produtos_ciclos" +
		" ADD CONSTRAINT entidades_fkey FOREIGN KEY (entidade_id)" +
		" REFERENCES entidades (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_ciclos" +
		" ADD CONSTRAINT ciclos_fkey FOREIGN KEY (ciclo_id)" +
		" REFERENCES ciclos (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_ciclos" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_ciclos" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// PRODUTOS_PILARES
	db.Exec("ALTER TABLE ONLY produtos_pilares" +
		" ADD CONSTRAINT entidades_fkey FOREIGN KEY (entidade_id)" +
		" REFERENCES entidades (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_pilares" +
		" ADD CONSTRAINT ciclos_fkey FOREIGN KEY (ciclo_id)" +
		" REFERENCES ciclos (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_pilares" +
		" ADD CONSTRAINT pilares_fkey FOREIGN KEY (pilar_id)" +
		" REFERENCES pilares (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_pilares" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_pilares" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// PRODUTOS_COMPONENTES
	db.Exec("ALTER TABLE ONLY produtos_componentes" +
		" ADD CONSTRAINT entidades_fkey FOREIGN KEY (entidade_id)" +
		" REFERENCES entidades (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_componentes" +
		" ADD CONSTRAINT ciclos_fkey FOREIGN KEY (ciclo_id)" +
		" REFERENCES ciclos (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_componentes" +
		" ADD CONSTRAINT pilares_fkey FOREIGN KEY (pilar_id)" +
		" REFERENCES pilares (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_componentes" +
		" ADD CONSTRAINT componentes_fkey FOREIGN KEY (componente_id)" +
		" REFERENCES componentes (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_componentes" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_componentes" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// PRODUTOS_TIPOS_NOTAS
	db.Exec("ALTER TABLE ONLY produtos_tipos_notas" +
		" ADD CONSTRAINT entidades_fkey FOREIGN KEY (entidade_id)" +
		" REFERENCES entidades (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_tipos_notas" +
		" ADD CONSTRAINT ciclos_fkey FOREIGN KEY (ciclo_id)" +
		" REFERENCES ciclos (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_tipos_notas" +
		" ADD CONSTRAINT pilares_fkey FOREIGN KEY (pilar_id)" +
		" REFERENCES pilares (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_tipos_notas" +
		" ADD CONSTRAINT componentes_fkey FOREIGN KEY (componente_id)" +
		" REFERENCES componentes (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_tipos_notas" +
		" ADD CONSTRAINT tipos_notas_fkey FOREIGN KEY (tipo_nota_id)" +
		" REFERENCES tipos_notas (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_tipos_notas" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_tipos_notas" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// PRODUTOS_ELEMENTOS
	db.Exec("ALTER TABLE ONLY produtos_elementos" +
		" ADD CONSTRAINT entidades_fkey FOREIGN KEY (entidade_id)" +
		" REFERENCES entidades (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_elementos" +
		" ADD CONSTRAINT ciclos_fkey FOREIGN KEY (ciclo_id)" +
		" REFERENCES ciclos (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_elementos" +
		" ADD CONSTRAINT pilares_fkey FOREIGN KEY (pilar_id)" +
		" REFERENCES pilares (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_elementos" +
		" ADD CONSTRAINT componentes_fkey FOREIGN KEY (componente_id)" +
		" REFERENCES componentes (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_elementos" +
		" ADD CONSTRAINT tipos_notas_fkey FOREIGN KEY (tipo_nota_id)" +
		" REFERENCES tipos_notas (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_elementos" +
		" ADD CONSTRAINT elementos_fkey FOREIGN KEY (elemento_id)" +
		" REFERENCES elementos (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_elementos" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_elementos" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// PRODUTOS_ITENS
	db.Exec("ALTER TABLE ONLY produtos_itens" +
		" ADD CONSTRAINT entidades_fkey FOREIGN KEY (entidade_id)" +
		" REFERENCES entidades (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_itens" +
		" ADD CONSTRAINT ciclos_fkey FOREIGN KEY (ciclo_id)" +
		" REFERENCES ciclos (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_itens" +
		" ADD CONSTRAINT pilares_fkey FOREIGN KEY (pilar_id)" +
		" REFERENCES pilares (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_itens" +
		" ADD CONSTRAINT componentes_fkey FOREIGN KEY (componente_id)" +
		" REFERENCES componentes (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_itens" +
		" ADD CONSTRAINT elementos_fkey FOREIGN KEY (elemento_id)" +
		" REFERENCES elementos (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_itens" +
		" ADD CONSTRAINT itens_fkey FOREIGN KEY (item_id)" +
		" REFERENCES itens (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_itens" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY produtos_itens" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// TIPOS_NOTAS
	db.Exec("ALTER TABLE ONLY tipos_notas" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY tipos_notas" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// TIPOS_NOTAS_COMPONENTES
	db.Exec("ALTER TABLE ONLY tipos_notas_componentes" +
		" ADD CONSTRAINT tipos_notas_fkey FOREIGN KEY (tipo_nota_id)" +
		" REFERENCES tipos_notas (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")
	db.Exec("ALTER TABLE ONLY tipos_notas_componentes" +
		" ADD CONSTRAINT componentes_fkey FOREIGN KEY (componente_id)" +
		" REFERENCES componentes (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")
	db.Exec("ALTER TABLE ONLY tipos_notas" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")
	db.Exec("ALTER TABLE ONLY tipos_notas" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// USERS
	db.Exec("ALTER TABLE ONLY users " +
		" ADD CONSTRAINT roles_fkey FOREIGN KEY (role_id)" +
		" REFERENCES roles (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY users" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY users" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// JURISDICOES
	db.Exec("ALTER TABLE ONLY jurisdicoes" +
		" ADD CONSTRAINT entidades_fkey FOREIGN KEY (entidade_id)" +
		" REFERENCES entidades (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY jurisdicoes" +
		" ADD CONSTRAINT escritorios_fkey FOREIGN KEY (escritorio_id)" +
		" REFERENCES escritorios (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY jurisdicoes" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY jurisdicoes" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	// MEMBROS
	db.Exec("ALTER TABLE ONLY membros" +
		" ADD CONSTRAINT usuarios_fkey FOREIGN KEY (usuario_id)" +
		" REFERENCES users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY membros" +
		" ADD CONSTRAINT escritorios_fkey FOREIGN KEY (escritorio_id)" +
		" REFERENCES escritorios (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY membros" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY membros" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

}

func createPKey() {
	db.Exec("ALTER TABLE ONLY actions ADD CONSTRAINT actions_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY actions_status ADD CONSTRAINT actions_status_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY activities ADD CONSTRAINT activities_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY activities_roles ADD CONSTRAINT activities_roles_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY ciclos ADD CONSTRAINT ciclos_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY componentes ADD CONSTRAINT componentes_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY componentes_pilares ADD CONSTRAINT componentes_pilares_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY elementos ADD CONSTRAINT elementos_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY elementos_componentes ADD CONSTRAINT elementos_componentes_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY entidades ADD CONSTRAINT entidades_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY entidades_ciclos ADD CONSTRAINT entidades_ciclos_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY escritorios ADD CONSTRAINT escritorios_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY features ADD CONSTRAINT features_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY features_activities ADD CONSTRAINT features_activities_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY features_roles ADD CONSTRAINT features_roles_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY integrantes ADD CONSTRAINT integrantes_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY itens ADD CONSTRAINT itens_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY jurisdicoes ADD CONSTRAINT jurisdicoes_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY membros ADD CONSTRAINT membros_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY pilares ADD CONSTRAINT pilares_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY pilares_ciclos ADD CONSTRAINT pilares_ciclos_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY planos ADD CONSTRAINT planos_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY produtos_ciclos ADD CONSTRAINT produtos_ciclos_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY produtos_componentes ADD CONSTRAINT produtos_componentes_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY produtos_elementos ADD CONSTRAINT produtos_elementos_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY produtos_itens ADD CONSTRAINT produtos_itens_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY produtos_pilares ADD CONSTRAINT produtos_pilares_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY roles ADD CONSTRAINT roles_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY status ADD CONSTRAINT status_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY tipos_notas ADD CONSTRAINT tipos_notas_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY tipos_notas_componentes ADD CONSTRAINT tipos_notas_componentes_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY users ADD CONSTRAINT users_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY workflows ADD CONSTRAINT workflows_pkey PRIMARY KEY (id)")
}
