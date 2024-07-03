package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hashicorp-demoapp/product-api-go/data"
	"github.com/hashicorp-demoapp/product-api-go/data/model"
	"github.com/hashicorp/go-hclog"
)

// game -
type Game struct {
	con data.Connection
	log hclog.Logger
}

// Newgame
func NewGame(con data.Connection, l hclog.Logger) *Game {
	return &Game{con, l}
}

func (c *Game) ServeHTTP(gameID *int, rw http.ResponseWriter, r *http.Request) {
	c.log.Info("Handle Order | unknown", "path", r.URL.Path)
	http.NotFound(rw, r)
}


func (c *Game) CreateGame(_ int, rw http.ResponseWriter, r *http.Request) {
	c.log.Info("Handle Game | CreateGame")

	body := model.Game{}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		c.log.Error("Unable to decode JSON", "error", err)
		http.Error(rw, "Unable to parse request body", http.StatusInternalServerError)
		return
	}

	game, err := c.con.CreateGame(body)
	if err != nil {
		c.log.Error("Unable to create new game", "error", err)
		http.Error(rw, fmt.Sprintf("Unable to create new game: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	d, err := game.ToJSON()
	if err != nil {
		c.log.Error("Unable to convert game to JSON", "error", err)
		http.Error(rw, "Unable to create new game", http.StatusInternalServerError)
		return
	}

	rw.Write(d)
}


func (c *Game) GetGame(userID int, rw http.ResponseWriter, r *http.Request){
	c.log.Info("Handle Game | GetGame")

	vars := mux.Vars(r)

	gameID, err := strconv.Atoi(vars["id"])
	if err != nil {
		c.log.Error("orderID provided could not be converted to an integer", "error", err)
		http.Error(rw, "Unable to list order", http.StatusInternalServerError)
		return
	}

	game, err := c.con.GetGame(gameID)
	if err != nil {
		c.log.Error("Unable to get game from database", "error", err)
		http.Error(rw, "Unable to list game", http.StatusInternalServerError)
		return
	}

	g := model.Game{}

	g = game

	d, err := g.ToJSON()
	if err != nil {
		c.log.Error("Unable to convert game to JSON", "error", err)
		http.Error(rw, "Unable to list game", http.StatusInternalServerError)
		return
	}

	rw.Write(d)
}

func (c *Game) UpdateGame(userID int, rw http.ResponseWriter, r *http.Request){
	c.log.Info("Handle Game | UpdateGame")

	// Get orderID
	vars := mux.Vars(r)
	gameID, err := strconv.Atoi(vars["id"])
	if err != nil {
		c.log.Error("gameID provided could not be converted to an integer", "error", err)
		http.Error(rw, "Unable to update game", http.StatusInternalServerError)
		return
	}

	body := model.Game{}

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		c.log.Error("Unable to decode JSON", "error", err)
		http.Error(rw, "Unable to parse request body", http.StatusInternalServerError)
		return
	}

	game, err := c.con.UpdateGame(gameID, body)
	if err != nil {
		c.log.Error("Unable to create new game", "error", err)
		http.Error(rw, "Unable to update game", http.StatusInternalServerError)
		return
	}

	d, err := game.ToJSON()
	if err != nil {
		c.log.Error("Unable to convert game to JSON", "error", err)
		http.Error(rw, "Unable to update game", http.StatusInternalServerError)
	}

	rw.Write(d)
}

func (c *Game) DeleteGame(userID int, rw http.ResponseWriter, r *http.Request) {
	c.log.Info("Handle Game | DeleteGame")

	vars := mux.Vars(r)

	gameID, err := strconv.Atoi(vars["id"])
	if err != nil {
		c.log.Error("gameID provided could not be converted to an integer", "error", err)
		http.Error(rw, "Unable to delete order", http.StatusInternalServerError)
		return
	}

	err = c.con.DeleteGame(gameID)
	if err != nil {
		c.log.Error("Unable to delete game from database", "error", err)
		http.Error(rw, "Unable to delete game", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(rw, "%s", "Deleted game")
}
