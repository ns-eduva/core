package crm

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID                   primitive.ObjectID `json:"id"`
	AuthUserID           primitive.ObjectID `json:"auth_user_id"`
	FirstName            string             `json:"first_name"`
	LastName             string             `json:"last_name"`
	Type                 string             `json:"type"`             // "student" | "alumni" | "staff" | "teacher"
	EstablishmentID      string             `json:"establishment_id"` // UUID string
	IsEstablishmentAdmin bool               `json:"is_establishment_admin"`
	Status               string             `json:"status"` // "pending" | "approved" | "rejected"
	Meta                 map[string]any     `json:"meta,omitempty"`
	CreatedAt            time.Time          `json:"created_at"`
	UpdatedAt            time.Time          `json:"updated_at"`
}

type Establishment struct {
	ID              string             `json:"id"` // UUID string
	Name            string             `json:"name"`
	Type            string             `json:"type"` // "lycée" | "université" | etc.
	Location        Location           `json:"location"`
	Contact         Contact            `json:"contact"`
	Validated       bool               `json:"validated"`
	CreatedByUserID primitive.ObjectID `json:"created_by_user_id"`
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`
}

type Location struct {
	Address string `json:"address"`
	City    string `json:"city"`
	Zip     string `json:"zip"`
	Country string `json:"country"`
}

type Contact struct {
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
}
