package albums

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

type handlerAlbums struct {
	db   *sqlx.DB
	txID string
}

// CreateAlbums godoc
// @Summary Crear una instancia de Albums
// @Description Método que permite crear una instancia del objeto Albums en la base de datos
// @tags Albums
// @Accept json
// @Produce json
// @Param RequestAlbums body RequestAlbums true "Datos para crear Albums"
// @Success 201 {object} ResponseAlbums
// @Failure 400 {object} ResponseAlbums
// @Failure 202 {object} ResponseAlbums
// @Router /api/v1/albums [POST]
func (h *handlerAlbums) CreateAlbums(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponseAlbums{Error: true}
	req := RequestAlbums{}
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
	data, code, err := srv.SrvAlbums.CreateAlbums(uuid.New().String(), req.Name, req.Artist, req.ReleaseDate)
	if err != nil {
		logger.Error.Printf("couldn't create Albums, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusCreated).JSON(res)
}

// UpdateAlbums godoc
// @Summary Actualiza una instancia de Albums
// @Description Método que permite Actualiza una instancia del objeto Albums en la base de datos
// @tags Albums
// @Accept json
// @Produce json
// @Param RequestAlbums body RequestAlbums true "Datos para actualizar Albums"
// @Success 200 {object} ResponseAlbums
// @Failure 400 {object} ResponseAlbums
// @Failure 202 {object} ResponseAlbums
// @Router /api/v1/albums [PUT]
func (h *handlerAlbums) UpdateAlbums(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponseAlbums{Error: true}
	req := RequestAlbums{}
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
	data, code, err := srv.SrvAlbums.UpdateAlbums(req.ID, req.Name, req.Artist, req.ReleaseDate)
	if err != nil {
		logger.Error.Printf("couldn't update Albums, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}

// DeleteAlbums godoc
// @Summary Elimina una instancia de Albums
// @Description Método que permite eliminar una instancia del objeto Albums en la base de datos
// @tags Albums
// @Accept json
// @Produce json
// @Param	id	path string true "Albums ID"
// @Success 200 {object} ResponseAlbums
// @Failure 400 {object} ResponseAlbums
// @Failure 202 {object} ResponseAlbums
// @Router /api/v1/albums/{id} [DELETE]
func (h *handlerAlbums) DeleteAlbums(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponseAlbums{Error: true}
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
	code, err := srv.SrvAlbums.DeleteAlbums(id)
	if err != nil {
		logger.Error.Printf("couldn't delete Albums, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}

// GetAlbumsByID godoc
// @Summary Obtiene una instancia de Albums por su id
// @Description Método que permite obtener una instancia del objeto Albums en la base de datos por su id
// @tags Albums
// @Accept json
// @Produce json
// @Param	id	path string true "Albums ID"
// @Success 200 {object} ResponseAlbums
// @Failure 400 {object} ResponseAlbums
// @Failure 202 {object} ResponseAlbums
// @Router /api/v1/albums/{id} [GET]
func (h *handlerAlbums) GetAlbumsByID(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponseAlbums{Error: true}
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
	data, code, err := srv.SrvAlbums.GetAlbumsByID(id)
	if err != nil {
		logger.Error.Printf("couldn't get Albums by id, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}

// GetAllAlbums godoc
// @Summary Obtiene todas las instancias de Albums
// @Description Método que permite obtener todas las instancias del objeto Albums en la base de datos
// @tags Albums
// @Accept json
// @Produce json
// @Success 200 {object} ResponseAlbums
// @Failure 202 {object} ResponseAlbums
// @Router /api/v1/albums/ [GET]
func (h *handlerAlbums) GetAllAlbums(c *fiber.Ctx) error {
	res := ResponseAlbums{Error: true}
	msg := messages.NewMsgs(h.db)

	user := models.Users{
		ID: uuid.New().String(),
	}

	srv := public.NewServer(h.db, &user, h.txID)
	data, err := srv.SrvAlbums.GetAllAlbums()
	if err != nil {
		logger.Error.Printf("couldn't get all Albums, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(22)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}
