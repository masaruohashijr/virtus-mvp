package db

import ()

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

	db.Exec(" ALTER TABLE ONLY public.jurisdicoes" +
		" ADD CONSTRAINT jurisdicoes_unique_key UNIQUE (escritorio_id, entidade_id)")

	db.Exec(" ALTER TABLE ONLY public.membros" +
		" ADD CONSTRAINT membros_unique_key UNIQUE (escritorio_id, usuario_id)")

	db.Exec(" ALTER TABLE ONLY public.ciclos_entidades" +
		" ADD CONSTRAINT membros_unique_key UNIQUE (entidade_id, ciclo_id)")

	db.Exec(" ALTER TABLE ONLY public.pilares_ciclos" +
		" ADD CONSTRAINT membros_unique_key UNIQUE (ciclo_id, pilar_id)")

	db.Exec(" ALTER TABLE ONLY public.componentes_pilares" +
		" ADD CONSTRAINT membros_unique_key UNIQUE (pilar_id, componente_id)")

	db.Exec(" ALTER TABLE ONLY public.elementos_componentes" +
		" ADD CONSTRAINT membros_unique_key UNIQUE (componente_id, elemento_id)")
}

func createFKey() {
	// ACTIONS
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

	//  ACTIONS_STATUS
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

	// ACTIVITIES_ROLES
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

	// CICLOS ENTIDADES
	db.Exec("ALTER TABLE ONLY public.ciclos_entidades" +
		" ADD CONSTRAINT entidades_fkey FOREIGN KEY (entidade_id)" +
		" REFERENCES public.entidades (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY public.ciclos_entidades" +
		" ADD CONSTRAINT ciclos_fkey FOREIGN KEY (ciclo_id)" +
		" REFERENCES public.ciclos (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY public.ciclos_entidades" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES public.users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY public.ciclos_entidades" +
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

	// COMPONENTES PILARES
	db.Exec("ALTER TABLE ONLY public.componentes_pilares" +
		" ADD CONSTRAINT componentes_fkey FOREIGN KEY (componente_id)" +
		" REFERENCES public.componentes (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY public.componentes_pilares" +
		" ADD CONSTRAINT pilares_fkey FOREIGN KEY (pilar_id)" +
		" REFERENCES public.pilares (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY public.componentes_pilares" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES public.users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY public.componentes_pilares" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES public.status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

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

	// ELEMENTOS_COMPONENTES
	db.Exec("ALTER TABLE ONLY public.elementos_componentes" +
		" ADD CONSTRAINT tipos_notas_fkey FOREIGN KEY (tipo_nota_id)" +
		" REFERENCES public.tipos_notas (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY public.elementos_componentes" +
		" ADD CONSTRAINT elementos_fkey FOREIGN KEY (elemento_id)" +
		" REFERENCES public.elementos (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY public.elementos_componentes" +
		" ADD CONSTRAINT componentes_fkey FOREIGN KEY (componente_id)" +
		" REFERENCES public.componentes (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY public.elementos_componentes" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES public.users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY public.elementos_componentes" +
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

	// ESCRITÓRIOS
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

	// FEATURES_ACTIVITIES
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

	// FEATURES_ROLES
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

	// ITENS
	db.Exec("ALTER TABLE ONLY public.itens" +
		" ADD CONSTRAINT elementos_fkey FOREIGN KEY (elemento_id)" +
		" REFERENCES public.elementos (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
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

	// PILARES_CICLOS
	db.Exec("ALTER TABLE ONLY public.pilares_ciclos" +
		" ADD CONSTRAINT pilares_fkey FOREIGN KEY (pilar_id)" +
		" REFERENCES public.pilares (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY public.pilares_ciclos" +
		" ADD CONSTRAINT ciclos_fkey FOREIGN KEY (ciclo_id)" +
		" REFERENCES public.ciclos (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY public.pilares_ciclos" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES public.users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY public.pilares_ciclos" +
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
	// USERS
	db.Exec("ALTER TABLE ONLY public.users " +
		" ADD CONSTRAINT roles_fkey FOREIGN KEY (role_id)" +
		" REFERENCES public.roles (id) MATCH SIMPLE" +
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
