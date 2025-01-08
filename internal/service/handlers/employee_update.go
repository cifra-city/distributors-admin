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

func EmployeeUpdate(w http.ResponseWriter, r *http.Request) {
	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}

	log := Server.Logger

	req, err := requests.NewEmployeeUpdate(r)
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

	userForUpdateId, err := uuid.Parse(req.Data.Id)
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
	httpkit.Render(w, NewEmployeeResponse(newEmployee))
}
