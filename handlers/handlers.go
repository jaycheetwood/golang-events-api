package handlers

import (
	"golang-events-api/errors"
	"golang-events-api/objects"
	"golang-events-api/store"
	"net/http"
)

// IEventHandler is implement all the handlers
type IEventHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	UpdateDetails(w http.ResponseWriter, r *http.Request)
	Cancel(w http.ResponseWriter, r *http.Request)
	Reschedule(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	store store.IEventStore
}

// NewEventHandler return current IEventHandler Implementation
func NewEventHandler(store store.IEventStore) IEventHandler {
	return &handler{store: store}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		WriteError(w, errors.ErrValidEventIdIsRequired)
		return
	}
	evt, err := h.store.Get(r.Context(), &objects.GetRequest{ID: id})
	if err != nil {
		WriteError(w, err)
		return
	}
	WriteResponse(w, &objects.EventResponseWrapper{Event: evt})
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	// after
	after := values.Get("after")
	// name
	name := values.Get("name")
	//limit
	limit, err := IntFromString(w, values.Get("limit"))
	if err != nil {
		return
	}
	// list events
	list, err := h.store.List(r.Context(), &objects.ListRequest{
		Limit: limit,
		Name:  name,
		After: after,
	})
	if err != nil {
		WriteError(w, err)
		return
	}
	WriteResponse(w, &objects.EventResponseWrapper{Events: list})
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *handler) UpdateDetails(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *handler) Cancel(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *handler) Reschedule(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
