--
-- PostgreSQL database dump
--

-- Dumped from database version 15.4 (Homebrew)
-- Dumped by pg_dump version 15.4 (Homebrew)

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
-- Name: student_db; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.student_db (
    id bigint NOT NULL,
    create_time timestamp with time zone,
    update_time timestamp with time zone,
    name text,
    age bigint,
    gender text,
    address text,
    tags text[]
);


ALTER TABLE public.student_db OWNER TO root;

--
-- Name: COLUMN student_db.id; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.student_db.id IS '主键';


--
-- Name: COLUMN student_db.tags; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.student_db.tags IS '标签';


--
-- Name: student_db_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.student_db_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.student_db_id_seq OWNER TO root;

--
-- Name: student_db_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.student_db_id_seq OWNED BY public.student_db.id;


--
-- Name: tag_db; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.tag_db (
    id bigint NOT NULL,
    create_time timestamp with time zone,
    update_time timestamp with time zone,
    label text,
    content text,
    count bigint
);


ALTER TABLE public.tag_db OWNER TO root;

--
-- Name: COLUMN tag_db.id; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.tag_db.id IS '主键';


--
-- Name: tag_db_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.tag_db_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tag_db_id_seq OWNER TO root;

--
-- Name: tag_db_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.tag_db_id_seq OWNED BY public.tag_db.id;


--
-- Name: user_db; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.user_db (
    user_id bigint,
    tenant_id bigint,
    create_time timestamp with time zone,
    update_time timestamp with time zone,
    username text,
    password text,
    token text,
    menus text[]
);


ALTER TABLE public.user_db OWNER TO root;

--
-- Name: COLUMN user_db.user_id; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.user_db.user_id IS '主键';


--
-- Name: COLUMN user_db.tenant_id; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.user_db.tenant_id IS '租户ID';


--
-- Name: COLUMN user_db.menus; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.user_db.menus IS '菜单权限';


--
-- Name: user_db_user_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.user_db_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_db_user_id_seq OWNER TO root;

--
-- Name: user_db_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.user_db_user_id_seq OWNED BY public.user_db.user_id;


--
-- Name: student_db id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.student_db ALTER COLUMN id SET DEFAULT nextval('public.student_db_id_seq'::regclass);


--
-- Name: tag_db id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.tag_db ALTER COLUMN id SET DEFAULT nextval('public.tag_db_id_seq'::regclass);


--
-- Name: user_db user_id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.user_db ALTER COLUMN user_id SET DEFAULT nextval('public.user_db_user_id_seq'::regclass);


--
-- Data for Name: student_db; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.student_db (id, create_time, update_time, name, age, gender, address, tags) FROM stdin;
7	2023-12-22 09:57:33.542104+08	2023-12-22 10:14:30.663814+08	王卫国	18	male	江苏南京	{牛奶过敏,颠狂}
\.


--
-- Data for Name: tag_db; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.tag_db (id, create_time, update_time, label, content, count) FROM stdin;
1	2023-09-15 14:01:02.707285+08	2023-09-15 14:01:02.707286+08	花粉过敏	闻了狂打喷嚏	0
2	2023-12-20 17:01:20.324139+08	2023-12-20 17:01:20.32414+08	冷空气	冷空气过敏，吸多了易😷	0
3	2023-12-20 17:02:42.316242+08	2023-12-20 17:02:42.316242+08	颠狂	中举后直接疯癫	0
5	2023-12-20 18:08:53.953074+08	2023-12-20 18:08:53.953074+08	牛奶过敏	乳糖不耐受	0
\.


--
-- Data for Name: user_db; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.user_db (user_id, tenant_id, create_time, update_time, username, password, token, menus) FROM stdin;
1	0	2023-09-14 15:51:53.79581+08	2023-09-14 15:51:53.795811+08	lisi	f416771662e9e37841b8636c7966ae8a72df8d08daedffb9080a92130eef8d31	Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Imxpc2kiLCJpc3MiOiJzb21lYm9keSIsImV4cCI6MTcwMzIzMDkxMSwibmJmIjoxNzAzMTQ0NTExLCJpYXQiOjE3MDMxNDQ1MTF9.NTspfyTQJR4Ve9jlFmblUU1mOhCCebdKOkopEyznvEU	{class,student,tag,notice}
\.


--
-- Name: student_db_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.student_db_id_seq', 7, true);


--
-- Name: tag_db_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.tag_db_id_seq', 5, true);


--
-- Name: user_db_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.user_db_user_id_seq', 1, true);


--
-- Name: student_db student_db_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.student_db
    ADD CONSTRAINT student_db_pkey PRIMARY KEY (id);


--
-- Name: tag_db tag_db_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.tag_db
    ADD CONSTRAINT tag_db_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

