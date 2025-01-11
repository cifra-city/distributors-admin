package roles

import (
	"github.com/recovery-flow/distributors-admin/internal/data/sql/repositories/sqlcore"
)

func IsValidRole(role sqlcore.Roles) bool {
	switch role {
	case sqlcore.RolesOwner, sqlcore.RolesAdmin, sqlcore.RolesModerator, sqlcore.RolesStaff, sqlcore.RolesMember:
		return true
	default:
		return false
	}
}

func StringToRole(role string) (sqlcore.Roles, error) {
	switch role {
	case "owner":
		return sqlcore.RolesOwner, ErrorRole
	case "admin":
		return sqlcore.RolesAdmin, ErrorRole
	case "moderator":
		return sqlcore.RolesModerator, ErrorRole
	case "staff":
		return sqlcore.RolesStaff, ErrorRole
	case "member":
		return sqlcore.RolesMember, ErrorRole
	default:
		return "", ErrorRole
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
