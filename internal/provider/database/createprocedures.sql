-- Создание пользователя
CREATE OR REPLACE PROCEDURE create_user (
	IN p_name VARCHAR,
	IN p_email VARCHAR,
	IN p_password VARCHAR,
	OUT n_id integer
)
LANGUAGE plpgsql
AS $$
BEGIN
    -- Проверка на дублирование почты
    IF EXISTS (SELECT 1 FROM users WHERE email = p_email) THEN
        RAISE EXCEPTION 'Пользователь с такой почтой уже существует.';
    END IF;
    -- Вставка нового пользователя
    INSERT INTO users (name, email, password)
    VALUES (p_name, p_email, p_password)

	RETURNING id into n_id;
END;
$$;

CREATE OR REPLACE PROCEDURE update_user (
    IN p_id integer,
    IN p_name VARCHAR,
    IN p_password VARCHAR
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE public.users SET name=p_name, password=p_password WHERE id = p_id;
END;
$$;

CREATE OR REPLACE PROCEDURE update_user (
    IN p_id integer,
    IN p_admin integer
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE public.users SET admin = p_admin WHERE id = p_id;
END;
$$;

CREATE OR REPLACE PROCEDURE update_user (
    IN p_email VARCHAR,
    IN p_admin integer
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE public.users SET admin = p_admin WHERE email = p_email;
END;
$$;


-- Создание поста
CREATE OR REPLACE PROCEDURE create_post (
	IN p_title VARCHAR,
	IN p_body VARCHAR,
	IN p_createdate DATE,
    IN p_user_id integer,
	OUT n_id integer
)
LANGUAGE plpgsql
AS $$
BEGIN
    -- Вставка нового поста
    INSERT INTO posts (title, body, createdate, updatedate, user_id)
    VALUES (p_title, p_body, p_createdate, p_createdate, p_user_id)

	RETURNING id into n_id;
END;
$$;

-- Обновление поста
CREATE OR REPLACE PROCEDURE update_post (
    IN p_id integer,
    IN n_title VARCHAR,
    IN n_body text,
    IN n_updatedate date
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE public.posts
    SET title = n_title, body = n_body, updatedate = n_updatedate
    WHERE id = p_id;
END;
$$;

-- Удаление всех постов пользователя
CREATE OR REPLACE PROCEDURE delete_user_posts (
    IN p_user_id integer
)
LANGUAGE plpgsql
AS $$
BEGIN
    DELETE FROM public.posts WHERE user_id = p_user_id;
END;
$$;

-- Создание комментария
CREATE OR REPLACE PROCEDURE create_comment(
    IN p_body VARCHAR,
	IN p_post_id integer,
	IN p_user_id integer,
	OUT n_id integer
)
LANGUAGE plpgsql
AS $$
BEGIN
    INSERT INTO comments (body, post_id, user_id)
    VALUES (p_body, p_post_id, p_user_id)

    RETURNING id INTO n_id;
END;
$$;

CREATE OR REPLACE PROCEDURE update_comment (
    IN p_id integer,
    IN n_body VARCHAR
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE public.comments
    SET body = n_body
    WHERE id = p_id;
END;
$$;

CREATE OR REPLACE PROCEDURE delete_comment(
    IN p_id integer
)
LANGUAGE plpgsql
AS $$
BEGIN
    DELETE FROM public.comments 
    WHERE id = p_id; 
END;
$$;

CREATE OR REPLACE PROCEDURE delete_comments(
    IN p_post_id integer
)
LANGUAGE plpgsql
AS $$
BEGIN
    DELETE FROM public.comments 
    WHERE post_id = p_post_id; 
END;
$$;