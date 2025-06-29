package auth

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Email     string               `bson:"email" json:"email"`
	Password  string               `bson:"password" json:"-"`
	CrmID     string               `bson:"crm_id" json:"crm_id"`
	RoleIDs   []primitive.ObjectID `bson:"role_ids" json:"role_ids"`
	CreatedAt time.Time            `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time            `bson:"updated_at" json:"updated_at"`
}

type Role struct {
	ID            primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Name          string               `bson:"name" json:"name"`
	Description   string               `bson:"description" json:"description"`
	PermissionIDs []primitive.ObjectID `bson:"permission_ids" json:"permission_ids"`
}

type Permission struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
}
