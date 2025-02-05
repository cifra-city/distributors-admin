package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/recovery-flow/distributors-admin/internal/data/sql/repositories/sqlcore"
)

type Distributors interface {
	Create(ctx context.Context, ownerId uuid.UUID, name string) (sqlcore.Distributor, error)

	Get(ctx context.Context, id uuid.UUID) (sqlcore.Distributor, error)

	UpdateName(ctx context.Context, id uuid.UUID, name string) (sqlcore.Distributor, error)

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

func (d *distributors) UpdateName(ctx context.Context, id uuid.UUID, name string) (sqlcore.Distributor, error) {
	return d.queries.UpdateDistributorName(ctx, sqlcore.UpdateDistributorNameParams{
		ID:   id,
		Name: name,
	})
}

func (d *distributors) List(ctx context.Context) ([]sqlcore.Distributor, error) {
	return d.queries.ListDistributors(ctx)
}
