package game

import (
	"cards-110-api/pkg/api"
	"cards-110-api/pkg/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	S ServiceI
}

type CreateGameRequest struct {
	PlayerIDs []string `json:"players"`
	Name      string   `json:"name"`
}

// Create @Summary Create a new game
// @Description Creates a new game with the given name and players
// @Tags Game
// @ID create-game
// @Accept json
// @Produce json
// @Param game body CreateGameRequest true "Game"
// @Security Bearer
// @Success 200 {object} Game
// @Failure 400 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /game [put]
func (h *Handler) Create(c *gin.Context) {
	// Check the user is correctly authenticated
	id, ok := auth.CheckValidated(c)
	if !ok {
		return
	}

	// Get the context from the request
	ctx := c.Request.Context()

	// Get the request body
	var req CreateGameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Message: err.Error()})
		return
	}

	// Create the game
	game, err := h.S.Create(ctx, req.PlayerIDs, req.Name, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, game)
}

// Get @Summary Get a game
// @Description Returns a game with the given ID
// @Tags Game
// @ID get-game
// @Produce json
// @Param gameId path string true "Game ID"
// @Security Bearer
// @Success 200 {object} Game
// @Failure 400 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /game/{gameId} [get]
func (h *Handler) Get(c *gin.Context) {
	// Check the user is correctly authenticated
	_, ok := auth.CheckValidated(c)
	if !ok {
		return
	}

	// Get the context from the request
	ctx := c.Request.Context()

	// Get the game ID from the request
	gameId := c.Param("gameId")

	// Get the game from the database
	game, has, err := h.S.Get(ctx, gameId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}
	if !has {
		c.JSON(http.StatusNotFound, api.ErrorResponse{Message: "Game not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, game)
}

// GetState @Summary Get the state of a game
// @Description Returns the state of a game with the given ID for the current user
// @Tags Game
// @ID get-game-state
// @Produce json
// @Param gameId path string true "Game ID"
// @Security Bearer
// @Success 200 {object} GameState
// @Failure 400 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /game/{gameId}/state [get]
func (h *Handler) GetState(c *gin.Context) {
	// Check the user is correctly authenticated
	id, ok := auth.CheckValidated(c)
	if !ok {
		return
	}

	// Get the context from the request
	ctx := c.Request.Context()

	// Get the game ID from the request
	gameId := c.Param("gameId")

	// Get the game from the database
	game, has, err := h.S.Get(ctx, gameId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}
	if !has {
		c.JSON(http.StatusNotFound, api.ErrorResponse{Message: "Game not found"})
		return
	}

	// Get the state for the current user
	state, err := game.GetState(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, state)
}

// GetAll @Summary Get all games
// @Description Returns all games
// @Tags Game
// @ID get-all-games
// @Produce json
// @Security Bearer
// @Success 200 {array} Game
// @Failure 400 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /game/all [get]
func (h *Handler) GetAll(c *gin.Context) {
	// Check the user is correctly authenticated
	_, ok := auth.CheckValidated(c)
	if !ok {
		return
	}

	// Get the context from the request
	ctx := c.Request.Context()

	// Get all games from the database
	games, err := h.S.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, games)
}

// Delete @Summary Delete a game
// @Description Deletes a game with the given ID
// @Tags Game
// @ID delete-game
// @Security Bearer
// @Param gameId path string true "Game ID"
// @Success 200
// @Failure 400 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /game/{gameId} [delete]
func (h *Handler) Delete(c *gin.Context) {
	// Check the user is correctly authenticated
	id, ok := auth.CheckValidated(c)
	if !ok {
		return
	}

	// Get the context from the request
	ctx := c.Request.Context()

	// Get the game ID from the request
	gameId := c.Param("gameId")

	// Cancel the game
	err := h.S.Delete(ctx, gameId, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// Call @Summary Make a call
// @Description Makes a call for the current user in the game with the given ID
// @Tags Game
// @ID call
// @Produce json
// @Security Bearer
// @Param gameId path string true "Game ID"
// @Param call query int true "Call"
// @Success 200 {object} Game
// @Failure 400 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /game/{gameId}/call [put]
func (h *Handler) Call(c *gin.Context) {
	// Check the user is correctly authenticated
	id, ok := auth.CheckValidated(c)
	if !ok {
		return
	}

	// Get the context from the request
	ctx := c.Request.Context()

	// Get the game ID from the request
	gameId := c.Param("gameId")

	// Get the call from the request
	ca, exists := c.GetQuery("call")
	if !exists {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Message: "Missing call"})
		return
	}
	// Check if is a valid call
	call, err := ParseCall(ca)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Message: err.Error()})
		return
	}

	// Make the call
	game, err := h.S.Call(ctx, gameId, id, call)

	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, game)
}
