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
	"github.com/recovery-flow/distributors-admin/internal/service/roles"
	"github.com/recovery-flow/tokens"
	"github.com/sirupsen/logrus"
)

func PlaceEmployeeUpdate(w http.ResponseWriter, r *http.Request) {
	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}

	log := Server.Logger

	req, err := requests.NewPlaceEmployeeUpdate(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	userId, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID) // TODO add auth for this user
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	placeId, err := uuid.Parse(chi.URLParam(r, "place_id"))
	if err != nil {
		log.Errorf("Failed to parse place id: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	userForUpdateId, err := uuid.Parse(chi.URLParam(r, "user_id"))
	if err != nil {
		log.Errorf("Failed to parse employee id: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	employees, err := Server.MongoDB.PlacesEmployees.FilterByPlaceId(placeId).FilterByUserId(userId).Get(r.Context())
	if err != nil {
		log.Errorf("Failed to get distributor employee: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	if len(employees) != 1 {
		if len(employees) == 0 {
			httpkit.RenderErr(w, problems.NotFound())
		} else {
			log.Errorf("More than one employee found %s %s %s", placeId, userForUpdateId, employees)
			httpkit.RenderErr(w, problems.InternalError())
		}
		return
	}
	curUser := employees[0]

	employees, err = Server.MongoDB.PlacesEmployees.FilterByPlaceId(placeId).FilterByUserId(userForUpdateId).Get(r.Context())
	if err != nil {
		log.Errorf("Failed to get place employee: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to get place employee"))
		return
	}

	if len(employees) != 1 {
		if len(employees) == 0 {
			httpkit.RenderErr(w, problems.NotFound())
		} else {
			log.Errorf("More than one employee found %s %s %s", placeId, userForUpdateId, employees)
			httpkit.RenderErr(w, problems.InternalError())
		}
		return
	}
	updatedUser := employees[0]

	curRole, err := roles.StringToRole(curUser.Role)
	if err != nil {
		log.Errorf("Failed to parse role: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	upRole, err := roles.StringToRole(updatedUser.Role)
	if err != nil {
		log.Errorf("Failed to parse role: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if roles.CompareRoles(curRole, upRole) == 0 {
		log.Warnf("User %s has no rights to add employee with role %s", userId, curRole)
		httpkit.RenderErr(w, problems.Forbidden("User has no rights to add employee with role"))
		return
	}

	_, err = Server.MongoDB.PlacesEmployees.FilterByPlaceId(placeId).FilterByUserId(updatedUser.ID).Update(r.Context(), req.Data.Attributes.Role)
	if err != nil {
		log.Errorf("Failed to get place employee: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to get place employee"))
		return
	}

	employees, err = Server.MongoDB.PlacesEmployees.FilterByPlaceId(placeId).FilterByUserId(userForUpdateId).Get(r.Context())
	if len(employees) != 1 {
		if len(employees) == 0 {
			httpkit.RenderErr(w, problems.NotFound())
		} else {
			log.Errorf("More than one employee found %s %s %s", placeId, userForUpdateId, employees)
			httpkit.RenderErr(w, problems.InternalError())
		}
		return
	}

	httpkit.Render(w, NewPlaceEmployeeResponse(employees[0]))
}
