package graph

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var _ http.Handler = (*Handler)(nil)

type matrix struct {
	Data [][]int `json:"data"`
}

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	var mtx matrix
	if err := d.Decode(&mtx); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		log.Print(err.Error())
		return
	}

	edges, err := GetEdges(mtx.Data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		log.Print(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(edges); err != nil {
		log.Print(err.Error())
	}
}
