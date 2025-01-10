package handlers

import (
	"net/http"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/comtools/httpkit"
	"github.com/cifra-city/comtools/httpkit/problems"
	"github.com/cifra-city/distributors-admin/internal/config"
	"github.com/cifra-city/distributors-admin/internal/service/requests"
	"github.com/cifra-city/distributors-admin/internal/service/roles"
	"github.com/cifra-city/tokens"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func PlaceEmployeeAdd(w http.ResponseWriter, r *http.Request) {
	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}

	log := Server.Logger

	req, err := requests.NewPlaceEmployeeAdd(r)
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

	distributorId, err := uuid.Parse(req.Data.Attributes.DistributorId)
	if err != nil {
		log.Errorf("Failed to parse distributorId id: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	placeId, err := uuid.Parse(chi.URLParam(r, "place_id"))
	if err != nil {
		log.Errorf("Failed to parse place id: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	curUser, err := Server.SqlDB.DistributorsEmployees.GetByUser(r.Context(), distributorId, userId)
	if err != nil {
		log.Errorf("Failed to get distributor employee: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	role, err := roles.StringToRole(req.Data.Attributes.Role)
	if err != nil {
		log.Errorf("Failed to parse role: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if roles.CompareRoles(curUser.Role, role) == 0 {
		log.Warnf("User %s has no rights to add employee with role %s", userId, req.Data.Attributes.Role)
		httpkit.RenderErr(w, problems.Forbidden("User has no rights to add employee with role %s", req.Data.Attributes.Role))
		return
	}

	username := req.Data.Attributes.Username
	newUserId, err := FetchUserIDFromUsername(Server.Config.Url.UserStorage.GetUser + username)
	if err != nil {
		log.Errorf("Failed to fetch user id for username %s: %v", username, err)
		httpkit.RenderErr(w, problems.NotFound("User not found"))
		return
	}

	distributorEmployee, err := Server.SqlDB.DistributorsEmployees.GetByUser(r.Context(), distributorId, newUserId)
	if err != nil {
		log.Errorf("Failed to get distributorId employee: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	employee, err := Server.MongoDB.PlacesEmployees.Create(
		r.Context(),
		placeId,
		distributorEmployee.ID,
		newUserId,
		req.Data.Attributes.Role,
	)
	if err != nil {
		log.Errorf("Failed to add place employee: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to create employee"))
		return
	}

	httpkit.Render(w, NewPlaceEmployeeResponse(employee))
}
