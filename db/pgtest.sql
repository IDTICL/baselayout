--
-- PostgreSQL database dump
--

-- Dumped from database version 13.2
-- Dumped by pg_dump version 13.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: users; Type: TABLE; Schema: public; Owner: rethel
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    username text,
    password text,
    role text,
    age bigint
);


ALTER TABLE public.users OWNER TO rethel;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: rethel
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO rethel;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rethel
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: rethel
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: rethel
--

COPY public.users (id, created_at, updated_at, deleted_at, username, password, role, age) FROM stdin;
1	2021-05-20 22:37:45.854913+08	2021-05-20 22:37:45.854913+08	\N	a123456	123456	users	33
\.


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rethel
--

SELECT pg_catalog.setval('public.users_id_seq', 1, true);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: rethel
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: rethel
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- PostgreSQL database dump complete
--

