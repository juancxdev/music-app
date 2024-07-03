
-- +migrate Up
CREATE TABLE public.songs(
    [id] UNIQUEIDENTIFIER NOT NULL PRIMARY KEY,
    [name] [VARCHAR] (255) NOT NULL,
    [artist] [VARCHAR] (255) NOT NULL,
    [album] [VARCHAR] (255) NOT NULL,
    is_deleted bit NOT NULL DEFAULT false,
    user_deleter UNIQUEIDENTIFIER NULL,
    deleted_at [datetime] NULL,
    user_creator UNIQUEIDENTIFIER NOT NULL,
    created_at [datetime] NOT NULL DEFAULT (getdate()),
    updated_at [datetime] NOT NULL DEFAULT (getdate())
);

-- +migrate Down
DROP TABLE public.songs;
