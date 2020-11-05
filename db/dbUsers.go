package db

import ()

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
		" SELECT 3, 'arnaldo', '$2a$10$aiYcB.Q5DpE1ZBLvxHRMD.nGu32qBvb5EMwCJGiOACItLFbghdb4K', " +
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
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 31, 'fulano', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'fulano@gmail.com', '61 984385415', 'Fulano de Tal', 2, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'fulano')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 29, 'sicrano', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'sicrano@gmail.com', '61 984385415', 'Sicrano de Tal', 3, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'sicrano')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 30, 'beltrano', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'beltrano@gmail.com', '61 984385415', 'Beltrano de Tal', 4, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'beltrano')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 31, 'huguinho', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'huguinho@gmail.com', '61 984385415', 'Huguinho da Silva', 2, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'huguinho')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 32, 'zezinho', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'zezinho@gmail.com', '61 984385415', 'Zezinho da Silva', 3, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'zezinho')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 33, 'luisinho', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'luisinho@gmail.com', '61 984385415', 'Luisinho de Tal', 4, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'luisinho')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 34, 'athos', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'athos@gmail.com', '61 984385415', 'Athos Mosqueteiro', 2, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'athos')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 35, 'porthos', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'porthos@gmail.com', '61 984385415', 'Porthos Mosqueteiro', 3, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'porthos')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 36, 'aramis', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'aramis@gmail.com', '61 984385415', 'Aramis Mosqueteiro', 4, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'aramis')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 34, 'zuenir', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'zuenir@gmail.com', '61 984385415', 'Zuenir Ventura', 2, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'zuenir')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 35, 'verissimo', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'verissimo@gmail.com', '61 984385415', 'Luís Fernando Veríssimo', 3, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'verissimo')"
	db.Exec(sql)
	sql = "INSERT INTO users (id, username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 36, 'ariano', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', " +
		" 'ariano@gmail.com', '61 984385415', 'Ariano Suassuna', 4, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'ariano')"
	db.Exec(sql)
}
