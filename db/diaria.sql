--
-- PostgreSQL database dump
--

-- Dumped from database version 11.1
-- Dumped by pg_dump version 11.1

-- Started on 2020-08-21 13:34:19

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
-- TOC entry 2874 (class 1262 OID 642500)
-- Name: aria; Type: DATABASE; Schema: -; Owner: vindixit
--
DROP DATABASE aria
CREATE DATABASE aria WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'Portuguese_Brazil.1252' LC_CTYPE = 'Portuguese_Brazil.1252';


ALTER DATABASE aria OWNER TO vindixit;

\connect aria

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
-- TOC entry 204 (class 1259 OID 692172)
-- Name: foods_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.foods_id_seq
    START WITH 1
    INCREMENT BY 1
    MINVALUE 0
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.foods_id_seq OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 203 (class 1259 OID 692157)
-- Name: foods; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.foods (
    id integer DEFAULT nextval('public.foods_id_seq'::regclass) NOT NULL,
    name character varying(255) NOT NULL,
    measure_id integer,
    qtd double precision,
    cho double precision,
    kcal double precision,
    measure character varying(255),
    CONSTRAINT teste CHECK ((measure_id = id))
);


ALTER TABLE public.foods OWNER TO postgres;

--
-- TOC entry 206 (class 1259 OID 692207)
-- Name: items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.items_id_seq
    START WITH 1
    INCREMENT BY 1
    MINVALUE 0
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.items_id_seq OWNER TO postgres;

--
-- TOC entry 205 (class 1259 OID 692179)
-- Name: items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.items (
    id integer DEFAULT nextval('public.items_id_seq'::regclass) NOT NULL,
    meal_id integer,
    quantidade_medida_usual double precision,
    quantidade_g_ml double precision,
    cho double precision,
    kcal double precision,
    food_id integer
);


ALTER TABLE public.items OWNER TO postgres;

--
-- TOC entry 199 (class 1259 OID 692123)
-- Name: meal_type_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.meal_type_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.meal_type_id_seq OWNER TO postgres;

--
-- TOC entry 200 (class 1259 OID 692132)
-- Name: meal_types; Type: TABLE; Schema: public; Owner: vindixit
--

CREATE TABLE public.meal_types (
    id integer DEFAULT nextval('public.meal_type_id_seq'::regclass) NOT NULL,
    name character varying(20) NOT NULL,
    start_at time without time zone,
    end_at time without time zone
);


ALTER TABLE public.meal_types OWNER TO vindixit;

--
-- TOC entry 197 (class 1259 OID 642683)
-- Name: meals; Type: TABLE; Schema: public; Owner: vindixit
--

CREATE TABLE public.meals (
    id integer NOT NULL,
    meal_type_id integer,
    bolus double precision,
    start_at time without time zone,
    end_at time without time zone,
    date date
);


ALTER TABLE public.meals OWNER TO vindixit;

--
-- TOC entry 196 (class 1259 OID 642681)
-- Name: meals_id_seq; Type: SEQUENCE; Schema: public; Owner: vindixit
--

CREATE SEQUENCE public.meals_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    MINVALUE 0
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.meals_id_seq OWNER TO vindixit;

--
-- TOC entry 2875 (class 0 OID 0)
-- Dependencies: 196
-- Name: meals_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: vindixit
--

ALTER SEQUENCE public.meals_id_seq OWNED BY public.meals.id;


--
-- TOC entry 202 (class 1259 OID 692143)
-- Name: measures_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.measures_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.measures_id_seq OWNER TO postgres;

--
-- TOC entry 201 (class 1259 OID 692138)
-- Name: measures; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.measures (
    id integer DEFAULT nextval('public.measures_id_seq'::regclass) NOT NULL,
    name character varying(255) NOT NULL
);


ALTER TABLE public.measures OWNER TO postgres;

--
-- TOC entry 198 (class 1259 OID 692112)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    "Id" integer NOT NULL,
    username character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    role character varying(255)
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 2715 (class 2604 OID 642686)
-- Name: meals id; Type: DEFAULT; Schema: public; Owner: vindixit
--

ALTER TABLE ONLY public.meals ALTER COLUMN id SET DEFAULT nextval('public.meals_id_seq'::regclass);


