package songs_play_list

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"music-app/internal/logger"
	"music-app/internal/messages"
	"music-app/internal/models"
	_ "music-app/internal/models"
	"music-app/pkg/public"
	"net/http"
)

type handlerSongsOfPlaylist struct {
	db   *sqlx.DB
	txID string
}

// CreateSongsOfPlaylist godoc
// @Summary Crear una instancia de SongsOfPlaylist
// @Description Método que permite crear una instancia del objeto SongsOfPlaylist en la base de datos
// @tags SongsOfPlaylist
// @Accept json
// @Produce json
// @Param RequestSongsOfPlaylist body RequestSongsOfPlaylist true "Datos para crear SongsOfPlaylist"
// @Success 201 {object} ResponseSongsOfPlaylist
// @Failure 400 {object} ResponseSongsOfPlaylist
// @Failure 202 {object} ResponseSongsOfPlaylist
// @Router /api/v1/songs-play-list [POST]
func (h *handlerSongsOfPlaylist) CreateSongsOfPlaylist(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponseSongsOfPlaylist{Error: true}
	req := RequestSongsOfPlaylist{}
	if err := c.BodyParser(&req); err != nil {
		logger.Error.Printf("couldn't parse body request, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	isValid, err := req.valid()
	if err != nil {
		logger.Error.Printf("couldn't validate body request, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	if !isValid {
		logger.Error.Println("couldn't validate body request")
		res.Code, res.Type, res.Msg = msg.GetByCode(1)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	user := models.Users{
		ID: uuid.New().String(),
	}
	srv := public.NewServer(h.db, &user, h.txID)
	data, code, err := srv.SrvSongsOfPlaylist.CreateSongsOfPlaylist(uuid.New().String(), req.Playlist, req.Song)
	if err != nil {
		logger.Error.Printf("couldn't create SongsOfPlaylist, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusCreated).JSON(res)
}

// UpdateSongsOfPlaylist godoc
// @Summary Actualiza una instancia de SongsOfPlaylist
// @Description Método que permite Actualiza una instancia del objeto SongsOfPlaylist en la base de datos
// @tags SongsOfPlaylist
// @Accept json
// @Produce json
// @Param RequestSongsOfPlaylist body RequestSongsOfPlaylist true "Datos para actualizar SongsOfPlaylist"
// @Success 200 {object} ResponseSongsOfPlaylist
// @Failure 400 {object} ResponseSongsOfPlaylist
// @Failure 202 {object} ResponseSongsOfPlaylist
// @Router /api/v1/songs-play-list [PUT]
func (h *handlerSongsOfPlaylist) UpdateSongsOfPlaylist(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponseSongsOfPlaylist{Error: true}
	req := RequestSongsOfPlaylist{}
	if err := c.BodyParser(&req); err != nil {
		logger.Error.Printf("couldn't parse body request, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	isValid, err := req.valid()
	if err != nil {
		logger.Error.Printf("couldn't validate body request, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	if !isValid {
		logger.Error.Println("couldn't validate body request")
		res.Code, res.Type, res.Msg = msg.GetByCode(1)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	user := models.Users{
		ID: uuid.New().String(),
	}
	srv := public.NewServer(h.db, &user, h.txID)
	data, code, err := srv.SrvSongsOfPlaylist.UpdateSongsOfPlaylist(req.ID, req.Playlist, req.Song)
	if err != nil {
		logger.Error.Printf("couldn't update SongsOfPlaylist, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}

// DeleteSongsOfPlaylist godoc
// @Summary Elimina una instancia de SongsOfPlaylist
// @Description Método que permite eliminar una instancia del objeto SongsOfPlaylist en la base de datos
// @tags SongsOfPlaylist
// @Accept json
// @Produce json
// @Param	id	path string true "SongsOfPlaylist ID"
// @Success 200 {object} ResponseSongsOfPlaylist
// @Failure 400 {object} ResponseSongsOfPlaylist
// @Failure 202 {object} ResponseSongsOfPlaylist
// @Router /api/v1/songs-play-list/{id} [DELETE]
func (h *handlerSongsOfPlaylist) DeleteSongsOfPlaylist(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponseSongsOfPlaylist{Error: true}
	id := c.Params("id")
	if id == "" {
		logger.Error.Println("id request no valid")
		res.Code, res.Type, res.Msg = msg.GetByCode(1)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	if !govalidator.IsUUIDv4(id) {
		logger.Error.Println("id request is not a valid uuid")
		res.Code, res.Type, res.Msg = msg.GetByCode(1)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	user := models.Users{
		ID: uuid.New().String(),
	}
	srv := public.NewServer(h.db, &user, h.txID)
	code, err := srv.SrvSongsOfPlaylist.DeleteSongsOfPlaylist(id)
	if err != nil {
		logger.Error.Printf("couldn't delete SongsOfPlaylist, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}

// GetSongsOfPlaylistByID godoc
// @Summary Obtiene una instancia de SongsOfPlaylist por su id
// @Description Método que permite obtener una instancia del objeto SongsOfPlaylist en la base de datos por su id
// @tags SongsOfPlaylist
// @Accept json
// @Produce json
// @Param	id	path string true "SongsOfPlaylist ID"
// @Success 200 {object} ResponseSongsOfPlaylist
// @Failure 400 {object} ResponseSongsOfPlaylist
// @Failure 202 {object} ResponseSongsOfPlaylist
// @Router /api/v1/songs-play-list/{id} [GET]
func (h *handlerSongsOfPlaylist) GetSongsOfPlaylistByID(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponseSongsOfPlaylist{Error: true}
	id := c.Params("id")
	if id == "" {
		logger.Error.Println("id request no valid")
		res.Code, res.Type, res.Msg = msg.GetByCode(1)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	if !govalidator.IsUUIDv4(id) {
		logger.Error.Println("id request is not a valid uuid")
		res.Code, res.Type, res.Msg = msg.GetByCode(1)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	user := models.Users{
		ID: uuid.New().String(),
	}
	srv := public.NewServer(h.db, &user, h.txID)
	data, code, err := srv.SrvSongsOfPlaylist.GetSongsOfPlaylistByID(id)
	if err != nil {
		logger.Error.Printf("couldn't get SongsOfPlaylist by id, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}

// GetAllSongsOfPlaylist godoc
// @Summary Obtiene todas las instancias de SongsOfPlaylist
// @Description Método que permite obtener todas las instancias del objeto SongsOfPlaylist en la base de datos
// @tags SongsOfPlaylist
// @Accept json
// @Produce json
// @Success 200 {object} ResponseSongsOfPlaylist
// @Failure 202 {object} ResponseSongsOfPlaylist
// @Router /api/v1/songs-play-list/ [GET]
func (h *handlerSongsOfPlaylist) GetAllSongsOfPlaylist(c *fiber.Ctx) error {
	res := ResponseSongsOfPlaylist{Error: true}
	msg := messages.NewMsgs(h.db)
	user := models.Users{
		ID: uuid.New().String(),
	}
	srv := public.NewServer(h.db, &user, h.txID)
	data, err := srv.SrvSongsOfPlaylist.GetAllSongsOfPlaylist()
	if err != nil {
		logger.Error.Printf("couldn't get all SongsOfPlaylist, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(22)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}
