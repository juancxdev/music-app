
-- +migrate Up
CREATE TABLE IF NOT EXISTS public.songs(
    id VARCHAR2(50) NOT NULL PRIMARY KEY,
    name VARCHAR2 (255) NOT NULL,
    artist VARCHAR2 (255) NOT NULL,
    album VARCHAR2 (255) NOT NULL,
    created_at DATE DEFAULT sysdate NOT NULL,
    updated_at DATE DEFAULT sysdate NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS public.songs;
