// internal/delivery/http/handlers/user_handler.go
package handlers

import (
	"net/http"
	"strconv"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/Turgho/Aluguei/internal/domain/repositories"
	"github.com/Turgho/Aluguei/internal/domain/usecases"
	"github.com/Turgho/Aluguei/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UserHandler agrupa os handlers HTTP relacionados a [entities.User].
type UserHandler struct {
	uc usecases.UserUseCase
}

// NewUserHandler retorna uma instância de [UserHandler].
func NewUserHandler(uc usecases.UserUseCase) *UserHandler {
	return &UserHandler{uc: uc}
}

// ── Request / Response ─────────────────────────────────────────────────────

// userResponse representa o corpo da resposta do usuário.
type userResponse struct {
	ID        string        `json:"id"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	CPF       string        `json:"-"`
	Email     string        `json:"email"`
	Phone     string        `json:"phone"`
	Role      entities.Role `json:"role"`
}

// updateRequest representa o corpo da requisição de atualização.
type updateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

// ── Helpers ────────────────────────────────────────────────────────────────

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

// parseUUID extrai e valida um UUID de um parâmetro de rota.
func parseUUID(c *gin.Context, param string) (uuid.UUID, bool) {
	id, err := uuid.Parse(c.Param(param))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_ID", "ID inválido")
		return uuid.Nil, false
	}
	return id, true
}

// ── Handlers ───────────────────────────────────────────────────────────────

// GetByID godoc
//
//	@Summary		Busca um usuário pelo ID
//	@Tags			users
//	@Produce		json
//	@Security		CookieAuth
//	@Param			id	path		string	true	"ID do usuário"
//	@Success		200	{object}	userResponse
//	@Failure		400	{object}	response.ErrorResponse
//	@Failure		404	{object}	response.ErrorResponse
//	@Router			/api/v1/users/{id} [get]
func (h *UserHandler) GetByID(c *gin.Context) {
	id, ok := parseUUID(c, "id")
	if !ok {
		return
	}

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
//	@Security		CookieAuth
//	@Param			id		path		string			true	"ID do usuário"
//	@Param			body	body		updateRequest	true	"Dados a atualizar"
//	@Success		200		{object}	userResponse
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		404		{object}	response.ErrorResponse
//	@Router			/api/v1/users/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	id, ok := parseUUID(c, "id")
	if !ok {
		return
	}

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
//	@Security		CookieAuth
//	@Param			id	path		string	true	"ID do usuário"
//	@Success		204
//	@Failure		400	{object}	response.ErrorResponse
//	@Failure		404	{object}	response.ErrorResponse
//	@Router			/api/v1/users/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	id, ok := parseUUID(c, "id")
	if !ok {
		return
	}

	if err := h.uc.Delete(id); err != nil {
		response.Error(c, http.StatusNotFound, "USER_NOT_FOUND", "usuário não encontrado")
		return
	}

	c.Status(http.StatusNoContent)
}

// Search godoc
//
//	@Summary		Busca usuários com filtros e paginação
//	@Tags			users
//	@Produce		json
//	@Security		CookieAuth
//	@Param			q				query	string	false	"Termo de busca (nome, email ou CPF)"
//	@Param			role			query	string	false	"Filtro por role"
//	@Param			is_active		query	bool	false	"Filtro por status ativo"
//	@Param			email_verified	query	bool	false	"Filtro por email verificado"
//	@Param			page			query	int		false	"Página (default: 1)"
//	@Param			page_size		query	int		false	"Itens por página (default: 20)"
//	@Success		200	{object}	pagination.Result[userResponse]
//	@Failure		400	{object}	response.ErrorResponse
//	@Router			/api/v1/users/search [get]
func (h *UserHandler) Search(c *gin.Context) {
	var filters repositories.UserFilters

	// Query params
	filters.Search = c.Query("search")

	// Paginação
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	filters.Page = page
	filters.PageSize = pageSize

	result, err := h.uc.Search(filters)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "INTERNAL_ERROR", err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}
