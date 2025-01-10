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

func GetPlacesEmployees(w http.ResponseWriter, r *http.Request) {
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

	employees, err := Server.MongoDB.PlacesEmployees.FilterByPlaceId(placeId).Get(r.Context())
	if err != nil {
		log.Errorf("Failed to get place staff: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to get place staff"))
		return
	}
	httpkit.Render(w, NewEmployeesPlaceCollectionResponse(employees))
}

func NewEmployeesPlaceCollectionResponse(employees []models.PlaceEmployee) *resources.PlaceEmployeeCollection {
	var collectionDataAttributes []resources.PlaceEmployeeCollectionDataAttributesInner

	for _, employee := range employees {
		collectionDataAttributes = append(collectionDataAttributes, resources.PlaceEmployeeCollectionDataAttributesInner{
			Data: NewPlaceEmployeeResponse(employee).Data,
		})
	}

	return &resources.PlaceEmployeeCollection{
		Data: resources.PlaceEmployeeCollectionData{
			Type:       resources.PlaceEmployeeCollectionType,
			Attributes: collectionDataAttributes,
		},
	}
}
