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

-- Получение admin по id или email
CREATE OR REPLACE FUNCTION get_admin(
    p_id INT DEFAULT NULL,
    p_email VARCHAR DEFAULT NULL
)
RETURNS TABLE (
    admin integer
)
LANGUAGE plpgsql
AS $$
BEGIN
    RETURN QUERY
    SELECT u.admin
    FROM public.users u
    WHERE 
        (p_id IS NULL OR u.id = p_id) AND
        (p_email IS NULL OR u.email = p_email);
END;
$$;


-- Получение поста по id
CREATE OR REPLACE FUNCTION get_post(
    p_id INT
)
RETURNS TABLE (
    id INT,
    title VARCHAR,
    body text,
    createdate text,
    updatedate text,
    hidden integer,
    user_id integer,
  user_name VARCHAR
)
LANGUAGE plpgsql
AS $$
BEGIN
    RETURN QUERY
    SELECT p.id, p.title, p.body, TO_CHAR(p.createdate, 'YYYY/MM/DD') AS createdate,  TO_CHAR(p.updatedate, 'YYYY/MM/DD') AS updatedate, p.hidden AS hidden, p.user_id, u.name
    FROM public.posts p
    JOIN public.users u ON u.id = p.user_id
    WHERE p.id = p_id;
END;
$$;

-- Получение всех постов
CREATE OR REPLACE FUNCTION get_posts(
    IN p_limit INT,
    IN p_offset INT
)
RETURNS TABLE (
    id INT,
    title VARCHAR,
    body text,
    createdate text,
    updatedate text,
    hidden integer,
    user_id integer,
    user_name VARCHAR
)
LANGUAGE plpgsql
AS $$
BEGIN
    RETURN QUERY
    SELECT p.id, p.title, p.body,
           TO_CHAR(p.createdate, 'YYYY/MM/DD'),
           TO_CHAR(p.updatedate, 'YYYY/MM/DD'),
           p.hidden,
           p.user_id,
           u.name
    FROM public.posts p
    JOIN public.users u ON u.id = p.user_id 
    WHERE u.admin >= 6 AND p.hidden = 0
    ORDER BY p.id DESC
    LIMIT p_limit OFFSET p_offset;
END;
$$;


-- Получение всех постов пользователя
CREATE OR REPLACE FUNCTION get_user_posts(
    IN p_user_id integer,
    IN p_limit INT,
    IN p_offset INT
)
RETURNS TABLE (
    id integer,
    title VARCHAR,
    body text,
    createdate text,
    updatedate text,
    hidden integer,
    user_id integer,
    user_name VARCHAR
)
LANGUAGE plpgsql
AS $$
BEGIN
    RETURN QUERY
    SELECT p.id, p.title, p.body,
           TO_CHAR(p.createdate, 'YYYY/MM/DD'),
           TO_CHAR(p.updatedate, 'YYYY/MM/DD'),
           p.hidden,
           p.user_id,
           u.name
    FROM public.posts p
    JOIN public.users u ON u.id = p.user_id
    WHERE p.user_id = p_user_id
    ORDER BY p.id DESC
    LIMIT p_limit OFFSET p_offset;
END;
$$;


CREATE OR REPLACE FUNCTION get_comments_by_post_id(
    IN p_post_id integer,
    IN p_limit INT,
    IN p_offset INT
)
RETURNS TABLE (
    id integer,
    bidy text,
    post_id integer,
    user_id integer,
    user_name VARCHAR
)
LANGUAGE plpgsql
AS $$
BEGIN
    SELECT c.id, c.body, c.post_id, c.user_id, u.name
  FROM public.comments c
    JOIN public.users u ON c.user_id = u.id
  WHERE post_id = p_post_id AND u.admin >= 0
    ORDER BY p.id DESC
    LIMIT p_limit OFFSET p_offset;
END;
$$;
