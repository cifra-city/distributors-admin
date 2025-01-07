package repositories

import (
	"context"

	"github.com/cifra-city/distributors-admin/internal/data/sql/repositories/sqlcore"
	"github.com/google/uuid"
)

const (
	RoleOwner = "owner"
	RoleAdmin = "admin"
	RoleStaff = "staff"
)

type DistributorsStaff interface {
	Create(ctx context.Context, distributorId uuid.UUID, userId uuid.UUID, role string) (sqlcore.DistributorsStaff, error)

	Get(ctx context.Context, StaffId uuid.UUID) (sqlcore.DistributorsStaff, error)
	GetByUser(ctx context.Context, distributorId uuid.UUID, userId uuid.UUID) (sqlcore.DistributorsStaff, error)

	Update(ctx context.Context, StaffId uuid.UUID, role string) (sqlcore.DistributorsStaff, error)

	Delete(ctx context.Context, StaffId uuid.UUID) error
	DeleteByUser(ctx context.Context, distributorId uuid.UUID, userId uuid.UUID) error

	ListByDistributor(ctx context.Context, distributorId uuid.UUID) ([]sqlcore.DistributorsStaff, error)
}

type distributorsStaff struct {
	queries *sqlcore.Queries
}

func NewDistributorsStaff(queries *sqlcore.Queries) DistributorsStaff {
	return &distributorsStaff{queries: queries}
}

func (d *distributorsStaff) Create(ctx context.Context, distributorId uuid.UUID, userId uuid.UUID, role string) (sqlcore.DistributorsStaff, error) {
	return d.queries.CreateDistributorStaff(ctx, sqlcore.CreateDistributorStaffParams{
		ID:             uuid.New(),
		DistributorsID: distributorId,
		UserID:         userId,
		Role:           role,
	})
}

func (d *distributorsStaff) Get(ctx context.Context, StaffId uuid.UUID) (sqlcore.DistributorsStaff, error) {
	return d.queries.GetDistributorStaffByID(ctx, StaffId)
}

func (d *distributorsStaff) GetByUser(ctx context.Context, distributorId uuid.UUID, userId uuid.UUID) (sqlcore.DistributorsStaff, error) {
	return d.queries.GetDistributorStaffByDistributorIDAndUserID(ctx, sqlcore.GetDistributorStaffByDistributorIDAndUserIDParams{
		DistributorsID: distributorId,
		UserID:         userId,
	})
}

func (d *distributorsStaff) Update(ctx context.Context, StaffId uuid.UUID, role string) (sqlcore.DistributorsStaff, error) {
	return d.queries.UpdateDistributorStaff(ctx, sqlcore.UpdateDistributorStaffParams{
		ID:   StaffId,
		Role: role,
	})
}

func (d *distributorsStaff) Delete(ctx context.Context, StaffId uuid.UUID) error {
	return d.queries.DeleteDistributorStaff(ctx, StaffId)
}

func (d *distributorsStaff) DeleteByUser(ctx context.Context, distributorId uuid.UUID, userId uuid.UUID) error {
	return d.queries.DeleteDistributorStaffByDistributorIDAndUserId(ctx, sqlcore.DeleteDistributorStaffByDistributorIDAndUserIdParams{
		DistributorsID: distributorId,
		UserID:         userId,
	})
}

func (d *distributorsStaff) ListByDistributor(ctx context.Context, distributorId uuid.UUID) ([]sqlcore.DistributorsStaff, error) {
	return d.queries.GetDistributorStaffByDistributorID(ctx, distributorId)
}
