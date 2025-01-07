package repositories

import (
	"context"

	"github.com/cifra-city/distributors-admin/internal/data/sql/repositories/sqlcore"
	"github.com/google/uuid"
)

type Distributors interface {
	Create(ctx context.Context, ownerId uuid.UUID, name string) (sqlcore.Distributor, error)

	Get(ctx context.Context, id uuid.UUID) (sqlcore.Distributor, error)

	UpdateName(ctx context.Context, id string, name string) (sqlcore.Distributor, error)

	List(ctx context.Context) ([]sqlcore.Distributor, error)
}

type distributors struct {
	queries *sqlcore.Queries
}

func NewDistributors(queries *sqlcore.Queries) Distributors {
	return &distributors{queries: queries}
}

func (d *distributors) Create(ctx context.Context, ownerId uuid.UUID, name string) (sqlcore.Distributor, error) {
	return d.queries.CreateDistributor(ctx, sqlcore.CreateDistributorParams{
		ID:      uuid.New(),
		Name:    name,
		OwnerID: ownerId,
	})
}

func (d *distributors) Get(ctx context.Context, id uuid.UUID) (sqlcore.Distributor, error) {
	return d.queries.GetDistributorByID(ctx, id)
}

func (d *distributors) UpdateName(ctx context.Context, id string, name string) (sqlcore.Distributor, error) {
	distributorId, err := uuid.Parse(id)
	if err != nil {
		return sqlcore.Distributor{}, err
	}
	return d.queries.UpdateDistributorName(ctx, sqlcore.UpdateDistributorNameParams{
		ID:   distributorId,
		Name: name,
	})
}

func (d *distributors) List(ctx context.Context) ([]sqlcore.Distributor, error) {
	return d.queries.ListDistributors(ctx)
}
