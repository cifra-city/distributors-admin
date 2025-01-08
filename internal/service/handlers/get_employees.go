package handlers

import (
	"net/http"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/comtools/httpkit"
	"github.com/cifra-city/comtools/httpkit/problems"
	"github.com/cifra-city/distributors-admin/internal/config"
	"github.com/cifra-city/distributors-admin/internal/data/sql/repositories/sqlcore"
	"github.com/cifra-city/distributors-admin/resources"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}

	log := Server.Logger

	distributorId, err := uuid.Parse(chi.URLParam(r, "distributor_id"))
	if err != nil {
		log.Errorf("Failed to parse distributor id: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	employees, err := Server.SqlDB.DistributorsEmployees.ListByDistributor(r.Context(), distributorId)
	if err != nil {
		log.Errorf("Failed to get distributor staff: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to get distributor staff"))
		return
	}

	httpkit.Render(w, NewEmployeesCollectionResponse(employees))
}

func NewEmployeesCollectionResponse(employees []sqlcore.DistributorsEmployee) *resources.EmployeeCollection {
	var collectionDataAttributes []resources.EmployeeCollectionDataAttributesInner

	for _, employee := range employees {
		collectionDataAttributes = append(collectionDataAttributes, resources.EmployeeCollectionDataAttributesInner{
			Data: resources.EmployeeData{
				Id:   employee.ID.String(),
				Type: resources.DistributorEmployeeType,
				Attributes: resources.EmployeeDataAttributes{
					UserId:        employee.UserID.String(),
					Role:          string(employee.Role),
					DistributorId: employee.DistributorsID.String(),
					CreatedAt:     employee.CreatedAt.Time,
				},
			},
		})
	}

	return &resources.EmployeeCollection{
		Data: resources.EmployeeCollectionData{
			Type:       resources.DistributorEmployeeCollectionType,
			Attributes: collectionDataAttributes,
		},
	}
}
