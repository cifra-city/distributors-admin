package repositories

import (
	"context"

	"github.com/cifra-city/distributors-admin/internal/data/sql/repositories/sqlcore"
	"github.com/cifra-city/distributors-admin/internal/service/roles"
	"github.com/google/uuid"
)

type DistributorsEmployees interface {
	Create(ctx context.Context, distributorId uuid.UUID, InitiatorId uuid.UUID, newUserId uuid.UUID, role string) (sqlcore.DistributorsEmployee, error)
	CreateOwner(ctx context.Context, distributorId uuid.UUID, userId uuid.UUID) (sqlcore.DistributorsEmployee, error)

	GetByUser(ctx context.Context, distributorId uuid.UUID, userId uuid.UUID) (sqlcore.DistributorsEmployee, error)
	GetByOwner(ctx context.Context, id uuid.UUID) (sqlcore.DistributorsEmployee, error)

	Update(ctx context.Context, EmployeeId uuid.UUID, role string) (sqlcore.DistributorsEmployee, error)
	UpdateByUser(ctx context.Context, distributorId uuid.UUID, InitiatorId uuid.UUID, userForUpdateId uuid.UUID, role string) (sqlcore.DistributorsEmployee, error)

	Delete(ctx context.Context, EmployeeId uuid.UUID) error
	DeleteByUser(ctx context.Context, distributorId uuid.UUID, InitiatorId uuid.UUID, DeletedUserId uuid.UUID) error

	ListByDistributor(ctx context.Context, distributorId uuid.UUID) ([]sqlcore.DistributorsEmployee, error)

	ValidateRoleChange(ctx context.Context, distributorId uuid.UUID, InitiatorId uuid.UUID, secondUserId uuid.UUID, role sqlcore.Roles) error
}

func (d *distributorsEmployees) ValidateRoleChange(ctx context.Context, distributorId uuid.UUID, InitiatorId uuid.UUID, secondUserId uuid.UUID, role sqlcore.Roles) error {
	InitiatorUser, err := d.GetByUser(ctx, distributorId, InitiatorId)
	if err != nil {
		return err
	}
	secondUser, err := d.GetByUser(ctx, distributorId, secondUserId)
	if err != nil {
		return err
	}

	if roles.CompareRoles(InitiatorUser.Role, secondUser.Role) != 1 || roles.CompareRoles(InitiatorUser.Role, role) == -1 {
		return roles.ErrorRolePriority
	}

	if role == sqlcore.RolesOwner || roles.CompareRoles(InitiatorUser.Role, sqlcore.RolesModerator) == -1 {
		return roles.ErrorNoPermission
	}
	return nil
}

type distributorsEmployees struct {
	queries *sqlcore.Queries
}

func NewDistributorsStaff(queries *sqlcore.Queries) DistributorsEmployees {
	return &distributorsEmployees{queries: queries}
}

func (d *distributorsEmployees) Create(ctx context.Context, distributorId uuid.UUID, InitiatorId uuid.UUID, newUserId uuid.UUID, role string) (sqlcore.DistributorsEmployee, error) {
	NewUserRole := sqlcore.Roles(role)
	if !roles.IsValidRole(NewUserRole) {
		return sqlcore.DistributorsEmployee{}, roles.ErrorRole
	}

	err := d.ValidateRoleChange(ctx, distributorId, InitiatorId, newUserId, NewUserRole)
	if err != nil {
		return sqlcore.DistributorsEmployee{}, err
	}

	return d.queries.CreateDistributorEmployees(ctx, sqlcore.CreateDistributorEmployeesParams{
		ID:             uuid.New(),
		DistributorsID: distributorId,
		UserID:         newUserId,
		Role:           NewUserRole,
	})
}

func (d *distributorsEmployees) CreateOwner(ctx context.Context, distributorId uuid.UUID, UserId uuid.UUID) (sqlcore.DistributorsEmployee, error) {
	return d.queries.CreateDistributorEmployees(ctx, sqlcore.CreateDistributorEmployeesParams{
		ID:             uuid.New(),
		DistributorsID: distributorId,
		UserID:         UserId,
		Role:           sqlcore.RolesOwner,
	})
}

func (d *distributorsEmployees) GetByUser(ctx context.Context, distributorId uuid.UUID, userId uuid.UUID) (sqlcore.DistributorsEmployee, error) {
	return d.queries.GetDistributorEmployeesByDistributorIDAndUserID(ctx, sqlcore.GetDistributorEmployeesByDistributorIDAndUserIDParams{
		DistributorsID: distributorId,
		UserID:         userId,
	})
}

func (d *distributorsEmployees) GetByOwner(ctx context.Context, id uuid.UUID) (sqlcore.DistributorsEmployee, error) {
	return d.queries.GetDistributorByOwner(ctx, id)
}

func (d *distributorsEmployees) Update(ctx context.Context, EmployeeId uuid.UUID, role string) (sqlcore.DistributorsEmployee, error) {
	NewUserRole := sqlcore.Roles(role)
	if !roles.IsValidRole(NewUserRole) {
		return sqlcore.DistributorsEmployee{}, roles.ErrorRole
	}
	return d.queries.UpdateDistributorEmployees(ctx, sqlcore.UpdateDistributorEmployeesParams{
		ID:   EmployeeId,
		Role: NewUserRole,
	})
}

func (d *distributorsEmployees) Delete(ctx context.Context, EmployeeId uuid.UUID) error {
	return d.queries.DeleteDistributorEmployees(ctx, EmployeeId)
}

func (d *distributorsEmployees) DeleteByUser(ctx context.Context, distributorId uuid.UUID, InitiatorId uuid.UUID, DeletedUserId uuid.UUID) error {
	InitiatorUser, err := d.GetByUser(ctx, distributorId, InitiatorId)
	if err != nil {
		return err
	}
	secondUser, err := d.GetByUser(ctx, distributorId, DeletedUserId)
	if err != nil {
		return err
	}
	if roles.CompareRoles(InitiatorUser.Role, secondUser.Role) != 1 {
		return roles.ErrorRolePriority
	}

	return d.queries.DeleteDistributorEmployeesByDistributorIDAndUserId(ctx, sqlcore.DeleteDistributorEmployeesByDistributorIDAndUserIdParams{
		DistributorsID: distributorId,
		UserID:         DeletedUserId,
	})
}

func (d *distributorsEmployees) UpdateByUser(ctx context.Context, distributorId uuid.UUID, InitiatorId uuid.UUID, userForUpdateId uuid.UUID, role string) (sqlcore.DistributorsEmployee, error) {
	NewUserRole := sqlcore.Roles(role)
	if !roles.IsValidRole(NewUserRole) {
		return sqlcore.DistributorsEmployee{}, roles.ErrorRole
	}

	err := d.ValidateRoleChange(ctx, distributorId, InitiatorId, userForUpdateId, NewUserRole)
	if err != nil {
		return sqlcore.DistributorsEmployee{}, err
	}

	return d.queries.UpdateDistributorEmployeesByDistributorIDAndUserID(ctx, sqlcore.UpdateDistributorEmployeesByDistributorIDAndUserIDParams{
		DistributorsID: distributorId,
		UserID:         userForUpdateId,
		Role:           NewUserRole,
	})
}

func (d *distributorsEmployees) ListByDistributor(ctx context.Context, distributorId uuid.UUID) ([]sqlcore.DistributorsEmployee, error) {
	return d.queries.GetDistributorEmployeesByDistributorID(ctx, distributorId)
}
