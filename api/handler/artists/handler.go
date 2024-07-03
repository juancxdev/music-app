package artists

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

type handlerArtists struct {
	db   *sqlx.DB
	txID string
}

// CreateArtists godoc
// @Summary Crear una instancia de Artists
// @Description Método que permite crear una instancia del objeto Artists en la base de datos
// @tags Artists
// @Accept json
// @Produce json
// @Param RequestArtists body RequestArtists true "Datos para crear Artists"
// @Success 201 {object} ResponseArtists
// @Failure 400 {object} ResponseArtists
// @Failure 202 {object} ResponseArtists
// @Router /api/v1/artists [POST]
func (h *handlerArtists) CreateArtists(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponseArtists{Error: true}
	req := RequestArtists{}
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
	data, code, err := srv.SrvArtists.CreateArtists(uuid.New().String(), req.Name, req.Surname, req.Nationality)
	if err != nil {
		logger.Error.Printf("couldn't create Artists, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusCreated).JSON(res)
}

// UpdateArtists godoc
// @Summary Actualiza una instancia de Artists
// @Description Método que permite Actualiza una instancia del objeto Artists en la base de datos
// @tags Artists
// @Accept json
// @Produce json
// @Param RequestArtists body RequestArtists true "Datos para actualizar Artists"
// @Success 200 {object} ResponseArtists
// @Failure 400 {object} ResponseArtists
// @Failure 202 {object} ResponseArtists
// @Router /api/v1/artists [PUT]
func (h *handlerArtists) UpdateArtists(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponseArtists{Error: true}
	req := RequestArtists{}
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
	data, code, err := srv.SrvArtists.UpdateArtists(req.ID, req.Name, req.Surname, req.Nationality)
	if err != nil {
		logger.Error.Printf("couldn't update Artists, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}

// DeleteArtists godoc
// @Summary Elimina una instancia de Artists
// @Description Método que permite eliminar una instancia del objeto Artists en la base de datos
// @tags Artists
// @Accept json
// @Produce json
// @Param	id	path string true "Artists ID"
// @Success 200 {object} ResponseArtists
// @Failure 400 {object} ResponseArtists
// @Failure 202 {object} ResponseArtists
// @Router /api/v1/artists/{id} [DELETE]
func (h *handlerArtists) DeleteArtists(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponseArtists{Error: true}
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
	code, err := srv.SrvArtists.DeleteArtists(id)
	if err != nil {
		logger.Error.Printf("couldn't delete Artists, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}

// GetArtistsByID godoc
// @Summary Obtiene una instancia de Artists por su id
// @Description Método que permite obtener una instancia del objeto Artists en la base de datos por su id
// @tags Artists
// @Accept json
// @Produce json
// @Param	id	path string true "Artists ID"
// @Success 200 {object} ResponseArtists
// @Failure 400 {object} ResponseArtists
// @Failure 202 {object} ResponseArtists
// @Router /api/v1/artists/{id} [GET]
func (h *handlerArtists) GetArtistsByID(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponseArtists{Error: true}
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
	data, code, err := srv.SrvArtists.GetArtistsByID(id)
	if err != nil {
		logger.Error.Printf("couldn't get Artists by id, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}

// GetAllArtists godoc
// @Summary Obtiene todas las instancias de Artists
// @Description Método que permite obtener todas las instancias del objeto Artists en la base de datos
// @tags Artists
// @Accept json
// @Produce json
// @Success 200 {object} ResponseArtists
// @Failure 202 {object} ResponseArtists
// @Router /api/v1/artists/ [GET]
func (h *handlerArtists) GetAllArtists(c *fiber.Ctx) error {
	res := ResponseArtists{Error: true}
	msg := messages.NewMsgs(h.db)
	user := models.Users{
		ID: uuid.New().String(),
	}
	srv := public.NewServer(h.db, &user, h.txID)
	data, err := srv.SrvArtists.GetAllArtists()
	if err != nil {
		logger.Error.Printf("couldn't get all Artists, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(22)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}
