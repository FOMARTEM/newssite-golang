# БД

```
BEGIN;

CREATE TABLE IF NOT EXISTS public.users
(
    id serial NOT NULL,
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    email character varying(255) COLLATE pg_catalog."default" NOT NULL,
    password character varying(255) COLLATE pg_catalog."default" NOT NULL,
    admin boolean NOT NULL DEFAULT false,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_email_key UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS public.posts
(
    post_id serial NOT NULL,
    title character varying(255) COLLATE pg_catalog."default" NOT NULL,
    body text COLLATE pg_catalog."default" NOT NULL,
    createdate date NOT NULL,
    updatedate date NOT NULL,
    user_id integer NOT NULL,
    CONSTRAINT posts_pkey PRIMARY KEY (post_id)
);


ALTER TABLE IF EXISTS public.posts
    ADD CONSTRAINT posts_user_id_fkey FOREIGN KEY (user_id)
    REFERENCES public.users (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

END;
```