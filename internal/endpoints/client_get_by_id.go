package endpoints

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	internalerrors "github.com/leonardogregoriocs/internal/internalErrors"
)

func (h *Handler) ClientGetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	id_obj, _ := strconv.Atoi(id)

	client, err := h.ClientService.GetBy(id_obj)
	if err != nil {
		if errors.Is(err, internalerrors.ErrInternal) {
			render.Status(r, 500)
			render.JSON(w, r, map[string]string{"Error": err.Error()})
		} else {
			render.Status(r, 400)
			render.JSON(w, r, map[string]string{"Error": err.Error()})
		}
		return
	}

	render.Status(r, 200)
	render.JSON(w, r, client)
}
