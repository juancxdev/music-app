package songs_played

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

type handlerSongsPlayed struct {
	db   *sqlx.DB
	txID string
}

// CreateSongsPlayed godoc
// @Summary Crear una instancia de SongsPlayed
// @Description Método que permite crear una instancia del objeto SongsPlayed en la base de datos
// @tags SongsPlayed
// @Accept json
// @Produce json
// @Param RequestSongsPlayed body RequestSongsPlayed true "Datos para crear SongsPlayed"
// @Success 201 {object} ResponseSongsPlayed
// @Failure 400 {object} ResponseSongsPlayed
// @Failure 202 {object} ResponseSongsPlayed
// @Router /api/v1/songsplayed [POST]
func (h *handlerSongsPlayed) CreateSongsPlayed(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponseSongsPlayed{Error: true}
	req := RequestSongsPlayed{}
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
	data, code, err := srv.SrvSongsPlayed.CreateSongsPlayed(uuid.New().String(), req.User, req.Song, req.Date)
	if err != nil {
		logger.Error.Printf("couldn't create SongsPlayed, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusCreated).JSON(res)
}

// UpdateSongsPlayed godoc
// @Summary Actualiza una instancia de SongsPlayed
// @Description Método que permite Actualiza una instancia del objeto SongsPlayed en la base de datos
// @tags SongsPlayed
// @Accept json
// @Produce json
// @Param RequestSongsPlayed body RequestSongsPlayed true "Datos para actualizar SongsPlayed"
// @Success 200 {object} ResponseSongsPlayed
// @Failure 400 {object} ResponseSongsPlayed
// @Failure 202 {object} ResponseSongsPlayed
// @Router /api/v1/songsplayed [PUT]
func (h *handlerSongsPlayed) UpdateSongsPlayed(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponseSongsPlayed{Error: true}
	req := RequestSongsPlayed{}
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
	data, code, err := srv.SrvSongsPlayed.UpdateSongsPlayed(req.ID, req.User, req.Song, req.Date)
	if err != nil {
		logger.Error.Printf("couldn't update SongsPlayed, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}

// DeleteSongsPlayed godoc
// @Summary Elimina una instancia de SongsPlayed
// @Description Método que permite eliminar una instancia del objeto SongsPlayed en la base de datos
// @tags SongsPlayed
// @Accept json
// @Produce json
// @Param	id	path string true "SongsPlayed ID"
// @Success 200 {object} ResponseSongsPlayed
// @Failure 400 {object} ResponseSongsPlayed
// @Failure 202 {object} ResponseSongsPlayed
// @Router /api/v1/songsplayed [DELETE]
func (h *handlerSongsPlayed) DeleteSongsPlayed(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponseSongsPlayed{Error: true}
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
	code, err := srv.SrvSongsPlayed.DeleteSongsPlayed(id)
	if err != nil {
		logger.Error.Printf("couldn't delete SongsPlayed, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}

// GetSongsPlayedByID godoc
// @Summary Obtiene una instancia de SongsPlayed por su id
// @Description Método que permite obtener una instancia del objeto SongsPlayed en la base de datos por su id
// @tags SongsPlayed
// @Accept json
// @Produce json
// @Param	id	path string true "SongsPlayed ID"
// @Success 200 {object} ResponseSongsPlayed
// @Failure 400 {object} ResponseSongsPlayed
// @Failure 202 {object} ResponseSongsPlayed
// @Router /api/v1/songsplayed/{id} [GET]
func (h *handlerSongsPlayed) GetSongsPlayedByID(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponseSongsPlayed{Error: true}
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
	data, code, err := srv.SrvSongsPlayed.GetSongsPlayedByID(id)
	if err != nil {
		logger.Error.Printf("couldn't get SongsPlayed by id, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}

// GetAllSongsPlayed godoc
// @Summary Obtiene todas las instancias de SongsPlayed
// @Description Método que permite obtener todas las instancias del objeto SongsPlayed en la base de datos
// @tags SongsPlayed
// @Accept json
// @Produce json
// @Success 200 {object} ResponseSongsPlayed
// @Failure 202 {object} ResponseSongsPlayed
// @Router /api/v1/songsplayed/ [GET]
func (h *handlerSongsPlayed) GetAllSongsPlayed(c *fiber.Ctx) error {
	res := ResponseSongsPlayed{Error: true}
	msg := messages.NewMsgs(h.db)
	user := models.Users{
		ID: uuid.New().String(),
	}
	srv := public.NewServer(h.db, &user, h.txID)
	data, err := srv.SrvSongsPlayed.GetAllSongsPlayed()
	if err != nil {
		logger.Error.Printf("couldn't get all SongsPlayed, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(22)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}
