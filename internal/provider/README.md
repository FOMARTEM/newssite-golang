# БД

```
BEGIN;

CREATE TABLE IF NOT EXISTS public.users
(
    id serial NOT NULL,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    admin boolean NOT NULL DEFAULT false,
    PRIMARY KEY (id),
    UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS public.posts
(
    post_id serial NOT NULL,
    title character varying(255) NOT NULL,
    body text NOT NULL,
    createdate date NOT NULL,
    updatedate date NOT NULL,
    user_id integer NOT NULL,
    PRIMARY KEY (post_id)
);


ALTER TABLE IF EXISTS public.posts
    ADD FOREIGN KEY (user_id)
    REFERENCES public.users (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

END;
```