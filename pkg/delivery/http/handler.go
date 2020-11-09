package http

import (
	"encoding/json"
	"net/http"

	"github.com/carlos-rodrigo/matching-app/pkg/infrastructure/storage"
	"github.com/carlos-rodrigo/matching-app/pkg/matching"
	"github.com/julienschmidt/httprouter"
)

//Handler interface represents a HTTP Hanlder that can perform a request
type Handler interface {
	Perform(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
}

func writeResponseWithoutData(w http.ResponseWriter, statusCode int, message string) {
	writeResponse(w, statusCode, message, nil)
}
func writeResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	body := ResponseBody{
		Code:    statusCode,
		Message: message,
		Data:    data,
	}
	errEncode := json.NewEncoder(w).Encode(body)
	if errEncode != nil {
		panic(errEncode)
	}
}

//ResponseBody struct represents an standard body to include in a Response
type ResponseBody struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func matchingParticipants() Handler {
	repo := storage.NewCsvParticipantsRepository("./respondents_data_test.csv")
	distance := matching.NewDistanceService()
	score := matching.NewScoreService()
	action := matching.NewMatchingParticipantsAction(repo, distance, score)
	handler := NewMatchingParticipantsHandler(action)

	return handler
}

//GetRouter returns a new configurated Router
func GetRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/matching/", matchingParticipants().Perform)
	return router
}
