package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/distributors-admin/internal/config"
	"github.com/recovery-flow/distributors-admin/internal/service/requests"
	"github.com/recovery-flow/tokens"
	"github.com/sirupsen/logrus"
)

func DistributorEmployeeUpdate(w http.ResponseWriter, r *http.Request) {
	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}

	log := Server.Logger

	req, err := requests.NewDistributorEmployeeUpdate(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	InitiatorId, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	distributorId, err := uuid.Parse(chi.URLParam(r, "distributor_id"))
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	userForUpdateId, err := uuid.Parse(chi.URLParam(r, "user_id"))
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	newEmployee, err := Server.SqlDB.DistributorsEmployees.UpdateByUser(r.Context(), distributorId, InitiatorId, userForUpdateId, req.Data.Attributes.Role)
	if err != nil {
		if err.Error() == "role must be one of: admin, moderator, staff, member" {
			log.Errorf("Failed to create distributor staff: %v", err)
			httpkit.RenderErr(w, problems.BadRequest(err)...)
			return
		}
		log.Errorf("Failed to update distributor employee: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to update distributor employee"))
		return
	}

	log.Infof("Staff added to distributor %v", distributorId)
	httpkit.Render(w, NewDistributorEmployeeResponse(newEmployee))
}