--
-- TOC entry 2865 (class 0 OID 692157)
-- Dependencies: 203
-- Data for Name: foods; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (2, 'Abacaxi', NULL, 75, 10, 44, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (3, 'Abacaxi em calda', NULL, 64, 19, 78, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (6, 'Abaráv', NULL, 170, 24, 414, 'unidade média ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (7, 'Abiu cru', NULL, 100, 15, 62, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (8, 'Abobora cabotian, cozida', NULL, 36, 3, 14, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (9, 'Abobora Cabotian, crua', NULL, 36, 4, 17, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (10, 'Abóbora dágua (picada)', NULL, 36, 0, 10, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (11, 'Abóbora doce (picada)', NULL, 36, 4, 18, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (12, 'Abóbora moranga (picada)', NULL, 36, 1, 7, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (13, 'Abobrinha recheada', NULL, 100, 6, 89, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (14, 'Abobrinha, italiana, cozida ', NULL, 20, 0, 3, 'colher de sopa ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (15, 'Abobrinha, italiana, crua ', NULL, 20, 1, 4, 'colher de sopa ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (16, 'Abobrinha, paulista, crua', NULL, 20, 2, 6, 'colher de sopa ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (17, 'Abricó-do-pará', NULL, 100, 13, 64, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (19, 'Açai natural Frooty ®', NULL, 60, 13, 62, '6 colheres de sopa rasas');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (20, 'Açaí (polpa)  com xarope de guarana e glicose', NULL, 100, 21, 110, 'taça pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (22, 'Açaí, suco de', NULL, 240, 72, 438, 'copo duplo cheio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (23, 'Acarajé', NULL, 100, 22, 282, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (24, 'Acelga (picada)', NULL, 6, 0, 2, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (25, 'Acém magro cozido', NULL, 100, 0, 215, 'pedaço médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (26, 'Acerola', NULL, 12, 1, 4, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (27, 'Achocolatado Diet Gold ®', NULL, 9, 6, 35, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (28, 'Achocolatado Diet Linea ®', NULL, 12, 5, 30, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (29, 'Açúcar branco renado', NULL, 30, 30, 116, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (30, 'Açúcar cristal', NULL, 24, 24, 96, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (31, 'Açúcar mascavo', NULL, 19, 17, 70, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (32, 'Agrião (picado)', NULL, 7, 0, 2, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (33, 'Agriao refogado', NULL, 25, 1, 19, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (34, 'Água-de-coco verde', NULL, 240, 10, 43, 'copo duplo');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (35, 'Aipim cozido', NULL, 100, 29, 120, 'pedaço médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (36, 'Aipo inteiro (picado)', NULL, 10, 0, 2, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (37, 'Alcachofra ', NULL, 20, 1, 6, 'pedaco médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (38, 'Alcaparra', NULL, 27, 1, 10, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (39, 'Alface, americana, crua', NULL, 10, 0, 1, 'folha média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (40, 'Alface, crespa, crua', NULL, 10, 0, 1, 'folha média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (41, 'Alface, lisa, crua', NULL, 10, 0, 1, 'folha média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (42, 'Alface, lisa, crua', NULL, 10, 0, 1, 'folha média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (43, 'Alface, roxa, crua', NULL, 10, 0, 1, 'folha média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (44, 'Alface, roxa, crua', NULL, 10, 0, 1, 'folha média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (45, 'Alfajor  ao leite Cacau Show®', NULL, 40, 23, 120, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (46, 'Alfajor® recheado com doce de leite', NULL, 25, 15, 98, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (47, 'Alfavaca crua', NULL, 10, 0, 3, 'folha média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (48, 'Algodão Doce', NULL, 30, 30, 116, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (49, 'Alho, cru', NULL, 10, 3, 11, '2 dentes');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (50, 'Alho-poró cru', NULL, 20, 3, 11, 'porcao pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (51, 'Ambrosia', NULL, 80, 33, 207, '2 colheres de sopa cheias');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (52, 'Ameixa de queijo', NULL, 12, 7, 36, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (53, 'Ameixa em calda', NULL, 7, 5, 20, '01 unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (54, 'Ameixa preta fresca', NULL, 42, 4, 18, 'unidade media');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (55, 'Amêndoa', NULL, 3, 1, 13, '01 unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (56, 'Amendoim caramelizado', NULL, 20, 14, 95, 'pacote pequeno');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (57, 'Amendoim cozido', NULL, 34, 7, 108, '2 colheres de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (58, 'Amendoim, grão, cru', NULL, 34, 5, 194, '2 colheres de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (59, 'Amendoim, torrado, salgado', NULL, 34, 5, 194, '2 colheres de sopa ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (60, 'Amido de arroz', NULL, 20, 17, 70, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (61, 'Amido de milho', NULL, 20, 17, 69, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (62, 'Amora (branca, preta e vermelha)', NULL, 8, 1, 5, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (63, 'Amora, geléia de', NULL, 40, 23, 93, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (64, 'Amora (branca, preta e vermelha)', NULL, 8, 1, 5, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (65, 'Bala Toe® Chocolate', NULL, 6, 22, 151, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (66, 'Banana pacova madura cozida sem casca', NULL, 35, 12, 48, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (67, 'Banana pacova madura frita (picada)', NULL, 35, 18, 96, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (68, 'Banana pacova madura in natura (picada)', NULL, 30, 11, 45, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (69, 'Banana pacova verde frita', NULL, 35, 30, 138, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (70, 'Banana-caturra ou Nanica', NULL, 86, 20, 79, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (71, 'Bananada', NULL, 40, 27, 115, 'unidade media');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (72, 'Banana-da-terra crua', NULL, 100, 27, 117, 'unidade grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (73, 'Banana-maçã', NULL, 65, 15, 72, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (74, 'Banana-ouro', NULL, 40, 9, 42, 'unidade pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (75, 'Banana-prata crua', NULL, 50, 13, 49, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (76, 'Banha de porco', NULL, 10, 0, 90, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (77, 'Barra de Cereal Diet Linea® Avela, Castanha e Chocolate', NULL, 25, 17, 75, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (78, 'Barra de cereal Diet Linea® banana com aveia', NULL, 25, 18, 65, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (79, 'Batata (Baked Potato®)', NULL, 470, 83, 369, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (80, 'Batata doce frita', NULL, 65, 39, 249, 'fatia media');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (81, 'Batata frita  (Bobs®)', NULL, 115, 41, 259, 'porção grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (82, 'Batata frita  (Bobs®)', NULL, 95, 34, 214, 'porção média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (83, 'Batata frita  (Bobs®)', NULL, 55, 20, 124, 'porção pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (84, 'Batata frita (Burguer King®)', NULL, 0, 37, 318, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (85, 'Batata frita (Habibs®)', NULL, 40, 10, 86, 'porção pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (86, 'Batata frita chips ', NULL, 13, 6, 70, 'punhado');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (87, 'Batata fritas  (Mc Donalds®)', NULL, 0, 25, 206, 'porção pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (88, 'Batata fritas (Mc Donalds®)', NULL, 0, 35, 288, 'porção média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (89, 'Batata fritas (Mc Donalds®)', NULL, 0, 49, 412, 'porção grande ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (90, 'Batata fritas McFritas Kids (Mc Donalds®)', NULL, 0, 11, 87, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (91, 'Batata inglesa cozida', NULL, 80, 6, 68, 'colher de sopa cheia picada');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (92, 'Batata inglesa fritta®', NULL, 65, 23, 182, 'escumadeira media cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (93, 'Batata inglesa Saute®', NULL, 25, 4, 37, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (94, 'Batata smiles Mccain®', NULL, 21, 7, 42, '01 unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (95, 'Batata, amido de', NULL, 16, 13, 53, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (96, 'Batata, fécula de', NULL, 20, 16, 66, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (97, 'Batata-baroa ou mandioquinha (picada)', NULL, 35, 10, 44, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (98, 'Batata-doce amarela assada (picada)', NULL, 30, 10, 43, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (99, 'Batata-doce branca cozida (picada)', NULL, 30, 8, 38, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (100, 'Batata-doce cozida', NULL, 42, 10, 43, 'colher sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (101, 'Batata-doce, doce de', NULL, 40, 24, 94, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (102, 'Baton (chocolate ao leite Garoto®)', NULL, 16, 9, 86, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (103, 'Beijinho', NULL, 25, 11, 105, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (104, 'Beijinho', NULL, 6, 3, 25, 'unidade pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (105, 'Beijinho diet', NULL, 6, 2, 20, 'unidade pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (106, 'Beijinho-de-coco Nestle®', NULL, 20, 11, 63, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (107, 'Beiju', NULL, 100, 87, 359, 'unidade media');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (108, 'Beiju com Coco', NULL, 125, 77, 622, 'unidade grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (109, 'Beiju de queijo com manteiga', NULL, 150, 87, 518, 'unidade média ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (110, 'Beirute Habibs®', NULL, 415, 51, 714, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (111, 'Bekleua', NULL, 100, 20, 290, '1 porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (112, 'Bem casado', NULL, 40, 25, 160, '01 unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (113, 'Bergamota ou Mexerica', NULL, 100, 15, 58, 'unidade grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (114, 'Berinjela cozida sem sal', NULL, 25, 2, 8, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (115, 'Berinjela frita', NULL, 13, 1, 10, 'rodela média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (116, 'Beterraba cozida (picada)', NULL, 20, 2, 9, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (117, 'Beterraba crua', NULL, 16, 1, 7, 'colher de sopa cheia ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (118, 'Bibs dog  (Habibs®)', NULL, 105, 21, 246, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (119, 'Bibs salad (Habibs®)', NULL, 300, 36, 209, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (120, 'Bife à milanesa', NULL, 80, 8, 230, 'unidade pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (121, 'Bife de boi', NULL, 100, 0, 228, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (122, 'Bife de fígado ', NULL, 100, 6, 216, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (123, 'Bife Role', NULL, 150, 4, 268, 'unidade grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (124, 'Big Bob picanha 100g (Bobs®)', NULL, 235, 35, 513, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (125, 'Big Bob picanha 200g (Bobs®)', NULL, 344, 38, 728, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (126, 'Big king  (Burguer King®)', NULL, 0, 31, 555, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (127, 'Big tasty (Mc Donalds®)', NULL, 0, 43, 841, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (128, 'Biscoito de água e sal ', NULL, 8, 5, 34, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (129, 'Biscoito de aveia e mel Nestle®', NULL, 6, 4, 26, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (130, 'Biscoito de polvilho (rosquinha)', NULL, 3, 2, 13, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (131, 'Biscoito de queijo', NULL, 12, 6, 51, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (132, 'Biscoito leite Nestlé®', NULL, 8, 5, 38, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (133, 'Biscoito maisena Nestlé®', NULL, 5, 4, 22, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (134, 'Biscoito maria Nestlé®', NULL, 6, 4, 26, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (135, 'Biscoito milho verde Nestlé®', NULL, 6, 4, 27, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (136, 'Biscoito passatempo® sem recheio', NULL, 6, 4, 22, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (137, 'Biscoito prestígio recheado São Luiz  Nestlé®', NULL, 15, 10, 71, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (138, 'Biscoito prestígio® wafer Nestlé®', NULL, 8, 5, 43, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (139, 'Biscoito recheado chocolate', NULL, 13, 9, 62, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (140, 'Biscoito salclic aperitivo São Luiz Nestle®', NULL, 5, 3, 24, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (141, 'Biscoito salgado integral gergelim Piraquê®', NULL, 7.5, 5, 38, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (142, 'Biscoito suíço creme de avelã São Luiz  Nestlé®', NULL, 13, 7, 73, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (143, 'Biscoito tipo cookies Baunilha Toddy®', NULL, 10, 6, 46, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (144, 'Biscoito tostines cream cracker São Luiz Nestlé®', NULL, 8, 5, 37, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (145, 'Biscoito tostines® leite ', NULL, 8, 6, 37, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (146, 'Biscoito tostines® recheado Chocolate', NULL, 13, 9, 64, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (147, 'Biscoito tostines® surpresa fun', NULL, 8, 6, 37, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (148, 'Biscoito tostines® wafer chocolate', NULL, 8, 5, 42, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (149, 'Biscoitos de farinha integral', NULL, 10, 7, 40, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (150, 'Biscoitos de glúten a 40%', NULL, 10, 3, 29, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (151, 'Biscoitos de glúten puro', NULL, 10, 8, 35, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (152, 'Bisnaguinha Seven Boys®', NULL, 20, 11, 62, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (153, 'BKTM batata suprema (Burguer King®)', NULL, 0, 59, 604, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (154, 'BKTM cheddar duplo (Burguer King®)', NULL, 0, 54, 660, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (155, 'BKTM chicken crisp (Burguer King®)', NULL, 0, 68, 667, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (156, 'BKTM chicken crisp furioso (Burguer King®)', NULL, 0, 80, 930, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (157, 'BKTM chicken sandwich (Burguer King®)', NULL, 0, 53, 609, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (158, 'BKTM sh (Burguer King®)', NULL, 0, 63, 573, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (159, 'BKTM grilled chicken (Burguer King®)', NULL, 0, 56, 508, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (160, 'BKTM nuggets (Burguer King®)', NULL, 0, 12, 148, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (161, 'BKTM pepperoni (Burguer King®)', NULL, 0, 30, 270, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (162, 'BKTM picanha (Burguer King®)', NULL, 0, 53, 922, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (163, 'BKTM stacker triplo com bacon (Burguer King®)', NULL, 0, 35, 1529, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (164, 'Blis Frutas Vermelhas ', NULL, 180, 29, 158, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (165, 'Bliss morango', NULL, 200, 31, 168, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (166, 'Bobs crispy (Bobs®)', NULL, 251, 48, 641, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (167, 'Bobs picanha gourmet 100g (Bobs®)', NULL, 258, 35, 480, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (168, 'Bobs picanha gourmet 200g (Bobs®)', NULL, 373, 36, 727, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (169, 'Bolacha de nata Panco®', NULL, 5, 4, 22, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (170, 'Bolacha sabor mais morango Trakinas®', NULL, 12, 8, 61, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (171, 'Bolinha de queijo ', NULL, 10, 3, 27, 'unidade pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (172, 'Bolinho de aipim com carne seca', NULL, 45, 12, 86, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (173, 'Bolinho de arroz', NULL, 40, 20, 164, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (174, 'Bolinho de bacalhau', NULL, 15, 3, 42, 'unidade media');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (175, 'Bolinho de bacalhau (Habibs®)', NULL, 30, 5, 87, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (176, 'Bolinho de carne', NULL, 50, 8, 67, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (177, 'Bolinho de chuva ', NULL, 30, 13, 81, 'unidade pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (178, 'Bolinho de estudante (Bolinho Ana Maria®)', NULL, 80, 49, 292, 'unidade média ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (179, 'Bolinho de Presunto e Queijo', NULL, 50, 12, 73, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (180, 'Bolinho de soja tipo salsicha 350g (Taamti) - Albee®', NULL, 90, 28, 232, '4 unidades');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (181, 'Bolo alemão', NULL, 60, 30, 227, 'fatia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (182, 'Bolo branco simples', NULL, 100, 55, 318, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (183, 'Bolo de aipim com coco', NULL, 80, 37, 243, 'fatia media');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (184, 'Bolo de arroz', NULL, 70, 39, 197, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (185, 'Bolo de banana', NULL, 70, 33, 211, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (186, 'Bolo de batata-doce', NULL, 90, 43, 292, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (187, 'Bolo de Brigadeiro', NULL, 50, 20, 147, 'fatia pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (188, 'Bolo de brigadeiro (festa) Amor aos Pedaços®', NULL, 60, 33, 211, 'fatia pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (189, 'Bolo de casamento', NULL, 75, 42, 285, '01 fatia pequena ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (190, 'Bolo de cenoura', NULL, 60, 38, 227, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (191, 'Bolo de cenoura com chocolate Dr. Otcker®', NULL, 60, 38, 227, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (192, 'Bolo de chocolate (recheio/cobertura)', NULL, 100, 54, 320, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (193, 'Bolo de Chocolate e Nozes', NULL, 50, 17, 175, 'fatia pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (194, 'Bolo de chocolate sem glacê', NULL, 60, 34, 306, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (195, 'Bolo de festa (recheio/cobertura)', NULL, 100, 54, 320, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (196, 'Bolo de festa diet (recheio/cobertura)', NULL, 80, 20, 220, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (197, 'Bolo de fubá', NULL, 50, 20, 160, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (198, 'Bolo de limão', NULL, 60, 37, 233, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (199, 'Bolo de mandioca (aipim)', NULL, 100, 48, 324, 'pedaço grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (200, 'Bolo de milho ', NULL, 100, 54, 290, 'fatia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (201, 'Bolo de nozes ', NULL, 50, 28, 200, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (202, 'Bolo Floresta Negra com Morango', NULL, 50, 10, 106, 'fatia pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (203, 'Bolo Merengue', NULL, 50, 11, 116, 'fatia peqquena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (204, 'Bolo Mousse de Chocolate', NULL, 50, 20, 173, 'fatia pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (205, 'Bolo mousse de chocolate Amor aos Pedaços®', NULL, 60, 23, 185, 'fatia pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (206, 'Bolo simples', NULL, 60, 33, 263, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (207, 'Bombom Alpino', NULL, 13, 8, 70, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (208, 'Bombom banana Caribe Garoto®', NULL, 17, 12, 65, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (209, 'Bombom bopinho de torrone', NULL, 10, 6, 54, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (210, 'Bombom charge Nestlé®', NULL, 40, 24, 187, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (211, 'Bombom Copinho de Torrone', NULL, 10, 6, 54, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (212, 'Bombom de Brigadeiro', NULL, 10, 6, 48, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (213, 'Bombom diet Cacau Show®', NULL, 13, 7, 56, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (214, 'Bombom Ferrero Rocher®', NULL, 12, 6, 72, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (215, 'Bombom It coco Garoto®', NULL, 17, 11, 80, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (216, 'Bombom Ouro Branco®', NULL, 22, 13, 117, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (217, 'Bombom prestígio Nestle®', NULL, 33, 21, 154, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (218, 'Bombom Sonho de Valsa®', NULL, 22, 13, 113, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (219, 'Bombom Trufa de Cereja', NULL, 10, 6, 45, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (220, 'Bombom trua de creme de cereja Cacau Show®', NULL, 25, 12, 123, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (221, 'Brezel (Pretzel)', NULL, 28, 23, 110, '17 unidades pequenos');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (222, 'Brigadeiro', NULL, 30, 16, 103, '01 unidade médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (223, 'Brigadeiro de Flocos', NULL, 15, 9, 60, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (224, 'Brigadeiro de Morango', NULL, 15, 9, 60, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (225, 'Broa de Milho', NULL, 50, 25, 128, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (226, 'Brócolis com cottage (Baked Potato®)', NULL, 80, 3, 31, 'concha');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (227, 'Brócolis cozido (picado)', NULL, 10, 0, 4, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (228, 'Cacau em pó Mae Terra®', NULL, 10, 3, 31, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (229, 'Cachorro-quente (média)', NULL, 125, 31, 330, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (230, 'Café com leite sem açúcar', NULL, 200, 7, 88, 'xícara de chá cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (231, 'Café infusão 10%', NULL, 50, 1, 4, 'xicara de café ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (232, 'Café solúvel pó', NULL, 4, 1, 6, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (233, 'Caipirinha sem açúcar', NULL, 200, 54, 436, 'copo');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (234, 'Cajá manga', NULL, 75, 9, 38, 'unidade grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (235, 'Cajá polpa congelada', NULL, 100, 12, 46, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (236, 'Cajá-umbu', NULL, 22, 2, 10, 'unidade media ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (237, 'Cajú', NULL, 100, 10, 43, 'unidade grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (238, 'Caju polpa congelada', NULL, 100, 12, 46, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (239, 'Caju, suco concentrado', NULL, 240, 25, 108, 'copo duplo cheio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (240, 'Cajuzinho', NULL, 12, 6, 44, 'unidade pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (241, 'Calda de chocolate', NULL, 10, 5, 31, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (242, 'Caldo de carne Knorr®', NULL, 9, 2, 24, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (243, 'Caldo de galinha Knorr®', NULL, 9, 1, 24, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (244, 'Caldo verde', NULL, 130, 7, 79, 'concha média cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (245, 'Caldo-de-cana', NULL, 240, 49, 201, 'copo duplo');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (246, 'Camarão (Baked Potato®)', NULL, 90, 0, 42, 'concha');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (247, 'Camarão cozido', NULL, 30, 0, 27, 'unidade grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (248, 'Camarão cozido', NULL, 20, 0, 18, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (249, 'Camarão frito', NULL, 20, 0, 37, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (250, 'Camarão grande cru', NULL, 30, 0, 25, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (251, 'Camembert (Baked Potato®)', NULL, 60, 0, 188, '4 fatias');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (252, 'Camu-camu', NULL, 48, 3, 15, '6 unidades');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (253, 'Canapé de capaccio', NULL, 3, 2, 20, '01 torrada');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (254, 'Canela em pau Kitano®', NULL, 2, 1, 5, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (255, 'Canelone de frango', NULL, 45, 9, 87, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (256, 'Canelone de ricota ', NULL, 30, 7, 74, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (257, 'Canja de galinha', NULL, 130, 12, 110, 'concha média cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (258, 'Canjica pronta', NULL, 25, 5, 29, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (259, 'Capelete de carne Frescarini®', NULL, 50, 26, 141, 'escumadeira');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (260, 'Capelete de frango Frescarini®', NULL, 50, 26, 140, 'escumadeira');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (261, 'Cappuccino classic 3 corações® po', NULL, 20, 14, 84, '2 colheres de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (262, 'Cappuccino classic diet 3 corações® pó', NULL, 14, 6, 66, '2 colheres de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (263, 'Caqui', NULL, 110, 22, 95, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (264, 'Cará cozido', NULL, 35, 7, 28, 'colher de sopa cheia picado');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (265, 'Carambola', NULL, 100, 11, 46, 'unidade grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (266, 'Caranguejo cozido', NULL, 100, 0, 83, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (267, 'Carne assada', NULL, 90, 0, 259, 'pedaco medio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (268, 'Carne bovina, acem, moído, cru', NULL, 25, 0, 49, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (269, 'Carne bovina, costela assada', NULL, 100, 0, 373, 'pedaço médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (270, 'Carne bovina, picanha, com gordura, grelhada', NULL, 100, 0, 289, 'pedaço médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (271, 'Carne com chilli (Baked Potato®)', NULL, 65, 28, 98, 'concha');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (272, 'Carne de bezerro', NULL, 100, 0, 115, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (273, 'Carne de boi moída', NULL, 25, 0, 49, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (274, 'Carne de boi, lagarto cozido', NULL, 100, 0, 222, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (275, 'Carne de boi, maminha', NULL, 100, 0, 199, 'lé médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (276, 'Carne de boi, paleta', NULL, 100, 0, 307, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (277, 'Carne de boi, paleta cozida', NULL, 100, 0, 194, 'lé médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (278, 'Carne de boi, picanha (Friboi®)', NULL, 100, 1.89999999999999991, 323, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (279, 'Carne de cabrito magra', NULL, 100, 0, 179, 'pedaço médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (280, 'Carne de cordeiro magra', NULL, 100, 0, 162, 'pedaço médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (281, 'Carne de porco, lombo assado', NULL, 100, 0, 210, 'pedaço médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (282, 'Carne de vaca, maminha', NULL, 100, 0, 159, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (283, 'Carne ensopada com legumes', NULL, 35, 3, 57, 'colher sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (284, 'Carne vegetal (de soja)', NULL, 25, 2, 29, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (285, 'Carpa assada', NULL, 100, 0, 110, 'lé médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (286, 'Carpaccio de Haddock', NULL, 15, 0, 17, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (287, 'Carpaccio de salmão', NULL, 15, 0, 21, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (288, 'Caruru (hortaliça crua picada)', NULL, 25, 2, 9, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (289, 'Caruru (prato baiano)', NULL, 30, 1, 64, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (290, 'Casquinha de siri ', NULL, 55, 1, 92, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (291, 'Castanha de caju ', NULL, 3, 1, 13, '01 unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (292, 'Castanha de pequi', NULL, 14, 3, 12, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (293, 'Castanha do pará', NULL, 6, 1, 42, '01 unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (294, 'Castanha portuguesa', NULL, 10, 5, 21, '01 unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (295, 'Catchup', NULL, 15, 4, 15, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (296, 'Catupiry', NULL, 35, 0, 88, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (297, 'CBO (McDonalds®)', NULL, 0, 56, 665, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (298, 'Cebola picada', NULL, 10, 1, 4, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (299, 'Cebolinha crua (picada)', NULL, 8, 0, 2, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (300, 'Cenoura cozida', NULL, 25, 2, 8, 'colher de sopa cheia ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (301, 'Cenoura crua ralada', NULL, 12, 1, 4, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (302, 'Cenouritas (Mc Donalds®)', NULL, 70, 3, 18, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (303, 'Centeio, farinha clara de ', NULL, 102, 79, 355, 'xícara');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (304, 'Cereal barra Nestlé® banana, aveia e mel', NULL, 20, 11, 64, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (305, 'Cereal barra Nutry® coco com chocolate', NULL, 22, 16, 90, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (306, 'Cereal barra Trio® light morango com chocolate', NULL, 20, 15, 78, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (307, 'Cereal de arroz e milho Nestlé®', NULL, 7, 6, 27, 'colher de sopa ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (308, 'Cereal infantil 3 cereais Nestlé®', NULL, 7, 5, 29, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (309, 'Cereal infantil 3 frutas Nestlé®', NULL, 7, 6, 27, 'colher de sopa ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (310, 'Cereal matinal  chocokrispis Kelloggs® ', NULL, 30, 25, 110, '3/4 de xícara');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (311, 'Cereal matinal all bran Kelloggs® ', NULL, 40, 18, 108, '3/4 de xícara');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (312, 'Cereal matinal corn flakes Kelloggs®', NULL, 30, 24, 108, '1 xícara');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (313, 'Cereal matinal crunch Nestlé®', NULL, 30, 23, 120, '3/4 xícara');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (314, 'Cereal matinal Kellness® granola tradicional ', NULL, 40, 27, 147, '1/2 xícara');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (315, 'Cereja em conserva', NULL, 5, 2, 8, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (316, 'Cereja fresca', NULL, 4, 1, 4, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (317, 'Cerveja ', NULL, 350, 12, 143, 'lata');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (318, 'Cevada, infuso de', NULL, 200, 5, 24, 'xícara de chá cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (319, 'Chá mate (infusão sem açúcar)', NULL, 200, 1, 6, 'xícara de chá cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (320, 'Chá sem açúcar (média)', NULL, 200, 1, 4, 'xícara de chá cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (321, 'Chambinho', NULL, 45, 4, 40, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (322, 'Champanhe', NULL, 100, 3, 82, 'taça');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (323, 'Champanhe tipo sidra', NULL, 100, 3, 82, 'taça');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (324, 'Champignon (Backed Potato®)', NULL, 60, 1, 6, 'concha');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (325, 'Champignon (conserva)', NULL, 10, 0, 2, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (326, 'Chandelle chocolate', NULL, 75, 18, 151, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (327, 'Chantilly', NULL, 20, 4, 60, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (328, 'Charutinho de folha de uva', NULL, 100, 3, 99, 'porcao');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (329, 'Charutinho de repolho', NULL, 100, 2, 67, 'porcao');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (330, 'Charuto folha de uva Habibs®', NULL, 245, 19, 263, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (331, 'Charuto repolho Habibs®', NULL, 285, 23, 295, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (332, 'Cheddar McMelt (Mc Donalds®)', NULL, 0, 29, 481, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (333, 'Cheeseburger', NULL, 140, 40, 358, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (334, 'Cheeseburguer (Bobs®)', NULL, 120, 30, 329, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (335, 'Cheeseburguer (Burguer King®)', NULL, 116, 32, 280, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (336, 'Cheeseburguer (Mc Donalds®)', NULL, 0, 30, 295, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (337, 'Cheeseburguer com bacon (Burguer King®)', NULL, 119, 32, 310, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (338, 'Cheeseburguer duplo com bacon (Burguer King®)', NULL, 160, 32, 440, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (339, 'Cheesecake com Mirtilo® (sorvete)', NULL, 60, 18, 140, '01 bola');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (340, 'Cheetos sabor natural', NULL, 25, 17, 121, '1 e 1/2 xícara');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (341, 'Chester', NULL, 15, 0, 31, 'fatia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (342, 'Chicken bacon crispy McWrap (Mc Donalds®)', NULL, 249, 54, 565, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (343, 'Chicken bacon grilled McWrap (Mc Donalds®)', NULL, 235, 32, 440, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (344, 'Chicken classic crispy Sandwich (Mc Donalds®)', NULL, 0, 55, 510, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (345, 'Chicken Junior Sandwich (Mc Donalds®)', NULL, 0, 38, 380, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (346, 'Chiclete', NULL, 5, 4, 16, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (347, 'Chiclete sabor sortido Bubbaloo®', NULL, 5, 4, 16, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (348, 'Chiclete Trident® tuti-frutti', NULL, 2, 1, 3, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (349, 'Chicória refogada', NULL, 45, 3, 39, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (350, 'Chimarrão', NULL, 200, 0, 12, 'cuia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (351, 'Chocolate alpino diet Nestlé®', NULL, 30, 17, 143, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (352, 'Chocolate ao leite', NULL, 30, 18, 160, 'barra pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (353, 'Chocolate ao leite diet', NULL, 30, 13, 157, 'barra pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (354, 'Chocolate ao leite diet Gold®', NULL, 25, 14, 122, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (355, 'Chocolate Batom®', NULL, 16, 10, 86, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (356, 'Chocolate Bis®', NULL, 7, 5, 37, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (357, 'Chocolate Classic Zero açúcar', NULL, 22, 8, 94, '1 unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (358, 'Chocolate Diamante Negro®', NULL, 30, 19, 156, 'barra pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (359, 'Chocolate em pó', NULL, 15, 9, 52, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (360, 'Chocolate Ferrero Rocher®', NULL, 13, 6, 71, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (361, 'Chocolate quente com leite (Mc Donalds®)', NULL, 250, 24, 178, 'copo grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (362, 'Chocolate Suflair® ', NULL, 30, 15, 132, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (363, 'Chocolate Talento®', NULL, 25, 13, 137, '1/4 barra');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (364, 'Chocolate Talento® diet', NULL, 25, 12, 129, 'barra pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (365, 'Chocotone Bauducco®', NULL, 100, 49, 396, '01 fatia pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (366, 'Chopp escuro', NULL, 290, 11, 148, 'tulipa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (367, 'Chouriço', NULL, 60, 1, 227, 'gomo');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (368, 'Chuchu à milanesa', NULL, 70, 10, 127, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (369, 'Chuchu ao molho branco', NULL, 30, 3, 28, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (370, 'Chucrute', NULL, 120, 5, 24, 'xícara');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (371, 'Churros com doce de leite', NULL, 30, 13, 97, 'unidade aperitivo');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (372, 'Cidra', NULL, 100, 10, 40, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (373, 'Ciriguela', NULL, 20, 4, 15, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (374, 'Coalhada integral Itambé®', NULL, 130, 17, 120, 'pote');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (375, 'Coalhada seca natural Alibey®', NULL, 15, 1, 19, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (376, 'Cobertura de chocolate ao leite Garoto®', NULL, 25, 14, 140, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (377, 'Coca-Cola®', NULL, 200, 20, 80, 'copo');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (378, 'Cocada light Santa Helena®', NULL, 70, 37, 405, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (379, 'Cocada queimada Brasil Caipira®', NULL, 20, 11, 86, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (380, 'Coco fresco ralado', NULL, 100, 15, 354, 'xícara');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (381, 'Coco-da-baía, água-de', NULL, 200, 10, 44, 'copo');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (382, 'Cogumelo em conserva', NULL, 27, 1, 5, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (383, 'Colomba Pascal gotas de chocolate (Bauducco®)', NULL, 80, 40, 350, 'fatia ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (384, 'Colomba Pascal salgada', NULL, 100, 20, 228, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (385, 'Conhaque', NULL, 50, 0, 115, 'dose');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (386, 'Cookie aveia e passas (Subway®)', NULL, 33, 28, 190, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (387, 'Cookie cheescake com framboesa (Subway®)', NULL, 33, 28, 200, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (388, 'Cookie chocolate  (Subway®)', NULL, 33, 29, 210, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (389, 'Cookies integral quinua e banana Taeq®', NULL, 30, 17, 123, '6 unidades');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (390, 'Coração de frango', NULL, 5, 0, 7, '01 unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (391, 'Coração de galinha cozido', NULL, 5, 0, 8, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (392, 'Costela de boi assada', NULL, 30, 0, 136, 'pedaço pequeno');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (393, 'Couve crua', NULL, 20, 1, 10, 'folha média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (394, 'Couve refogada', NULL, 25, 4, 36, 'folha média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (395, 'Couve-flor à milanesa', NULL, 90, 11, 136, 'ramo médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (396, 'Couve-flor cozida', NULL, 60, 4, 25, 'ramo médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (397, 'Coxa de frango assada', NULL, 100, 0, 215, 'média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (398, 'Cream Cheese ', NULL, 30, 1, 84, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (399, 'Creme de abacate sem açúcar', NULL, 25, 3, 43, 'colher sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (400, 'Creme de amendoim', NULL, 20, 4, 125, 'colher de sopa rasa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (401, 'Creme de leite', NULL, 15, 1, 46, 'colher de sopa rasa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (402, 'Creme de milho', NULL, 33, 5, 35, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (403, 'Creme vegetal de chantilly Vigor®', NULL, 20, 2, 56, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (404, 'Cremogema chocolate', NULL, 20, 17, 72, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (405, 'Crepe de maçã diet', NULL, 100, 18, 110, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (406, 'Croissant', NULL, 80, 32, 328, 'unidade grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (407, 'Croissant de chocolate', NULL, 57, 24, 237, 'unidade pequena ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (408, 'Croissant de queijo', NULL, 80, 32, 388, 'unidade média ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (409, 'Croquete de carne/milho', NULL, 10, 4, 35, 'unidade pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (410, 'Cuca alemã', NULL, 100, 33, 209, 'fatia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (411, 'Cup noodles®', NULL, 65, 38, 300, 'unidade comercial');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (412, 'Cupuaçu (creme)', NULL, 20, 5, 135, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (413, 'Curau de milho', NULL, 140, 33, 162, '2/3 xícara de chá');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (414, 'Curry', NULL, 9, 5, 34, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (415, 'Cuscuz de milho ', NULL, 85, 34, 161, 'pedaço pequeno');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (416, 'Cuscuz paulista', NULL, 80, 34, 161, '01 fatia pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (417, 'Damasco seco', NULL, 7, 2, 9, '01 unidade  pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (418, 'Damasco, geléia de', NULL, 20, 13, 52, 'colher de sopa ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (419, 'Danette®', NULL, 110, 25, 166, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (420, 'Dobradinha', NULL, 100, 8, 111, 'concha pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (421, 'Doce de abacaxi', NULL, 40, 32, 130, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (422, 'Doce de abóbora cremoso', NULL, 40, 22, 80, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (423, 'Doce de abóbora e coco', NULL, 40, 18, 83, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (424, 'Doce de arroz de leite', NULL, 40, 13, 66, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (425, 'Doce de batata doce', NULL, 30, 22, 121, 'porção pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (426, 'Doce de buriti', NULL, 50, 41, 165, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (427, 'Doce de cidra', NULL, 50, 20, 81, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (428, 'Doce de coco', NULL, 40, 21, 219, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (429, 'Doce de cupuaçu', NULL, 40, 30, 118, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (430, 'Doce de goiaba', NULL, 50, 21, 86, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (431, 'Doce de goiaba em calda Diet House®', NULL, 20, 6, 23, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (432, 'Doce de laranja azeda', NULL, 50, 40, 153, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (433, 'Doce de laranja Diet House®', NULL, 20, 2, 9, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (434, 'Doce de leite', NULL, 40, 22, 116, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (435, 'Doce de limão', NULL, 50, 27, 107, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (436, 'Doce de mamão verde', NULL, 40, 19, 78, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (437, 'Doce de manga', NULL, 50, 27, 140, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (438, 'Doce de ovos', NULL, 40, 6, 52, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (439, 'Docinhos variados (média)', NULL, 30, 16, 125, '01 unidade média ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (440, 'Doriana® cremosa', NULL, 10, 0, 59, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (441, 'Doriana® light', NULL, 10, 0, 32, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (442, 'Double cheddar (Bobs®)', NULL, 210, 32, 484, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (443, 'Double cheddar (Habibs®)', NULL, 335, 35, 668, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (444, 'Double cheeseburguer (Bobs®)', NULL, 185, 33, 512, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (445, 'Double grill bacon (Bobs®)', NULL, 488, 58, 1314, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (446, 'Drops comum', NULL, 3, 3, 11, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (447, 'Empada de frango', NULL, 55, 14, 197, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (448, 'Empada de Palmito', NULL, 55, 14, 129, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (449, 'Empadão goiano', NULL, 300, 48, 618, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (450, 'Empadinha', NULL, 12, 4, 56, 'unidade pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (451, 'Enrolado de presunto e queijo', NULL, 40, 15, 114, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (452, 'Enrolado de salsicha', NULL, 27, 3, 79, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (453, 'Ervilha em conserva ', NULL, 27, 4, 20, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (454, 'Ervilha verde cozida', NULL, 30, 4, 26, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (455, 'Esha de carne aberta', NULL, 75, 18, 125, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (456, 'Esha de carne aberta', NULL, 75, 18, 125, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (457, 'Esha de carne fechada', NULL, 80, 29, 203, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (458, 'Esha de espinafre (Habibs®)', NULL, 80, 17, 122, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (459, 'Esha folhada de cheddar com pepperoni (Habibs®)', NULL, 55, 11, 201, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (460, 'Esha folhada de cheddar com pepperoni (Habibs®)', NULL, 55, 11, 201, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (461, 'Esha folhada de chocolate (Habibs®)', NULL, 45, 16, 199, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (462, 'Esha folhada de chocolate (Habibs®)', NULL, 45, 16, 199, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (463, 'Esha folhada de chocolate com M&Ms® (Habibs®)', NULL, 65, 28, 295, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (464, 'Espinafre cru', NULL, 20, 0, 4, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (465, 'Espinafre refogado', NULL, 25, 2, 27, 'colher do sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (466, 'Estrogono de lé mignon (Backed Potato®)', NULL, 85, 0, 146, 'concha');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (467, 'Estrogonoe de carne', NULL, 25, 1, 39, 'colher do sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (468, 'Estrogonoe de frango', NULL, 25, 1, 43, 'colher do sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (469, 'Extrato de malte', NULL, 15, 10, 42, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (470, 'Falafel (bolinho de grão de bico)', NULL, 100, 20, 155, 'porção ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (471, 'Fanta® laranja', NULL, 200, 22, 90, 'copo');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (472, 'Farelo de trigo', NULL, 9, 5, 28, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (473, 'Farinha dágua-do-pará', NULL, 50, 41, 166, 'copo descartável pequeno');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (474, 'Farinha de arroz', NULL, 17, 14, 60, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (475, 'Farinha de centeio integral', NULL, 15, 11, 54, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (476, 'Farinha de mandioca', NULL, 16, 14, 57, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (477, 'Farinha de milho', NULL, 15, 12, 54, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (478, 'Farinha de milho integral', NULL, 15, 11, 53, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (479, 'Farinha de rosca', NULL, 15, 11, 61, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (480, 'Farinha de tapioca com coco e açúcar', NULL, 25, 21, 107, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (481, 'Farinha de trigo ', NULL, 20, 15, 71, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (482, 'Farinha láctea (Nestlé®)', NULL, 8, 6, 30, 'colher de sopa rasa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (483, 'Farofa', NULL, 15, 12, 71, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (484, 'Farofa com Bacon', NULL, 45, 26, 206, '01 colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (485, 'Farofa com lingüiça', NULL, 15, 7, 54, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (486, 'Farofa com tempero/óleo', NULL, 15, 12, 71, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (487, 'Farofa de farinha de mandioca', NULL, 25, 20, 96, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (488, 'Fatouche', NULL, 100, 5, 163, '1 porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (489, 'Feijão branco cozido', NULL, 17, 4, 20, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (490, 'Feijão carioquinha cozido', NULL, 17, 3, 19, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (491, 'Feijão preto cozido', NULL, 17, 2, 12, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (492, 'Feijão tropeiro', NULL, 15, 7, 50, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (493, 'Feijão-fradinho', NULL, 17, 2, 13, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (494, 'Feijoada caseira', NULL, 225, 24, 346, 'concha média cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (495, 'Fermento  biológico', NULL, 15, 1, 14, 'tablete');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (496, 'Fermento em pó', NULL, 10, 4, 17, 'colher de sopa rasa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (497, 'Fettuccine', NULL, 110, 24, 150, '01 pegador');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (498, 'Fettuccine a bolonhesa (Perdigão®)', NULL, 350, 49, 458, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (499, 'Fígado de boi grelhado', NULL, 100, 4, 225, 'bife médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (500, 'Fígado de galinha cru', NULL, 30, 1, 53, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (501, 'Figo', NULL, 55, 8, 38, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (502, 'Figo cristalizado', NULL, 30, 22, 86, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (503, 'Figo enlatado em calda', NULL, 20, 10, 38, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (504, 'Filé a parmeggiana', NULL, 150, 13, 490, '01 unidade pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (505, 'File ao molho  madeira', NULL, 180, 12, 227, '01 unidade pequena ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (506, 'Flocos de milho', NULL, 10, 9, 38, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (507, 'Focaccia', NULL, 50, 22, 136, 'fatia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (508, 'Fogazza de calabresa (Habibs®)', NULL, 70, 22, 186, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (509, 'Fogazza de mussarela (Habibs®)', NULL, 70, 21, 200, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (510, 'Folhado de frango', NULL, 40, 9, 109, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (511, 'Fondue de carne', NULL, 90, 0, 180, 'bife médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (512, 'Fondue de chocolate ', NULL, 30, 15, 105, 'colher sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (513, 'Fondue de queijo (Gramado®)', NULL, 30, 1, 70, 'colher de sopa (molho)');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (514, 'Framboesa', NULL, 15, 2, 8, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (515, 'Framboesa, doce em pasta', NULL, 50, 35, 143, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (516, 'Framboesa, geléia de', NULL, 34, 24, 97, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (517, 'Frango à milanesa', NULL, 100, 14, 311, 'lé pequeno ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (518, 'Frango assado', NULL, 65, 0, 78, 'sobrecoxa média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (519, 'Frango com requeijão (Backed Potato®)', NULL, 85, 0, 110, 'concha');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (520, 'Frango cozido', NULL, 65, 0, 82, 'sobrecoxa média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (521, 'Frango lé cozido', NULL, 100, 0, 163, 'lé pequeno ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (522, 'Frango frito', NULL, 65, 0, 94, 'sobrecoxa média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (523, 'Franlitos (Bobs®)', NULL, 120, 28, 272, 'porção de 6 unidades');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (524, 'Frigideira de repolho com camarão seco', NULL, 25, 1, 25, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (525, 'Fruta - pão', NULL, 100, 24, 96, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (526, 'Fruta-de-conde ou pinha', NULL, 60, 8, 41, 'unidade pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (527, 'Frutas cristalizadas industrializadas', NULL, 15, 12, 48, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (528, 'Fubá', NULL, 20, 16, 71, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (529, 'Funghi', NULL, 20, 10, 71, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (530, 'Galeto assado', NULL, 95, 0, 115, 'pedaço pequeno');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (531, 'Galinha ao molho pardo', NULL, 95, 1, 150, 'pedaço médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (532, 'Galinhada com pequi', NULL, 60, 13, 108, 'colher de servir');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (533, 'Gatorade® - média sabores', NULL, 500, 24, 120, 'garrafa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (534, 'Gelatina de frutas em pó', NULL, 25, 22, 97, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (535, 'Gelatina diet em pó ', NULL, 14, 0, 1, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (536, 'Gelatina em pó com açúcar', NULL, 14, 12, 55, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (537, 'Gelatina light  morango preparada', NULL, 25, 0, 3, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (538, 'Geléia de damasco diet', NULL, 22, 6, 19, 'colher de sopa rasa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (539, 'Geléia de frutas (média)', NULL, 30, 18, 74, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (540, 'Geléia de mocotó', NULL, 40, 12, 64, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (541, 'Geléia de mocotó dietética', NULL, 40, 2, 22, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (542, 'Geléia de morango diet', NULL, 22, 6, 25, 'colher de sopa rasa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (543, 'Germen de trigo', NULL, 10, 4, 37, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (544, 'Glucose de milho', NULL, 9, 6, 26, 'colher de sobremesa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (545, 'Goiaba', NULL, 170, 20, 96, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (546, 'Goiabada', NULL, 40, 27, 109, 'fatia pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (547, 'Goma de tapioca', NULL, 20, 11, 46, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (548, 'Gordura vegetal hidrogenada', NULL, 14, 0, 126, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (549, 'Gorgonzola (Backed Potato®)', NULL, 60, 0, 288, 'concha');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (550, 'Gran picanha 100g (Bobs®)', NULL, 233, 27, 526, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (551, 'Gran picanha 200g (Bobs®)', NULL, 348, 27, 755, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (552, 'Granola', NULL, 11, 7, 51, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (553, 'Grão-de-bico cozido', NULL, 22, 4, 25, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (554, 'Graviola', NULL, 100, 16, 62, 'porção ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (555, 'Guaraná diet', NULL, 240, 0, 0, 'copo duplo');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (556, 'Guaraná Jesus®', NULL, 350, 42, 168, 'lata');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (557, 'Guaraná Kuat®', NULL, 240, 25, 98, 'copo duplo');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (558, 'Guaraná refrigerante', NULL, 240, 24, 96, 'copo duplo');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (559, 'Guariroba', NULL, 15, 2, 10, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (560, 'Guariroba refogada', NULL, 60, 1, 22, '2 colheres de sopa cheias');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (561, 'Hamburguer - sanduíche', NULL, 125, 40, 301, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (562, 'Hamburguer (Mc Donalds®)', NULL, 0, 29, 245, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (563, 'Hamburguer de carne bovina', NULL, 56, 2, 139, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (564, 'Hamburguer de frango', NULL, 56, 2, 131, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (565, 'Hamburguer de peru', NULL, 56, 0, 82, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (566, 'Hamburguer simples ', NULL, 0, 29, 245, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (567, 'Harumaki de carne', NULL, 50, 8, 103, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (568, 'Harumaki de carne', NULL, 50, 8, 103, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (569, 'Harumaki ladéla', NULL, 50, 8, 158, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (570, 'Harumaki ladéla', NULL, 50, 8, 158, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (571, 'Herbalife de baunilha', NULL, 15, 7, 51, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (572, 'Hipoglosso (peixe) cozido', NULL, 100, 0, 127, 'lé médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (573, 'Homus', NULL, 30, 14, 94, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (574, 'Homus (Habibs®)', NULL, 240, 34, 360, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (575, 'Hossomaki de atum', NULL, 110, 26, 164, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (576, 'Hossomaki de pepino', NULL, 110, 26, 123, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (577, 'Hossomaki de pepino', NULL, 110, 26, 123, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (578, 'Hossomaki de salmão', NULL, 110, 26, 155, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (579, 'Inhame, raiz sem casca de (picado)', NULL, 35, 8, 37, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (580, 'Iogurte Activia® light', NULL, 100, 11, 62, 'pote');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (581, 'Iogurte Activia® natural', NULL, 170, 12, 132, 'pote');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (582, 'Iogurte Activia® original', NULL, 100, 16, 110, 'pote');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (583, 'Iogurte Corpus® de morango light', NULL, 170, 8, 69, 'garranha');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (584, 'Iogurte Dan up®', NULL, 180, 28, 157, 'garranha');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (585, 'Iogurte de frutas (média)', NULL, 100, 17, 90, 'pote');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (586, 'Iogurte de frutas light (média)', NULL, 100, 7, 42, 'pote');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (587, 'Iogurte Grego Nestlé® light', NULL, 90, 10, 78, 'pote');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (588, 'Iogurte Grego Nestlé® tradicional', NULL, 100, 15, 113, 'pote');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (589, 'Iogurte Grego Vigor®', NULL, 100, 16, 151, 'pote');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (590, 'Iogurte Grego Vigor® zero', NULL, 100, 5, 49, 'pote');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (591, 'Iogurte Molico® total cálcio líquido (média)', NULL, 170, 8, 73, 'garranha');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (592, 'Iogurte Molico® total cálcio polpa (média)', NULL, 100, 6, 42, 'pote');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (593, 'Iogurte natural c/ mel', NULL, 200, 36, 225, 'copo');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (594, 'Iogurte natural desnatado (Fiore®) ', NULL, 140, 6, 45, 'copo');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (595, 'Iogurte natural desnatado (média)', NULL, 185, 12, 85, 'copo');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (596, 'Iogurte natural integral (média)', NULL, 200, 12, 142, 'copo');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (597, 'Iogurte petit suisse (média)', NULL, 45, 9, 71, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (598, 'Isca de carne de porco', NULL, 30, 0, 67, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (599, 'Isca frita de pirarucu', NULL, 100, 16, 258, 'escumadeira cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (600, 'Jabuticaba', NULL, 5, 1, 2, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (601, 'Jaca  ', NULL, 12, 1, 3, 'bago  médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (602, 'Jambo', NULL, 40, 5, 22, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (603, 'Jambu cozido', NULL, 15, 2, 8, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (604, 'Jamelão ou jambolão', NULL, 10, 1, 4, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (605, 'Jaraqui cru', NULL, 100, 0, 129, 'pedaço médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (606, 'Jenipapo', NULL, 100, 26, 113, 'pedaço médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (607, 'Jiló', NULL, 60, 5, 31, 'colher (sopa) cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (608, 'Joelho de porco (Einsbein®)', NULL, 100, 0, 215, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (609, 'Kafta na bandeja (Habibs®)', NULL, 250, 14, 300, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (610, 'Kani kama cru', NULL, 16, 0, 13, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (611, 'Kanimaki ', NULL, 15, 4, 20, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (612, 'Kappamaki de pepino', NULL, 15, 4, 18, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (613, 'Karo', NULL, 15, 12, 49, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (614, 'Kasespatzle (macarrão com queijo)', NULL, 130, 20, 198, 'pegador');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (615, 'Ketchup de tomate', NULL, 20, 5, 23, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (616, 'Kinder ovo® maxi (Ovo de Páscoa)', NULL, 25, 13, 143, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (617, 'Kiwi', NULL, 76, 11, 51, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (618, 'Lagosta à Thermidor', NULL, 100, 13, 460, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (619, 'Laranja lima', NULL, 90, 10, 41, 'unidade media');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (620, 'Lasanha à bolonhesa', NULL, 170, 27, 355, 'escumadeira');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (621, 'Lasanha de bacalhau e espinafre', NULL, 200, 42, 503, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (622, 'Leite condensado', NULL, 15, 8, 49, 'colher de sopa ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (623, 'Leite condensado desnatado', NULL, 15, 9, 42, 'colher de sopa ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (624, 'Leite de cabra integral', NULL, 240, 13, 223, 'copo duplo cheio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (625, 'Leite de coco industrializado ', NULL, 200, 11, 516, 'garrafa pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (626, 'Leite de coco industrializado light', NULL, 200, 6, 250, 'garrafa pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (627, 'Leite de soja Ades®', NULL, 200, 8, 82, 'copo ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (628, 'Leite de soja zero Ades®', NULL, 200, 3, 59, 'copo');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (629, 'Leite de vaca desnatado', NULL, 240, 12, 84, 'copo duplo cheio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (630, 'Leite de vaca desnatado em pó', NULL, 10, 5, 35, 'colher (sopa) cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (631, 'Leite de vaca integral  em pó ', NULL, 16, 6, 80, 'colher (sopa) cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (632, 'Leite de vaca integral pasteurizado', NULL, 240, 12, 141, 'copo duplo cheio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (633, 'Lentilha cozida', NULL, 18, 3, 19, 'colher sopa (cheia)');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (634, 'Licor', NULL, 20, 6, 62, 'cálice');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (635, 'Limão', NULL, 12, 1, 4, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (636, 'Limão, geléia de', NULL, 15, 1, 7, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (637, 'Limonada s/ açúcar ', NULL, 200, 10, 28, 'copo');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (638, 'Língua de boi cozida', NULL, 30, 0, 94, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (639, 'Língua de gato Kompenhagen®', NULL, 6, 3, 36, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (640, 'Linguado assado', NULL, 120, 0, 140, 'posta média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (641, 'Linguiça calabresa', NULL, 40, 0, 146, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (642, 'Linguiça calabresa defumada', NULL, 40, 0, 146, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (643, 'Linguiça de frango', NULL, 60, 0, 146, 'gomo');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (644, 'Linguiça de peru', NULL, 60, 0, 83, 'gomo');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (645, 'Linguiça paio', NULL, 160, 0, 363, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (646, 'Lombo de porco assado', NULL, 120, 0, 252, 'pedaço grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (647, 'Lombo de vitela assado/ cozido', NULL, 50, 0, 113, 'pedaço médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (648, 'Lula cozida', NULL, 80, 0, 74, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (649, 'M&M®', NULL, 12, 9, 61, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (650, 'M&M´s® amendoim', NULL, 12, 8, 64, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (651, 'M&M´s® chocolate', NULL, 12, 9, 61, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (652, 'Maçã com casca', NULL, 90, 14, 58, 'unidade pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (653, 'Maçã, suco de', NULL, 200, 22, 97, 'copo ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (654, 'Macadâmica natural ', NULL, 15, 3, 97, 'colher de sopa ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (655, 'Macarrão à bolonhesa', NULL, 110, 22, 136, 'escumadeira cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (656, 'Macarrão ao alho e óleo', NULL, 110, 35, 241, 'escumadeira cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (657, 'Macarrão caseiro cozido', NULL, 50, 11, 54, 'colher de arroz cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (658, 'Macarrão instantâneo Maggi® lámen queijo', NULL, 80, 50, 370, 'pacote');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (659, 'Macarrão mini fusilli picolini Barrila®', NULL, 80, 59, 280, 'prato');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (660, 'Macaúba crua', NULL, 100, 14, 404, '2 unidades');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (661, 'Maizena', NULL, 20, 17, 70, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (662, 'Mamão formosa', NULL, 100, 12, 45, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (663, 'Mamão papaya', NULL, 160, 16, 64, 'meia unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (664, 'Mamão papaya', NULL, 40, 3, 14, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (665, 'Mamão verde, doce de', NULL, 40, 19, 78, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (666, 'Mandioca cozida (picada)', NULL, 30, 9, 37, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (667, 'Mandioca frita', NULL, 100, 50, 300, 'pedaço grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (668, 'Mandioquinha', NULL, 25, 5, 20, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (669, 'Mané pelado', NULL, 70, 37, 226, 'pedaço médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (670, 'Manga espada', NULL, 140, 22, 91, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (671, 'Manjar ', NULL, 90, 36, 255, 'fatia media');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (672, 'Manteiga ', NULL, 32, 0, 235, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (673, 'Maracujá', NULL, 45, 6, 31, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (674, 'Margarina', NULL, 32, 0, 219, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (675, 'Maria mole', NULL, 44, 33, 132, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (676, 'Marmelada', NULL, 60, 40, 158, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (677, 'Marreco recheado assado', NULL, 100, 3, 258, '1/4 unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (678, 'Marron-glacê', NULL, 60, 35, 149, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (679, 'Marshmelow Dr. Oetker®', NULL, 10, 9, 40, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (680, 'Martini', NULL, 50, 2, 41, 'dose');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (681, 'Massa fresca semipronta para pizza', NULL, 140, 80, 441, 'unidade (inteira)');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (682, 'Massa fresca semipronta para pizza brotinho', NULL, 40, 23, 126, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (683, 'Massa pronta para pastel ', NULL, 30, 16, 83, '2 discos de massa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (684, 'Matrinxã inteiro cru', NULL, 100, 0, 246, 'pedaço médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (685, 'McChicken (Mc Donalds®)', NULL, 177, 40, 424, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (686, 'McColosso caramelo (Mc Donalds®)', NULL, 142, 52, 291, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (687, 'McColosso chocolate  (Mc Donalds®)', NULL, 137, 47, 274, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (688, 'McFish (Mc Donalds®)', NULL, 151, 37, 362, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (689, 'McFlurry ovomaltine (Mc Donalds®)', NULL, 200, 68, 425, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (690, 'McNíco (Mc Donalds®)', NULL, 282, 37, 600, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (691, 'McNuggets 10 (Mc Donalds®)', NULL, 166, 27, 431, 'porção de 10 unidades');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (692, 'McNuggets 4 (Mc Donalds®)', NULL, 66, 11, 173, 'porção de 4 unidades');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (693, 'McNuggets 6 (Mc Donalds®)', NULL, 0, 16, 259, 'porção de 6 unidades');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (694, 'Melabie', NULL, 100, 25, 162, '1 porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (695, 'Melado', NULL, 16, 12, 47, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (696, 'Melancia', NULL, 200, 11, 48, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (697, 'Melão', NULL, 115, 7, 32, 'fatia grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (698, 'Melão cantalupo', NULL, 100, 8, 34, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (699, 'Melão gaúcho', NULL, 100, 7, 29, 'fatia pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (700, 'Merengue de morango', NULL, 60, 52, 214, 'porção média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (701, 'Merluza assada', NULL, 100, 0, 122, 'lé médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (702, 'Michui de Filé Mingnon', NULL, 100, 2, 88, '1 porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (703, 'Michui de frango', NULL, 100, 1, 132, ' porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (704, 'Michui de Frango', NULL, 100, 2, 90, '1 porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (705, 'Milho em espiga com 1 c.sobremesa de  manteiga ', NULL, 100, 32, 233, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (706, 'Milho verde em conserva enlatado', NULL, 24, 5, 23, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (707, 'Milk shake clássico de chocolate (Bobs®)', NULL, 700, 177, 846, 'copo grande ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (708, 'Milk shake clássico de morango (Bobs®)', NULL, 700, 136, 734, 'copo grande ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (709, 'Milk shake crocante de ovomaltine (Bobs®)', NULL, 700, 157, 852, 'copo grande ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (710, 'Milk shake de baunilha', NULL, 290, 51, 330, 'copo de milk shake');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (711, 'Milk shake de chocolate', NULL, 290, 61, 350, 'copo de milk shake');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (712, 'Minestrone (sopa)', NULL, 300, 23, 283, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (713, 'Mingau (média)', NULL, 37, 8, 50, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (714, 'Mini bolacha sabor chocolate Trakinas®', NULL, 12, 8, 58, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (715, 'Mini bomba de avelã Cristallo®', NULL, 15, 6, 41, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (716, 'Mini Bomba de Caramelo', NULL, 10, 3, 25, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (717, 'Mini Bomba de Chocolate', NULL, 10, 2, 23, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (718, 'Mini bomba de chocolate Cristallo®', NULL, 15, 6, 43, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (719, 'Mini cenouras', NULL, 125, 13, 54, 'xícara de café');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (720, 'Mini churro de doce de Leite (Habibs®)', NULL, 20, 9, 56, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (721, 'Mini hamburguinho', NULL, 60, 15, 123, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (722, 'Mini hot dog', NULL, 60, 15, 165, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (723, 'Mini kibe frito de cremilly (Habibs®)', NULL, 45, 8, 101, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (724, 'Mini pizza', NULL, 60, 20, 182, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (725, 'Mini Torta de Limão', NULL, 10, 5, 32, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (726, 'Mini Torta de Maçã', NULL, 10, 3, 21, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (727, 'Mini Torta de Morango', NULL, 10, 3, 21, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (728, 'Miojo sabor carne suave Nissin®', NULL, 85, 49, 374, 'pacote');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (729, 'Miojo sabor galinha Nissin®', NULL, 85, 50, 369, 'pacote');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (730, 'Miojo sabor legumes Nissin®', NULL, 85, 51, 372, 'pacote');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (731, 'Miolos', NULL, 25, 0, 30, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (732, 'Missoshiro', NULL, 300, 9, 70, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (733, 'Misto-quente', NULL, 85, 29, 283, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (734, 'Miúdos de boi (dobradinha, livrelho)', NULL, 35, 0, 44, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (735, 'Mix de frutas desidratadas Frutolla®', NULL, 25, 6, 34, '2 colheres de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (736, 'Mjadra', NULL, 52, 52, 314, '1 porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (737, 'Moela', NULL, 18, 0, 29, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (738, 'Molho à bolonhesa', NULL, 22, 2, 41, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (739, 'Molho agridoce (Mc Donalds®)', NULL, 28, 10, 45, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (740, 'Molho barbecue  MasterFoods®', NULL, 25, 5, 18, '2 colheres de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (741, 'Molho barbecue (Mc Donalds®)', NULL, 28, 10, 46, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (742, 'Molho branco ', NULL, 35, 5, 70, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (743, 'Molho caipira (Mc Donalds®)', NULL, 28, 6, 50, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (744, 'Molho caseiro azeite/vinagre', NULL, 13, 0, 113, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (745, 'Molho de alho calve', NULL, 15, 2, 55, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (746, 'Molho de iogurte', NULL, 15, 2, 12, '01 colher de sopa ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (747, 'Molho de mostarda light', NULL, 15, 1, 10, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (748, 'Molho de pimenta', NULL, 35, 3, 12, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (749, 'Molho de queijo com ervas light', NULL, 13, 1, 8, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (750, 'Molho de tomate sabor pizza Pomarola®', NULL, 20, 2, 8, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (751, 'Molho de tomate tradicional  340g Fugini®', NULL, 20, 3, 13, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (752, 'Molho de tomate tradicional  340g Tarantella®', NULL, 20, 2, 8, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (753, 'Molho de tomate tradicional 340g Pomarola®', NULL, 20, 2, 8, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (754, 'Molho inglês', NULL, 6, 1, 6, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (755, 'Molho madeira Uncle Bens®', NULL, 30, 3, 13, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (756, 'Molho para carnes madeira 340g Predilecta® ', NULL, 60, 5, 23, '3 colheres de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (757, 'Molho para salada tipo caesar  Masterfoods®', NULL, 13, 2, 25, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (758, 'Molho pesto Genovese 190 g Barilla® ', NULL, 60, 2, 320, '3 colheres de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (759, 'Molho pra salada tipo caesar ( Masterfoods®)', NULL, 13, 2, 25, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (760, 'Molho pronto de pesto Hemmer®', NULL, 20, 2, 105, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (761, 'Molho shoyo', NULL, 6, 1, 5, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (762, 'Molho tártaro gourmet', NULL, 20, 1, 105, 'colher de sopa ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (763, 'Molho vermelho light', NULL, 14, 1, 10, 'colher de sobremesa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (764, 'Mondongo (dobradinha)', NULL, 100, 8, 111, 'concha pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (765, 'Moqueca de ovos', NULL, 85, 3, 116, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (766, 'Moqueca de peixe', NULL, 185, 4, 218, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (767, 'Morango', NULL, 12, 1, 5, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (768, 'Mortadela', NULL, 15, 0, 41, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (769, 'Mostarda folha cozida ', NULL, 45, 2, 12, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (770, 'Mousse de chocolate', NULL, 25, 8, 79, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (771, 'Mousse de maracujá', NULL, 35, 11, 100, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (772, 'Mucilon de arroz', NULL, 9, 8, 34, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (773, 'Mucilon de milho', NULL, 9, 8, 34, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (774, 'Muffin de banana Suavipan®', NULL, 40, 20, 141, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (775, 'Mungunzá', NULL, 150, 29, 150, 'concha');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (776, 'Nabo cozido (picado)', NULL, 18, 1, 5, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (777, 'Namorado cozido', NULL, 100, 0, 121, 'lé médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (778, 'Néctar de manga Dell Valle®', NULL, 200, 30, 118, 'caixinha');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (779, 'Nectarina', NULL, 100, 12, 49, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (780, 'Nescau®', NULL, 16, 14, 61, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (781, 'Nêspera', NULL, 40, 4, 20, 'unidade grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (782, 'Nesquick em pó Nestlé®', NULL, 16, 15, 63, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (783, 'Nesquick® preparado, caixinha de', NULL, 200, 29, 170, '1 und');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (784, 'Neston®', NULL, 8, 6, 29, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (785, 'Nhoque', NULL, 100, 21, 120, '01 escumadeira cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (786, 'Niguiri sushi de atum', NULL, 30, 10, 59, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (787, 'Niguiri sushi de kani', NULL, 30, 10, 51, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (788, 'Niguiri sushi de kani', NULL, 30, 10, 51, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (789, 'Niguiri sushi de salmão', NULL, 30, 10, 57, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (790, 'Ninho® soleil morango (bebida láctea)', NULL, 200, 31, 170, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (791, 'Nozes', NULL, 5, 1, 35, '01 unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (792, 'Nuggets de frango tradicional Sadia®', NULL, 25, 3, 53, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (793, 'Nuggets de legumes Sadia®', NULL, 21, 5, 36, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (794, 'Nuggets de peixe Sadia®', NULL, 23, 4, 41, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (795, 'Nutella®', NULL, 20, 11, 106, 'colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (796, 'Óleos vegetais', NULL, 8, 0, 72, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (797, 'Olho de Sogra', NULL, 10, 5, 34, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (798, 'Omelete', NULL, 65, 2, 110, 'um ovo ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (799, 'Onion rings (Burguer King®)', NULL, 0, 5, 319, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (800, 'Ovo de codorna', NULL, 10, 0, 16, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (801, 'Ovo de galinha inteiro (cozido)', NULL, 45, 0, 71, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (802, 'Ovo de galinha, clara cozida', NULL, 30, 0, 15, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (803, 'Ovo de galinha, gema cozida', NULL, 15, 0, 53, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (804, 'Ovo de Páscoa (média)', NULL, 25, 15, 134, '1 pedaço médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (805, 'Ovo de Páscoa Amandita®', NULL, 25, 15, 135, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (806, 'Ovo de Páscoa Bis®', NULL, 25, 15, 130, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (807, 'Ovo de páscoa Diamante Negro®', NULL, 25, 16, 128, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (808, 'Ovo de Páscoa Ferrero Rocher®', NULL, 25, 9, 151, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (809, 'Ovo de Páscoa Kinder Ovo Maxi® ', NULL, 25, 13, 143, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (810, 'Ovo de Páscoa Lacta®', NULL, 25, 15, 131, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (811, 'Ovo de Páscoa Lacta® Diet', NULL, 25, 15, 111, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (812, 'Ovo de Páscoa Shot®', NULL, 25, 14, 134, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (813, 'Ovo de Páscoa Sonho de Valsa®', NULL, 25, 15, 134, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (814, 'Ovomaltine®', NULL, 14, 12, 53, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (815, 'Paçoca', NULL, 30, 20, 115, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (816, 'Paçoca diet Airon®', NULL, 20, 7, 79, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (817, 'Pacu inteiro cru', NULL, 100, 0, 292, 'pedaço médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (818, 'Paio', NULL, 160, 5, 363, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (819, 'Palmito em conserva', NULL, 15, 1, 3, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (820, 'Pamonha', NULL, 160, 69, 413, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (821, 'Pamonha doce', NULL, 130, 47, 335, 'unidade pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (822, 'Panettone frutas', NULL, 25, 14, 88, 'fatia pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (823, 'Panqueca carne', NULL, 80, 18, 229, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (824, 'Panqueca de batata (Kartoelpuer®)', NULL, 30, 22, 97, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (825, 'Pão alemão integral  (Wickbold®)', NULL, 50, 17, 99, '1/2 unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (826, 'Pão careca doce', NULL, 50, 28, 134, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (827, 'Pão colonial italiano', NULL, 50, 27, 128, 'fatia grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (828, 'Pão com tucumã', NULL, 85, 28, 241, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (829, 'Pão de batata-inglesa', NULL, 50, 29, 137, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (830, 'Pão de cachorro-quente', NULL, 58, 31, 170, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (831, 'Pão de centeio integral', NULL, 50, 23, 116, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (832, 'Pão de forma branco ', NULL, 25, 12, 62, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (833, 'Pão de forma de aveia', NULL, 25, 11, 59, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (834, 'Pão de forma de centeio', NULL, 27, 13, 75, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (835, 'Pão de hambúrguer', NULL, 70, 40, 188, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (836, 'Pão de leite', NULL, 54, 30, 149, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (837, 'Pão de mel', NULL, 15, 13, 76, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (838, 'Pão de mel doce de leite Cacau Show®', NULL, 50, 30, 190, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (839, 'Pão de milho caseiro', NULL, 70, 43, 200, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (840, 'Pão de milho com 50% de farinha de trigo', NULL, 70, 42, 204, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (841, 'Pão de milho de forma Panco®', NULL, 32, 19, 96, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (842, 'Pão de milho industrializado', NULL, 70, 43, 201, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (843, 'Pão de passas', NULL, 50, 26, 136, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (844, 'Pão de queijo', NULL, 20, 9, 87, 'unidade pequena ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (845, 'Pão de queijo light Forno de Minas®', NULL, 27, 9, 55, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (846, 'Pão de torresmo', NULL, 70, 33, 375, 'fatia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (847, 'Pão doce', NULL, 50, 28, 134, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (848, 'Pão francês', NULL, 50, 28, 135, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (849, 'Pão italiano', NULL, 50, 28, 125, '01 fatia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (850, 'Pão ligth de quinua Sonda®', NULL, 50, 21, 114, '2 fatias');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (851, 'Pão sírio integral Pita Bread®', NULL, 53, 11, 53, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (852, 'Pão sírio Pita Bread®', NULL, 58, 33, 158, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (853, 'Papinha de banana e aveia Nestlé®', NULL, 120, 24, 104, 'pote');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (854, 'Papinha de carne com legumes Nestlé®', NULL, 155, 9, 83, 'pote');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (855, 'Pasta de atum', NULL, 35, 2, 112, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (856, 'Pasta de tofu', NULL, 26, 2, 45, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (857, 'Pastel de belém (Habib´s®)', NULL, 50, 14, 107, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (858, 'Pastel de carne', NULL, 32, 5, 84, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (859, 'Pastel de queijo', NULL, 25, 5, 75, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (860, 'Pastel de queijo de forno', NULL, 40, 15, 185, 'unidade grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (861, 'Pastel português', NULL, 35, 10, 149, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (862, 'Pastelzinho', NULL, 8, 3, 24, 'unidade pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (863, 'Patê de atum Coqueiro®', NULL, 10, 0, 17, 'colher de chá');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (864, 'Patê de atum ligth Gomes da Costa®', NULL, 10, 1, 12, 'colher de chá');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (865, 'Patê de Foie Gras, enlatado, defumado', NULL, 8, 0, 22, '01 colher de cha cheia ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (866, 'Patê de frango com ervas nas  Excelsior®', NULL, 10, 1, 20, 'colher de chá');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (867, 'Patê de peito de peru defumado Sadia®', NULL, 10, 0, 22, 'colher de chá');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (868, 'Patê de salmão Gomes da Costa®', NULL, 10, 0, 15, 'colher de chá');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (869, 'Pato no tucupi assado', NULL, 100, 3, 303, 'sobrecoxa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (870, 'Pavê de chocolate', NULL, 85, 16, 154, 'pedaço médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (871, 'Pavê de chocolate branco Amor aos Pedaços®', NULL, 60, 20, 208, 'fatia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (872, 'Pave de cholocale ao leite', NULL, 85, 34, 200, '2 colheres de copa cheias');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (873, 'Pé-de-moleque', NULL, 17, 10, 88, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (874, 'Pé-de-moleque (Norte)', NULL, 100, 45, 336, 'pedaço médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (875, 'Pé-de-moleque diet Airon®', NULL, 14, 6, 61, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (876, 'Pé-de-moleque Santa Helena*®', NULL, 17, 10, 88, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (877, 'Peito de peru defumado ligth Sadia®', NULL, 20, 0, 21, 'fatia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (878, 'Peito de peru defumado Perdigão®', NULL, 25, 0, 41, 'fatia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (879, 'Peixe de água doce cozido (média)', NULL, 120, 0, 117, 'lé médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (880, 'Peixe do mar cozido (média)', NULL, 120, 0, 117, 'lé médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (881, 'Peixe grelhado', NULL, 100, 0, 180, '1 posta ou lé');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (882, 'Peixe na telha', NULL, 230, 7, 262, 'posta média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (883, 'Pepino com casca (picles)', NULL, 50, 2, 12, '1/4 xícara de chá');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (884, 'Pepino cru', NULL, 3, 0, 1, 'fatia pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (885, 'Pequi refogado', NULL, 70, 9, 143, '4 unidades médias');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (886, 'Pêra crua ', NULL, 130, 19, 80, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (887, 'Pernil assado', NULL, 100, 1.69999999999999996, 94, '1 fatia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (888, 'Pernil suíno temperado seara®', NULL, 100, 2, 140, '2 ¹/² fatia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (889, 'Peru (carne branca assada)', NULL, 32, 0, 57, 'pedaço médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (890, 'Pescada inteira crua', NULL, 100, 0, 111, 'pedaço médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (891, 'Pêssego amarelo', NULL, 60, 6, 29, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (892, 'Pessego em calda', NULL, 30, 6, 25, '01 metade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (893, 'Picolé abacaxi Kibon®', NULL, 59, 16, 65, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (894, 'Picolé chicabon Kibon®', NULL, 65, 19, 106, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (895, 'Picolé de coco la frutta Nestlé®', NULL, 58, 13, 90, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (896, 'Picolé de morango la frutta Nestlé®', NULL, 60, 8, 36, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (897, 'Picolé de uva la frutta Nestlé®', NULL, 61, 15, 62, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (898, 'Picolé fruttare limão Kibon®', NULL, 58, 13, 51, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (899, 'Picolé mega clássico Nestlé®', NULL, 77, 22, 238, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (900, 'Picolé mega trufas Nestlé®', NULL, 77, 27, 258, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (901, 'Pimenta-malagueta', NULL, 15, 1, 6, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (902, 'Pimentão ', NULL, 13, 1, 7, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (903, 'Pinhão', NULL, 10, 5, 25, '01 unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (904, 'Pipoca doce', NULL, 20, 20, 95, 'saco medio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (905, 'Pipoca no óleo/manteiga com sal', NULL, 20, 14, 90, 'saco médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (906, 'Pipoca salgada', NULL, 20, 14, 90, 'saco médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (907, 'Piquiá', NULL, 50, 15, 179, 'unidade pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (908, 'Pirão de farinha de mandioca', NULL, 30, 9, 36, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (909, 'Pirarucu ', NULL, 100, 0, 120, 'lé médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (910, 'Pirarucu de casaca', NULL, 110, 25, 293, 'escumadeira cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (911, 'Pirulito', NULL, 20, 19, 74, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (912, 'Pistache', NULL, 1, 0, 4, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (913, 'Pitanga', NULL, 15, 1, 7, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (914, 'Pitaya', NULL, 200, 14, 84, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (915, 'Pizza Brasileira', NULL, 111, 25, 268, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (916, 'Pizza Califórnia', NULL, 100, 26, 213, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (917, 'Pizza de Alcachofra', NULL, 100, 21, 273, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (918, 'Pizza de atum', NULL, 100, 21, 258, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (919, 'Pizza de bacon', NULL, 100, 22, 288, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (920, 'Pizza de bacon com requeijão ', NULL, 123, 26, 244, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (921, 'Pizza de banana com canela', NULL, 100, 40, 207, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (922, 'Pizza de brigadeiro', NULL, 100, 59, 402, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (923, 'Pizza de calabresa', NULL, 100, 21, 240, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (924, 'Pizza de camarão', NULL, 100, 21, 268, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (925, 'Pizza de champignon', NULL, 100, 22, 222, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (926, 'Pizza de chocolate ao leite', NULL, 100, 54, 483, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (927, 'Pizza de chocolate branco', NULL, 100, 41, 443, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (928, 'Pizza de doce de leite', NULL, 100, 51, 423, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (929, 'Pizza de Escarola', NULL, 100, 22, 294, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (930, 'Pizza de frango', NULL, 100, 22, 229, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (931, 'Pizza de frango com catupiry ', NULL, 111, 25, 274, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (932, 'Pizza de marguerita ', NULL, 119, 27, 274, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (933, 'Pizza de mussarela', NULL, 100, 22, 278, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (934, 'Pizza de mussarela de búfala, rúcula e tomate seco', NULL, 100, 22, 165, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (935, 'Pizza de palmito com mussarela', NULL, 100, 25, 220, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (936, 'Pizza de pepperoni ', NULL, 102, 26, 274, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (937, 'Pizza de portuguesa', NULL, 125, 24, 246, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (938, 'Pizza de Prestígio', NULL, 100, 55, 420, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (939, 'Pizza de quatro queijos', NULL, 100, 27, 277, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (940, 'Pizza de Romeu e Julieta', NULL, 100, 27, 266, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (941, 'Pizza de Rúcula', NULL, 100, 22, 255, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (942, 'Pizza Napolitana', NULL, 100, 20, 207, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (943, 'Pizza Toscana', NULL, 100, 20, 227, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (944, 'Pizza Vegetariana ', NULL, 111, 22, 192, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (945, 'Polenguinho', NULL, 20, 0, 67, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (946, 'Polenta com molho de carne', NULL, 100, 8, 59, 'escumadeira');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (947, 'Polenta cozida', NULL, 30, 3, 20, 'pedaço médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (948, 'Polenta frita', NULL, 20, 2, 18, 'pedaço pequeno');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (949, 'Polenta mole', NULL, 100, 8, 59, '01 concha pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (950, 'Polpetone a parmeggiana', NULL, 150, 12, 280, '01 unidade  ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (951, 'Polvilho', NULL, 16, 14, 56, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (952, 'Polvo cru', NULL, 200, 4, 160, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (953, 'Porco frito com rodelas de limão', NULL, 100, 0, 311, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (954, 'Prato do chef picanha 160g (Bobs®)', NULL, 300, 29, 608, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (955, 'Prato primavera (Habibs®)', NULL, 435, 27, 460, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (956, 'Prato verão (Habibs®)', NULL, 330, 35, 420, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (957, 'Premium salad (McDonalds®)', NULL, 0, 3, 104, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (958, 'Premium salad crispy (McDonalds®)', NULL, 0, 17, 317, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (959, 'Premium salad grill (McDonalds®)', NULL, 0, 3, 227, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (960, 'Presunto cozido', NULL, 15, 0, 22, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (961, 'Presunto de peru', NULL, 15, 0, 15, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (962, 'Pringles® original', NULL, 2, 1, 10, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (963, 'Pringles® sabor cebola', NULL, 2, 1, 10, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (964, 'Pudim de leite', NULL, 50, 12, 91, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (965, 'Pudim de passas', NULL, 50, 14, 95, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (966, 'Pudim de tapioca', NULL, 110, 19, 263, 'fatia grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (967, 'Pudim de tapioca com leite condensado', NULL, 90, 40, 150, 'fatia grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (968, 'Pupunha', NULL, 25, 5, 41, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (969, 'Purê de batata', NULL, 45, 8, 56, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (970, 'Purê de inhame', NULL, 40, 7, 49, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (971, 'Purê de tomate', NULL, 45, 5, 22, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (972, 'Quarteirão (Mc Donalds®)', NULL, 0, 35, 533, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (973, 'Queijadinha de coco', NULL, 35, 9, 72, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (974, 'Queijo de coalho em espeto', NULL, 30, 0, 100, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (975, 'Queijo gorgonzola nacional', NULL, 38, 0, 151, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (976, 'Queijo mussarela', NULL, 15, 0, 42, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (977, 'Queijo parmesão', NULL, 15, 0, 54, '01 pedaço pequeno');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (978, 'Queijo parmesão nacional', NULL, 15, 0, 61, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (979, 'Queijo prato', NULL, 15, 0, 60, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (980, 'Queijo provolone nacional', NULL, 15, 0, 51, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (981, 'Queijo roquefort nacional', NULL, 30, 0, 120, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (982, 'Queijo suiço', NULL, 30, 0, 120, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (983, 'Queijo tipo requeijão', NULL, 30, 0, 89, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (984, 'Queijo tipo ricota nacional', NULL, 30, 0, 54, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (985, 'Queijo tofu', NULL, 10, 0, 7, 'pedaço pequeno');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (986, 'Queijo-de-minas frescal', NULL, 30, 0, 73, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (987, 'Queijo-de-minas frescal light', NULL, 30, 0, 46, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (988, 'Quentão ', NULL, 100, 34, 295, '1/2 copo');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (989, 'Quiabo cozido sem sal', NULL, 40, 3, 15, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (990, 'Quibe assado', NULL, 100, 15, 172, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (991, 'Quibe frito (Habibs®)', NULL, 85, 14, 152, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (992, 'Quiche de Presunto e Queijo', NULL, 50, 10, 259, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (993, 'Quiche de queijo ', NULL, 30, 11, 227, '01 unidade pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (994, 'Quindim', NULL, 35, 15, 111, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (995, 'Quinua em flocos Mãe Terra®', NULL, 40, 29, 159, '1/2 xícara de chá');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (996, 'Rã, carne de (desada)', NULL, 10, 0, 7, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (997, 'Rabada crua', NULL, 40, 0, 155, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (998, 'Rabanada', NULL, 60, 48, 249, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (999, 'Rapadura', NULL, 55, 48, 194, 'pedaço médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1000, 'Ratatouille', NULL, 107, 6, 77, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1001, 'Ravióli de carne', NULL, 50, 26, 141, 'escumadeira');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1002, 'Ravióli de carne', NULL, 50, 26, 141, 'escumadeira');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1003, 'Ravióli de mussarela', NULL, 100, 20, 210, '01 escumadeira cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1004, 'Ravióli de queijo', NULL, 50, 19, 149, 'escumadeira');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1005, 'Refrigerante', NULL, 240, 24, 96, 'copo duplo');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1006, 'Refrigerante citrus (Schweppes®)', NULL, 350, 42, 170, 'lata');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1007, 'Refrigerante dietético', NULL, 240, 0, 1, 'copo duplo cheio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1008, 'Repolho cozido (picado)', NULL, 10, 0, 1, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1009, 'Repolho cru (picado)', NULL, 10, 0, 2, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1010, 'Requeijão comum', NULL, 30, 0, 106, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1011, 'Requeijão cremoso', NULL, 30, 1, 106, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1012, 'Requeijão cremoso light', NULL, 30, 1, 54, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1013, 'Risole de Queijo', NULL, 50, 8, 112, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1014, 'Risoto', NULL, 25, 5, 41, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1015, 'Risoto a milanes ', NULL, 25, 19, 95, '01 colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1016, 'Risoto de camarão', NULL, 25, 6, 45, '01 colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1017, 'Risoto de frango', NULL, 25, 6, 45, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1018, 'Roll cake de chocolate', NULL, 38, 22, 143, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1019, 'Romã', NULL, 150, 23, 84, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1020, 'Rosquinhas de leite', NULL, 7, 4, 27, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1021, 'Sagu', NULL, 20, 17, 70, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1022, 'Sagu com creme de baunilha', NULL, 100, 37, 159, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1023, 'Sakemaki', NULL, 15, 3, 22, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1024, 'Salad caesar (Burguer King®)', NULL, 0, 2, 53, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1025, 'Salada (McDonalds®)', NULL, 0, 1, 8, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1026, 'Salada agadir', NULL, 100, 3, 167, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1027, 'Salada Agadir', NULL, 100, 3, 167.22999999999999, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1028, 'Salada almanara®', NULL, 100, 12, 292, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1029, 'Salada Almanara®', NULL, 100, 12, 293, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1030, 'Salada caesar (Bobs®)', NULL, 100, 5, 74, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1031, 'Salada de atum (sanduiche) -Bobs®', NULL, 154, 25, 377, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1032, 'Salada de batata com maionese', NULL, 38, 7, 58, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1033, 'Salada de batatas (Kartoelsalat®)', NULL, 30, 5, 82, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1034, 'Salada de frango (sanduiche) - Bobs®', NULL, 154, 25, 356, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1035, 'Salada de frutas', NULL, 20, 3, 10, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1036, 'Salada de frutas', NULL, 120, 18, 65, '01 potinho pequeno');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1037, 'Salada de frutas completa ', NULL, 100, 13, 52, 'taça pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1038, 'Salada de legumes com maionese', NULL, 35, 5, 34, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1039, 'Salada tropical (Bobs®)', NULL, 99, 7, 26, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1040, 'Salada tropical com tiras de hamburguer (Bobs®)', NULL, 277, 10, 289, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1041, 'Salada verde (Burguer King®)', NULL, 0, 1, 115, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1042, 'Salame', NULL, 20, 0, 59, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1043, 'Salgadinho de soja Jasmine®', NULL, 20, 4, 98, ' colher de sopa');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1044, 'Salgadinho Pingo douro sabor baicon Elma Chips®', NULL, 20, 11, 100, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1045, 'Salmão defumado', NULL, 100, 20, 233, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1046, 'Salmão grelhado', NULL, 100, 0, 171, '1 lé pequeno');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1047, 'Salpicão de frango', NULL, 25, 2, 61, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1048, 'Salsão/aipo cru', NULL, 15, 0, 3, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1049, 'Salsicha comum', NULL, 35, 0, 116, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1050, 'Salsicha de frango Sadia®', NULL, 35, 1, 74, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1051, 'Salsicha de peru light Sadia®', NULL, 35, 1, 58, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1052, 'Salsicha envasada (em conserva)', NULL, 35, 0, 64, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1053, 'Salsichão', NULL, 100, 3, 312, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1054, 'Salsichão (Wurst®)', NULL, 100, 3, 312, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1055, 'Sanduíche americano', NULL, 190, 28, 278, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1056, 'Sanduiche de almondegas  (Subway®)', NULL, 379, 61, 498, 'porção de 15cm');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1057, 'Sanduíche de atum', NULL, 120, 33, 372, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1058, 'Sanduiche de atum (Subway®)', NULL, 260, 56, 414, 'porção de 15cm');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1059, 'Sanduiche de BNT (Subway®)', NULL, 237, 47, 420, 'porção de 15cm');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1060, 'Sanduiche de carne (Subway®)', NULL, 249, 46, 386, 'porção de 15cm');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1061, 'Sanduiche de churrasco (Subway®)', NULL, 266, 48, 447, 'porção de 15cm');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1062, 'Sanduíche de frango', NULL, 120, 33, 299, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1063, 'Sanduiche de frango (Subway®)', NULL, 241, 46, 315, 'porção de 15cm');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1064, 'Sanduiche de frango defumado c/ cream cheese (Subway®)', NULL, 259, 47, 450, 'porção de 15cm');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1065, 'Sanduiche de frango teriyaki (Subway®)', NULL, 283, 61, 379, 'porção de 15cm');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1066, 'Sanduiche de frutos do mar (Subway®)', NULL, 260, 63, 413, 'porção de 15cm');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1067, 'Sanduiche de italiano (Subway®)', NULL, 222, 47, 404, 'porção de 15cm');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1068, 'Sanduiche de mussarela de búfala  (Subway®)', NULL, 261, 55, 450, 'porção de 15cm');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1069, 'Sanduiche de peito de peru (Bobs®)', NULL, 161, 25, 278, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1070, 'Sanduiche de peito de peru (Subway®)', NULL, 234, 47, 300, 'porção de 15cm');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1071, 'Sanduiche de pizza sub (Subway®)', NULL, 250, 50, 418, 'porção de 15cm');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1072, 'Sanduiche de pizzaiolo (Subway®)', NULL, 215, 48, 331, 'porção de 15cm');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1073, 'Sanduiche de presunto (Subway®)', NULL, 222, 46, 286, 'porção de 15cm');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1074, 'Sanduiche de presunto e peito de peru (Subway®)', NULL, 245, 47, 312, 'porção de 15cm');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1075, 'Sanduiche de rosbife (Subway®)', NULL, 237, 46, 315, 'porção de 15cm');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1076, 'Sanduiche de subway club TM (Subway®)', NULL, 275, 47, 350, 'porção de 15cm');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1077, 'Sanduiche de subway melt TM (Subway®)', NULL, 270, 47, 419, 'porção de 15cm');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1078, 'Sanduíche natural', NULL, 120, 29, 265, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1079, 'Sanduíche queijo quente', NULL, 85, 33, 300, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1080, 'Sanduiche vegetariano  (Subway®)', NULL, 177, 46, 239, 'porção de 15cm');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1081, 'Sapoti', NULL, 60, 16, 58, 'unidade pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1082, 'Saquê', NULL, 50, 3, 11, 'dose');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1083, 'Sardinha assada', NULL, 100, 0, 164, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1084, 'Sardinha enlatada em molho de tomate', NULL, 33, 0, 63, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1085, 'Sardinha enlatada em óleo', NULL, 33, 0, 65, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1086, 'Sardinha frita', NULL, 33, 1, 120, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1087, 'Sardinha inteira crua', NULL, 100, 0, 144, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1088, 'Sashimi de atum', NULL, 10, 0, 15, 'fatia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1089, 'Sashimi de salmão', NULL, 10, 0, 15, 'fatia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1090, 'Shitake frito', NULL, 10, 7, 50, '01 colher de sopa ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1091, 'Shoyo', NULL, 12, 1, 9, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1092, 'Siri', NULL, 16, 0, 15, 'unidade pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1093, 'Soja cozida', NULL, 17, 2, 29, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1094, 'Sopa caldo verde', NULL, 130, 7, 80, 'concha média ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1095, 'Sopa de cebola (creme)', NULL, 130, 6, 58, 'concha média ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1096, 'Sopa de cogumelo (creme)', NULL, 130, 12, 139, 'concha média ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1097, 'Sopa de creme de ervilha enlatada', NULL, 130, 27, 152, 'concha média ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1098, 'Sopa de ervilha', NULL, 130, 26, 165, 'concha média ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1099, 'Sopa de espinafre (creme)', NULL, 130, 5, 110, 'concha média ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1100, 'Sopa de feijão branco', NULL, 130, 18, 126, 'concha média ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1101, 'Sopa de frango', NULL, 130, 3, 46, 'concha média ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1102, 'Sopa de legumes com carne', NULL, 130, 8, 100, 'concha média ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1103, 'Sopa de lentilhas enlatada', NULL, 130, 16, 108, 'concha média ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1104, 'Sopa de macarrão', NULL, 130, 19, 132, 'concha média ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1105, 'Sopa de milho na mateiga 19g  Knorr Quick®', NULL, 19, 12, 77, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1106, 'Sorvete  massa de morango Nestlé®', NULL, 100, 27, 186, 'bola grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1107, 'Sorvete (casquinha) baunilha  (Bobs®)', NULL, 124, 39, 220, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1108, 'Sorvete copo light Kibon®', NULL, 63, 4, 39, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1109, 'Sorvete copo sundae morango Kibon®', NULL, 84, 25, 164, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1110, 'Sorvete cornetto brigadeiro Kibon®', NULL, 79, 23, 239, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1111, 'Sorvete cornetto caramelo Kibon®', NULL, 75, 28, 224, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1112, 'Sorvete cornetto crocante Kibon®', NULL, 72, 27, 247, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1113, 'Sorvete de casquinha baunilha (McDonalds®)', NULL, 0, 32, 192, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1114, 'Sorvete de casquinha chocolate (McDonalds®)', NULL, 0, 31, 192, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1115, 'Sorvete de casquinha mista (McDonalds®)', NULL, 0, 31, 192, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1116, 'Sorvete de flocos Nestlé®', NULL, 60, 14, 121, 'bola');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1117, 'Sorvete de frutas', NULL, 50, 15, 63, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1118, 'Sorvete de massa crocante crunch Nestlé®', NULL, 100, 36, 187, 'bola grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1119, 'Sorvete de massa de abacaxi Nestlé®', NULL, 80, 27, 161, 'bola média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1120, 'Sorvete de massa de chiclete Nestlé®', NULL, 100, 25, 195, 'bola grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1121, 'Sorvete de massa de creme Nestlé®', NULL, 80, 20, 144, 'bola média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1122, 'Sorvete eskibon  Kibon®', NULL, 48, 12, 159, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1123, 'Sorvete la fruta Nestlé®', NULL, 100, 27, 108, 'bola grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1124, 'Sorvete massa galak Nestlé®', NULL, 80, 20, 162, 'bola média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1125, 'Sorvete moça doce de leite Nestlé®', NULL, 100, 27, 181, 'bola grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1126, 'Sorvete moça esta brigadeiro Neslté®', NULL, 100, 31, 221, 'bola grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1127, 'Sorvete prestígio Nestlé®', NULL, 100, 24, 229, 'bola grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1128, 'Sorvete sedução manjar branco Nestlé®', NULL, 100, 27, 195, 'bola grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1129, 'Sorvete sedução mousse chocolate Nestlé®', NULL, 100, 36, 230, 'bola grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1130, 'Sorvete sedução romeu e julieta Nestlé®', NULL, 100, 29, 185, 'bola grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1131, 'Sorvete sem parar Nestlé®', NULL, 110, 39, 217, 'unidade copo');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1132, 'Sorvete troppo chocolate Nestlé®', NULL, 79, 25, 224, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1133, 'Sorvete troppo crocante Nestlé®', NULL, 76, 29, 229, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1134, 'Sorvete troppo flocos Nestlé®', NULL, 76, 28, 214, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1135, 'Sorvete troppo morango Nestlé®', NULL, 79, 30, 942, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1136, 'Spaghetti ao sugo', NULL, 110, 22, 170, '01 pegador');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1137, 'Steak tartare', NULL, 100, 2, 199, '01 porção pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1138, 'SteakhouseTM burguer (Burguer King®)', NULL, 0, 57, 936, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1139, 'SteakhouseTM junior (Burguer King®)', NULL, 0, 33, 560, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1140, 'Strogono de carne', NULL, 25, 0, 43, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1141, 'Strogono de frango', NULL, 25, 0, 50, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1142, 'Strudel de Amêndoas  (Hungaria®)', NULL, 100, 45, 285, 'porção pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1143, 'Strudel de Frango com Catupiry  (Hungaria®)', NULL, 100, 31, 546, 'porção pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1144, 'Strudel de frango com catupiry  Hungaria®', NULL, 100, 31, 546, 'porção pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1145, 'Strudel de maçã ', NULL, 220, 59, 396, 'fatia grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1146, 'Strudel de Maçã (Hungaria®)', NULL, 100, 30, 157, 'porção pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1147, 'Suco de abacaxi com açúcar', NULL, 240, 25, 103, 'copo duplo cheio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1148, 'Suco de acerola com açúcar', NULL, 240, 14, 62, 'copo duplo cheio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1149, 'Suco de laranja (envasado)', NULL, 240, 26, 116, 'copo duplo cheio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1150, 'Suco de laranja (fresco)', NULL, 240, 31, 140, 'copo duplo cheio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1151, 'Suco de laranja Del Valle®', NULL, 300, 29, 116, 'copo grande ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1152, 'Suco de laranja,cenoura sem açúcar', NULL, 240, 31, 137, 'copo duplo cheio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1153, 'Suco de mamão com açúcar', NULL, 240, 22, 91, 'copo duplo cheio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1154, 'Suco de manga com açúcar', NULL, 240, 23, 96, 'copo duplo cheio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1155, 'Suco de manga light Dell Valle®', NULL, 200, 9, 37, 'copo');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1156, 'Suco de maracujá com açúcar', NULL, 240, 17, 70, 'copo duplo cheio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1157, 'Suco de melão com açúcar', NULL, 240, 20, 82, 'copo duplo cheio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1158, 'Suco de morango', NULL, 240, 11, 52, 'copo duplo cheio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1159, 'Suco de morango com açúcar', NULL, 240, 22, 96, 'copo duplo cheio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1160, 'Suco de pessego', NULL, 240, 8, 38, 'copo duplo cheio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1161, 'Suco de tomate', NULL, 240, 11, 58, 'copo duplo cheio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1162, 'Suco de tomate enlatado', NULL, 240, 10, 49, 'copo duplo cheio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1163, 'Suco de uva engarrafado', NULL, 240, 36, 151, 'copo duplo cheio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1164, 'Suco de uva Kapo ®', NULL, 200, 23, 106, 'caixinha');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1165, 'Sucrilhos', NULL, 5, 5, 19, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1166, 'Suflê de espinafre', NULL, 55, 1, 89, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1167, 'Suflê de legumes', NULL, 55, 5, 70, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1168, 'Sundae caramelo (McDonalds®)', NULL, 0, 51, 323, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1169, 'Sundae chocolate (Bobs®)', NULL, 172, 80, 463, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1170, 'Sundae chocolate (McDonalds®)', NULL, 0, 40, 290, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1171, 'Sundae morango  (McDonalds®)', NULL, 0, 47, 292, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1172, 'Super McShake coco (Mc Donalds®)', NULL, 300, 50, 312, 'copo pequeno');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1173, 'Super McShake flocos (Mc Donalds®)', NULL, 300, 63, 334, 'copo pequeno');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1174, 'Super McShake morango (Mc Donalds®)', NULL, 300, 64, 336, 'copo pequeno');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1175, 'Sushi', NULL, 22, 14, 57, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1176, 'Sushi de Atum', NULL, 30, 7, 40, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1177, 'Sushi de Salmao', NULL, 30, 7, 38, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1178, 'Suspiro', NULL, 10, 9, 38, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1179, 'Sustagem', NULL, 16, 10, 62, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1180, 'Tablete Alpino® Diet', NULL, 30, 17, 143, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1181, 'Taça Habibs®', NULL, 330, 69, 554, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1182, 'Tainha', NULL, 100, 0, 204, 'posta pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1183, 'Tâmara seca', NULL, 10, 7, 28, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1184, 'Tambaqui (lé cru)', NULL, 100, 0, 151, 'pedaço médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1185, 'Tangerina', NULL, 135, 15, 67, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1186, 'Tapioca pronta', NULL, 100, 43, 174, 'unidade pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1187, 'Tapioquinha com queijo e coco ralado', NULL, 100, 62, 430, 'unidade pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1188, 'Tapioquinha com tucumã', NULL, 100, 48, 360, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1189, 'Tartar de salmão', NULL, 63, 1, 178, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1190, 'Tartar de salmão com molho sour cream', NULL, 100, 5, 265, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1191, 'Tarte de bacalhau', NULL, 90, 27, 282, 'unidade média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1192, 'Tatu (carne) recheado com lingüiça', NULL, 100, 3, 259, 'fatia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1193, 'Temaki (média)', NULL, 120, 24, 258, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1194, 'Temaki de Atum', NULL, 100, 18, 157, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1195, 'Temaki de Salmao', NULL, 100, 18, 190, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1196, 'Tempero Maggi® amaciante de carnes', NULL, 10, 2, 14, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1197, 'Tempero Maggi® fondor', NULL, 10, 2, 14, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1198, 'Tender ', NULL, 100, 0, 146, '01 fatia na');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1199, 'Terrine ligth (peito de peru)', NULL, 130, 14, 131, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1200, 'Tiramissu ', NULL, 100, 21, 364, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1201, 'Tiramissú de frutas', NULL, 200, 23, 138, 'taça');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1202, 'Tomate, purê de (enlatado)', NULL, 20, 1, 8, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1203, 'Top sundae caramelo (Mc Donalds®)', NULL, 0, 77, 508, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1204, 'Top sundae chocolate (Mc Donalds®)', NULL, 0, 65, 475, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1205, 'Top sundae morango (Mc Donalds®)', NULL, 0, 73, 478, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1206, 'Torradas', NULL, 8, 5, 25, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1207, 'Torradas alho,água e sal Casa Victoriana®', NULL, 30, 16, 150, '10 unidades');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1208, 'Torradas aperitivo', NULL, 2, 1, 7, '01 unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1209, 'Torresmo', NULL, 10, 0, 54, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1210, 'Torta Alemã', NULL, 50, 16, 192, 'fatia pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1211, 'Torta de banana (Mc Donalds®)', NULL, 0, 34, 228, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1212, 'Torta de carne ', NULL, 85, 16, 209, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1213, 'Torta de frango', NULL, 85, 16, 200, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1214, 'Torta de limão', NULL, 90, 34, 217, 'fatia pequena ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1215, 'Torta de maçã (Mc Donalds®)', NULL, 0, 26, 198, 'unidade ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1216, 'Torta holandesa', NULL, 60, 22, 181, 'porção pequena');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1217, 'Torta mousse de limão  Miss Daisy®', NULL, 60, 16, 121, '1/8 do pacote');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1218, 'Torta quente (Burguer King®)', NULL, 0, 24, 142, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1219, 'Tortellini com recheio de carne', NULL, 190, 33, 367, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1220, 'Tremoço em conserva', NULL, 20, 3, 24, 'colher de sopa ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1221, 'Trigo cozido', NULL, 25, 6, 28, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1222, 'Trigo em grão', NULL, 25, 20, 91, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1223, 'Trigo, bolo de', NULL, 60, 36, 203, 'fatia média');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1224, 'Triplo cheese (Bobs®)', NULL, 245, 35, 694, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1225, 'Trufa de chocolate', NULL, 30, 15, 140, 'unidade grande');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1226, 'Tucumã descascado', NULL, 30, 2, 142, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1227, 'Tucunaré (lé cru)', NULL, 100, 0, 102, 'pedaço médio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1228, 'Tutu de feijão', NULL, 20, 7, 43, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1229, 'Uramaki califórnia ', NULL, 20, 10, 47, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1230, 'Uramaki salmão', NULL, 22, 10, 62, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1231, 'Uramaki salmão', NULL, 22, 10, 62, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1232, 'Uva comum', NULL, 8, 1, 6, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1233, 'Uva do tipo européia', NULL, 100, 18, 79, 'cacho pequeno');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1234, 'Uva itália (uva verde)', NULL, 8, 1, 6, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1235, 'Uva passa', NULL, 18, 14, 55, '01 colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1236, 'Vaca atolada', NULL, 100, 3, 259, 'concha média cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1237, 'Vagem comum em conserva', NULL, 20, 1, 4, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1238, 'Vagem cozida', NULL, 20, 2, 8, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1239, 'Vatapá', NULL, 100, 9, 127, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1240, 'Vinagre', NULL, 10, 0, 2, 'colher de sopa cheia');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1241, 'Vinho (média)', NULL, 150, 6, 26, 'taça');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1242, 'Vinho branco', NULL, 150, 5, 21, 'taça');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1243, 'Vinho branco seco', NULL, 150, 0, 99, 'taça');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1244, 'Vinho de jenipapo', NULL, 150, 38, 152, 'taça');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1245, 'Vinho quente', NULL, 150, 24, 98, '01 copo ');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1246, 'Vitamina de fruta com suco', NULL, 240, 49, 223, 'copo duplo cheio');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1247, 'Waffer®- média sabores', NULL, 7.5, 3, 21, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1248, 'Whopper JR® com queijo (Burguer King®)', NULL, 0, 31, 464, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1249, 'Whopper® com queijo (Burguer King®)', NULL, 0, 52, 757, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1250, 'Whopper® duplo com queijo  (Burguer King®)', NULL, 0, 52, 978, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1251, 'Xinxim de galinha', NULL, 150, 1, 358, 'porção');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1252, 'Yakult®', NULL, 80, 11, 50, 'unidade');
INSERT INTO public.foods (id, name, measure_id, qtd, cho, kcal, measure) VALUES (1, 'Abacate (picado)', NULL, 45, 3, 79, 'colher de sopa cheia');


--
-- TOC entry 2867 (class 0 OID 692179)
-- Dependencies: 205
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.items (id, meal_id, quantidade_medida_usual, quantidade_g_ml, cho, kcal, food_id) VALUES (25, 10, 0, 12, 0.800000000000000044, 21.0700000000000003, 1);
INSERT INTO public.items (id, meal_id, quantidade_medida_usual, quantidade_g_ml, cho, kcal, food_id) VALUES (28, 15, 0, 0, 0, 32, 15);
INSERT INTO public.items (id, meal_id, quantidade_medida_usual, quantidade_g_ml, cho, kcal, food_id) VALUES (29, 14, 0, 43, 4.29999999999999982, 19.3500000000000014, 112);
INSERT INTO public.items (id, meal_id, quantidade_medida_usual, quantidade_g_ml, cho, kcal, food_id) VALUES (30, 14, 0, 0, 32, 22.8599999999999994, 49);
INSERT INTO public.items (id, meal_id, quantidade_medida_usual, quantidade_g_ml, cho, kcal, food_id) VALUES (32, 15, 0, 22, 0, 3.29999999999999982, 12);
INSERT INTO public.items (id, meal_id, quantidade_medida_usual, quantidade_g_ml, cho, kcal, food_id) VALUES (33, 15, 0, 321, 67.4099999999999966, 353.100000000000023, 17);
INSERT INTO public.items (id, meal_id, quantidade_medida_usual, quantidade_g_ml, cho, kcal, food_id) VALUES (38, 16, 0, 200, 26.6700000000000017, 117.329999999999998, 2);
INSERT INTO public.items (id, meal_id, quantidade_medida_usual, quantidade_g_ml, cho, kcal, food_id) VALUES (41, 16, 0, 111, 7.40000000000000036, 194.870000000000005, 1);
INSERT INTO public.items (id, meal_id, quantidade_medida_usual, quantidade_g_ml, cho, kcal, food_id) VALUES (42, 17, 0, 0, 12, 1.33000000000000007, 9);


--
-- TOC entry 2862 (class 0 OID 692132)
-- Dependencies: 200
-- Data for Name: meal_types; Type: TABLE DATA; Schema: public; Owner: vindixit
--

INSERT INTO public.meal_types (id, name, start_at, end_at) VALUES (1, 'Café da Manhã', '06:00:00', '09:00:00');
INSERT INTO public.meal_types (id, name, start_at, end_at) VALUES (2, 'Almoço', '11:30:00', '14:00:00');
INSERT INTO public.meal_types (id, name, start_at, end_at) VALUES (3, 'Jantar', '17:00:00', '19:00:00');
INSERT INTO public.meal_types (id, name, start_at, end_at) VALUES (4, 'Lanche da Manhã', '09:00:00', '11:30:00');
INSERT INTO public.meal_types (id, name, start_at, end_at) VALUES (5, 'Lanche da Tarde', '14:00:00', '17:00:00');
INSERT INTO public.meal_types (id, name, start_at, end_at) VALUES (6, 'Ceia', '19:00:00', '21:00:00');
INSERT INTO public.meal_types (id, name, start_at, end_at) VALUES (7, 'Lanche Noturno', '21:00:00', '06:00:00');


--
-- TOC entry 2859 (class 0 OID 642683)
-- Dependencies: 197
-- Data for Name: meals; Type: TABLE DATA; Schema: public; Owner: vindixit
--

INSERT INTO public.meals (id, meal_type_id, bolus, start_at, end_at, date) VALUES (16, 1, 0, '07:15:00', NULL, '2020-08-18');
INSERT INTO public.meals (id, meal_type_id, bolus, start_at, end_at, date) VALUES (17, 7, 0, '21:37:00', NULL, '2020-08-19');
INSERT INTO public.meals (id, meal_type_id, bolus, start_at, end_at, date) VALUES (14, 1, 0, '07:03:00', NULL, '2020-08-18');
INSERT INTO public.meals (id, meal_type_id, bolus, start_at, end_at, date) VALUES (10, 2, 0, '13:32:00', NULL, '2020-08-17');
INSERT INTO public.meals (id, meal_type_id, bolus, start_at, end_at, date) VALUES (15, 1, 0, '07:08:00', NULL, '2020-08-18');


--
-- TOC entry 2863 (class 0 OID 692138)
-- Dependencies: 201
-- Data for Name: measures; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.measures (id, name) VALUES (2, 'colher de sopa cheia');
INSERT INTO public.measures (id, name) VALUES (3, 'colher de chá');


--
-- TOC entry 2860 (class 0 OID 692112)
-- Dependencies: 198
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users ("Id", username, password, role) VALUES (1, 'aria', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', 'admin');


--
-- TOC entry 2876 (class 0 OID 0)
-- Dependencies: 204
-- Name: foods_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.foods_id_seq', 1254, true);


--
-- TOC entry 2877 (class 0 OID 0)
-- Dependencies: 206
-- Name: items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.items_id_seq', 42, true);


--
-- TOC entry 2878 (class 0 OID 0)
-- Dependencies: 199
-- Name: meal_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.meal_type_id_seq', 7, true);


--
-- TOC entry 2879 (class 0 OID 0)
-- Dependencies: 196
-- Name: meals_id_seq; Type: SEQUENCE SET; Schema: public; Owner: vindixit
--

SELECT pg_catalog.setval('public.meals_id_seq', 17, true);


--
-- TOC entry 2880 (class 0 OID 0)
-- Dependencies: 202
-- Name: measures_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.measures_id_seq', 3, true);


--
-- TOC entry 2732 (class 2606 OID 692183)
-- Name: items Items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT "Items_pkey" PRIMARY KEY (id);


--
-- TOC entry 2730 (class 2606 OID 692162)
-- Name: foods foods_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.foods
    ADD CONSTRAINT foods_pk PRIMARY KEY (id);


--
-- TOC entry 2726 (class 2606 OID 692185)
-- Name: meal_types meal_types_pkey; Type: CONSTRAINT; Schema: public; Owner: vindixit
--

ALTER TABLE ONLY public.meal_types
    ADD CONSTRAINT meal_types_pkey PRIMARY KEY (id);


--
-- TOC entry 2722 (class 2606 OID 642688)
-- Name: meals meals_pkey; Type: CONSTRAINT; Schema: public; Owner: vindixit
--

ALTER TABLE ONLY public.meals
    ADD CONSTRAINT meals_pkey PRIMARY KEY (id);


--
-- TOC entry 2728 (class 2606 OID 692142)
-- Name: measures measures_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.measures
    ADD CONSTRAINT measures_pkey PRIMARY KEY (id);


--
-- TOC entry 2724 (class 2606 OID 692119)
-- Name: users pk_Id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT "pk_Id" PRIMARY KEY ("Id");


--
-- TOC entry 2733 (class 2606 OID 692186)
-- Name: meals fk_meal_types; Type: FK CONSTRAINT; Schema: public; Owner: vindixit
--

ALTER TABLE ONLY public.meals
    ADD CONSTRAINT fk_meal_types FOREIGN KEY (meal_type_id) REFERENCES public.meal_types(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- TOC entry 2736 (class 2606 OID 692201)
-- Name: items foods_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT foods_fkey FOREIGN KEY (food_id) REFERENCES public.foods(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- TOC entry 2735 (class 2606 OID 692191)
-- Name: items meals_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT meals_fkey FOREIGN KEY (meal_id) REFERENCES public.meals(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- TOC entry 2734 (class 2606 OID 692163)
-- Name: foods measure_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.foods
    ADD CONSTRAINT measure_id FOREIGN KEY (measure_id) REFERENCES public.measures(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


-- Completed on 2020-08-21 13:34:20

--
-- PostgreSQL database dump complete
--

