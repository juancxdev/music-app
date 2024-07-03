package public

import (
	"github.com/jmoiron/sqlx"
	"music-app/internal/models"
	"music-app/pkg/public/albums"
	"music-app/pkg/public/artists"
	"music-app/pkg/public/play_lists"
	"music-app/pkg/public/songs"
	"music-app/pkg/public/songs_play_list"
	"music-app/pkg/public/songs_played"
	"music-app/pkg/public/users"
)

type Server struct {
	SrvAlbums          albums.PortsServerAlbums
	SrvArtists         artists.PortsServerArtists
	SrvPlaylists       play_lists.PortsServerPlaylists
	SrvSongs           songs.PortsServerSongs
	SrvSongsOfPlaylist songs_play_list.PortsServerSongsOfPlaylist
	SrvSongsPlayed     songs_played.PortsServerSongsPlayed
	SrvUsers           users.PortsServerUsers
}

func NewServer(db *sqlx.DB, usr *models.Users, txID string) *Server {
	return &Server{
		SrvAlbums:          albums.NewAlbumsService(albums.FactoryStorage(db, usr, txID), usr, txID),
		SrvArtists:         artists.NewArtistsService(artists.FactoryStorage(db, usr, txID), usr, txID),
		SrvPlaylists:       play_lists.NewPlaylistsService(play_lists.FactoryStorage(db, usr, txID), usr, txID),
		SrvSongs:           songs.NewSongsService(songs.FactoryStorage(db, usr, txID), usr, txID),
		SrvSongsOfPlaylist: songs_play_list.NewSongsOfPlaylistService(songs_play_list.FactoryStorage(db, usr, txID), usr, txID),
		SrvSongsPlayed:     songs_played.NewSongsPlayedService(songs_played.FactoryStorage(db, usr, txID), usr, txID),
		SrvUsers:           users.NewUsersService(users.FactoryStorage(db, usr, txID), usr, txID),
	}
}
