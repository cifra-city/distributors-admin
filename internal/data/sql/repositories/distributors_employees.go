package repositories

import (
	"context"

	"github.com/cifra-city/distributors-admin/internal/data/sql/repositories/sqlcore"
	"github.com/google/uuid"
	"github.com/pkg/errors"
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
}

type distributorsEmployees struct {
	queries *sqlcore.Queries
}

func NewDistributorsStaff(queries *sqlcore.Queries) DistributorsEmployees {
	return &distributorsEmployees{queries: queries}
}

var ErrorRole = errors.New("role must be one of: owner, admin, moderator, staff, member")
var ErrorNoPermission = errors.New("User is have not permission")
var ErrorRolePriority = errors.New("You can't change/delete user with the same or higher role")

func isValidRole(role sqlcore.Roles) bool {
	switch role {
	case sqlcore.RolesOwner, sqlcore.RolesAdmin, sqlcore.RolesModerator, sqlcore.RolesStaff, sqlcore.RolesMember:
		return true
	default:
		return false
	}
}

//	1, if first role is higher priority
//
// -1, if second role is higher priority
//
//	0, if roles are equal
func CompareRoles(role1, role2 sqlcore.Roles) int {
	// Маппинг ролей на их приоритеты
	priority := map[sqlcore.Roles]int{
		sqlcore.RolesOwner:     5,
		sqlcore.RolesAdmin:     4,
		sqlcore.RolesModerator: 3,
		sqlcore.RolesStaff:     2,
		sqlcore.RolesMember:    1,
	}

	p1, ok1 := priority[role1]
	p2, ok2 := priority[role2]

	// Если какая-то из ролей не существует в маппинге
	if !ok1 || !ok2 {
		panic("Invalid role provided")
	}

	// Сравниваем приоритеты
	if p1 > p2 {
		return 1
	} else if p1 < p2 {
		return -1
	}
	return 0
}

func (d *distributorsEmployees) Create(ctx context.Context, distributorId uuid.UUID, InitiatorId uuid.UUID, newUserId uuid.UUID, role string) (sqlcore.DistributorsEmployee, error) {
	InitiatorUser, err := d.GetByUser(ctx, distributorId, InitiatorId)
	if err != nil {
		return sqlcore.DistributorsEmployee{}, err
	}

	NewUserRole := sqlcore.Roles(role)
	if !isValidRole(NewUserRole) {
		return sqlcore.DistributorsEmployee{}, ErrorRole
	}

	if NewUserRole == sqlcore.RolesOwner || CompareRoles(InitiatorUser.Role, sqlcore.RolesModerator) == -1 || CompareRoles(InitiatorUser.Role, NewUserRole) == -1 {
		return sqlcore.DistributorsEmployee{}, ErrorNoPermission
	}

	UserRole := sqlcore.Roles(role)
	if !isValidRole(UserRole) {
		return sqlcore.DistributorsEmployee{}, ErrorRole
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
	UserRole := sqlcore.Roles(role)
	if !isValidRole(UserRole) {
		return sqlcore.DistributorsEmployee{}, ErrorRole
	}
	return d.queries.UpdateDistributorEmployees(ctx, sqlcore.UpdateDistributorEmployeesParams{
		ID:   EmployeeId,
		Role: UserRole,
	})
}

func (d *distributorsEmployees) Delete(ctx context.Context, EmployeeId uuid.UUID) error {
	return d.queries.DeleteDistributorEmployees(ctx, EmployeeId)
}

func (d *distributorsEmployees) DeleteByUser(ctx context.Context, distributorId uuid.UUID, InitiatorId uuid.UUID, DeletedUserId uuid.UUID) error {
	userFoDelete, err := d.GetByUser(ctx, distributorId, DeletedUserId)
	if err != nil {
		return err
	}
	InitiatorUser, err := d.GetByUser(ctx, distributorId, InitiatorId)
	if err != nil {
		return err
	}
	if CompareRoles(InitiatorUser.Role, userFoDelete.Role) != 1 {
		return ErrorRolePriority
	}

	return d.queries.DeleteDistributorEmployeesByDistributorIDAndUserId(ctx, sqlcore.DeleteDistributorEmployeesByDistributorIDAndUserIdParams{
		DistributorsID: distributorId,
		UserID:         DeletedUserId,
	})
}

func (d *distributorsEmployees) UpdateByUser(ctx context.Context, distributorId uuid.UUID, InitiatorId uuid.UUID, userForUpdateId uuid.UUID, role string) (sqlcore.DistributorsEmployee, error) {
	InitiatorUser, err := d.GetByUser(ctx, distributorId, InitiatorId)
	if err != nil {
		return sqlcore.DistributorsEmployee{}, err
	}
	UserFoUpdate, err := d.GetByUser(ctx, distributorId, userForUpdateId)
	if err != nil {
		return sqlcore.DistributorsEmployee{}, err
	}
	NewUserRole := sqlcore.Roles(role)
	if !isValidRole(NewUserRole) {
		return sqlcore.DistributorsEmployee{}, ErrorRole
	}

	if CompareRoles(InitiatorUser.Role, UserFoUpdate.Role) != 1 || CompareRoles(InitiatorUser.Role, NewUserRole) == -1 {
		return sqlcore.DistributorsEmployee{}, ErrorRolePriority
	}

	if NewUserRole == sqlcore.RolesOwner || CompareRoles(InitiatorUser.Role, sqlcore.RolesModerator) == -1 {
		return sqlcore.DistributorsEmployee{}, ErrorNoPermission
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
