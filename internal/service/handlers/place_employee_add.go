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

	_, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID) // TODO add auth for this user
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

	username := req.Data.Attributes.Username
	userId, err := FetchUserIDFromUsername(Server.Config.Url.UserStorage.GetUser + username)
	if err != nil {
		log.Errorf("Failed to fetch user id for username %s: %v", username, err)
		httpkit.RenderErr(w, problems.NotFound("User not found"))
		return
	}

	distributor, err := uuid.Parse(req.Data.Attributes.DistributorId)
	if err != nil {
		log.Errorf("Failed to parse distributor id: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	distributorEmployee, err := Server.SqlDB.DistributorsEmployees.GetByUser(r.Context(), distributor, userId)
	if err != nil {
		log.Errorf("Failed to get distributor employee: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	employee, err := Server.MongoDB.PlacesEmployees.Add(
		r.Context(),
		placeId,
		distributorEmployee.ID,
		userId,
		req.Data.Attributes.Role,
	)
	if err != nil {
		log.Errorf("Failed to add place employee: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, NewPlaceEmployeeResponse(employee))
}
