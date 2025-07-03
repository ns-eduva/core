package auth

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID   `json:"id"`
	Email     string               `json:"email"`
	Password  string               `json:"-"`
	CrmID     string               `json:"crm_id"`
	RoleIDs   []primitive.ObjectID `json:"role_ids"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
}

type Role struct {
	ID            primitive.ObjectID   `json:"id"`
	Name          string               `json:"name"`
	Description   string               `json:"description"`
	PermissionIDs []primitive.ObjectID `json:"permission_ids"`
}

type Permission struct {
	ID          primitive.ObjectID `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
}
