package play_lists

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

type handlerPlaylists struct {
	db   *sqlx.DB
	txID string
}

// CreatePlaylists godoc
// @Summary Crear una instancia de Playlists
// @Description Método que permite crear una instancia del objeto Playlists en la base de datos
// @tags Playlists
// @Accept json
// @Produce json
// @Param RequestPlaylists body RequestPlaylists true "Datos para crear Playlists"
// @Success 201 {object} ResponsePlaylists
// @Failure 400 {object} ResponsePlaylists
// @Failure 202 {object} ResponsePlaylists
// @Router /api/v1/playlists [POST]
func (h *handlerPlaylists) CreatePlaylists(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponsePlaylists{Error: true}
	req := RequestPlaylists{}
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
	data, code, err := srv.SrvPlaylists.CreatePlaylists(uuid.New().String(), req.Name, req.User)
	if err != nil {
		logger.Error.Printf("couldn't create Playlists, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusCreated).JSON(res)
}

// UpdatePlaylists godoc
// @Summary Actualiza una instancia de Playlists
// @Description Método que permite Actualiza una instancia del objeto Playlists en la base de datos
// @tags Playlists
// @Accept json
// @Produce json
// @Param RequestPlaylists body RequestPlaylists true "Datos para actualizar Playlists"
// @Success 200 {object} ResponsePlaylists
// @Failure 400 {object} ResponsePlaylists
// @Failure 202 {object} ResponsePlaylists
// @Router /api/v1/playlists [PUT]
func (h *handlerPlaylists) UpdatePlaylists(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponsePlaylists{Error: true}
	req := RequestPlaylists{}
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
	data, code, err := srv.SrvPlaylists.UpdatePlaylists(req.ID, req.Name, req.User)
	if err != nil {
		logger.Error.Printf("couldn't update Playlists, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}

// DeletePlaylists godoc
// @Summary Elimina una instancia de Playlists
// @Description Método que permite eliminar una instancia del objeto Playlists en la base de datos
// @tags Playlists
// @Accept json
// @Produce json
// @Param	id	path string true "Playlists ID"
// @Success 200 {object} ResponsePlaylists
// @Failure 400 {object} ResponsePlaylists
// @Failure 202 {object} ResponsePlaylists
// @Router /api/v1/playlists/{id} [DELETE]
func (h *handlerPlaylists) DeletePlaylists(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponsePlaylists{Error: true}
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
	code, err := srv.SrvPlaylists.DeletePlaylists(id)
	if err != nil {
		logger.Error.Printf("couldn't delete Playlists, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}

// GetPlaylistsByID godoc
// @Summary Obtiene una instancia de Playlists por su id
// @Description Método que permite obtener una instancia del objeto Playlists en la base de datos por su id
// @tags Playlists
// @Accept json
// @Produce json
// @Param	id	path string true "Playlists ID"
// @Success 200 {object} ResponsePlaylists
// @Failure 400 {object} ResponsePlaylists
// @Failure 202 {object} ResponsePlaylists
// @Router /api/v1/playlists/{id} [GET]
func (h *handlerPlaylists) GetPlaylistsByID(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponsePlaylists{Error: true}
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
	data, code, err := srv.SrvPlaylists.GetPlaylistsByID(id)
	if err != nil {
		logger.Error.Printf("couldn't get Playlists by id, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}

// GetAllPlaylists godoc
// @Summary Obtiene todas las instancias de Playlists
// @Description Método que permite obtener todas las instancias del objeto Playlists en la base de datos
// @tags Playlists
// @Accept json
// @Produce json
// @Success 200 {object} ResponsePlaylists
// @Failure 202 {object} ResponsePlaylists
// @Router /api/v1/playlists/ [GET]
func (h *handlerPlaylists) GetAllPlaylists(c *fiber.Ctx) error {
	res := ResponsePlaylists{Error: true}
	msg := messages.NewMsgs(h.db)
	user := models.Users{
		ID: uuid.New().String(),
	}
	srv := public.NewServer(h.db, &user, h.txID)
	data, err := srv.SrvPlaylists.GetAllPlaylists()
	if err != nil {
		logger.Error.Printf("couldn't get all Playlists, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(22)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}
