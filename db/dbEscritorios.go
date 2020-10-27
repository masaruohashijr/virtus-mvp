package db

import ()

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
