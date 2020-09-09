--
-- PostgreSQL database dump
--

-- Dumped from database version 11.1
-- Dumped by pg_dump version 11.1

-- Started on 2020-09-09 11:41:07

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
-- TOC entry 209 (class 1259 OID 700531)
-- Name: actions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.actions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.actions_id_seq OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 213 (class 1259 OID 700554)
-- Name: actions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.actions (
    id integer DEFAULT nextval('public.actions_id_seq'::regclass) NOT NULL,
    name character varying(255) NOT NULL,
    workflow_id integer NOT NULL,
    origin_status_id integer,
    destination_status_id integer
);


ALTER TABLE public.actions OWNER TO postgres;

--
-- TOC entry 197 (class 1259 OID 700433)
-- Name: beers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.beers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.beers_id_seq OWNER TO postgres;

--
-- TOC entry 198 (class 1259 OID 700439)
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
-- TOC entry 204 (class 1259 OID 700497)
-- Name: features_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.features_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.features_id_seq OWNER TO postgres;

--
-- TOC entry 205 (class 1259 OID 700499)
-- Name: features; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.features (
    id integer DEFAULT nextval('public.features_id_seq'::regclass) NOT NULL,
    name character varying(255),
    code character varying(255)
);


ALTER TABLE public.features OWNER TO postgres;

--
-- TOC entry 206 (class 1259 OID 700508)
-- Name: features_roles_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.features_roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.features_roles_id_seq OWNER TO postgres;

--
-- TOC entry 207 (class 1259 OID 700510)
-- Name: features_roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.features_roles (
    id integer DEFAULT nextval('public.features_roles_id_seq'::regclass) NOT NULL,
    feature_id integer,
    role_id integer
);


ALTER TABLE public.features_roles OWNER TO postgres;

--
-- TOC entry 200 (class 1259 OID 700451)
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
-- TOC entry 201 (class 1259 OID 700456)
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
-- TOC entry 199 (class 1259 OID 700449)
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
-- TOC entry 196 (class 1259 OID 700427)
-- Name: orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.orders (
    id integer DEFAULT nextval('public.orders_id_seq'::regclass) NOT NULL,
    user_id integer,
    ordered_at timestamp without time zone,
    take_out_at timestamp without time zone
);


ALTER TABLE public.orders OWNER TO postgres;

--
-- TOC entry 203 (class 1259 OID 700485)
-- Name: roles_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.roles_id_seq OWNER TO postgres;

--
-- TOC entry 202 (class 1259 OID 700480)
-- Name: roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.roles (
    id integer DEFAULT nextval('public.roles_id_seq'::regclass) NOT NULL,
    name character varying(255)
);


ALTER TABLE public.roles OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 700533)
-- Name: status_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.status_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.status_id_seq OWNER TO postgres;

--
-- TOC entry 212 (class 1259 OID 700547)
-- Name: status; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.status (
    id integer DEFAULT nextval('public.status_id_seq'::regclass) NOT NULL,
    status character varying(255) NOT NULL,
    stereotype character varying(5),
    workflow_id integer
);


ALTER TABLE public.status OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 700621)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 700623)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer DEFAULT nextval('public.users_id_seq'::regclass) NOT NULL,
    username character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    mobile character varying(255) NOT NULL,
    name character varying(255)
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 208 (class 1259 OID 700529)
-- Name: workflows_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.workflows_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.workflows_id_seq OWNER TO postgres;

--
-- TOC entry 211 (class 1259 OID 700541)
-- Name: workflows; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.workflows (
    id integer DEFAULT nextval('public.workflows_id_seq'::regclass) NOT NULL,
    name character varying(255) NOT NULL
);


ALTER TABLE public.workflows OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 700580)
-- Name: workflows_entities; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.workflows_entities (
    entity_name character varying(255) NOT NULL,
    workflow_id integer NOT NULL
);


ALTER TABLE public.workflows_entities OWNER TO postgres;

