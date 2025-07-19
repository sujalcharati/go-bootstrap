package models

type Player struct {
	Username string `json:"username"`
}

type MatchResponse struct {
	Match   bool   `json:"match"`
	Message string `json:"message"`
}

var WaitingPlayer []string
