// internal/delivery/http/handlers/user_handler.go
package handlers

import (
	"net/http"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/Turgho/Aluguei/internal/domain/usecases"
	"github.com/Turgho/Aluguei/pkg/response"
	"github.com/gin-gonic/gin"
)

// UserHandler agrupa os handlers HTTP relacionados a [models.User].
type UserHandler struct {
	uc usecases.UserUseCase
}

// NewUserHandler retorna uma instância de [UserHandler].
func NewUserHandler(uc usecases.UserUseCase) *UserHandler {
	return &UserHandler{uc: uc}
}

// userResponse representa o corpo da resposta do user.
type userResponse struct {
	ID        string        `json:"id"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	CPF       string        `json:"cpf"`
	Email     string        `json:"email"`
	Phone     string        `json:"phone"`
	Role      entities.Role `json:"role"`
}

// registerRequest representa o corpo da requisição de registro.
type registerRequest struct {
	FirstName string        `json:"first_name" binding:"required"`
	LastName  string        `json:"last_name" binding:"required"`
	CPF       string        `json:"cpf" binding:"required"`
	Email     string        `json:"email" binding:"required,email"`
	Phone     string        `json:"phone"`
	Password  string        `json:"password" binding:"required,min=8"`
	Role      entities.Role `json:"role" binding:"required"`
}

// loginResponse representa a resposta do login.
type loginResponse struct {
	Token string `json:"token"`
}

// loginRequest representa o corpo da requisição de login.
type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// updateRequest representa o corpo da requisição de atualização.
type updateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

func toUserResponse(user *entities.User) userResponse {
	return userResponse{
		ID:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CPF:       user.CPF,
		Email:     user.Email,
		Phone:     user.Phone,
		Role:      user.Role,
	}
}

func toUserResponses(users []*entities.User) []userResponse {
	res := make([]userResponse, 0, len(users))
	for _, u := range users {
		res = append(res, toUserResponse(u))
	}
	return res
}

// Register godoc
//
//	@Summary		Registra um novo usuário
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			body	body		registerRequest	true	"Dados do usuário"
//	@Success		201		{object}	userResponse
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		409		{object}	response.ErrorResponse
//	@Router			/api/v1/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	user, err := h.uc.Create(req.FirstName, req.LastName, req.CPF, req.Email, req.Phone, req.Password, req.Role)
	if err != nil {
		response.Error(c, http.StatusConflict, "CONFLICT", err.Error())
		return
	}

	c.JSON(http.StatusCreated, toUserResponse(user))
}

// Login godoc
//
//	@Summary		Autentica um usuário e retorna um token JWT
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			body	body		loginRequest	true	"Credenciais"
//	@Success		200		{object}	loginResponse
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		401		{object}	response.ErrorResponse
//	@Router			/api/v1/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	token, err := h.uc.Login(req.Email, req.Password)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "INVALID_CREDENTIALS", "credenciais inválidas")
		return
	}

	c.JSON(http.StatusOK, loginResponse{Token: token})
}

// GetByID godoc
//
//	@Summary		Busca um usuário pelo ID
//	@Tags			users
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		string	true	"ID do usuário"
//	@Success		200	{object}	userResponse
//	@Failure		404	{object}	response.ErrorResponse
//	@Router			/api/v1/users/{id} [get]
func (h *UserHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	user, err := h.uc.GetByID(id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "USER_NOT_FOUND", "usuário não encontrado")
		return
	}

	c.JSON(http.StatusOK, toUserResponse(user))
}

// Update godoc
//
//	@Summary		Atualiza dados de um usuário
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id		path		string			true	"ID do usuário"
//	@Param			body	body		updateRequest	true	"Dados a atualizar"
//	@Success		200		{object}	userResponse
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		404		{object}	response.ErrorResponse
//	@Router			/api/v1/users/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	id := c.Param("id")

	user, err := h.uc.GetByID(id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "USER_NOT_FOUND", "usuário não encontrado")
		return
	}

	var req updateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}

	if err := h.uc.Update(user); err != nil {
		response.Error(c, http.StatusInternalServerError, "INTERNAL_ERROR", err.Error())
		return
	}

	c.JSON(http.StatusOK, toUserResponse(user))
}

// Delete godoc
//
//	@Summary		Remove um usuário pelo ID
//	@Tags			users
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		string	true	"ID do usuário"
//	@Success		204
//	@Failure		404		{object}	response.ErrorResponse
//	@Router			/api/v1/users/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.uc.Delete(id); err != nil {
		response.Error(c, http.StatusNotFound, "USER_NOT_FOUND", "usuário não encontrado")
		return
	}

	c.Status(http.StatusNoContent)
}

// Search godoc
//
//	@Summary		Busca usuários por nome, email ou CPF
//	@Tags			users
//	@Produce		json
//	@Security		BearerAuth
//	@Param			q	query		string	true	"Termo de busca"
//	@Success		200	{array}	userResponse
//	@Failure		400	{object}	response.ErrorResponse
//	@Router			/api/v1/users/search [get]
func (h *UserHandler) Search(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		response.Error(c, http.StatusBadRequest, "BAD_REQUEST", "parâmetro q é obrigatório")
		return
	}

	users, err := h.uc.Search(query)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "INTERNAL_ERROR", err.Error())
		return
	}

	c.JSON(http.StatusOK, toUserResponses(users))
}
