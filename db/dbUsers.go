package db

import ()

func createUsers() {
	sql := "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'aria', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'aria@gmail.com', '61 984385415', 'Ária Ohashi', 1, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'aria')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'masaru', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'masaru@gmail.com', '61 984385415', 'Masaru Ohashi Jr', 1, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'masaru')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'arnaldo', '$2a$10$aiYcB.Q5DpE1ZBLvxHRMD.nGu32qBvb5EMwCJGiOACItLFbghdb4K', " +
		" 'arnaldo@gmail.com', '61 984385415', 'Arnaldo Burle', 1, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'arnaldo')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'ana', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'ana@gmail.com', '61 984385415', 'Ana Carolina Baasch', 2, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'ana')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'annette', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'annette@gmail.com', '61 984385415', 'Annette Lopes Pinto ', 2, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'annette')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'carlos', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'carlos@gmail.com', '61 984385415', 'Carlos Marne Dias Alves', 2, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'carlos')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'christian', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'christian@gmail.com', '61 984385415', 'Christian Aggensteiner Catunda', 2, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'christian')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'dagomar', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'dagomar@gmail.com', '61 984385415', 'Dagomar Alécio Anhê', 2, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'dagomar')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'david', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'david@gmail.com', '61 984385415', 'David Prates Coutinho', 3, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'david')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'elthon', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'elthon@gmail.com', '61 984385415', 'Elthon Baier Nunes', 3, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'elthon')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'fabio', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'fábio@gmail.com', '61 984385415', 'Fábio Lucas de Albuquerque Lima', 3, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'fabio')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'fabricio', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'fabricio@gmail.com', '61 984385415', 'Fabricio Cardoso de Meneses', 3, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'fabricio')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'felipe', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'felipe@gmail.com', '61 984385415', 'Felipe Spolavori Martins', 3, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'felipe')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'fernando', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'fernando@gmail.com', '61 984385415', 'Fernando Duarte Folle', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'fernando')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'hilton', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'hilton@gmail.com', '61 984385415', 'Hilton de Enzo Mitsunaga', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'hilton')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'chedeak', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'josé@gmail.com', '61 984385415', 'José Carlos Sampaio Chedeak', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'chedeak')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'jose', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'josé@gmail.com', '61 984385415', 'José de Arimatéria Pinheiro Torres', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'jose')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'luciano', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'luciano@gmail.com', '61 984385415', 'Luciano Draghetti', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'luciano')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'lucio', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'lucio@gmail.com', '61 984385415', 'Lucio Rodrigues Capelletto', 1, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'lucio')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'luis', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'luis@gmail.com', '61 984385415', 'Luis Ronaldo Martins Angoti', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'luis')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'manoel', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'manoel@gmail.com', '61 984385415', 'Manoel Robson Aguiar', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'manoel')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'mauricio', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'maurício@gmail.com', '61 984385415', 'Maurício de Aguirre Nakata', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'mauricio')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'milton', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'milton@gmail.com', '61 984385415', 'Milton Santos', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'milton')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'otavio', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'otávio@gmail.com', '61 984385415', 'Otávio Lima Reis', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'otavio')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'paulo', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'paulo@gmail.com', '61 984385415', 'Paulo Roberto Pereira De Macedo', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'paulo')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'peterson', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'peterson@gmail.com', '61 984385415', 'Peterson Gonçalves', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'peterson')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'rita', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'rita@gmail.com', '61 984385415', 'Rita de Cassia Correa da Silva', 1, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'rita')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'sergio', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'sergio@gmail.com', '61 984385415', 'Sergio Djundi Taniguchi', 4, 1, now()::timestamp  " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'sergio')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'fulano', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'fulano@gmail.com', '61 984385415', 'Fulano de Tal', 2, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'fulano')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'sicrano', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'sicrano@gmail.com', '61 984385415', 'Sicrano de Tal', 3, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'sicrano')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'beltrano', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'beltrano@gmail.com', '61 984385415', 'Beltrano de Tal', 4, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'beltrano')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'huguinho', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'huguinho@gmail.com', '61 984385415', 'Huguinho da Silva', 2, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'huguinho')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'zezinho', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'zezinho@gmail.com', '61 984385415', 'Zezinho da Silva', 3, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'zezinho')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'luisinho', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'luisinho@gmail.com', '61 984385415', 'Luisinho de Tal', 4, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'luisinho')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'athos', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'athos@gmail.com', '61 984385415', 'Athos Mosqueteiro', 2, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'athos')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'porthos', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'porthos@gmail.com', '61 984385415', 'Porthos Mosqueteiro', 3, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'porthos')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'aramis', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'aramis@gmail.com', '61 984385415', 'Aramis Mosqueteiro', 4, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'aramis')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'zuenir', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'zuenir@gmail.com', '61 984385415', 'Zuenir Ventura', 2, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'zuenir')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'verissimo', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'verissimo@gmail.com', '61 984385415', 'Luís Fernando Veríssimo', 3, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'verissimo')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'james', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'james@gmail.com', '61 984385415', 'James Taylor', 1, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'james')"
	db.Exec(sql)
	sql = "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'leonardo', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'leofiuza@gmail.com', '61 984385415', 'Leonardo Fiuza', 1, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'leonardo')"
	db.Exec(sql)
}
