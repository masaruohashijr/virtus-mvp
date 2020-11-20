package db

import ()

func createEscritorios() {
	sql := "INSERT INTO public.escritorios( " +
		" id, nome, descricao, abreviatura, author_id, criado_em) " +
		" SELECT 1, 'Escritório de Representação - Pernambuco', 'Escritório de Representação - Pernambuco','ERPE', 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM escritorios WHERE id = 1)"
	db.Exec(sql)
	sql = "INSERT INTO public.escritorios( " +
		" id, nome, descricao, abreviatura, author_id, criado_em) " +
		" SELECT 2, 'Escritório de Representação - São Paulo', 'Escritório de Representação - São Paulo','ERSP', 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM escritorios WHERE id = 2)"
	db.Exec(sql)
	sql = "INSERT INTO public.escritorios( " +
		" id, nome, descricao, abreviatura, author_id, criado_em) " +
		" SELECT 3, 'Escritório de Representação - Minas Gerais', 'Escritório de Representação - Minas Gerais','ERMG', 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM escritorios WHERE id = 3)"
	db.Exec(sql)
	sql = "INSERT INTO public.escritorios( " +
		" id, nome, descricao, abreviatura, author_id, criado_em) " +
		" SELECT 4, 'Escritório de Representação - Rio Grande do Sul', 'Escritório de Representação - Rio Grande do Sul','ERRS', 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM escritorios WHERE id = 4)"
	db.Exec(sql)
	sql = "INSERT INTO public.escritorios( " +
		" id, nome, descricao, abreviatura, author_id, criado_em) " +
		" SELECT 5, 'Escritório de Representação - Rio de Janeiro', 'Escritório de Representação - Rio de Janeiro','ERRJ', 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM escritorios WHERE id = 5)"
	db.Exec(sql)
	sql = "INSERT INTO public.escritorios( " +
		" id, nome, descricao, abreviatura, author_id, criado_em) " +
		" SELECT 6, 'Escritório de Representação - Distrito Federal', 'Escritório de Representação - Distrito Federal','ERDF', 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM escritorios WHERE id = 6)"
	db.Exec(sql)
}

func createMembros() {
	sql := " INSERT INTO membros " +
		" (  escritorio_id,  usuario_id,  author_id,  criado_em  ) " +
		" SELECT 6, 29, 1, now()::timestamp WHERE NOT EXISTS " +
		" (SELECT 1 FROM membros WHERE escritorio_id = 6 AND usuario_id = 29) "
	db.Exec(sql)
	sql = " INSERT INTO membros " +
		" (  escritorio_id,  usuario_id,  author_id,  criado_em  ) " +
		" SELECT 6, 30, 1, now()::timestamp WHERE NOT EXISTS " +
		" (SELECT 1 FROM membros WHERE escritorio_id = 6 AND usuario_id = 30) "
	db.Exec(sql)
	sql = " INSERT INTO membros " +
		" (  escritorio_id,  usuario_id,  author_id,  criado_em  ) " +
		" SELECT 6, 31, 1, now()::timestamp WHERE NOT EXISTS " +
		" (SELECT 1 FROM membros WHERE escritorio_id = 6 AND usuario_id = 31) "
	db.Exec(sql)
	sql = " INSERT INTO membros " +
		" (  escritorio_id,  usuario_id,  author_id,  criado_em  ) " +
		" SELECT 6, 37, 1, now()::timestamp WHERE NOT EXISTS " +
		" (SELECT 1 FROM membros WHERE escritorio_id = 6 AND usuario_id = 37) "
	db.Exec(sql)
	sql = " UPDATE escritorios SET chefe_id = 29 WHERE id = 6 "
	db.Exec(sql)
}
