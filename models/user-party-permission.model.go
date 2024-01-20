package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role string 

const (
	ADMIN 	Role = "ADMIN"
	VIEWER	Role = "VIEWER"
)

type UserPartyPermission struct {
	ID				string		`gorm:"primaryKey" json:"id"`
	UserId			string		`json:"user_id"`
	User			User		`json:"user"`
	PartyId			string		`json:"party_id"`
	Party			Party		`json:"party"`
	Permission		Role		`json:"permission"`
}

func (u *UserPartyPermission) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}