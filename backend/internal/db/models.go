// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserRole string

const (
	UserRoleUser  UserRole = "user"
	UserRoleAdmin UserRole = "admin"
)

func (e *UserRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserRole(s)
	case string:
		*e = UserRole(s)
	default:
		return fmt.Errorf("unsupported scan type for UserRole: %T", src)
	}
	return nil
}

type NullUserRole struct {
	UserRole UserRole `json:"userRole"`
	Valid    bool     `json:"valid"` // Valid is true if UserRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserRole) Scan(value interface{}) error {
	if value == nil {
		ns.UserRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserRole), nil
}

type Category struct {
	Category string `json:"category"`
}

type Offer struct {
	CreatedAt   pgtype.Timestamp `json:"createdAt"`
	OfferID     int32            `json:"offerId"`
	UserID      int32            `json:"userId"`
	Skill       string           `json:"skill"`
	Description string           `json:"description"`
}

type Report struct {
	CreatedAt       pgtype.Timestamp `json:"createdAt"`
	ReportID        int32            `json:"reportId"`
	ReportingUserID int32            `json:"reportingUserId"`
	ReportedOfferID pgtype.Int4      `json:"reportedOfferId"`
	Reason          string           `json:"reason"`
	Description     string           `json:"description"`
	Status          string           `json:"status"`
}

type ReportReason struct {
	Reason string `json:"reason"`
}

type Review struct {
	CreatedAt       pgtype.Timestamp `json:"createdAt"`
	ReviewID        int32            `json:"reviewId"`
	ReviewingUserID int32            `json:"reviewingUserId"`
	ReviewedUserID  int32            `json:"reviewedUserId"`
	StarCount       int32            `json:"starCount"`
	Review          string           `json:"review"`
}

type Skill struct {
	Skill string `json:"skill"`
}

type SkillCategory struct {
	Skill    string `json:"skill"`
	Category string `json:"category"`
}

type User struct {
	CreatedAt       pgtype.Timestamp `json:"createdAt"`
	UserID          int32            `json:"userId"`
	Username        string           `json:"username"`
	DiscordUsername pgtype.Text      `json:"discordUsername"`
	PasswordHash    string           `json:"-"`
	Role            UserRole         `json:"role"`
}
