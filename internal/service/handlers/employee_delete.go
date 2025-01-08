package handlers

import (
	"net/http"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/comtools/httpkit"
	"github.com/cifra-city/comtools/httpkit/problems"
	"github.com/cifra-city/distributors-admin/internal/config"
	"github.com/cifra-city/distributors-admin/internal/service/requests"
	"github.com/cifra-city/tokens"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func EmployeeDelete(w http.ResponseWriter, r *http.Request) {
	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}

	log := Server.Logger

	req, err := requests.NewEmployeeDelete(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	InitiatorId, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	distributorId, err := uuid.Parse(chi.URLParam(r, "distributor_id"))
	if err != nil {
		log.Errorf("Failed to parse distributor id: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	userForDeleteId, err := uuid.Parse(req.Data.Attributes.UserId)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	err = Server.SqlDB.DistributorsEmployees.DeleteByUser(r.Context(), distributorId, InitiatorId, userForDeleteId)
	if err != nil {
		log.Errorf("Failed to delete employee: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to delete employee"))
		return
	}

	log.Infof("Employee %s deleted", userForDeleteId)
	httpkit.Render(w, http.StatusOK)
}
