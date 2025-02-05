package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/distributors-admin/internal/config"
	"github.com/recovery-flow/tokens"
	"github.com/sirupsen/logrus"
)

func DistributorEmployeeDelete(w http.ResponseWriter, r *http.Request) {
	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}

	log := Server.Logger

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

	userForDeleteId, err := uuid.Parse(chi.URLParam(r, "user_id"))
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