--
-- TOC entry 2930 (class 0 OID 700554)
-- Dependencies: 213
-- Data for Name: actions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.actions (id, name, workflow_id, origin_status_id, destination_status_id) FROM stdin;
\.


--
-- TOC entry 2915 (class 0 OID 700439)
-- Dependencies: 198
-- Data for Name: beers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.beers (id, name, qtd, price) FROM stdin;
3	Antartica	200	200
1	Molson	100	100
\.


--
-- TOC entry 2922 (class 0 OID 700499)
-- Dependencies: 205
-- Data for Name: features; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.features (id, name, code) FROM stdin;
1	Listar Cervejas	listBeers
2	Criar Cerveja	createBeer
\.


--
-- TOC entry 2924 (class 0 OID 700510)
-- Dependencies: 207
-- Data for Name: features_roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.features_roles (id, feature_id, role_id) FROM stdin;
\.


--
-- TOC entry 2918 (class 0 OID 700456)
-- Dependencies: 201
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.items (id, quantity, beer_id, price, item_value, order_id) FROM stdin;
\.


--
-- TOC entry 2913 (class 0 OID 700427)
-- Dependencies: 196
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.orders (id, user_id, ordered_at, take_out_at) FROM stdin;
\.


--
-- TOC entry 2919 (class 0 OID 700480)
-- Dependencies: 202
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.roles (id, name) FROM stdin;
2	admin
3	cliente
4	vendedor
6	admin2
7	admin3
8	admin4
9	admin6
\.


--
-- TOC entry 2929 (class 0 OID 700547)
-- Dependencies: 212
-- Data for Name: status; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.status (id, status, stereotype, workflow_id) FROM stdin;
\.


--
-- TOC entry 2933 (class 0 OID 700623)
-- Dependencies: 216
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, username, password, email, mobile, name) FROM stdin;
1	aria	$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy	aria@vindixit.com	61 984385415	Ária Ohashi
\.


--
-- TOC entry 2928 (class 0 OID 700541)
-- Dependencies: 211
-- Data for Name: workflows; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.workflows (id, name) FROM stdin;
\.


--
-- TOC entry 2931 (class 0 OID 700580)
-- Dependencies: 214
-- Data for Name: workflows_entities; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.workflows_entities (entity_name, workflow_id) FROM stdin;
\.


--
-- TOC entry 2939 (class 0 OID 0)
-- Dependencies: 209
-- Name: actions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.actions_id_seq', 1, false);


--
-- TOC entry 2940 (class 0 OID 0)
-- Dependencies: 197
-- Name: beers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.beers_id_seq', 3, true);


--
-- TOC entry 2941 (class 0 OID 0)
-- Dependencies: 204
-- Name: features_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.features_id_seq', 2, true);


--
-- TOC entry 2942 (class 0 OID 0)
-- Dependencies: 206
-- Name: features_roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.features_roles_id_seq', 1, false);


--
-- TOC entry 2943 (class 0 OID 0)
-- Dependencies: 200
-- Name: items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.items_id_seq', 30, true);


--
-- TOC entry 2944 (class 0 OID 0)
-- Dependencies: 199
-- Name: orders_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.orders_id_seq', 25, true);


--
-- TOC entry 2945 (class 0 OID 0)
-- Dependencies: 203
-- Name: roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.roles_id_seq', 9, true);


--
-- TOC entry 2946 (class 0 OID 0)
-- Dependencies: 210
-- Name: status_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.status_id_seq', 1, false);


--
-- TOC entry 2947 (class 0 OID 0)
-- Dependencies: 215
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 1, false);


--
-- TOC entry 2948 (class 0 OID 0)
-- Dependencies: 208
-- Name: workflows_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.workflows_id_seq', 1, false);


--
-- TOC entry 2774 (class 2606 OID 700559)
-- Name: actions actions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.actions
    ADD CONSTRAINT actions_pkey PRIMARY KEY (id);


