package songs

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

type handlerSongs struct {
	db   *sqlx.DB
	txID string
}

// CreateSongs godoc
// @Summary Crear una instancia de Songs
// @Description Método que permite crear una instancia del objeto Songs en la base de datos
// @tags Songs
// @Accept json
// @Produce json
// @Param RequestSongs body RequestSongs true "Datos para crear Songs"
// @Success 201 {object} ResponseSongs
// @Failure 400 {object} ResponseSongs
// @Failure 202 {object} ResponseSongs
// @Router /api/v1/songs [POST]
func (h *handlerSongs) CreateSongs(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponseSongs{Error: true}
	req := RequestSongs{}
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
	data, code, err := srv.SrvSongs.CreateSongs(uuid.New().String(), req.Name, req.Artist, req.Album)
	if err != nil {
		logger.Error.Printf("couldn't create Songs, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusCreated).JSON(res)
}

// UpdateSongs godoc
// @Summary Actualiza una instancia de Songs
// @Description Método que permite Actualiza una instancia del objeto Songs en la base de datos
// @tags Songs
// @Accept json
// @Produce json
// @Param RequestSongs body RequestSongs true "Datos para actualizar Songs"
// @Success 200 {object} ResponseSongs
// @Failure 400 {object} ResponseSongs
// @Failure 202 {object} ResponseSongs
// @Router /api/v1/songs [PUT]
func (h *handlerSongs) UpdateSongs(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponseSongs{Error: true}
	req := RequestSongs{}
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
	data, code, err := srv.SrvSongs.UpdateSongs(req.ID, req.Name, req.Artist, req.Album)
	if err != nil {
		logger.Error.Printf("couldn't update Songs, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}

// DeleteSongs godoc
// @Summary Elimina una instancia de Songs
// @Description Método que permite eliminar una instancia del objeto Songs en la base de datos
// @tags Songs
// @Accept json
// @Produce json
// @Param	id	path string true "Songs ID"
// @Success 200 {object} ResponseSongs
// @Failure 400 {object} ResponseSongs
// @Failure 202 {object} ResponseSongs
// @Router /api/v1/songs/{id} [DELETE]
func (h *handlerSongs) DeleteSongs(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponseSongs{Error: true}
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
	code, err := srv.SrvSongs.DeleteSongs(id)
	if err != nil {
		logger.Error.Printf("couldn't delete Songs, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}

// GetSongsByID godoc
// @Summary Obtiene una instancia de Songs por su id
// @Description Método que permite obtener una instancia del objeto Songs en la base de datos por su id
// @tags Songs
// @Accept json
// @Produce json
// @Param	id	path string true "Songs ID"
// @Success 200 {object} ResponseSongs
// @Failure 400 {object} ResponseSongs
// @Failure 202 {object} ResponseSongs
// @Router /api/v1/songs/{id} [GET]
func (h *handlerSongs) GetSongsByID(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponseSongs{Error: true}
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
	data, code, err := srv.SrvSongs.GetSongsByID(id)
	if err != nil {
		logger.Error.Printf("couldn't get Songs by id, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}

// GetAllSongs godoc
// @Summary Obtiene todas las instancias de Songs
// @Description Método que permite obtener todas las instancias del objeto Songs en la base de datos
// @tags Songs
// @Accept json
// @Produce json
// @Success 200 {object} ResponseSongs
// @Failure 202 {object} ResponseSongs
// @Router /api/v1/songs/ [GET]
func (h *handlerSongs) GetAllSongs(c *fiber.Ctx) error {
	res := ResponseSongs{Error: true}
	msg := messages.NewMsgs(h.db)
	user := models.Users{
		ID: uuid.New().String(),
	}
	srv := public.NewServer(h.db, &user, h.txID)
	data, err := srv.SrvSongs.GetAllSongs()
	if err != nil {
		logger.Error.Printf("couldn't get all Songs, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(22)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}
