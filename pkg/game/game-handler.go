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

// GetStats @Summary Get the user's stats
// @Description Returns stats for the current user
// @Tags Stats
// @ID get-stats
// @Produce json
// @Security Bearer
// @Success 200 {object} PlayerStats
// @Failure 400 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /stats [get]
func (h *Handler) GetStats(c *gin.Context) {
	// Check the user is correctly authenticated
	id, ok := auth.CheckValidated(c)
	if !ok {
		return
	}

	// Get the context from the request
	ctx := c.Request.Context()

	// Get the stats from the database
	stats, err := h.S.GetStats(ctx, id)
	if err != nil {
		c.JSON(http.StatusOK, api.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, stats)
}

// GetStatsForPlayer @Summary Get the stats for a player
// @Description Returns stats for a player
// @Tags Stats
// @ID get-stats-for-player
// @Produce json
// @Param playerId path string true "Player ID"
// @Success 200 {object} PlayerStats
// @Failure 400 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /stats/{playerId} [get]
func (h *Handler) GetStatsForPlayer(c *gin.Context) {
	// Check the user is correctly authenticated
	_, ok := auth.CheckValidated(c)
	if !ok {
		return
	}

	// Get the context from the request
	ctx := c.Request.Context()

	// Get the player ID from the request
	playerId := c.Param("playerId")

	// Get the stats from the database
	stats, err := h.S.GetStats(ctx, playerId)
	if err != nil {
		c.JSON(http.StatusOK, api.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, stats)
}
