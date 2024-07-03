
-- +migrate Up
CREATE TABLE IF NOT EXISTS public.songs_play_list(
    id VARCHAR2(50) NOT NULL PRIMARY KEY,
    playlist CHANGE-THIS-TYPE  NOT NULL,
    song CHANGE-THIS-TYPE  NOT NULL,
    created_at DATE DEFAULT sysdate NOT NULL,
    updated_at DATE DEFAULT sysdate NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS public.songs_play_list;
