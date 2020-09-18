--
-- PostgreSQL database dump
--

-- Dumped from database version 11.1
-- Dumped by pg_dump version 11.1

-- Started on 2020-09-17 15:46:21

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 222 (class 1259 OID 702079)
-- Name: features_roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.features_roles (
    id integer DEFAULT nextval('public.features_roles_id_seq'::regclass) NOT NULL,
    feature_id integer,
    role_id integer
);


ALTER TABLE public.features_roles OWNER TO postgres;

--
-- TOC entry 2880 (class 0 OID 702079)
-- Dependencies: 222
-- Data for Name: features_roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.features_roles (id, feature_id, role_id) VALUES (159, 1, 1);
INSERT INTO public.features_roles (id, feature_id, role_id) VALUES (160, 2, 1);
INSERT INTO public.features_roles (id, feature_id, role_id) VALUES (161, 3, 1);
INSERT INTO public.features_roles (id, feature_id, role_id) VALUES (162, 4, 1);
INSERT INTO public.features_roles (id, feature_id, role_id) VALUES (163, 5, 1);
INSERT INTO public.features_roles (id, feature_id, role_id) VALUES (164, 6, 1);
INSERT INTO public.features_roles (id, feature_id, role_id) VALUES (165, 7, 1);
INSERT INTO public.features_roles (id, feature_id, role_id) VALUES (166, 8, 1);
INSERT INTO public.features_roles (id, feature_id, role_id) VALUES (167, 9, 1);
INSERT INTO public.features_roles (id, feature_id, role_id) VALUES (168, 10, 1);
INSERT INTO public.features_roles (id, feature_id, role_id) VALUES (169, 11, 1);
INSERT INTO public.features_roles (id, feature_id, role_id) VALUES (170, 12, 1);
INSERT INTO public.features_roles (id, feature_id, role_id) VALUES (171, 13, 1);
INSERT INTO public.features_roles (id, feature_id, role_id) VALUES (172, 14, 1);
INSERT INTO public.features_roles (id, feature_id, role_id) VALUES (173, 15, 1);
INSERT INTO public.features_roles (id, feature_id, role_id) VALUES (174, 16, 1);
INSERT INTO public.features_roles (id, feature_id, role_id) VALUES (175, 17, 1);


--
-- TOC entry 2754 (class 2606 OID 702115)
-- Name: features_roles feature_role_unique_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.features_roles
    ADD CONSTRAINT feature_role_unique_key UNIQUE (feature_id, role_id);


--
-- TOC entry 2756 (class 2606 OID 702093)
-- Name: features_roles features_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.features_roles
    ADD CONSTRAINT features_roles_pkey PRIMARY KEY (id);


--
-- TOC entry 2757 (class 2606 OID 702099)
-- Name: features_roles features_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.features_roles
    ADD CONSTRAINT features_fkey FOREIGN KEY (feature_id) REFERENCES public.features(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- TOC entry 2758 (class 2606 OID 702104)
-- Name: features_roles roles_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.features_roles
    ADD CONSTRAINT roles_fkey FOREIGN KEY (role_id) REFERENCES public.roles(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


-- Completed on 2020-09-17 15:46:21

--
-- PostgreSQL database dump complete
--

