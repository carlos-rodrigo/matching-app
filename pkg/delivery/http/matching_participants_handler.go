package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/carlos-rodrigo/matching-app/pkg/matching"
	"github.com/julienschmidt/httprouter"
)

type matchingHandler struct {
	Action matching.Action
}

func (h *matchingHandler) Perform(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	project := matching.Project{}

	body, errReadBody := ioutil.ReadAll(r.Body)
	if errReadBody != nil {
		log.Println(errReadBody)
		writeResponseWithoutData(w, http.StatusBadRequest, "Can't read body from request")
		return
	}
	errCloseReadBody := r.Body.Close()
	if errCloseReadBody != nil {
		log.Println(errCloseReadBody)
		writeResponseWithoutData(w, http.StatusBadRequest, "Can't read body from request")
		return
	}
	errUnmarshalProject := json.Unmarshal(body, &project)
	if errUnmarshalProject != nil {
		log.Println(errUnmarshalProject)
		writeResponseWithoutData(w, http.StatusUnprocessableEntity, "Incorrect Body")
		return
	}
	participants, errMatching := h.Action.GetMatchingParticipantsForProject(project)
	if errMatching != nil {
		log.Println(errMatching)
		writeResponseWithoutData(w, http.StatusInternalServerError, errMatching.Error())
		return
	}

	log.Println("Results count: %q", len(participants))
	writeResponse(w, http.StatusOK, "Successful Login!", participants)
}

//NewMatchingParticipantsHandler returns an initialized LoginUserHandler
func NewMatchingParticipantsHandler(action matching.Action) Handler {
	return &matchingHandler{
		Action: action,
	}
}
