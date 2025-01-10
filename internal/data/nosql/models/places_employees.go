package models

import (
	"time"

	"github.com/google/uuid"
)

type PlaceEmployee struct {
	ID         uuid.UUID `bson:"_id,omitempty"`        // Уникальный идентификатор записи
	PlaceID    uuid.UUID `bson:"place_id"`             // Ссылка на place
	EmployeeID uuid.UUID `bson:"employee_id"`          // Ссылка на пользователя
	UserID     uuid.UUID `bson:"user_id"`              // ID пользователя
	Role       string    `bson:"role"`                 // Роль сотрудника (например, owner, admin)
	CreatedAt  time.Time `bson:"created_at,omitempty"` // Время добавления сотрудника
	UpdatedAt  time.Time `bson:"updated_at,omitempty"` // Время обновления записи
}
