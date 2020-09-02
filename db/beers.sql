--
-- PostgreSQL database dump
--

-- Dumped from database version 11.1
-- Dumped by pg_dump version 11.1

-- Started on 2020-09-02 11:26:59

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 198 (class 1259 OID 700433)
-- Name: beers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.beers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.beers_id_seq OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 199 (class 1259 OID 700439)
-- Name: beers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.beers (
    id integer DEFAULT nextval('public.beers_id_seq'::regclass) NOT NULL,
    name character varying(255) NOT NULL,
    qtd integer,
    price double precision
);


ALTER TABLE public.beers OWNER TO postgres;

--
-- TOC entry 200 (class 1259 OID 700446)
-- Name: clients_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.clients_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.clients_id_seq OWNER TO postgres;

--
-- TOC entry 196 (class 1259 OID 700419)
-- Name: clients; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.clients (
    id integer DEFAULT nextval('public.clients_id_seq'::regclass) NOT NULL,
    username character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    mobile character varying(255) NOT NULL,
    name character varying(255)
);


ALTER TABLE public.clients OWNER TO postgres;

--
-- TOC entry 202 (class 1259 OID 700451)
-- Name: items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.items_id_seq OWNER TO postgres;

--
-- TOC entry 203 (class 1259 OID 700456)
-- Name: items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.items (
    id integer DEFAULT nextval('public.items_id_seq'::regclass) NOT NULL,
    quantity double precision,
    beer_id integer,
    price double precision,
    item_value double precision,
    order_id integer
);


ALTER TABLE public.items OWNER TO postgres;

--
-- TOC entry 201 (class 1259 OID 700449)
-- Name: orders_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.orders_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.orders_id_seq OWNER TO postgres;

--
-- TOC entry 197 (class 1259 OID 700427)
-- Name: orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.orders (
    id integer DEFAULT nextval('public.orders_id_seq'::regclass) NOT NULL,
    client_id integer,
    ordered_at timestamp without time zone,
    take_out_at timestamp without time zone
);


ALTER TABLE public.orders OWNER TO postgres;

--
-- TOC entry 2842 (class 0 OID 700439)
-- Dependencies: 199
-- Data for Name: beers; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.beers (id, name, qtd, price) VALUES (1, 'Molson', 100, 100);
INSERT INTO public.beers (id, name, qtd, price) VALUES (3, 'Antartica', 200, 200);


--
-- TOC entry 2839 (class 0 OID 700419)
-- Dependencies: 196
-- Data for Name: clients; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.clients (id, username, password, email, mobile, name) VALUES (1, 'aria', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', 'aria@vindixit.com', '61 984385415', 'Ária Ohashi');
INSERT INTO public.clients (id, username, password, email, mobile, name) VALUES (3, 'aquino', '$2a$10$y0lE08uyKPlhiORjkyKSf.qwEjyTJqaaZZYfcSsFX6bAH6MkfvtWu', 'fabiochr@gmail.com', '61 984381019', 'Fábio Aquino');


--
-- TOC entry 2846 (class 0 OID 700456)
-- Dependencies: 203
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.items (id, quantity, beer_id, price, item_value, order_id) VALUES (2, 10, 1, 100, 1000, 22);
INSERT INTO public.items (id, quantity, beer_id, price, item_value, order_id) VALUES (3, 10, 1, 100, 1000, 23);


--
-- TOC entry 2840 (class 0 OID 700427)
-- Dependencies: 197
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.orders (id, client_id, ordered_at, take_out_at) VALUES (22, 1, '2020-08-31 07:59:00', '2020-09-01 07:59:00');
INSERT INTO public.orders (id, client_id, ordered_at, take_out_at) VALUES (23, 1, '2020-08-31 08:18:00', '2020-09-01 08:18:00');


--
-- TOC entry 2852 (class 0 OID 0)
-- Dependencies: 198
-- Name: beers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.beers_id_seq', 3, true);


--
-- TOC entry 2853 (class 0 OID 0)
-- Dependencies: 200
-- Name: clients_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.clients_id_seq', 3, true);


--
-- TOC entry 2854 (class 0 OID 0)
-- Dependencies: 202
-- Name: items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.items_id_seq', 28, true);


--
-- TOC entry 2855 (class 0 OID 0)
-- Dependencies: 201
-- Name: orders_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.orders_id_seq', 23, true);


--
-- TOC entry 2713 (class 2606 OID 700444)
-- Name: beers beers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.beers
    ADD CONSTRAINT beers_pkey PRIMARY KEY (id);


--
-- TOC entry 2715 (class 2606 OID 700461)
-- Name: items items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);


--
-- TOC entry 2711 (class 2606 OID 700431)
-- Name: orders order_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT order_pkey PRIMARY KEY (id);


--
-- TOC entry 2709 (class 2606 OID 700426)
-- Name: clients pk_Id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.clients
    ADD CONSTRAINT "pk_Id" PRIMARY KEY (id);


--
-- TOC entry 2716 (class 2606 OID 700462)
-- Name: items beers_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT beers_fkey FOREIGN KEY (beer_id) REFERENCES public.beers(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- TOC entry 2717 (class 2606 OID 700470)
-- Name: items orders_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT orders_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


-- Completed on 2020-09-02 11:27:00

--
-- PostgreSQL database dump complete
--

