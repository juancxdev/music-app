
-- +migrate Up
CREATE TABLE IF NOT EXISTS public.play_lists(
    id uuid NOT NULL PRIMARY KEY,
    name VARCHAR (255) NOT NULL,
    user VARCHAR (32)  NOT NULL,
    is_deleted bool NOT NULL DEFAULT false,
    user_deleter uuid NULL,
    deleted_at TIMESTAMP NULL,
    user_creator uuid NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

-- +migrate Down
DROP TABLE IF EXISTS public.play_lists;
