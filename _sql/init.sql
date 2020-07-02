-- Database: kanban

-- DROP DATABASE kanban;

CREATE DATABASE kanban
    WITH
    OWNER = admin
    ENCODING = 'UTF8'
    LC_COLLATE = 'C'
    LC_CTYPE = 'C'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;


-- Table: public.boards

-- DROP TABLE public.boards;

CREATE TABLE public.boards
(
    id          bigint                                              NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    name        character varying(500) COLLATE pg_catalog."default" NOT NULL,
    description character varying(1000) COLLATE pg_catalog."default",
    CONSTRAINT boards_pkey PRIMARY KEY (id)
)
    TABLESPACE pg_default;

ALTER TABLE public.boards
    OWNER to admin;


-- Table: public.columns

-- DROP TABLE public.columns;

CREATE TABLE public.columns
(
    id         bigint                                              NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    board_id   bigint,
    name       character varying(255) COLLATE pg_catalog."default" NOT NULL,
    "position" integer                                             NOT NULL,
    CONSTRAINT columns_pkey PRIMARY KEY (id),
    CONSTRAINT columns_board_id_fkey FOREIGN KEY (board_id)
        REFERENCES public.boards (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)
    TABLESPACE pg_default;

ALTER TABLE public.columns
    OWNER to admin;


-- Table: public.tasks

-- DROP TABLE public.tasks;

CREATE TABLE public.tasks
(
    id          bigint                                              NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    column_id   bigint,
    name        character varying(500) COLLATE pg_catalog."default" NOT NULL,
    description character varying(5000) COLLATE pg_catalog."default",
    priority    integer                                             NOT NULL,
    CONSTRAINT tasks_pkey PRIMARY KEY (id),
    CONSTRAINT tasks_column_id_fkey FOREIGN KEY (column_id)
        REFERENCES public.columns (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)
    TABLESPACE pg_default;

ALTER TABLE public.tasks
    OWNER to admin;


-- Table: public.comment

-- DROP TABLE public.comment;

CREATE TABLE public.comment
(
    id         bigint                                               NOT NULL,
    created_at time with time zone                                  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    parent_id  bigint                                               NOT NULL,
    body       character varying(5000) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT comment_pkey PRIMARY KEY (id)
)
    TABLESPACE pg_default;

ALTER TABLE public.comment
    OWNER to admin;