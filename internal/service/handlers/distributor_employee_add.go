package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/comtools/httpkit"
	"github.com/cifra-city/comtools/httpkit/problems"
	"github.com/cifra-city/distributors-admin/internal/config"
	"github.com/cifra-city/distributors-admin/internal/data/sql/repositories/sqlcore"
	"github.com/cifra-city/distributors-admin/internal/service/requests"
	"github.com/cifra-city/distributors-admin/internal/service/roles"
	"github.com/cifra-city/distributors-admin/resources"
	"github.com/cifra-city/tokens"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func DistributorEmployeeAdd(w http.ResponseWriter, r *http.Request) {
	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}

	log := Server.Logger

	req, err := requests.NewDistributorEmployeeAdd(r)
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

	username := req.Data.Attributes.Username
	NewUserId, err := FetchUserIDFromUsername(Server.Config.Url.UserStorage.GetUser + username)
	if err != nil {
		log.Errorf("Failed to fetch user id for username %s: %v", username, err)
		httpkit.RenderErr(w, problems.NotFound("User not found"))
		return
	}

	_, err = Server.SqlDB.DistributorsEmployees.GetByUser(r.Context(), distributorId, InitiatorId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Errorf("Failed to get distributor staff: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to get distributor staff"))
		return
	}

	newEmployee, err := Server.SqlDB.DistributorsEmployees.Create(r.Context(), distributorId, InitiatorId, NewUserId, req.Data.Attributes.Role)
	if err != nil {
		if errors.Is(err, roles.ErrorRole) {
			httpkit.RenderErr(w, problems.BadRequest(roles.ErrorRole)...)
			return
		}
		if errors.Is(err, roles.ErrorRolePriority) || errors.Is(err, roles.ErrorNoPermission) {
			httpkit.RenderErr(w, problems.Forbidden())
			return
		}
		log.Errorf("Failed to create distributor staff: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to create distributor staff"))
		return
	}

	log.Infof("Staff added to distributor %v", distributorId)
	httpkit.Render(w, NewDistributorEmployeeResponse(newEmployee))
}

func NewDistributorEmployeeResponse(employees sqlcore.DistributorsEmployee) *resources.DistributorEmployee {
	return &resources.DistributorEmployee{
		Data: resources.DistributorEmployeeData{
			Id:   employees.ID.String(),
			Type: resources.DistributorEmployeeType,
			Attributes: resources.DistributorEmployeeDataAttributes{
				UserId:        employees.UserID.String(),
				Role:          string(employees.Role),
				DistributorId: employees.DistributorsID.String(),
				CreatedAt:     employees.CreatedAt.Time,
			},
		},
	}
}

func FetchUserIDFromUsername(url string) (uuid.UUID, error) {
	resp, err := http.Get(url)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return uuid.Nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var responseData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return uuid.Nil, fmt.Errorf("failed to decode response: %w", err)
	}

	userIDStr, ok := responseData["data"].(map[string]interface{})["attributes"].(map[string]interface{})["id"].(string)
	if !ok {
		return uuid.Nil, fmt.Errorf("user_id not found in response")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid user ID format: %w", err)
	}

	return userID, nil
}
