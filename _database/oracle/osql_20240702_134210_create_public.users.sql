
-- +migrate Up
CREATE TABLE IF NOT EXISTS public.users(
    id VARCHAR2(50) NOT NULL PRIMARY KEY,
    name VARCHAR2 (255) NOT NULL,
    email VARCHAR2 (255) NOT NULL,
    creationDate CHANGE-THIS-TYPE  NOT NULL,
    created_at DATE DEFAULT sysdate NOT NULL,
    updated_at DATE DEFAULT sysdate NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS public.users;
