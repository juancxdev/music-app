
-- +migrate Up
CREATE TABLE IF NOT EXISTS public.songs_played(
    id uuid NOT NULL PRIMARY KEY,
    user VARCHAR (32)  NOT NULL,
    song VARCHAR (32)  NOT NULL,
    date TIMESTAMP NOT NULL,
    is_deleted bool NOT NULL DEFAULT false,
    user_deleter uuid NULL,
    deleted_at TIMESTAMP NULL,
    user_creator uuid NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

-- +migrate Down
DROP TABLE IF EXISTS public.songs_played;
