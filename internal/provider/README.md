# БД

```
BEGIN;


CREATE TABLE IF NOT EXISTS public.users
(
    id serial NOT NULL,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    password text NOT NULL,
    admin integer NOT NULL DEFAULT 0,
    PRIMARY KEY (id),
    CONSTRAINT uniq_email UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS public.posts
(
    id serial NOT NULL,
    title character varying(255) NOT NULL,
    body text NOT NULL,
    createdate date NOT NULL,
    updatedate date,
    user_id integer NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.comments
(
    id serial NOT NULL,
    body text NOT NULL,
    post_id integer NOT NULL,
    user_id integer NOT NULL,
    PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public.posts
    ADD CONSTRAINT user_id FOREIGN KEY (user_id)
    REFERENCES public.users (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.comments
    ADD CONSTRAINT post_id FOREIGN KEY (post_id)
    REFERENCES public.posts (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.comments
    ADD CONSTRAINT user_id FOREIGN KEY (user_id)
    REFERENCES public.users (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

END;
```