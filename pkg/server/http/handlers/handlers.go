package handlers

import (
	"github.com/AnkushJadhav/kamaji-root/logger"
)

// Response500 is the action to performon HTTP 500 internal server error
type Response500 struct {
	ID      string `json:"id"`
	Message string `json:"msg"`
}

// Handle500InternalServerError handles a generic HTTP 500 event
func Handle500InternalServerError(id string, err error) Response500 {
	logger.Errorln(err)
	return Response500{
		ID:      id,
		Message: "Oops! Something went wrong!",
	}
}
