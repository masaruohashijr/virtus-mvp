package db

import (
	hd "beerwh/handlers"
	"database/sql"
	"log"
)
var db  *sql.DB
func Initialize(){
	db = hd.Db
	createSeq()
	createTable()
	createAdmin()
	createPKey()
	createFKey()
	createBeer()
}

func createFKey(){
db.Exec("ALTER TABLE ONLY public.items" +
	    " ADD CONSTRAINT beers_fkey FOREIGN KEY (beer_id)" +
	    " REFERENCES public.beers(id) ON UPDATE RESTRICT ON DELETE RESTRICT")

db.Exec("ALTER TABLE ONLY public.items"+
	    " ADD CONSTRAINT orders_fkey FOREIGN KEY (order_id)" +
	    " REFERENCES public.orders(id) ON UPDATE RESTRICT ON DELETE RESTRICT")

db.Exec("ALTER TABLE ONLY public.orders" +
		" ADD CONSTRAINT clients_fkey FOREIGN KEY (client_id)" +
        " REFERENCES public.clients (id) MATCH SIMPLE" +
        " ON UPDATE RESTRICT ON DELETE RESTRICT" +
        " NOT VALID")
}

func createPKey(){
	db.Exec("ALTER TABLE ONLY public.beers ADD CONSTRAINT beers_pkey PRIMARY KEY (id)") 
    db.Exec("ALTER TABLE ONLY public.items ADD CONSTRAINT items_pkey PRIMARY KEY (id)") 
    db.Exec("ALTER TABLE ONLY public.orders ADD CONSTRAINT order_pkey PRIMARY KEY (id)")
    db.Exec("ALTER TABLE ONLY public.clients ADD CONSTRAINT clients_pkey PRIMARY KEY (id)")

}

func createAdmin(){
	query:="INSERT INTO clients (id, username, password, email, mobile, name)" +
		" SELECT 1, 'aria', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', 'aria@vindixit.com', '61 984385415', 'Ária Ohashi'" +
		" WHERE NOT EXISTS (SELECT id FROM clients WHERE username = 'aria')"
	log.Println(query)
	db.Exec(query)
}

func createBeer(){
	query:="INSERT INTO beers (id, name, qtd, price)" +
		" SELECT 1, 'Molson', 100, 100" +
		" WHERE NOT EXISTS (SELECT name FROM beers WHERE name = 'Molson')"
	log.Println(query)
	db.Exec(query)
	
}

func createSeq(){
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.beers_id_seq " +
	    " START WITH 1" +
	    " INCREMENT BY 1" +
	    " NO MINVALUE" +
	    " NO MAXVALUE" +
	    " CACHE 1")	
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.clients_id_seq " +
	    " START WITH 1" +
	    " INCREMENT BY 1" +
	    " NO MINVALUE" +
	    " NO MAXVALUE" +
	    " CACHE 1")	
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.items_id_seq " +
	    " START WITH 1" +
	    " INCREMENT BY 1" +
	    " NO MINVALUE" +
	    " NO MAXVALUE" +
	    " CACHE 1")	
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.orders_id_seq " +
	    " START WITH 1" +
	    " INCREMENT BY 1" +
	    " NO MINVALUE" +
	    " NO MAXVALUE" +
	    " CACHE 1")
}

func createTable(){
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.beers  (" +
		" id integer DEFAULT nextval('public.beers_id_seq'::regclass) NOT NULL,"+
	    " name character varying(255) NOT NULL,"+
	    " qtd integer,"+
	    " price double precision)")
	
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.clients (" +
	    " id integer DEFAULT nextval('public.clients_id_seq'::regclass) NOT NULL," +
	    " username character varying(255) NOT NULL," +
	    " password character varying(255) NOT NULL," +
	    " email character varying(255) NOT NULL," +
	    " mobile character varying(255) NOT NULL," +
	    " name character varying(255))")
	
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.items (" +
	    " id integer DEFAULT nextval('public.items_id_seq'::regclass) NOT NULL," +
	    " quantity double precision," +
	    " beer_id integer," +
	    " price double precision," +
	    " item_value double precision," +
	    " order_id integer)")
		
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.orders (" +
	    " id integer DEFAULT nextval('public.orders_id_seq'::regclass) NOT NULL," +
	    " client_id integer," +
	    " ordered_at timestamp without time zone," +
	    " take_out_at timestamp without time zone)")
}