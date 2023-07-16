package endpoints

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/leonardogregoriocs/internal/contract"
	internalerrors "github.com/leonardogregoriocs/internal/internalErrors"
)

func (h *Handler) TransactionsPost(w http.ResponseWriter, r *http.Request) {
	var request contract.NewTransaction
	err := render.DecodeJSON(r.Body, &request)
	id, err := h.TransactionsService.CreateNewTransactions(request)

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
	render.JSON(w, r, map[string]int{"id": id})
}
