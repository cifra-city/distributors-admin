package handlers

import (
	"net/http"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/comtools/httpkit"
	"github.com/cifra-city/comtools/httpkit/problems"
	"github.com/cifra-city/distributors-admin/internal/config"
	"github.com/cifra-city/distributors-admin/internal/data/sql/repositories/sqlcore"
	"github.com/cifra-city/distributors-admin/internal/service/requests"
	"github.com/cifra-city/distributors-admin/resources"
	"github.com/cifra-city/tokens"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func DistributorCreate(w http.ResponseWriter, r *http.Request) {
	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}

	log := Server.Logger

	req, err := requests.NewDistributorCreate(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	userID, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	distributor, err := Server.SqlDB.Distributors.Create(r.Context(), userID, req.Data.Attributes.Name)
	if err != nil {
		log.Errorf("Failed to create distributor: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to create distributor"))
		return
	}

	_, err = Server.SqlDB.DistributorsEmployees.Create(r.Context(), distributor.ID, userID, string(sqlcore.RolesOwner))
	if err != nil {
		log.Errorf("Failed to create distributor staff: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to create distributor staff"))
		return
	}

	httpkit.Render(w, NewDistributorResponse(distributor))
}

func NewDistributorResponse(distributor sqlcore.Distributor) *resources.Distributor {
	return &resources.Distributor{
		Data: resources.DistributorData{
			Id:   distributor.ID.String(),
			Type: "distributor",
			Attributes: resources.DistributorDataAttributes{
				Name:    distributor.Name,
				OwnerId: distributor.OwnerID.String(),
			},
		},
	}
}
