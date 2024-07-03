package users

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

type handlerUsers struct {
	db   *sqlx.DB
	txID string
}

// CreateUsers godoc
// @Summary Crear una instancia de Users
// @Description Método que permite crear una instancia del objeto Users en la base de datos
// @tags Users
// @Accept json
// @Produce json
// @Param RequestUsers body RequestUsers true "Datos para crear Users"
// @Success 201 {object} ResponseUsers
// @Failure 400 {object} ResponseUsers
// @Failure 202 {object} ResponseUsers
// @Router /api/v1/users [POST]
func (h *handlerUsers) CreateUsers(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponseUsers{Error: true}
	req := RequestUsers{}
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
	data, code, err := srv.SrvUsers.CreateUsers(uuid.New().String(), req.Name, req.Email, req.CreationDate)
	if err != nil {
		logger.Error.Printf("couldn't create Users, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusCreated).JSON(res)
}

// UpdateUsers godoc
// @Summary Actualiza una instancia de Users
// @Description Método que permite Actualiza una instancia del objeto Users en la base de datos
// @tags Users
// @Accept json
// @Produce json
// @Param RequestUsers body RequestUsers true "Datos para actualizar Users"
// @Success 200 {object} ResponseUsers
// @Failure 400 {object} ResponseUsers
// @Failure 202 {object} ResponseUsers
// @Router /api/v1/users [PUT]
func (h *handlerUsers) UpdateUsers(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponseUsers{Error: true}
	req := RequestUsers{}
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
	data, code, err := srv.SrvUsers.UpdateUsers(req.ID, req.Name, req.Email, req.CreationDate)
	if err != nil {
		logger.Error.Printf("couldn't update Users, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}

// DeleteUsers godoc
// @Summary Elimina una instancia de Users
// @Description Método que permite eliminar una instancia del objeto Users en la base de datos
// @tags Users
// @Accept json
// @Produce json
// @Param	id	path string true "Users ID"
// @Success 200 {object} ResponseUsers
// @Failure 400 {object} ResponseUsers
// @Failure 202 {object} ResponseUsers
// @Router /api/v1/users/{id} [DELETE]
func (h *handlerUsers) DeleteUsers(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponseUsers{Error: true}
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
	code, err := srv.SrvUsers.DeleteUsers(id)
	if err != nil {
		logger.Error.Printf("couldn't delete Users, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}

// GetUsersByID godoc
// @Summary Obtiene una instancia de Users por su id
// @Description Método que permite obtener una instancia del objeto Users en la base de datos por su id
// @tags Users
// @Accept json
// @Produce json
// @Param	id	path string true "Users ID"
// @Success 200 {object} ResponseUsers
// @Failure 400 {object} ResponseUsers
// @Failure 202 {object} ResponseUsers
// @Router /api/v1/users/{id} [GET]
func (h *handlerUsers) GetUsersByID(c *fiber.Ctx) error {
	msg := messages.NewMsgs(h.db)
	res := ResponseUsers{Error: true}
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
	data, code, err := srv.SrvUsers.GetUsersByID(id)
	if err != nil {
		logger.Error.Printf("couldn't get Users by id, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}

// GetAllUsers godoc
// @Summary Obtiene todas las instancias de Users
// @Description Método que permite obtener todas las instancias del objeto Users en la base de datos
// @tags Users
// @Accept json
// @Produce json
// @Success 200 {object} ResponseUsers
// @Failure 202 {object} ResponseUsers
// @Router /api/v1/users/ [GET]
func (h *handlerUsers) GetAllUsers(c *fiber.Ctx) error {
	res := ResponseUsers{Error: true}
	msg := messages.NewMsgs(h.db)
	user := models.Users{
		ID: uuid.New().String(),
	}
	srv := public.NewServer(h.db, &user, h.txID)
	data, err := srv.SrvUsers.GetAllUsers()
	if err != nil {
		logger.Error.Printf("couldn't get all Users, error: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(22)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = msg.GetByCode(29)
	return c.Status(http.StatusOK).JSON(res)
}
