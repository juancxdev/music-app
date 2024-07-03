
-- +migrate Up
CREATE TABLE public.songs_played(
    [id] UNIQUEIDENTIFIER NOT NULL PRIMARY KEY,
    [user] [CHANGE-THIS-TYPE]  NOT NULL,
    [song] [CHANGE-THIS-TYPE]  NOT NULL,
    [date] [CHANGE-THIS-TYPE]  NOT NULL,
    is_deleted bit NOT NULL DEFAULT false,
    user_deleter UNIQUEIDENTIFIER NULL,
    deleted_at [datetime] NULL,
    user_creator UNIQUEIDENTIFIER NOT NULL,
    created_at [datetime] NOT NULL DEFAULT (getdate()),
    updated_at [datetime] NOT NULL DEFAULT (getdate())
);

-- +migrate Down
DROP TABLE public.songs_played;
