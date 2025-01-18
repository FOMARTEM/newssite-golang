-- Получение пользователя по id или email
CREATE OR REPLACE FUNCTION get_user(
    p_id INT DEFAULT NULL,
    p_email VARCHAR DEFAULT NULL
)
RETURNS TABLE (
    id INT,
    name VARCHAR,
    email VARCHAR,
    password text,
    admin integer
)
LANGUAGE plpgsql
AS $$
BEGIN
    RETURN QUERY
    SELECT u.id, u.name, u.email, u.password, u.admin
    FROM public.users u
    WHERE 
        (p_id IS NULL OR u.id = p_id) AND
        (p_email IS NULL OR u.email = p_email);
END;
$$;

-- Получение пароля
CREATE OR REPLACE FUNCTION get_password(
    p_email VARCHAR
)
RETURNS TABLE (
    password text
)
LANGUAGE plpgsql
AS $$
BEGIN
    RETURN QUERY
    SELECT u.password
    FROM public.users u
    WHERE 
        u.email = p_email;
END;
$$;