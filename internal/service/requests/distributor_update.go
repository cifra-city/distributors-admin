package requests

import (
	"encoding/json"
	"net/http"

	"github.com/cifra-city/comtools"
	"github.com/cifra-city/distributors-admin/resources"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func NewDistributorUpdate(r *http.Request) (req resources.DistributorUpdate, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = comtools.NewDecodeError("body", err)
		return
	}

	errs := validation.Errors{
		"data/id":         validation.Validate(req.Data.Id, validation.Required),
		"data/type":       validation.Validate(req.Data.Type, validation.Required, validation.In(resources.DistributorUpdateType)),
		"data/attributes": validation.Validate(req.Data.Attributes, validation.Required),
	}
	return req, errs.Filter()
}
