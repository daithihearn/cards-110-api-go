package game

import (
	"cards-110-api/pkg/api"
	"cards-110-api/pkg/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
// @Param revision query int false "Revision"
// @Security Bearer
// @Success 200 {object} State
// @Success 204 "No Content"
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

	// Get the revision from the request
	revision, exists := c.GetQuery("revision")

	// Get the game from the database
	state, has, err := h.S.GetState(ctx, gameId, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}
	if !has {
		c.JSON(http.StatusNotFound, api.ErrorResponse{Message: "Game not found"})
		return
	}

	// Check if the game has been updated
	if exists {
		// First convert the revision to an int
		// If the revision is less than or equal to the current revision, return no content
		rev, err := strconv.Atoi(revision)
		if err != nil {
			c.JSON(http.StatusBadRequest, api.ErrorResponse{Message: "Invalid revision"})
			return
		}
		if state.Revision <= rev {
			c.Status(http.StatusNoContent)
			return
		}
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
// @Success 200 {object} State
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
	state, err := game.GetState(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, state)
}

type SelectSuitRequest struct {
	Suit  Suit       `json:"suit"`
	Cards []CardName `json:"cards"`
}

// SelectSuit @Summary Select the suit
// @Description When in the Called state, the Goer can select the suit and what cards they want to keep from their hand and the dummy hand
// @Tags Game
// @ID select-suit
// @Produce json
// @Security Bearer
// @Param gameId path string true "Game ID"
// @Para body SelectSuitRequest true "Select Suit Request"
// @Success 200 {object} State
// @Failure 400 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /game/{gameId}/suit [put]
func (h *Handler) SelectSuit(c *gin.Context) {
	// Check the user is correctly authenticated
	id, ok := auth.CheckValidated(c)
	if !ok {
		return
	}

	// Get the context from the request
	ctx := c.Request.Context()

	// Get the game ID from the request
	gameId := c.Param("gameId")

	// Get the request body
	var req SelectSuitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Message: err.Error()})
		return
	}

	// Select the suit
	game, err := h.S.SelectSuit(ctx, gameId, id, req.Suit, req.Cards)

	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}
	state, err := game.GetState(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, state)
}

type BuyRequest struct {
	Cards []CardName `json:"cards"`
}

// Buy @Summary Buy cards
// @Description When in the Buying state, the Goer can buy cards from the deck
// @Tags Game
// @ID buy
// @Produce json
// @Security Bearer
// @Param gameId path string true "Game ID"
// @Para body BuyRequest true "Buy Request"
// @Success 200 {object} State
// @Failure 400 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /game/{gameId}/buy [put]
func (h *Handler) Buy(c *gin.Context) {
	// Check the user is correctly authenticated
	id, ok := auth.CheckValidated(c)
	if !ok {
		return
	}

	// Get the context from the request
	ctx := c.Request.Context()

	// Get the game ID from the request
	gameId := c.Param("gameId")

	// Get the request body
	var req BuyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Message: err.Error()})
		return
	}

	// Buy the cards
	game, err := h.S.Buy(ctx, gameId, id, req.Cards)

	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}
	state, err := game.GetState(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, state)
}

// Play @Summary Play a card
// @Description When in the Playing state, the current player can play a card
// @Tags Game
// @ID play
// @Produce json
// @Security Bearer
// @Param gameId path string true "Game ID"
// @Param card query string true "Card"
// @Success 200 {object} State
// @Failure 400 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /game/{gameId}/play [put]
func (h *Handler) Play(c *gin.Context) {
	// Check the user is correctly authenticated
	id, ok := auth.CheckValidated(c)
	if !ok {
		return
	}

	// Get the context from the request
	ctx := c.Request.Context()

	// Get the game ID from the request
	gameId := c.Param("gameId")

	// Get the card from the request
	card, exists := c.GetQuery("card")
	if !exists {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Message: "Missing card"})
		return
	}
	// Check if is a valid card
	cn, err := ParseCardName(card)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Message: err.Error()})
		return
	}

	// Play the card
	game, err := h.S.Play(ctx, gameId, id, cn)

	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}
	state, err := game.GetState(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, state)
}
