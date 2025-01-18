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
    UPDATE public.users SET admin=p_admin WHERE id = p_id;
END;
$$;

CREATE OR REPLACE PROCEDURE update_user (
    IN p_email VARCHAR,
    IN p_admin integer
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE public.users SET email=p_email WHERE id = p_id;
END;
$$;