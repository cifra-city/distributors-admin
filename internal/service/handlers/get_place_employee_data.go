package handlers

import (
	"net/http"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/comtools/httpkit"
	"github.com/cifra-city/comtools/httpkit/problems"
	"github.com/cifra-city/distributors-admin/internal/config"
	"github.com/cifra-city/distributors-admin/internal/data/nosql/models"
	"github.com/cifra-city/distributors-admin/resources"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func GetPlaceEmployee(w http.ResponseWriter, r *http.Request) {
	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}

	log := Server.Logger

	placeId, err := uuid.Parse(chi.URLParam(r, "place_id"))
	if err != nil {
		log.Errorf("Failed to parse place id: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	userId, err := uuid.Parse(chi.URLParam(r, "user_id"))
	if err != nil {
		log.Errorf("Failed to parse employee id: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	var employee models.PlaceEmployee
	err = Server.MongoDB.PlacesEmployees.Filter().ByPlaceId(placeId).ByUserId(userId).Execute(r.Context(), &employee)
	if err != nil {
		log.Errorf("Failed to get place employee: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to get place employee"))
		return
	}

	httpkit.Render(w, NewPlaceEmployeeResponse(employee))
}

func NewPlaceEmployeeResponse(employee models.PlaceEmployee) *resources.PlaceEmployee {
	return &resources.PlaceEmployee{
		Data: resources.PlaceEmployeeData{
			Id:   employee.ID.String(),
			Type: resources.PlaceEmployeeType,
			Attributes: resources.PlaceEmployeeDataAttributes{
				UserId:     employee.UserID.String(),
				PlaceId:    employee.PlaceID.String(),
				EmployeeId: employee.EmployeeID.String(),
				Role:       employee.Role,
				CreatedAt:  employee.CreatedAt,
				UpdatedAt:  employee.UpdatedAt,
			},
		},
	}
}