--
-- TOC entry 2758 (class 2606 OID 700444)
-- Name: beers beers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.beers
    ADD CONSTRAINT beers_pkey PRIMARY KEY (id);


--
-- TOC entry 2766 (class 2606 OID 700639)
-- Name: features_roles feature_role_unique_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.features_roles
    ADD CONSTRAINT feature_role_unique_key UNIQUE (feature_id, role_id);


--
-- TOC entry 2764 (class 2606 OID 700507)
-- Name: features features_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.features
    ADD CONSTRAINT features_pkey PRIMARY KEY (id);


--
-- TOC entry 2768 (class 2606 OID 700637)
-- Name: features_roles features_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.features_roles
    ADD CONSTRAINT features_roles_pkey PRIMARY KEY (id);


--
-- TOC entry 2760 (class 2606 OID 700461)
-- Name: items items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);


--
-- TOC entry 2756 (class 2606 OID 700431)
-- Name: orders order_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT order_pkey PRIMARY KEY (id);


--
-- TOC entry 2776 (class 2606 OID 700586)
-- Name: workflows_entities order_unique_index; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.workflows_entities
    ADD CONSTRAINT order_unique_index UNIQUE (entity_name);


--
-- TOC entry 2762 (class 2606 OID 700496)
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- TOC entry 2772 (class 2606 OID 700552)
-- Name: status status_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.status
    ADD CONSTRAINT status_pkey PRIMARY KEY (id);


--
-- TOC entry 2780 (class 2606 OID 700646)
-- Name: users username_unique_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT username_unique_key UNIQUE (username);


--
-- TOC entry 2782 (class 2606 OID 700631)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 2778 (class 2606 OID 700584)
-- Name: workflows_entities workflows_entities_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.workflows_entities
    ADD CONSTRAINT workflows_entities_pkey PRIMARY KEY (entity_name, workflow_id);


--
-- TOC entry 2770 (class 2606 OID 700546)
-- Name: workflows workflows_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.workflows
    ADD CONSTRAINT workflows_pkey PRIMARY KEY (id);


--
-- TOC entry 2784 (class 2606 OID 700462)
-- Name: items beers_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT beers_fkey FOREIGN KEY (beer_id) REFERENCES public.beers(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- TOC entry 2789 (class 2606 OID 700592)
-- Name: actions destination_status_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.actions
    ADD CONSTRAINT destination_status_fkey FOREIGN KEY (destination_status_id) REFERENCES public.status(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- TOC entry 2786 (class 2606 OID 700519)
-- Name: features_roles features_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.features_roles
    ADD CONSTRAINT features_fkey FOREIGN KEY (feature_id) REFERENCES public.features(id);


--
-- TOC entry 2785 (class 2606 OID 700470)
-- Name: items orders_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT orders_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- TOC entry 2790 (class 2606 OID 700597)
-- Name: actions origin_status_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.actions
    ADD CONSTRAINT origin_status_fkey FOREIGN KEY (origin_status_id) REFERENCES public.status(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- TOC entry 2787 (class 2606 OID 700524)
-- Name: features_roles roles_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.features_roles
    ADD CONSTRAINT roles_fkey FOREIGN KEY (role_id) REFERENCES public.roles(id);


--
-- TOC entry 2783 (class 2606 OID 700640)
-- Name: orders users_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT users_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- TOC entry 2788 (class 2606 OID 700560)
-- Name: status workflows_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.status
    ADD CONSTRAINT workflows_fkey FOREIGN KEY (workflow_id) REFERENCES public.workflows(id);


--
-- TOC entry 2791 (class 2606 OID 700602)
-- Name: actions workflows_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.actions
    ADD CONSTRAINT workflows_fkey FOREIGN KEY (workflow_id) REFERENCES public.workflows(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


-- Completed on 2020-09-09 11:41:07

--
-- PostgreSQL database dump complete
--

