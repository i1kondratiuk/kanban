--
-- PostgreSQL database dump
--

-- Dumped from database version 12.3
-- Dumped by pg_dump version 12.2

-- Started on 2020-07-06 01:37:07 EEST

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
-- TOC entry 202 (class 1259 OID 16393)
-- Name: boards; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.boards (
    id bigint NOT NULL,
    name character varying(500) NOT NULL,
    description character varying(1000)
);


ALTER TABLE public.boards OWNER TO admin;

--
-- TOC entry 203 (class 1259 OID 16401)
-- Name: boards_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

ALTER TABLE public.boards ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.boards_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 204 (class 1259 OID 16413)
-- Name: columns; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.columns (
    id bigint NOT NULL,
    board_id bigint,
    name character varying(255) NOT NULL,
    "position" integer NOT NULL
);


ALTER TABLE public.columns OWNER TO admin;

--
-- TOC entry 208 (class 1259 OID 16470)
-- Name: columns_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

ALTER TABLE public.columns ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.columns_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 207 (class 1259 OID 16452)
-- Name: comments; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.comments (
    id bigint NOT NULL,
    parent_id bigint NOT NULL,
    body character varying(5000) NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.comments OWNER TO admin;

--
-- TOC entry 209 (class 1259 OID 16472)
-- Name: comments_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

ALTER TABLE public.comments ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.comments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 206 (class 1259 OID 16433)
-- Name: tasks; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.tasks (
    id bigint NOT NULL,
    column_id bigint,
    name character varying(500) NOT NULL,
    description character varying(5000),
    priority integer NOT NULL
);


ALTER TABLE public.tasks OWNER TO admin;

--
-- TOC entry 205 (class 1259 OID 16431)
-- Name: tasks_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

ALTER TABLE public.tasks ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.tasks_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 3082 (class 2606 OID 16405)
-- Name: boards boards_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.boards
    ADD CONSTRAINT boards_pkey PRIMARY KEY (id);


--
-- TOC entry 3084 (class 2606 OID 16417)
-- Name: columns columns_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.columns
    ADD CONSTRAINT columns_pkey PRIMARY KEY (id);


--
-- TOC entry 3088 (class 2606 OID 16456)
-- Name: comments comment_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comment_pkey PRIMARY KEY (id);


--
-- TOC entry 3086 (class 2606 OID 16440)
-- Name: tasks tasks_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_pkey PRIMARY KEY (id);


--
-- TOC entry 3089 (class 2606 OID 16426)
-- Name: columns columns_board_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.columns
    ADD CONSTRAINT columns_board_id_fkey FOREIGN KEY (board_id) REFERENCES public.boards(id);


--
-- TOC entry 3090 (class 2606 OID 16441)
-- Name: tasks tasks_column_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_column_id_fkey FOREIGN KEY (column_id) REFERENCES public.columns(id);


-- Completed on 2020-07-06 01:37:07 EEST

--
-- PostgreSQL database dump complete
--

