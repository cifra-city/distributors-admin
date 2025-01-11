package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/distributors-admin/internal/config"
	"github.com/recovery-flow/distributors-admin/internal/data/sql/repositories/sqlcore"
	"github.com/recovery-flow/distributors-admin/internal/service/requests"
	"github.com/recovery-flow/distributors-admin/resources"
	"github.com/recovery-flow/tokens"
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

	_, err = Server.SqlDB.DistributorsEmployees.GetByOwner(r.Context(), userID)
	if err != sql.ErrNoRows {
		log.Warn("User already has a distributor")
		httpkit.RenderErr(w, problems.BadRequest(errors.New("one user, one distributor"))...)
		return
	}
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Errorf("Failed to get distributor staff: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to get distributor staff"))
		return
	}

	distributor, err := Server.SqlDB.Distributors.Create(r.Context(), userID, req.Data.Attributes.Name)
	if err != nil {
		log.Errorf("Failed to create distributor: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to create distributor"))
		return
	}

	_, err = Server.SqlDB.DistributorsEmployees.CreateOwner(r.Context(), distributor.ID, userID)
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
