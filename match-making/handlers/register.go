package handlers

import (
	"encoding/json"
	"match-making/models"
	"net/http"
)

func RegisterPlayer(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newplayer models.Player

	err := json.NewDecoder(r.Body).Decode(&newplayer)

	if err != nil || newplayer.Username == "" {
		http.Error(w, "username missing", http.StatusBadRequest)
		return
	}

	models.WaitingPlayer = append(models.WaitingPlayer, newplayer.Username)

	var response models.MatchResponse

	if len(models.WaitingPlayer) >= 2 {

		player1 := models.WaitingPlayer[0]
		player2 := models.WaitingPlayer[1]

		models.WaitingPlayer = models.WaitingPlayer[2:]

		response = models.MatchResponse{
			Match:   true,
			Message: "Match found! " + player1 + " vs " + player2,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	response = models.MatchResponse{
		Match:   false,
		Message: "still waiting for another player..",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
