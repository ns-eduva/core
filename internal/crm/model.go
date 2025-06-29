package crm

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	AuthUserID           primitive.ObjectID `bson:"auth_user_id" json:"auth_user_id"`
	FirstName            string             `bson:"first_name" json:"first_name"`
	LastName             string             `bson:"last_name" json:"last_name"`
	Type                 string             `bson:"type" json:"type"`                         // "student" | "alumni" | "staff" | "teacher"
	EstablishmentID      string             `bson:"establishment_id" json:"establishment_id"` // UUID string
	IsEstablishmentAdmin bool               `bson:"is_establishment_admin" json:"is_establishment_admin"`
	Status               string             `bson:"status" json:"status"` // "pending" | "approved" | "rejected"
	Meta                 map[string]any     `bson:"meta,omitempty" json:"meta,omitempty"`
	CreatedAt            time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt            time.Time          `bson:"updated_at" json:"updated_at"`
}

type Establishment struct {
	ID              string             `bson:"_id" json:"id"` // UUID string
	Name            string             `bson:"name" json:"name"`
	Type            string             `bson:"type" json:"type"` // "lycée" | "université" | etc.
	Location        Location           `bson:"location" json:"location"`
	Contact         Contact            `bson:"contact" json:"contact"`
	Validated       bool               `bson:"validated" json:"validated"`
	CreatedByUserID primitive.ObjectID `bson:"created_by_user_id" json:"created_by_user_id"`
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updated_at"`
}

type Location struct {
	Address string `bson:"address" json:"address"`
	City    string `bson:"city" json:"city"`
	Zip     string `bson:"zip" json:"zip"`
	Country string `bson:"country" json:"country"`
}

type Contact struct {
	Email   string `bson:"email" json:"email"`
	Phone   string `bson:"phone" json:"phone"`
	Website string `bson:"website" json:"website"`
}
