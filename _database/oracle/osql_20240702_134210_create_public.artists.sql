
-- +migrate Up
CREATE TABLE IF NOT EXISTS public.artists(
    id VARCHAR2(50) NOT NULL PRIMARY KEY,
    name VARCHAR2 (255) NOT NULL,
    surname VARCHAR2 (255) NOT NULL,
    nationality VARCHAR2 (255) NOT NULL,
    created_at DATE DEFAULT sysdate NOT NULL,
    updated_at DATE DEFAULT sysdate NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS public.artists;
